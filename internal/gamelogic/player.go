package gamelogic

import (
	"encoding/json"
	"github.com/fatih/color"
	"time"
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
	HeroName      string `json:"HeroName"`
	Health        int    `json:"Health"`
	MaxHealth     int
	Attack        int      `json:"Attack"`
	Defense       int      `json:"Defense"`
	Inventory     []Object `json:"Inventory"`
	HealingPotion int
	SuperPower    Object `json:"superPower"`
}

func (p *Player) healingpotion() {
	const PV = 50
	if p.HealingPotion > 0 {
		p.Health += PV
		color.Green("💖 Vous récupérez %d PV grâce à une potion ! Il vous reste %d PV.\n", PV, p.Health)
		p.HealingPotion--
	} else {
		color.Red("❌ Plus de potion de soin ! Vous passez votre tour.\n")
	}
	time.Sleep(1 * time.Second)
}

type Boss struct {
	Name      string `json:"Name"`
	Health    int    `json:"Health"`
	MaxHealth int
	Attack    int `json:"Attack"`
	Defense   int `json:"Defense"`
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
	game.Players.MaxHealth = game.Players.Health
	for i := range game.Bosses {
		game.Bosses[i].MaxHealth = game.Bosses[i].Health
	}
	return game, nil
}
