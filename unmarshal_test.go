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
	}{}
	v := url.Values{}
	v.Add("a", "whatever")
	v.Add("b", "100")
	v.Add("c", "42")

	err := Unmarshal(v, &q)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", q)
}
