package main

import (
	"."
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTodoListCanAddTodo(t *testing.T) {
	lst := main.NewToDoList()

	doingCount := func() int {
		return len(lst.GetByStatus(main.Doing))
	}

	assert.Equal(t, doingCount(), 0, "ToDoリスト生成時は0件")
	lst.Add(time.Now(), "test")
	assert.Equal(t, doingCount(), 1, "ToDoを追加したので1件")
}

func TestTodoCanComplete(t *testing.T) {
	lst := main.NewToDoList()
	lst.Add(time.Now(), "test")

	err := lst.Complete(0)
	assert.Equal(t, err, nil, "ERROR")
	assert.Equal(t, len(lst.GetByStatus(main.Doing)), 0, "ToDo完了時はステータスはDoingではない")
	assert.Equal(t, len(lst.GetByStatus(main.Complete)), 1, "ToDo完了時はステータスはComplete")
}

func TestCleanShuldPurgeCompleteTodos(t *testing.T) {
	lst := main.NewToDoList()
	lst.Add(time.Now(), "todo1")
	lst.Add(time.Now(), "todo2")
	lst.Complete(0)
	lst.Clean()
	assert.Equal(t, len(lst.GetByStatus(main.Complete)), 0, "Clean後は完了ToDoは削除されている")
	assert.Equal(t, len(lst.GetByStatus(main.Doing)), 1, "未完了ToDoは削除されない")
}
