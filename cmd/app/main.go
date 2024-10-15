package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/antremis/go-cli-todos/internal"
)

type CommandLineFlags struct {
	File       *string
	Create     *string
	Update     *string
	Delete     *string
	List       *bool
	Get        *string
	Complete   *string
	Uncomplete *string
	Help       *string
}

func main() {
	cmf := CommandLineFlags{}
	cmf.File = flag.String("file", "todos.json", "Path to the JSON file")
	cmf.Create = flag.String("create", "", "Create a new todo")
	cmf.Update = flag.String("update", "", "Update a todo")
	cmf.Delete = flag.String("delete", "", "Delete a todo")
	cmf.Get = flag.String("get", "", "Get a todo")
	cmf.List = flag.Bool("list", false, "List all todos")
	cmf.Complete = flag.String("complete", "", "Complete a todo")
	cmf.Uncomplete = flag.String("uncomplete", "", "Uncomplete a todo")

	flag.Parse()

	todos := internal.NewFileStorage(*cmf.File)
	switch {
	case strings.Compare(*cmf.Create, "") != 0:
		if Ok := todos.Create(internal.TodoCreate{Title: *cmf.Create}); Ok != nil {
			fmt.Println("Error occured while creating todo")
			return
		}
	case strings.Compare(*cmf.Update, "") != 0:
		args := strings.Split(*cmf.Update, ":")
		if Ok := todos.Update(args[0], internal.TodoUpdate{Title: args[1]}); Ok != nil {
			fmt.Println("Error occured while updating todo")
			return
		}
	case strings.Compare(*cmf.Delete, "") != 0:
		if Ok := todos.Delete(*cmf.Delete); Ok != nil {
			fmt.Println("Error occured while deleting todo")
			return
		}
	case strings.Compare(*cmf.Get, "") != 0:
		todos.Get(*cmf.Get)
	case *cmf.List:
		todos.List()
	case strings.Compare(*cmf.Complete, "") != 0:
		if Ok := todos.Complete(*cmf.Complete); Ok != nil {
			fmt.Println("Error occured while completing todo")
			return
		}
	case strings.Compare(*cmf.Uncomplete, "") != 0:
		if Ok := todos.Uncomplete(*cmf.Uncomplete); Ok != nil {
			fmt.Println("Error occured while uncompleting todo")
			return
		}
	default:
		todos.List()
	}
}
