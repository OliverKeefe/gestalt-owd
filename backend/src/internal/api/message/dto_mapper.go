package message

import (
	"log"
	"reflect"
)

func MapToResponse[T any](src any) (T, error) {
	var response T

	dstValue := reflect.ValueOf(&response).Elem()
	srcValue := reflect.Indirect(reflect.ValueOf(src))
	dstType := dstValue.Type()

	for i := 0; i < dstType.NumField(); i++ {
		df := dstType.Field(i)
		dv := dstValue.Field(i)

		if !dv.CanSet() {
			log.Printf("could not map value: %v, is unexported or read-only?", dv)
			continue
		}

		name := df.Tag.Get("map")
		if name == "" {
			name = df.Name
		}

		sf := srcValue.FieldByName(name)
		if !sf.IsValid() || sf.Type() != dv.Type() {
			log.Printf("src field invalid %v", sf)
			continue
		}

		dv.Set(sf)
	}

	return response, nil
}
