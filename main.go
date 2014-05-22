package main

import (
	"fmt"
	"time"
)

type ToDo struct {
	Registerd time.Time // 登録日時
	Content   string    // ToDoの内容
}

func NewToDo(t time.Time, content string) *ToDo {
	return &ToDo{t, content}
}

type ToDoList []ToDo

func Sum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello world")
}
