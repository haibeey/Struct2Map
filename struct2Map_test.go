package struct2map

import (
	"fmt"
	"testing"
)

type emptyStruct struct{}

type testStruct struct {
	Number  int
	number  int
	Str     string
	str     string
	Map     map[interface{}]interface{}
	mmap     map[interface{}]interface{}
	Array   []interface{}
	array   []interface{}
	Struct  emptyStruct
	sstruct  emptyStruct
	Pointer *emptyStruct
	pointer *emptyStruct
}

func TestStruct2Map(t *testing.T) {

	Number := 10
	Str := "test"
	Struct := testStruct{
		Number:  Number,
		number:  Number,
		Str:     Str,
		str:     Str,
		Map:     make(map[interface{}]interface{}),
		mmap:     make(map[interface{}]interface{}),
		Array:   []interface{}{},
		array:   []interface{}{},
		Struct:  emptyStruct{},
		sstruct:  emptyStruct{},
		Pointer: new(emptyStruct),
		pointer: new(emptyStruct),
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

	var a interface{}
	a = testStruct{}
	v, err = Struct2Map(a)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(v)

}
