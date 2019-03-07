package main

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
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
	res, err := simplejson.NewJson([]byte(src))
	if err != nil {
		return "err"
	}
	var a Info
	a.Page = res.Get("page").MustInt()
	newFruits :=append(res.Get("fruits").MustStringArray(), "banana")
	a.Fruits = newFruits
	jsonBytes, _ := json.Marshal(a)
	return string(jsonBytes)
}
