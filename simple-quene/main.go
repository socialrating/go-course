package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Queue struct {
	queue []string
}

func (q *Queue) Add(newPerson string) {
	q.queue = append(q.queue, newPerson)
}

func (q *Queue) Remove() string {
	if len(q.queue) == 0 {
		return ""
	}
	removedPerson := q.queue[0]
	q.queue = q.queue[1:]
	return removedPerson
}

func (q *Queue) View() {
	if len(q.queue) == 0 {
		fmt.Println("Очередь пуста")
		return
	}
	fmt.Println("Очередь:")
	for _, person := range q.queue {
		fmt.Println(person)
	}
}

const queueIsEmpty = "Очередь пуста."

func (q *Queue) Clear() {
	q.queue = []string{}
}

func main() {
	queue := Queue{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Добро пожаловать в Симулятор очереди!")
	fmt.Println("Доступные команды:")
	fmt.Println("1. add [имя] - Добавить человека в очередь.")
	fmt.Println("2. serve      - Обслужить следующего в очереди.")
	fmt.Println("3. list       - Показать текущую очередь.")
	fmt.Println("4. clear      - Очистить очередь.")
	fmt.Println("5. exit       - Выйти из программы.")

	for {
		fmt.Print("> ")
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)
		command := parts[0]
		if len(parts) == 0 {
			continue
		}
		switch command {
		case "add":
			queue.Add(parts[1])
			fmt.Println("Добавлен:", parts[1])
		case "serve":
			removed := queue.Remove()
			if removed == "" {
				fmt.Println("Очередь пуста!")
			} else {
				fmt.Println("Обслужен:", removed)
			}
		case "list":
			queue.View()
		case "clear":
			queue.Clear()
			fmt.Println("Очередь очищена.")
		case "exit":
			fmt.Println("До свидания!")
			return
		}
	}
}
