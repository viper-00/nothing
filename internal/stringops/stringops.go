package stringops

import "strings"

// StringArrToJSONArr returns json arr of given string arr
func StringArrToJSONArr(data []string) string {
	var builder strings.Builder
	builder.WriteString("[")

	for index, value := range data {
		builder.WriteString(value)

		if index < (len(data) - 1) {
			builder.WriteString(",")
		}
	}

	builder.WriteString("]")
	return builder.String()
}
