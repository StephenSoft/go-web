package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("Tempalte Parse faild ,err:%v", err)
		return
	}

	user1 := User{
		Name:   "小明",
		Gender: "男",
		Age:    8,
	}

	tmpl.Execute(w, user1)

}

func main() {

	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":9292", nil)
	if err != nil {
		fmt.Println("Http server start faild ,err:%v", err)
		return
	}
}
