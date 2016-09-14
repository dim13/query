package query

import (
	"errors"
	"log"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// Unmarshal url.Values to struct
func Unmarshal(q url.Values, v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return errors.New("non-pointer passed to Unmarshal")
	}
	// convert url.Values keys to lowercase
	ql := url.Values{}
	for k, v := range q {
		kl := strings.ToLower(k)
		ql[kl] = append(ql[kl], v...)
	}
	log.Println(">>>", ql)
	return unmarshal(ql, val.Elem())
}

func unmarshal(q url.Values, v reflect.Value) error {
	if v.Kind() != reflect.Struct {
		log.Println(v.Kind())
		return errors.New("must be a struct")
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		name := strings.ToLower(t.Field(i).Name)
		arg, ok := q[name]
		if !ok {
			continue
		}
		log.Println("arg", arg)
		f := v.Field(i)
		switch v.Field(i).Kind() {
		case reflect.String:
			log.Println("string", name)
			f.SetString(arg[0])
		case reflect.Int, reflect.Int32, reflect.Int64:
			log.Println("int", name)
			i, err := strconv.ParseInt(arg[0], 10, 64)
			if err != nil {
				return err
			}
			f.SetInt(i)
		case reflect.Uint, reflect.Uint32, reflect.Uint64:
			log.Println("uint", name)
			u, err := strconv.ParseUint(arg[0], 10, 64)
			if err != nil {
				return err
			}
			f.SetUint(u)
		case reflect.Slice:
			log.Println("slice", name)
		default:
			log.Println("???", name, v.Field(i).Kind())
			return errors.New("unsupported type")
		}
	}
	return nil
}
