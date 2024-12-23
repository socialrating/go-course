package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	c := NewClient()
	valCurs, err := c.GetCbrDaily()
	if err != nil {
		panic(err)
	}

	sort.Slice(valCurs.Valute, func(i, j int) bool {
		valueI := strings.Replace(valCurs.Valute[i].Value, ",", ".", -1)
		floatValueI, err := strconv.ParseFloat(valueI, 64)
		if err != nil {
			fmt.Println(err)
		}
		valueJ := strings.Replace(valCurs.Valute[j].Value, ",", ".", -1)
		floatValueJ, err := strconv.ParseFloat(valueJ, 64)
		if err != nil {
			fmt.Println(err)
		}
		return floatValueI > floatValueJ
	})

	fmt.Println("Список валют")
	for _, valute := range valCurs.Valute {
		fmt.Printf("%s - %s; курс - %v\n", valute.CharCode, valute.Name, valute.Value)
	}
}
