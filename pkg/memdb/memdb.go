package memdb

import (
	"fmt"
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
