package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/genai"
	"log"
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

func Llm_choices() (Game, error) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	config := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText(`
You are a JSON Game Story Generator.

You must ALWAYS respond in **valid JSON format only**. No explanations, no comments, no extra text. Only valid JSON.

Your task is to generate a complete adventure game story using the following EXACT JSON structure and naming conventions optimized for Go parsers with PascalCase keys:
make the history inspired by the populair anime 
func (seedrandomizer())
{
  "WorldIntro": "A short and engaging introduction to the game world.",
  "Player": {
	"HeroName": "The name of the main character.",
    "Health": Integer value (starting health),
    "Attack": Integer value (starting attack power),
    "Defense": Integer value (starting defense power)
	"Inventory": [
    {
		// ItemType 0 == heal, 1 == attack, 2 == defense
		// those item are inspired by the anime
      "ItemType": Integer value (starting defense power),
      "NameItem": "item are inspired by the anime",
      "Health": Integer value (starting defense power),
      "Attack": Integer value (starting defense power),
      "Defense": Integer value (starting defense power)
    },
    {
      "ItemType": Integer value (starting defense power),
      "NameItem": "item are inspired by the anime",
      "Health": Integer value (starting defense power),
      "Attack": Integer value (starting defense power),
      "Defense": Integer value (starting defense power)
    },
    {
      "ItemType": Integer value (starting defense power),
      "NameItem": "item are inspired by the anime",
      "Health": Integer value (starting defense power),
      "Attack": Integer value (starting defense power),
      "Defense": Integer value (starting defense power)
    }
  ]
  },
  "Bosses": [
    {
      "Name": "Boss name",
      "Health": Integer value (boss health),
      "Attack": Integer value (boss attack power),
      "Defense": Integer value (boss defense power)
    }
    // Exactly 3 unique bosses must be provided
  ],
  "Levels": [
    "Level 1 name",
    "Level 2 name",
    "Level 3 name"
    // Each level should match the boss's territory
  ],
  "Dialogues": [
    {
      "Type": "Intro",
      "Content": "Game starting dialogue that sets the quest."
    },
    {
      "Type": "Boss",
      "BossName": "Boss 1 name",
      "Content": "Dialogue when encountering this boss."
    },
    {
      "Type": "Boss",
      "BossName": "Boss 2 name",
      "Content": "Dialogue when encountering this boss."
    },
    {
      "Type": "Boss",
      "BossName": "Boss 3 name",
      "Content": "Dialogue when encountering this boss."
    },
    {
      "Type": "Victory",
      "Content": "Final dialogue after defeating all bosses."
    }
  ]
}

⚙️ Important Rules:
- Use PascalCase for all keys.
- The 'Bosses' list must have EXACTLY 3 bosses.
- The 'Dialogues' array must always include:
    - One intro dialogue (Type: "Intro").
    - One boss dialogue for EACH boss (Type: "Boss", BossName must match exactly).
    - One victory dialogue (Type: "Victory").
- All fields must be fully populated. No empty fields.
- The boss names in the 'Bosses' array must EXACTLY match the 'BossName' fields in the boss dialogues.
- Always return strictly valid JSON. No other text or explanations.

This format is fully compatible with Go struct parsing without needing json tags or additional mapping.

`, genai.RoleUser),
	}

	result, _ := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text("generate it"),
		config,
	)
	fmt.Println(result.Text())

	formated := result.Text()[7 : len(result.Text())-3]
	var game Game
	err = json.Unmarshal([]byte(formated), &game)
	if err != nil {
		return Game{}, err
	}
	return game, nil
}
