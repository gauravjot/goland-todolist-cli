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
			records, err := readCSVRecords(file)
			if err != nil {
				fmt.Println(err)
				return
			}
			// Print the file
			for _, line := range records {
				for _, value := range line {
					fmt.Print(value + ", ")
				}
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
			task, _ := writeTaskToFile(
				true,
				Task{
					id:            0,
					text:          msg,
					dateDue:       0,
					dateCreated:   time.Now().Unix(),
					dateCompleted: 0,
				})

			color.Green("\nTodo added!")
			fmt.Println("Todo text > " + task.text)
			fmt.Println("Todo id > " + strconv.Itoa(task.id))
		} else if primaryArg == "edit" {
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

			task, err := writeTaskToFile(
				false,
				Task{
					id:            id,
					text:          msg,
					dateDue:       0,
					dateCreated:   time.Now().Unix(),
					dateCompleted: 0,
				})

			if err != nil {
				fmt.Println(err)
				return
			}
			color.Green("\nTodo " + strconv.Itoa(task.id) + " updated > " + task.text)
		} else if primaryArg == "del" {
			// get next argument containing id
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println(err)
				return
			}
			err = deleteTaskFromFile(id)
			if err != nil {
				fmt.Println(err)
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
			fmt.Println("\n\ttodolist edit <id>")
			fmt.Println("\n- Delete a todo")
			fmt.Println("\n\ttodolist del <id>")
		}
	} else {
		fmt.Println("Run following command for help.\n\ntodolist help")
	}
}
