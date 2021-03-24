package struct2map

import (
	"testing"
	"fmt"
)

type emptyStruct struct{}

type testStruct struct {
	Number  int
	Str     string
	Map     map[interface{}]interface{}
	Array   []interface{}
	Struct  emptyStruct
	Pointer *emptyStruct
}

func TestStruct2Map(t *testing.T) {

	Number := 0
	Str := ""
	Struct := testStruct{
		Number:  Number,
		Str:     Str,
		Map:     make(map[interface{}]interface{}),
		Array:   []interface{}{},
		Struct:  emptyStruct{},
		Pointer: new(emptyStruct),
	}

	_, err := Struct2Map(Number)
	if err == nil {
		t.Errorf(err.Error())
	}
	_, err = Struct2Map(Str)
	if err == nil {
		t.Errorf(err.Error())
	}
	v, err := Struct2Map(&Struct)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(v)


}
