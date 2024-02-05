package main

import (
	"encoding/json"
	"fmt"
)

type Optional[T any] struct {
	value  T
	wasSet bool
}

func (o *Optional[T]) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &o.value); err != nil {
		return err
	}

	o.wasSet = true
	return nil
}

func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if !o.wasSet {
		return []byte("null"), nil
	}
	fmt.Print("hii")
	return json.Marshal(o.value)
}

func (o *Optional[T]) Value() (T, bool) {
	return o.value, o.wasSet
}

type Person struct {
	Name     string           `json:"name"`
	NickName Optional[string] `json:"nick_name,omitempty"`
}

func main() {
	var p1 Person
	_ = json.Unmarshal([]byte(`{"name":"Parikshit"}`), &p1)
	fmt.Println(p1)

	val, ok := p1.NickName.Value()
	if ok {
		fmt.Println(val)
	}

	bytes, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))

	var p2 Person
	_ = json.Unmarshal([]byte(`{"name":"Parikshit","nick_name":"PSKP"}`), &p2)
	fmt.Println(p2)

	val, ok = p2.NickName.Value()
	if ok {
		fmt.Println(val)
	}

	bytes, _ = json.Marshal(p2)
	fmt.Println(string(bytes))
}
