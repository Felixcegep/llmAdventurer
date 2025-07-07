package gamelogic

import (
	"encoding/json"
)

type Game struct {
	Intro     string     `json:"WorldIntro"`
	Players   Player     `json:"Player"`
	Bosses    []Boss     `json:"Bosses"`
	Levels    []string   `json:"Levels"`
	Dialogues []Dialogue `json:"Dialogues"`
}
type itemType int

const (
	Heal itemType = iota
	Attack
	Defense
)

type Object struct {
	ItemType itemType `json:"ItemType"`
	NameItem string   `json:"NameItem"`
	Health   int      `json:"Health"`
	Attack   int      `json:"Attack"`
	Defense  int      `json:"Defense"`
}
type Player struct {
	HeroName      string   `json:"HeroName"`
	Health        int      `json:"Health"`
	Attack        int      `json:"Attack"`
	Defense       int      `json:"Defense"`
	Inventory     []Object `json:"Inventory"`
	HealingPotion int
}

func (p *Player) attack(attack, defense int) {

}
func (p *Player) itemdefense(defenseadd int) {
	p.Defense += defenseadd
}
func (p *Player) healingpotion(healingpotion int) {
	const PV = 50
	p.Health += PV
}

type Boss struct {
	Name    string `json:"Name"`
	Health  int    `json:"Health"`
	Attack  int    `json:"Attack"`
	Defense int    `json:"Defense"`
}
type Dialogue struct {
	Timetype string `json:"Type"`
	Content  string `json:"Content"`
}

func GameInitialization(generated string) (Game, error) {
	var game Game
	err := json.Unmarshal([]byte(generated), &game)
	if err != nil {
		return Game{}, err
	}
	return game, nil
}
