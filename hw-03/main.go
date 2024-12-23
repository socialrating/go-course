package main

import "fmt"

type Man struct {
	Name     string
	LastName string
	Age      int
	Gender   string
	Crimes   int
}

func main() {
	people := map[string]Man{
		"Анна":      {Name: "Анна", LastName: "Воронина", Age: 33, Gender: "female", Crimes: 3},
		"Дмитрий":   {Name: "Дмитрий", LastName: "Соколов", Age: 29, Gender: "male", Crimes: 0},
		"Михаил":    {Name: "Михаил", LastName: "Журавлев", Age: 45, Gender: "male", Crimes: 5},
		"Елизавета": {Name: "Елизавета", LastName: "Ковальская", Age: 25, Gender: "female", Crimes: 2},
		"Ольга":     {Name: "Ольга", LastName: "Ромашова", Age: 16, Gender: "female", Crimes: 1},
	}

	suspects := []string{"Анна", "Елизавета", "Михаил"}

	var mostCriminal Man
	for _, name := range suspects {
		if person, ok := people[name]; ok {
			if person.Crimes > mostCriminal.Crimes {
				mostCriminal = person
			}
		}
	}

	if mostCriminal.Name != "" {
		fmt.Println("Наиболее криминальная личность:" , mostCriminal.Name, mostCriminal.LastName, "количество преступлений:", mostCriminal.Crimes)
	} else {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым.")
	}
}