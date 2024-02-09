package main

import (
	"fmt"
	"github.com/persi-man/cli-task/models"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

var (
	tasks = []models.Task{}
)

func main() {

	err := models.LoadTasks()
	if err != nil {
		log.Fatalf("Failed to load tasks: %v", err)
	}

	tasks = append(tasks, models.Task{
		ID:        1,
		Title:     "Manger",
		Status:    models.StatusDone,
		CreatedAt: time.Now(),
	})
	tasks = append(tasks, models.Task{
		ID:        2,
		Title:     "Bouger",
		Status:    models.StatusPending,
		CreatedAt: time.Now(),
	})
	tasks = append(tasks, models.Task{
		ID:        3,
		Title:     "Dormir",
		Status:    models.StatusDone,
		CreatedAt: time.Now(),
	})
	//Print all tasks
	app := &cli.App{
		Name:  "TodoList",
		Usage: "TodoList is a CLI for managing your TODOs.",
		Commands: []*cli.Command{
			// Print all tasks
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "Print all tasks",
				Action: func(c *cli.Context) error {
					for i := 0; i < len(tasks); i++ {
						fmt.Println(tasks[i].ID, " : ", tasks[i].Title, " - ", tasks[i].Status, " - ", tasks[i].CreatedAt)
					}
					return nil
				},
			},
			// Print specific task by their id
			{
				Name:    "print",
				Aliases: []string{"p"},
				Usage:   "Print specific task by id",
				Action: func(c *cli.Context) error {
					id := c.Args().First()
					for i := 0; i < len(tasks); i++ {
						if fmt.Sprint(tasks[i].ID) == id {
							fmt.Println(tasks[i].ID, " : ", tasks[i].Title, " - ", tasks[i].Status, " - ", tasks[i].CreatedAt)
						}
					}
					return nil
				},
			},
			// Add new task to the list
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "create a new task",
				Action: func(c *cli.Context) error {
					title := c.Args().First()
					tasks = append(tasks, models.Task{
						ID:        len(tasks) + 1,
						Title:     title,
						Status:    models.StatusPending,
						CreatedAt: time.Now(),
					})
					fmt.Println("The task n°", tasks[len(tasks)-1].ID, " : ", tasks[len(tasks)-1].Title, ", at status ", tasks[len(tasks)-1].Status, " is create at ", tasks[len(tasks)-1].CreatedAt)

					return nil
				},
			},
			//Update an existing task
			{
				Name:    "update",
				Aliases: []string{"u"},
				Usage:   "update an existing task",
				Subcommands: []*cli.Command{
					{ //Update the title of a task
						Name:    "Title",
						Aliases: []string{"t"},
						Usage:   "update the title of a task",
						Action: func(c *cli.Context) error {
							id := c.Args().First()
							title := c.Args().Get(1)
							for i := 0; i < len(tasks); i++ {
								if fmt.Sprint(tasks[i].ID) == id {
									tasks[i].Title = title
									fmt.Println("The task n°", tasks[i].ID, " : ", tasks[i].Title, " is updated")
								}
							}
							return nil
						},
					},
					{ //Update the status of a task
						Name:    "Status",
						Aliases: []string{"s"},
						Usage:   "update the status of a task",
						Action: func(c *cli.Context) error {
							id := c.Args().First()
							status := c.Args().Get(1)
							for i := 0; i < len(tasks); i++ {
								if fmt.Sprint(tasks[i].ID) == id {
									tasks[i].Status = models.Status(status)
									fmt.Println("The task n°", tasks[i].ID, " : ", tasks[i].Title, " is now ", tasks[i].Status)
								}
							}
							return nil
						},
					},
				},
			},
			{ //Delete a task
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "delete a task",
				Subcommands: []*cli.Command{
					{ //Delete a task by their id
						Name:  "id",
						Usage: "delete a task by their id",
						Action: func(c *cli.Context) error {
							id := c.Args().First()
							for i := 0; i < len(tasks); i++ {
								if fmt.Sprint(tasks[i].ID) == id {
									tasks = append(tasks[:i], tasks[i+1:]...)
									fmt.Println("The task n°", id, " is deleted. The list of tasks is now :", len(tasks))
								}
							}
							return nil
						},
					},
					{ //Delete task by title
						Name:    "title",
						Aliases: []string{"t"},
						Usage:   "delete task by title",
						Action: func(c *cli.Context) error {
							title := c.Args().First()
							for i := 0; i < len(tasks); i++ {
								if tasks[i].Title == title {
									tasks = append(tasks[:i], tasks[i+1:]...)
									fmt.Println("The task n°", tasks[i].ID, " : ", title, " is deleted. The list of tasks is now :", len(tasks))
								}
							}
							return nil
						},
					},
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
