package main

import "fmt"

type Sigma struct {
	Level   int
	IsSigma bool
}

func (s *Sigma) GetInfo() {
	if s.Level == 0 {
		fmt.Println("either you are Abylay or Tamerlan(((")
	} else if s.Level > 50 && s.Level < 100 {
		fmt.Println("You are Zhandos and you are sigma!")
		s.IsSigma = true
	} else if s.Level == 100 {
		s.IsSigma = true
		fmt.Println("You are Aigali Bratan and you are king of sigmas!")
	} else {
		fmt.Println("Idi nahui Nura, peace of shit")
	}
}

type Character struct {
	Nick string
	Sigma
}

func (c *Character) SetLevel(level int) {
	c.Level = level
}
func main() {
	user1 := Character{Nick: "zhakyndama"}
	user1.SetLevel(0)
	user1.GetInfo()
	user2 := Character{Nick: "aigalisultankul"}
	user2.SetLevel(100)
	user2.GetInfo()
}
