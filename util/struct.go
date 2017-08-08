package util

import (
	"reflect"
	"fmt"
)

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		// we just skip unknown field names
		return nil
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return fmt.Errorf(
			"Provided value for %s type (%s)didn't match obj field type (%s)",
			name,
			structFieldType,
			val.Type())
	}

	structFieldValue.Set(val)
	return nil
}
