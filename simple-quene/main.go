package main

import "fmt"

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

func (q *Queue) Clear() {
	q.queue = []string{}
}

func main() {
	queue := Queue{}

	for {
		fmt.Println("Добро пожаловать в Симулятор очереди!")
		fmt.Println("Доступные команды:")
		fmt.Println("1. Добавить человека в очередь")
		fmt.Println("2. Удалить человека из очереди")
		fmt.Println("3. Просмотреть очередь")
		fmt.Println("4. Очистить очередь")
		fmt.Println("5. Выйти")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var name string
			fmt.Print("Имя: ")
			fmt.Scanln(&name)
			queue.Add(name)
		case 2:
			removed := queue.Remove()
			if removed == "" {
			 fmt.Println("Очередь пуста!")
			} else {
			 fmt.Println("Удален:", removed)
			}
		case 3:
			queue.View()
		case 4:
			queue.Clear()
			fmt.Println("Очередь очищена")
		case 5:
			return
		}
	}
}
