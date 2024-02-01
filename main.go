package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func main() {
	// Get args if user types "list"
	if len(os.Args) > 1 {
		primaryArg := os.Args[1]
		if primaryArg == "list" {
			file, err := os.Open("userdata.csv")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()
			records, err := TaskList{}.getTasks()
			if err != nil {
				fmt.Println(err)
				return
			}
			// Print the file
			for _, task := range records {
				fmt.Print("ID: " + strconv.Itoa(task.id) + " | ")
				fmt.Print("Todo: " + task.text)
				fmt.Print(" | Date created: " + time.Unix(task.dateCreated, 0).String())
				fmt.Print(" | Date due: " + time.Unix(task.dateDue, 0).String())
				fmt.Print(" | Date completed: " + time.Unix(task.dateCompleted, 0).String())
				fmt.Print("\n")
			}
		} else if primaryArg == "add" {
			reader := bufio.NewReader((os.Stdin))
			fmt.Print("Enter todo: ")
			text, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}

			msg := fmt.Sprintln(text)
			task, _ := Task{
				id:            0,
				text:          msg,
				dateDue:       0,
				dateCreated:   time.Now().Unix(),
				dateCompleted: 0,
			}.addTask()

			color.Green("\nTodo added!")
			fmt.Println("Todo text > " + task.text)
			fmt.Println("Todo id > " + strconv.Itoa(task.id))
		} else if primaryArg == "edit" || primaryArg == "update" {
			// get next argument containing id
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println(err)
				return
			}
			reader := bufio.NewReader((os.Stdin))
			fmt.Print("Edit todo title: ")
			text, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
			}

			msg := fmt.Sprintln(text)

			task, err := Task{
				id:            id,
				text:          msg,
				dateDue:       0,
				dateCreated:   time.Now().Unix(),
				dateCompleted: 0,
			}.editTask()

			if err != nil {
				fmt.Println(err)
				return
			}
			color.Green("\nTodo " + strconv.Itoa(task.id) + " updated > " + task.text)
		} else if primaryArg == "del" || primaryArg == "delete" || primaryArg == "remove" {
			// get next argument containing id
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println(err)
				return
			}
			error := Task{
				id:            id,
				text:          "",
				dateDue:       0,
				dateCreated:   0,
				dateCompleted: 0,
			}.deleteTask()
			if error != nil {
				fmt.Println(error)
				return
			}
			color.Green("\nTodo " + strconv.Itoa(id) + " deleted")
		} else {
			fmt.Println("Here are the available commands:")
			fmt.Println("\n- List all todos")
			fmt.Println("\n\ttodolist list")
			fmt.Println("\n- Add a todo")
			fmt.Println("\n\ttodolist add")
			fmt.Println("\n- Edit a todo")
			fmt.Println("\n\ttodolist edit|update <id>")
			fmt.Println("\n- Delete a todo")
			fmt.Println("\n\ttodolist del|delete|remove <id>")
		}
	} else {
		fmt.Println("Run following command for help.\n\ntodolist help")
	}
}
