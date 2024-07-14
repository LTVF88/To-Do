package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "strings"
)

func input(title string) string {
	fmt.Print(title)
	var s string
	_, err := fmt.Scanln(&s)
	if err != nil {
		fmt.Println(err)
	}
	return s
}
func Menu() {
	fmt.Printf(`
1.Добавить задачу
2.Список задач
3.Удалить задачу
` + "\n")
}
func AddTask() {
	NewTask := input("Напишите новую задачу: ")
	readFile, err := os.OpenFile(`c:\To-Do\Tasks.txt`, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()
	FileLen, err := os.ReadFile(`c:\To-Do\Tasks.txt`)
	if err != nil {
		fmt.Println(err)
	}
	if len(FileLen) == 0 {
		if _, err := readFile.WriteString(NewTask); err != nil {
			fmt.Println(err)
		}
	} else {
		if _, err := readFile.WriteString("\n" + NewTask); err != nil {
			fmt.Println(err)
		}

	}
}
func Tasks() {
	readFile, err := os.Open(`c:\To-Do\Tasks.txt`)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()
	for i, line := range lines {
		fmt.Printf("%d: %s\n", i+1, line)
	}
}

//функция для удаления Задчач
// func DeleteTask() {
// 	Tasks()
// 	TaskToDelete := input("Какую задачу удалить?: ")
// 	readFile, err := os.Open(`c:\To-Do\Tasks.txt`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fileScanner := bufio.NewScanner(readFile)
// 	fileScanner.Split(bufio.ScanLines)
// 	var lines []string
// 	for fileScanner.Scan() {
// 		lines = append(lines, fileScanner.Text())
// 	}
// 	fmt.Println(lines)
// }

type Task struct {
	Description string
	Priority    string
	Completed   bool
}

func main() {
	for {
		Menu()
		switch input("Выберите действие: ") {
		//Добавить задачу
		case "1":
			AddTask()
		//Список задач
		case "2":
			Tasks()
			// case "3":
			// 	DeleteTask()
		}
	}

}
