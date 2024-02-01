package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	id            int
	text          string
	dateDue       int64
	dateCreated   int64
	dateCompleted int64
}

type TaskList []Task

var filename string = "userdata.csv"

/**
 * Task methods
 */

func (task Task) addTask() (Task, error) {
	tasks, _ := TaskList{}.getTasks()
	var new_id int
	if len(tasks) > 0 {
		id := tasks[len(tasks)-1].id
		new_id = id + 1
	} else {
		new_id = 0
	}
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return Task{}, err
	}
	defer f.Close()
	// Append to file
	csvWriter := csv.NewWriter(f)
	csvWriter.Write([]string{
		strconv.Itoa(new_id),
		strings.TrimSpace(task.text),
		strconv.FormatInt(task.dateDue, 10),
		strconv.FormatInt(task.dateCreated, 10),
		strconv.FormatInt(task.dateCompleted, 10),
	})
	csvWriter.Flush()

	return Task{
		id:            new_id,
		text:          strings.TrimSpace(task.text),
		dateDue:       task.dateDue,
		dateCreated:   task.dateCreated,
		dateCompleted: task.dateCompleted,
	}, nil
}

func (alteredTask Task) editTask() (Task, error) {
	allTasks, _ := TaskList{}.getTasks()
	os.Remove(filename)
	file, _ := os.Create(filename)
	csvWriter := csv.NewWriter(file)
	_writeHeadingToFile(file)
	for _, task := range allTasks {
		newText := task.text
		if task.id == alteredTask.id {
			newText = strings.TrimSpace(alteredTask.text)
		}
		csvWriter.Write([]string{
			strconv.Itoa(task.id),
			newText,
			strconv.FormatInt(task.dateDue, 10),
			strconv.FormatInt(task.dateCreated, 10),
			strconv.FormatInt(task.dateCompleted, 10),
		})
		csvWriter.Flush()
	}

	return alteredTask, nil
}

func (taskToDelete Task) deleteTask() error {
	allTasks, err := TaskList{}.getTasks()
	if err != nil {
		fmt.Println(err)
		return err
	}
	// Write to file
	os.Remove(filename)
	file, _ := os.Create(filename)
	csvWriter := csv.NewWriter(file)
	_writeHeadingToFile(file)
	for _, task := range allTasks {
		if task.id != taskToDelete.id {
			csvWriter.Write([]string{
				strconv.Itoa(task.id),
				task.text,
				strconv.FormatInt(task.dateDue, 10),
				strconv.FormatInt(task.dateCreated, 10),
				strconv.FormatInt(task.dateCompleted, 10),
			})
		}
	}
	return nil
}

/**
 * TaskList methods
 */

func (tasklist TaskList) getTasks() (TaskList, error) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return []Task{}, err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	// Omit first line
	records = records[1:]
	// Convert to Task
	var tasks []Task
	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		dateDue, _ := strconv.ParseInt(record[2], 10, 64)
		dateCreated, _ := strconv.ParseInt(record[3], 10, 64)
		dateCompleted, _ := strconv.ParseInt(record[4], 10, 64)
		tasks = append(tasks, Task{
			id:            id,
			text:          record[1],
			dateDue:       dateDue,
			dateCreated:   dateCreated,
			dateCompleted: dateCompleted,
		})
	}
	return tasks, err
}

func _writeHeadingToFile(f *os.File) {
	csvWriter := csv.NewWriter(f)
	csvWriter.Write([]string{"id", "text", "dateDue", "dateCreated", "dateCompleted"})
	csvWriter.Flush()
}
