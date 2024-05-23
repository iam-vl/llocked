package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml") // receive template.Template(s)
	if err != nil {
		panic(err)
	}
	user := User{
		Name: "VL",
		Age:  123,
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
