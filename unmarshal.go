package query

import (
	"errors"
	"log"
	"net/url"
	"reflect"
	"strings"
)

func Unmarshal(q url.Values, v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return errors.New("non-pointer passed to Unmarshal")
	}
	return unmarshal(q, val.Elem())
}

func unmarshal(q url.Values, v reflect.Value) error {
	if v.Kind() != reflect.Struct {
		log.Println(v.Kind())
		return errors.New("must be a struct")
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		name := strings.ToLower(t.Field(i).Name)
		switch v.Field(i).Kind() {
		case reflect.String:
			log.Println("string", name)
		case reflect.Int, reflect.Int32, reflect.Int64:
			log.Println("int", name)
		case reflect.Uint, reflect.Uint32, reflect.Uint64:
			log.Println("uint", name)
		default:
			log.Println("???", name, v.Field(i).Kind())
		}
	}
	return nil
}
