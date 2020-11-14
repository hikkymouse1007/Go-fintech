package main

import (
	"encoding/json"
	"fmt"
)

type T struct{}

type Person struct {
	//Name      string   `json:"-"`
	Name      string   `json:"name"`
	Age       int      `json:"age,omitempty"`
	Nicknames []string `json:"nicknames,omitempty"`
	T         *T       `json:"T,omitempty"`
}

func (p *Person) UnmarshalJSON(b []byte) error {
	type Person2 struct{
		Name string
	}
	var p2 Person2
	err := json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
	}
	p.Name = p2.Name + "!!"
	return err
}

//func (p Person) MarshalJSON() ([]byte, error) {
//	//a := &struct{Name string}{Name: "test"} Personを使わず関数の中のみで直接宣言も可能
//	v, err := json.Marshal(&struct {
//		Name string
//	}{
//		Name: "Mr." + p.Name,
//	})
//	return v, err
//}

func main() {
	// jsonのキーは大文字でも実行される
	b := []byte(`{"name":"mike","age":20,"nicknames":[]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	v, _ := json.Marshal(p)
	fmt.Println(string(v))
}
