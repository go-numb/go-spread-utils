package spreads

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func SetError(err error, msg any) error {
	return fmt.Errorf("%v, %s", msg, err.Error())
}

func ConvertInterfaceToString(sheetValues [][]interface{}) [][]string {
	rows := make([][]string, len(sheetValues))
	for i := 0; i < len(sheetValues); i++ {
		var row []string
		for j := 0; j < len(sheetValues[i]); j++ {
			row = append(row, strings.TrimSpace(fmt.Sprintf("%v", sheetValues[i][j])))
		}
		rows[i] = row
	}

	return rows
}

func ConvertStringToInterface(records [][]string) [][]interface{} {
	in := make([][]interface{}, len(records))
	for i := 0; i < len(records); i++ {
		in[i] = make([]interface{}, len(records[i]))
		for j, cell := range records[i] {
			// Important!! []string -> []interface{}
			in[i][j] = fmt.Sprintf("%v", trim(cell))
		}
	}
	return in
}

func Bind(records [][]string, binder any) error {
	// Reflect on the binder to verify it's a pointer to a slice of structs.
	binderVal := reflect.ValueOf(binder)
	if binderVal.Kind() != reflect.Ptr || binderVal.Elem().Kind() != reflect.Slice {
		return errors.New("binder must be a pointer to a slice of structs")
	}

	// Dig further to verify element type is struct.
	structType := binderVal.Elem().Type().Elem()
	if structType.Kind() != reflect.Struct {
		return errors.New("binder must be a pointer to a slice of structs")
	}

	// Assuming first row is headers.
	headers := records[0]

	// Create a map from headers (column names) to struct field indexes.
	fieldMap := make(map[string]int)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		tag := field.Tag.Get("csv")
		for _, h := range headers {
			if h == tag {
				fieldMap[h] = i
				break
			}
		}
	}

	// Iterate over data rows, starting from 1 as we assume 0 is header.
	for _, row := range records[1:] {
		newStructPtr := reflect.New(structType).Elem()
		for col, value := range row {
			fieldIndex, exists := fieldMap[headers[col]]
			if !exists {
				continue // Skip if no matching struct field
			}
			fieldVal := newStructPtr.Field(fieldIndex)

			// Handle basic data types - expand as necessary.
			switch fieldVal.Kind() {
			case reflect.String:
				fieldVal.SetString(value)
			case reflect.Int, reflect.Int32, reflect.Int64:
				if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
					fieldVal.SetInt(intValue)
				}
				// Add more type handlers as necessary.
			}
		}
		binderVal.Elem().Set(reflect.Append(binderVal.Elem(), newStructPtr))
	}
	return nil
}

func rangekey(sheetName, rangeKey string) string {
	return fmt.Sprintf("%s!%s", sheetName, rangeKey)
}

func trim(s string) string {
	// Trim "["&"]" from string
	s = strings.TrimLeft(s, "[")
	s = strings.TrimRight(s, "]")
	return s
}
