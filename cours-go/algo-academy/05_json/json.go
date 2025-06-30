package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Go offers built-in support for Json encoding and decoding, including to and from built-in and custom data types

// ====== Encoding ======
// To encode JSON data we use the Marshal function
// func Marshal(v interface{}) ([]byte, error)

// only the exported (public) fields of a struct will be present in the JSON output. other fields are ignored
// a feild with json: struct tag is stored with its tag name instead of this variable name
// pointer will be encoded as the values the point to, or null if the pointer ie nil
// channel, complex, and function types cannot be encoded

//  ====== Decoding =====
// To dencode JSON datawe use the Unmarshal function
// func Unmarshal(data []byte, v interface{}) error

// ==== arbitrary objects and arrays =====
// The encoding/json package uses

// map[string]interface{} to store arbitrary JSON objects, and []interfac{} to store arbitrary JSON arrays.

type FruitBasket struct {
	Name    string
	Fruit   []string
	ID      int64  `json:"Ref"` // feilds appears in JSON as key "Ref"
	private string // An unexported field is not ecoded
	Created time.Time
}

func main() {
	data := FruitBasket{
		Name:    "Fruit1",
		Fruit:   []string{"ed", "eded", "edede"},
		ID:      8465454454,
		private: "code",
		Created: time.Now(),
	}
	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	// func MarshalIndenet(v interface{}, prefix, indent string) ([]byte, error)
	// b1, err := json.MarshalIndent(data, "-", "         ")
	// fmt.Println(b)
	fmt.Println(string(b))
	// fmt.Println(string(b1))
	var db1	FruitBasket

	err1 := json.Unmarshal(b, &db1)
	if err1 != nil {
		log.Println(err1)
	}
	fmt.Println(db1)
	fmt.Println(db1.Name, db1.Fruit, db1.ID)
	child := []byte(`{"Name":"Eve", "Age":6, "Parents": ["Alice", "Bob"]}`)
	var x any
	json.Unmarshal(child, &x)
	fmt.Println(x)

	// type data assertions

	// var i any = "jhbjbjkb"
	// s := i.(string)
	// fmt.Println(s)
	// db2 := x.(map[string]interface{})
	
	var z interface{} = 7
	fmt.Println(z.(int))

	v, ok := z.(int)
	fmt.Println(v, ok)

	for k, v := range data {
		fmt.Println(k, v)
		switch v := v.(type) {
		case condition:
			
		}
	}

}
