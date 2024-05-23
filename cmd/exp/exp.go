package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
	Age  int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml") // receive template.Template(s)
	if err != nil {
		panic(err)
	}
	user := User{
		Name: "VL",
		Bio:  `<script>alert("Haha, you have nbeen h4x0r3d!");</script>`,
		Age:  47,
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
