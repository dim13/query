package query

import "testing"

func TestMarshal(t *testing.T) {
	q := struct {
		A string `query:"x"`
		B int
		C string `query:",optional"`
		D [3]byte
		E []byte
		F uint16
	}{
		A: "test",
		B: 100,
		D: [3]byte{1, 2, 3},
		E: []byte{'A', 'B'},
		F: 65535,
	}
	v, err := Marshal(q)
	if err != nil {
		t.Error(err)
	}
	if v != "?b=100&d=%01%02%03&e=AB&f=65535&x=test" {
		t.Error("wrong result", v)
	}

}
