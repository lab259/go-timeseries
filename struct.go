package timesrs

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func GetFieldValue(field string, data interface{}) (interface{}, error) {
	return nil, errors.New("not implemented")
	fields := strings.Split(field, ".")
	v := reflect.ValueOf(data)
	for _, key := range fields {
		if v.Kind() == reflect.Ptr {
			// v = reflect.map
		} else if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		// we only accept structs
		if v.Kind() == reflect.Map {
			continue
		} else if v.Kind() != reflect.Struct {
			return nil, fmt.Errorf("only accepts structs; got %T", v)
		}

		v = v.FieldByName(key)
	}
	return v.Interface(), nil
}
