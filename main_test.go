package main

import (
	"."
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTodoListCanAddTodo(t *testing.T) {
	lst := make(main.ToDoList, 0)

	assert.Equal(t, len(lst), 0, "ToDoリスト生成時は0件")

	td := main.NewToDo(time.Now(), "test")
	lst = append(lst, *td)

	assert.Equal(t, len(lst), 1, "ToDoを追加したので1件")
	assert.Equal(t, lst[0].Content, "test", "追加したToDo内容を参照できる")
}
