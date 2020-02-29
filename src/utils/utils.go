package utils

import (
	"reflect"
	"strconv"
	"strings"
)

//=====================

//TODO implement for multi level
func StructToMapAsTag(data interface{}, tagName string) map[string]interface{} {
	result := make(map[string]interface{})
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)
	numField := val.NumField()
	for i := 0; i < numField; i++ {
		field := val.Field(i)
		fieldType := field.Type()
		if fieldType.Kind() == reflect.Struct {
			for j := 0; j < fieldType.NumField(); j++ {
				subField := fieldType.Field(j)
				fieldName := strings.Split(subField.Tag.Get(tagName), ",")[0]
				f := field.Field(j)
				result[fieldName] = reflectToNativeTypeMapping(f)

			}
		} else {
			fieldName := strings.Split(typ.Field(i).Tag.Get(tagName), ",")[0]
			f := val.Field(i)
			result[fieldName] = reflectToNativeTypeMapping(f)
		}
	}
	return result
}

//func ListStructToMapAsTag(i []interface{}, tagName string) []map[string]interface{} {
//	result := make([]map[string]interface{}, 0)
//	for index := range i {
//		result = append(result, StructToMapAsTag(i[index], tagName))
//	}
//	return result
//}

/*
 * source : map data, key mapping struct's tag if it existed, orElse fieldName
 * tagName: tagName for binding
 * target: struct's pointer
 */

//TODO: need to use struct, not use pointer
//func MapToStruct(source map[string]interface{}, tagName string, target interface{}) {
//	val := reflect.ValueOf(target).Elem()
//	numField := val.NumField()
//	for i := 0; i < numField; i++ {
//		field := val.Field(i)
//		if !field.CanSet() {
//			continue
//		}
//
//		if field.Kind() == reflect.Struct {
//			for j := 0; j < field.NumField(); j++ {
//				fieldName := strings.Split(field.Type().Field(j).Tag.Get(tagName), ",")[0]
//				if item, ok := source[fieldName]; ok {
//					if field.Field(j).CanSet() {
//						typeMapping(item, field.Field(j))
//					}
//				}
//			}
//			continue
//		}
//		fieldName := strings.Split(val.Type().Field(i).Tag.Get(tagName), ",")[0]
//		if item, ok := source[fieldName]; ok {
//			typeMapping(item, val.Field(i))
//		}
//	}
//}

func reflectToNativeTypeMapping(f reflect.Value) string {
	var v string
	switch f.Interface().(type) {
	case int, int8, int16, int32, int64:
		v = strconv.FormatInt(f.Int(), 10)
	case uint, uint8, uint16, uint32, uint64:
		v = strconv.FormatUint(f.Uint(), 10)
	case float32:
		v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
	case float64:
		v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
	case []byte:
		v = string(f.Bytes())
	case bool:
		v = strconv.FormatBool(f.Bool())
	case string:
		v = f.String()
	}
	return v
}
//
//func typeMapping(item interface{}, field reflect.Value) {
//	if item != nil {
//		switch field.Kind() {
//		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
//			field.SetInt(item.(int64))
//		case reflect.String:
//			field.SetString(item.(string))
//		case reflect.Float32, reflect.Float64:
//			field.SetFloat(item.(float64))
//		case reflect.Bool:
//			field.SetBool(item.(bool))
//		case reflect.Ptr:
//			if reflect.ValueOf(item).Kind() == reflect.Bool {
//				itemBool := item.(bool)
//				field.Set(reflect.ValueOf(&itemBool))
//			}
//		case reflect.Struct:
//			field.Set(reflect.ValueOf(item))
//		}
//	}
//}
