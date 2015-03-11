package main

import (
	"fmt"
	"time"
	"v8runner"
)

type Foo struct {
	Foo []int `json:"foo"`
	Bar Bar   `json:"bar"`
}

type Bar struct {
	Baz  bool   `json:"baz"`
	Hoge string `json:"hoge"`
}

func main() {
	scripts := []string{
		"null",
		"true",
		"123",
		"457.78",
		"[10, 20, 30]",
		"'Hello, World'",
		"new Date()",
		`obj = {"foo": [1, 2], "bar": {"baz": true, "hoge": "fuga"}}`,
	}

	go func() {
		for {
			fmt.Println("dummy script")
			time.Sleep(1 * time.Second)
		}
	}()

	for _, s := range scripts {
		var res interface{}
		v8runner.RunV8(s, &res)
		time.Sleep(1 * time.Second)
		fmt.Printf("Script -> %s\n", s)
		fmt.Printf("Result -> %T: %#+v\n\n", res, res)
	}
}
