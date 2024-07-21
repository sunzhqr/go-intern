package main

import "fmt"

type Permission struct {
	Role   string
	Access bool
}

type User struct {
	Name string
}

func main() {
	var users = []User{{"Harry"}, {"Voldemort"}}
	var hogwarts = make(map[User][]Permission, 2)
	for key := range users {
		fmt.Println(users[key])
	}
	for key := range users {
		if _, ok := hogwarts[users[key]]; !ok {
			hogwarts[users[key]] = make([]Permission, 0, 2)
		}
		hogwarts[users[key]] = append(hogwarts[users[key]], Permission{Role: "read_only", Access: true})
	}

	for _, user := range users {
		hogwarts[user] = append(hogwarts[user], Permission{Role: "avada_kadavra", Access: true})
	}

	hogwarts[User{"Harry"}] = append(hogwarts[User{"Harry"}], Permission{"auto_mirror", true})

	for ind, mag := range hogwarts {
		fmt.Print("mag = ", ind, mag, "->")
		fmt.Printf("%+v\n", hogwarts[ind])
	}
}
