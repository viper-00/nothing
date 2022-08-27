package memdb

import (
	"fmt"
	"reflect"
	"sync"
)

const (
	Int     string = "INT"
	Int64   string = "INT64"
	Float32 string = "FLOAT32"
	Float64 string = "FLOAT64"
	Bool    string = "BOOL"
	String  string = "STRING"
)

type Col struct {
	Name       string
	Type       string
	IntVal     int
	Int64Val   int64
	Float32Val float32
	Float64Val float64
	BoolVal    bool
	StringVal  string
}

type Row struct {
	Id      int
	Columns map[string]Col
}

type Table struct {
	Name string
	sync.Mutex
	Columns  map[string]Col
	Rows     map[int]Row
	RowCount int
	Db       *Database
}

type Database struct {
	Name   string
	Tables map[string]*Table
}

func CreateDatabase(name string) Database {
	return Database{Name: name}
}

func (db *Database) Create(tableName string, cols ...Col) error {
	if db.Tables == nil {
		db.Tables = make(map[string]*Table)
	}

	processedCols := make(map[string]Col)
	for _, c := range cols {
		if c.Type != Int && c.Type != Int64 && c.Type != Float32 && c.Type != Float64 && c.Type != Bool && c.Type != String {
			return fmt.Errorf("column type not supported")
		}
		processedCols[c.Name] = c
	}
	db.Tables[tableName] = &Table{Name: tableName, Columns: processedCols, RowCount: 0, Db: db}
	return nil
}

type Result struct {
	TableName string
	Db        *Database
	RowCount  int
	Rows      []Row
	Error     error
}

func (table *Table) Where(col string, op string, operand interface{}) *Result {
	res := &Result{TableName: table.Name, RowCount: 0, Db: table.Db}

	table.Lock()
	defer table.Unlock()

	for i := 0; i < len(table.Rows); i++ {
		c, ok := table.Rows[i].Columns[col]
		if !ok {
			continue
		}

		c2, err := buildColumn(operand, c.Type, Col{})
		if err != nil {
			res.Error = err
			return res
		}

		if c.compare(c2, op) {
			row := Row{Id: table.Rows[i].Id, Columns: make(map[string]Col)}
			for _, cl := range table.Rows[i].Columns {
				row.Columns[cl.Name] = cl
			}
			res.Rows = append(res.Rows, row)
		}
	}

	res.Error = nil
	res.RowCount = len(res.Rows)
	return res
}

func (table *Table) Insert(cols string, values ...interface{}) error {
	return nil
}

func buildColumn(value interface{}, dataType string, newCol Col) (Col, error) {
	errStr := "data type of value doesn't match. expected: %s, got: %v"

	switch dataType {
	case Int:
		val, ok := value.(int)
		if !ok {
			return newCol, fmt.Errorf(errStr, Int, reflect.TypeOf(value))
		}
		newCol.IntVal = val
	case Int64:
		val, ok := value.(int64)
		if !ok {
			return newCol, fmt.Errorf(errStr, Int64, reflect.TypeOf(value))
		}
		newCol.Int64Val = val
	case Float32:
		val, ok := value.(float32)
		if !ok {
			return newCol, fmt.Errorf(errStr, Float32, reflect.TypeOf(value))
		}
		newCol.Float32Val = val
	case Float64:
		val, ok := value.(float64)
		if !ok {
			return newCol, fmt.Errorf(errStr, Float64, reflect.TypeOf(value))
		}
		newCol.Float64Val = val
	case Bool:
		val, ok := value.(bool)
		if !ok {
			return newCol, fmt.Errorf(errStr, Bool, reflect.TypeOf(value))
		}
		newCol.BoolVal = val
	case String:
		val, ok := value.(string)
		if !ok {
			return newCol, fmt.Errorf(errStr, String, reflect.TypeOf(value))
		}
		newCol.StringVal = val
	}
	return newCol, nil
}

func (c *Col) compare(c2 Col, op string) bool {
	switch c.Type {
	case Int:
		switch op {
		case "==":
			return c.IntVal == c2.IntVal
		case "!=":
			return c.IntVal != c2.IntVal
		case ">":
			return c.IntVal > c2.IntVal
		case "<":
			return c.IntVal < c2.IntVal
		case ">=":
			return c.IntVal >= c2.IntVal
		case "<=":
			return c.IntVal <= c2.IntVal
		}
	case Int64:
		switch op {
		case "==":
			return c.Int64Val == c2.Int64Val
		case "!=":
			return c.Int64Val != c2.Int64Val
		case ">":
			return c.Int64Val > c2.Int64Val
		case "<":
			return c.Int64Val < c2.Int64Val
		case ">=":
			return c.Int64Val >= c2.Int64Val
		case "<=":
			return c.Int64Val <= c2.Int64Val
		}
	case Float64:
		switch op {
		case "==":
			return c.Float64Val == c2.Float64Val
		case "!=":
			return c.Float64Val != c2.Float64Val
		case ">":
			return c.Float64Val > c2.Float64Val
		case "<":
			return c.Float64Val < c2.Float64Val
		case ">=":
			return c.Float64Val >= c2.Float64Val
		case "<=":
			return c.Float64Val <= c2.Float64Val
		}
	case Bool:
		switch op {
		case "==":
			return c.BoolVal == c2.BoolVal
		case "!=":
			return c.BoolVal != c2.BoolVal
		}
	case String:
		switch op {
		case "==":
			return c.StringVal == c2.StringVal
		case "!=":
			return c.StringVal != c2.StringVal
		}
	}
	return false
}

func (r *Result) Add(col string, op string, operand interface{}) *Result {
	res := Result{}

	for i := 0; i < len(r.Rows); i++ {
		c, ok := r.Rows[i].Columns[col]
		if !ok {
			continue
		}

		c2, err := buildColumn(operand, c.Type, Col{})
		if err != nil {
			r.Error = err
			return r
		}

		if c.compare(c2, op) {
			row := Row{Id: r.Rows[i].Id, Columns: make(map[string]Col)}
			for _, cl := range r.Rows[i].Columns {
				row.Columns[cl.Name] = cl
			}
			res.Rows = append(res.Rows, row)
		}
	}

	r.Error = nil
	r.Rows = res.Rows
	r.RowCount = len(res.Rows)
	return r
}

func (r *Result) Delete() {
	if len(r.Rows) == 0 {
		return
	}

	r.Db.Tables[r.TableName].Lock()
	defer r.Db.Tables[r.TableName].Unlock()

	for _, resultRow := range r.Rows {
		delete(r.Db.Tables[r.TableName].Rows, resultRow.Id)
	}

	r.Db.Tables[r.TableName].RowCount = len(r.Db.Tables[r.TableName].Rows)
}

func (r *Result) Update(col string, value interface{}) {
	if len(r.Rows) == 0 {
		return
	}

	r.Db.Tables[r.TableName].Lock()
	defer r.Db.Tables[r.TableName].Unlock()
	for _, resultRow := range r.Rows {
		oldCol := r.Db.Tables[r.TableName].Rows[resultRow.Id].Columns[col]
		newCol, err := buildColumn(value, oldCol.Type, oldCol)
		if err != nil {
			r.Error = err
			return
		}
		r.Db.Tables[r.TableName].Rows[resultRow.Id].Columns[col] = newCol
	}
}
