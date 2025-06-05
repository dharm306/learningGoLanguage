package main

import "fmt"

type Cricketer interface {
	play()
	role()
}

type Cricket struct {
	Player Cricketer
}

type Batsman struct{}

func (b Batsman) play() {
	fmt.Println("Batsman is playing.")
}
func (b Batsman) role() {
	fmt.Println("Batsman is batting.")
}

type Bowler struct{}

func (b Bowler) play() {
	fmt.Println("Bowler is playing.")
}
func (b Bowler) role() {
	fmt.Println("Bowler is bowling.")
}

type AllRounder struct{}

func (a AllRounder) play() {
	fmt.Println("AllRounder is playing.")
}
func (a AllRounder) role() {
	fmt.Println("AllRounder is batting and bowling.")
}

func main() {
	fmt.Println("This program is based on Interfcae in Go.")
	cricket := Cricket{}
	role := ""
	fmt.Println("Enter the role of the player (batsman/bowler/allrounder):")
	fmt.Scanln(&role)
	switch role {
	case "batsman":
		cricket.Player = Batsman{}
	case "bowler":
		cricket.Player = Bowler{}
	case "allrounder":
		cricket.Player = AllRounder{}
	default:
		fmt.Println("Invalid role. Please enter batsman, bowler, or allrounder.")
		return
	}
	cricket.Player.play()
	cricket.Player.role()
	fmt.Println("Player role and action executed successfully.")

}
