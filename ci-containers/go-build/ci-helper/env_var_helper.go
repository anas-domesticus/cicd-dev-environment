package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func flattenStruct(v interface{}, prefix string, envVars map[string]string) {
	value := reflect.ValueOf(v)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	typ := value.Type()

	switch value.Kind() {
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			fieldType := typ.Field(i)
			tag := fieldType.Tag.Get("json")
			if tag == "" || tag == "-" {
				tag = fieldType.Name
			} else if idx := strings.Index(tag, ","); idx != -1 {
				tag = tag[:idx]
			}

			key := strings.ToUpper(fmt.Sprintf("%s_%s", prefix, tag))

			switch field.Kind() {
			case reflect.Struct:
				// Check if the field is of type time.Time
				if field.Type() == reflect.TypeOf(time.Time{}) {
					envVars[key] = field.Interface().(time.Time).Format(time.RFC3339)
				} else {
					flattenStruct(field.Interface(), key, envVars)
				}
			case reflect.Slice:
				flattenStruct(field.Interface(), key, envVars)
			default:
				envVars[key] = fmt.Sprintf("%v", field.Interface())
			}
		}
	default:
		return
	}

}

func ToEnvVars(payload GiteaPayload) map[string]string {
	envVars := make(map[string]string)
	flattenStruct(payload, "GIT", envVars)
	return envVars
}
