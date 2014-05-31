package main

import (
	"errors"
	"fmt"
	"time"
)

const (
	Doing    = iota
	Complete = iota
)

type ToDo struct {
	Registerd time.Time // 登録日時
	Content   string    // ToDoの内容
	Status    int       // 状態（実行中, 完了）
}

func newToDo(t time.Time, content string) *ToDo {
	return &ToDo{t, content, Doing}
}

type ToDoList struct {
	items []ToDo
}

func NewToDoList() *ToDoList {
	items := make([]ToDo, 0)
	return &ToDoList{items}
}

func (t *ToDoList) Add(regTime time.Time, content string) {
	todo := *(newToDo(regTime, content))
	t.items = append(t.items, todo)
}

func (t *ToDoList) GetByStatus(status int) []ToDo {
	items := make([]ToDo, 0)
	for _, item := range t.items {
		if item.Status == status {
			items = append(items, item)
		}
	}
	return items
}

func (t *ToDoList) Complete(no int) error {
	i := 0
	for j, todo := range t.items {
		if todo.Status != Doing {
			continue
		}
		if no != i {
			continue
		}

		t.items[j].Status = Complete
		return nil
	}

	return errors.New("noが範囲外")
}

func main() {
	fmt.Println("Hello world")
}
