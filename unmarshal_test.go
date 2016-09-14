package query

import (
	"net/url"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	q := struct {
		A string
		B int
		C uint32
		D []string
	}{}
	v := url.Values{}
	v.Add("a", "whatever")
	v.Add("A", "Whatever")
	v.Add("b", "100")
	v.Add("c", "42")
	v.Add("d", "xxx")
	v.Add("d", "yyy")
	v.Add("d", "zzz")

	err := Unmarshal(v, &q)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", q)
}
