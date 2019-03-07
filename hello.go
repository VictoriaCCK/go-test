package main

import (
	"encoding/json"
	"fmt"
)

type Info struct {
	Page   int `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	fmt.Println(addBanana(str))
}

func addBanana(src string) string{
	var a = Info{}
	err := json.Unmarshal([]byte(src), &a)
	if err != nil {
		return "err"
	}
	a.Fruits = append(a.Fruits, "banana")
	jsonBytes, err := json.Marshal(a)
	if err != nil {
		return "err"
	}
	return string(jsonBytes)
}
