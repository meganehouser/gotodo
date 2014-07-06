package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
	"path"
	"strconv"
	"time"
)

const (
	Doing    = iota
	Complete = iota
)

var StoreFile string = path.Join(getUserHome(), "tasks.json")

func getUserHome() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

type ToDo struct {
	Registerd time.Time // 登録日時
	Content   string    // ToDoの内容
	Status    int       // 状態（実行中, 完了）
}

// ユーザーディレクトリのパス取得
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

func LoadToDoList(path string) *ToDoList {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return NewToDoList()
	}

	todoitems := make([]ToDo, 0)
	err = json.Unmarshal(data, &todoitems)
	if err != nil {
		panic(err)
	}

	tdlist := NewToDoList()
	tdlist.items = todoitems
	return tdlist
}

func SaveToDoList(data *ToDoList, path string) {
	jsn, err := json.Marshal((*data).items)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = ioutil.WriteFile(path, jsn, 0600)
	if err != nil {
		log.Fatal(err)
		return
	}
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
			i = i + 1
			continue
		}
		if no != i {
			i = i + 1
			continue
		}

		t.items[j].Status = Complete
		return nil
	}

	return errors.New("noが範囲外")
}

func (t *ToDoList) Clean() {
	t.items = t.GetByStatus(Doing)
}

func main() {
	ls := flag.Bool("ls", false, "list up tasks")
	clean := flag.Bool("clean", false, "clean up completed tasks")
	completeNo := flag.String("finish", "", "complete the task")
	add := flag.String("add", "", "add task")

	flag.Parse()

	data := *LoadToDoList(StoreFile)
	defer SaveToDoList(&data, StoreFile)

	if *ls {
		items := data.GetByStatus(Doing) 
		for i, item := range items {
			fmt.Printf("[%d] [%s] %s\r\n", i, item.Registerd.Format("2006/01/02"), item.Content)
		}
	} else if *clean {
		data.Clean()
	} else if *completeNo != "" {
		no, err := strconv.Atoi(*completeNo)
		if err != nil {
			panic(err)
		}

		data.Complete(no)
	} else if *add != "" {
		data.Add(time.Now(), *add)
	}

}
