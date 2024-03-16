package main

import "fmt"

type Tasks struct {
	todoList []Task
}

type Task struct {
	title    string
	text     string
	priority int
}

func (t *Tasks) addTask(arg Task) {
	t.todoList = append(t.todoList, arg)
}

func (t *Tasks) all() {
	for _, v := range t.todoList {
		fmt.Println(v)
	}
}

func (t *Tasks) removeTask(index int) {
	t.todoList = append(t.todoList[:index], t.todoList[index+1:]...)
}

func (t *Task) editTitle(arg string) {
	t.title = arg
}

func (t *Task) editText(arg string) {
	t.text = arg
}
