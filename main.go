package main

import "fmt"

type Player struct {
	health int
}

func takeDamagefromExplosion(player *Player) {
	fmt.Println("player is taking damage form explosion")
	dmg := 10
	player.health -= dmg
}

// pointer 8 byte long integer to a specific slot in memory
func main() {
	player := Player{
		health: 100,
	}
	takeDamagefromExplosion(&player)
	fmt.Println(player.health)

}
func Super(s *Saiyan) {
	s.Power += 10000
}

func (s *Saiyan) NameChange(name string) {
	s.Name = name
}

type Saiyan struct {
	Name  string
	Power int
}
