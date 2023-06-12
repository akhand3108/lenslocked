package main

import (
	"html/template"
	"os"
)

type User struct {
	Name      string
	Age       int
	Relatives map[string]string
	Friends   []string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name:    "John Smith",
		Age:     23,
		Friends: []string{"Tony", "Natasha"},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

}
