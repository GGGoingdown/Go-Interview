# Embed

支援go在1.16版本以後


```go
package main

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed asset/*
var f embed.FS

type User struct {
	Name string `json:"name"`
}

func main() {
	file, err := f.ReadFile("asset/foo.json")
	if err != nil {
		panic(err.Error())
	}
	println(string(file)) 
    /*
        {
            "name": "eddie"
        }
    */
	var user User
	if er := json.Unmarshal(file, &user); er != nil {
		panic("unmarshal user failure")
	}
}

```