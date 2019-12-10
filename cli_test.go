package cli

import (
	"os"
)

func Example() {
	app := NewApp()
	app.Name = "todo"
	app.Usage = "task list on the command line"
	app.Commands = []Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			Action: func(c *Context) {
				println("added task: ", c.Args().First())
			},
		},
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *Context) {
				println("completed task: ", c.Args().First())
			},
		},
	}

	app.Run(os.Args)
}
