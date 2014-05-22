package main

import (
	"fmt"
	"time"
)

type ToDo struct {
	Registerd time.Time // 登録日時
	Content   string    // ToDoの内容
}

func newToDo(t time.Time, content string) *ToDo {
	return &ToDo{t, content}
}

type ToDoList struct {
	items *[]ToDo
}

func NewToDoList() *ToDoList {
	items := make([]ToDo, 0)
	return &ToDoList{&items}
}

func (t *ToDoList) Count() int {
	return len(*t.items)
}

func (t *ToDoList) Add(regTime time.Time, content string) {
	todo := newToDo(regTime, content)
	temp := append(*t.items, *todo)
	t.items = &temp
}

func (t *ToDoList) Get(no int) (bool, *ToDo) {
	if no <= t.Count() {
		return true, &(*t.items)[no]
	}

	return false, nil
}

func main() {
	fmt.Println("Hello world")
}
