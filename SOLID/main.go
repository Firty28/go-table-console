package main




func main() {

	todo := Tasks{}

	todo.addTask(Task{"Название 5", "Содержание 2", 3})
	todo.addTask(Task{"Название 4", "Содержание 234", 3})
	todo.addTask(Task{"Название 7", "Содержание 33", 3})
	todo.removeTask(1)
	todo.addTask(Task{"Название 73", "Содержание 33", 3})
	todo.all()
}
