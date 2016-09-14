package query

import (
	"errors"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func Marshal(v interface{}) (string, error) {
	val := reflect.ValueOf(v)
	return marshalQuery(val)
}

func parseTag(tag string) (string, string) {
	if i := strings.Index(tag, ","); i != -1 {
		return tag[:i], tag[i+1:]
	}
	return tag, ""
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String, reflect.Slice:
		return v.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Bool:
		return v.Bool() == false
	}
	return false
}

func marshalQuery(v reflect.Value) (string, error) {
	if v.Kind() != reflect.Struct {
		return "", errors.New("must be a struct")
	}
	q := url.Values{}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		name := strings.ToLower(t.Field(i).Name)
		tag, param := parseTag(t.Field(i).Tag.Get("query"))
		if tag != "" {
			name = tag
		}
		f := v.Field(i)
		if param == "optional" && isZero(f) {
			continue
		}
		switch f.Kind() {
		case reflect.Bool:
			if f.Bool() == true {
				q.Add(name, "1")
			} else {
				q.Add(name, "0")
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			q.Add(name, strconv.Itoa(int(f.Int())))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			q.Add(name, strconv.Itoa(int(f.Uint())))
		case reflect.String:
			q.Add(name, f.String())
		case reflect.Slice, reflect.Array:
			if f.Type().Elem().Kind() == reflect.Uint8 {
				tmp := make([]byte, f.Len())
				for i := 0; i < f.Len(); i++ {
					tmp[i] = byte(f.Index(i).Uint())
				}
				q.Add(name, string(tmp))
			}
		}
	}
	return "?" + q.Encode(), nil
}
