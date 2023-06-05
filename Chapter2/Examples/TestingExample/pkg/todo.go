package pkg

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type Todo struct {
	Text string
	Db *Db
}
func NewTodo() Todo {
	return Todo{
		Text: "",
		Db: NewDB(),
	}
}