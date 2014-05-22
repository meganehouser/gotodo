package main

import (
	"."
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTodoListCanAddTodo(t *testing.T) {
	lst := main.NewToDoList()

	assert.Equal(t, lst.Count(), 0, "ToDoリスト生成時は0件")

	lst.Add(time.Now(), "test")

	assert.Equal(t, lst.Count(), 1, "ToDoを追加したので1件")

	exists, todo := lst.Get(0)
	if !exists {
		t.Error("追加したToDoは取得できる")
		return
	}

	assert.Equal(t, todo.Content, "test", "追加したToDo内容を参照できる")
}
