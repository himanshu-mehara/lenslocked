package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio string
}


func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}


	user := User{
		Name: "himansu-mehara",
		Bio : `<script>alert("haha, your account is hacked")</script>`,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
