package main

import (
	"fmt"
)

type Relationship struct {
	IsStraight bool
	Partner    string
}

type Buddy struct {
	Name string
}

func main() {
	buddies := []Buddy{{"Abylay"}, {"Aigali"}}
	keksiki := make(map[Buddy][]Relationship, 2)
	for key := range buddies {
		if _, ok := keksiki[buddies[key]]; !ok {
			keksiki[buddies[key]] = make([]Relationship, 0, 3)
		}
		keksiki[buddies[key]] = append(keksiki[buddies[key]], Relationship{IsStraight: false, Partner: "Nur-Assyl"})
	}

	for _, v := range buddies {
		if v == struct{ Name string }{"Aigali"} { // or v.Name == "Aigali"
			keksiki[v] = append(keksiki[v], Relationship{IsStraight: true, Partner: "Atyn umytyp kaldym"})
			continue
		}
		keksiki[v] = append(keksiki[v], Relationship{IsStraight: true, Partner: "Molya"})
	}

	keksiki[Buddy{"Abylay"}] = append(keksiki[Buddy{"Abylay"}], Relationship{IsStraight: true, Partner: "Aizhan"})

	for kekser, kekserInfo := range keksiki {
		fmt.Print(kekser.Name, " -> ")
		fmt.Printf("%+v\n", kekserInfo)
	}
}
