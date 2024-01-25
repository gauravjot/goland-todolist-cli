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

func writeTaskToFile(isNew bool, task Task) (Task, error) {
	// Open file
	f, err := os.OpenFile("userdata.csv", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return Task{}, err
	}
	defer f.Close()

	lines, err := readCSVRecords(f)
	if err != nil {
		fmt.Println(err)
		return Task{}, err
	}
	if len(lines) == 0 {
		writeHeadingToFile(f)
	}
	if isNew {
		// If it is new task, read last line for id
		var new_id int
		if len(lines) > 0 {
			id, _ := strconv.Atoi(lines[len(lines)-1][0])
			new_id = id + 1
		} else {
			new_id = 0
		}
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
	} else {
		// If we are updating a task, read all lines and update the line with the id
		var isFound bool
		for i, line := range lines {
			if line[0] == strconv.Itoa(int(task.id)) {
				lines[i][1] = strings.TrimSpace(task.text)
				lines[i][2] = strconv.FormatInt(task.dateDue, 10)
				isFound = true
				break
			}
		}
		if !isFound {
			return Task{}, fmt.Errorf("Task with id %d not found", task.id)
		}

		// Write to file
		os.Remove("userdata.csv")
		f, _ := os.Create("userdata.csv")
		csvWriter := csv.NewWriter(f)
		csvWriter.WriteAll(lines)

		return task, nil
	}
}

func deleteTaskFromFile(id int) error {
	f, err := os.OpenFile("userdata.csv", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()
	lines, err := readCSVRecords(f)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var isFound bool
	for i, line := range lines {
		if line[0] == strconv.Itoa(int(id)) {
			lines = append(lines[:i], lines[i+1:]...)
			isFound = true
			break
		}
	}
	// Write to file
	if !isFound {
		return fmt.Errorf("Task with id %d not found", id)
	}

	os.Remove("userdata.csv")
	file, _ := os.Create("userdata.csv")
	csvWriter := csv.NewWriter(file)
	csvWriter.WriteAll(lines)
	return nil
}

func readCSVRecords(file *os.File) ([][]string, error) {
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	return records, err
}

func writeHeadingToFile(f *os.File) {
	csvWriter := csv.NewWriter(f)
	csvWriter.Write([]string{"id", "text", "dateDue", "dateCreated", "dateCompleted"})
	csvWriter.Flush()
}
