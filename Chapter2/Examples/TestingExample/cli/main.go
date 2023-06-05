package main

import "github.com/monegim/Book-FP/pkg"

func main() {
	t := pkg.NewTodo()
	t.Write("hello world")
}
