package main

import (
	"encoding/json"
	"fmt"
)

var _ json.Unmarshaler = (*Option[int])(nil)

type Option[T any] struct {
	value *T
}

func (o *Option[T]) UnmarshalJSON(bytes []byte) error {
	o.value = new(T)
	return json.Unmarshal(bytes, o.value)
}

func (o *Option[T]) Value() (out T, ok bool) {
	if o.value == nil {
		return
	}
	return *o.value, true
}

type Person struct {
	Name     string         `json:"name"`
	NickName Option[string] `json:"nick_name"`
}

func main() {
	var p1 Person
	_ = json.Unmarshal([]byte(`{"name":"Parikshit"}`), &p1)
	fmt.Println(p1)

	val, ok := p1.NickName.Value()
	if ok {
		fmt.Println(val)
	}

	var p2 Person
	_ = json.Unmarshal([]byte(`{"name":"Parikshit","nick_name":"PSKP"}`), &p2)
	fmt.Println(p2)

	val, ok = p2.NickName.Value()
	if ok {
		fmt.Println(val)
	}
}
