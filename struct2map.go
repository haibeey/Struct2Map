package struct2map

import (
	"fmt"
	"reflect"
)

//Struct2Map converts a go structs to a map
func Struct2Map(model interface{}) (map[string]interface{}, error) {
	if model == nil {
		return nil, nil
	}
	ret := make(map[string]interface{})

	modelReflect := reflect.ValueOf(model)

	switch modelReflect.Kind() {
	case reflect.Map:
		iter := modelReflect.MapRange()
		for iter.Next() {
			ret[iter.Key().String()] = iter.Value().Interface()
		}
		return ret, nil
	case reflect.Ptr:
		modelReflect = modelReflect.Elem()
	case reflect.Struct, reflect.Interface:
	default:
		return nil, fmt.Errorf("Passed value must be a map or pointer or a struct ?")
	}

	if modelReflect.Kind() != reflect.Struct {
		if modelReflect.Kind() != reflect.Interface {
			return nil, fmt.Errorf("Passed value must be a map or pointer or a struct %d", modelReflect.Kind())
		}
	}
	modelRefType := modelReflect.Type()
	fieldsCount := modelReflect.NumField()

	var fieldData interface{}

	var err error
loop:
	for i := 0; i < fieldsCount; i++ {
		field := modelReflect.Field(i)
		if field.IsZero() {
			switch field.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				ret[modelRefType.Field(i).Name] = 0
			case reflect.String:
				ret[modelRefType.Field(i).Name] = ""
			default:
				continue loop
			}
		}
		switch field.Kind() {
		case reflect.Struct:
			fallthrough
		case reflect.Ptr:
			fieldData, err = Struct2Map(field.Interface())
			if err != nil {
				return ret, err
			}
		default:
			fieldData = field.Interface()
		}

		ret[modelRefType.Field(i).Name] = fieldData
	}
	return ret, nil
}
