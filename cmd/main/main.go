package main

import (
	"fmt"
	"learn-mongo/api/ById"
	"learn-mongo/api/GetAll"
	"learn-mongo/api/Post"
	"net/http"
)

func main() {
	http.HandleFunc("/todos", GetAll.GetAll)
	http.HandleFunc("/todo", ById.ById)
	http.HandleFunc("/post/todo", Post.PostTodo)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server Run At http://localhost:3000")
	}

}
