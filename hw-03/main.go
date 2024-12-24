package main

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func NewMan(name, lastName string, age int, gender string, crimes int) Man {
	return Man{
		Name:     name,
		LastName: lastName,
		Age:      age,
		Gender:   gender,
		Crimes:   crimes,
	}
}

func main() {
	people := map[string]Man{
		"Анна":      NewMan("Анна", "Воронина", 33, "female", 3),
		"Дмитрий":   NewMan("Дмитрий", "Соколов", 29, "male", 0),
		"Михаил":    NewMan("Михаил", "Журавлев", 45, "male", 5),
		"Елизавета": NewMan("Елизавета", "Ковальская", 25, "female", 2),
		"Ольга":     NewMan("Ольга", "Ромашова", 16, "female", 1),
	}

	suspects := []string{"Анна", "Елизавета", "Михаил"}

	var mostCriminal Man
	for _, name := range suspects {
		person, ok := people[name]
		if !ok {
			continue 
		}

		if person.Crimes > mostCriminal.Crimes {
			mostCriminal = person
		}
	}

	if mostCriminal.Name != "" {
		fmt.Println("Наиболее криминальная личность:", mostCriminal.Name, mostCriminal.LastName, "количество преступлений:", mostCriminal.Crimes)
	} else {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым.")
	}
}
