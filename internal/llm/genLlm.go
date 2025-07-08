package llm

import (
	"context"
	"fmt"
	"google.golang.org/genai"
	"math/rand"
	"time"
)

var animeList = []string{
	"Attack on Titan",
	"Demon Slayer",
	"Sword Art Online",
	"Fullmetal Alchemist: Brotherhood",
	"My Hero Academia",
	"One Piece",
	"Naruto",
	"Bleach",
	"Jujutsu Kaisen",
	"Hunter x Hunter",
	"Fairy Tail",
	"Tokyo Ghoul",
	"Black Clover",
	"Blue Exorcist",
	"Fire Force",
	"Chainsaw Man",
	"Dragon Ball Z",
	"Dragon Ball Super",
	"Yu Yu Hakusho",
	"Inuyasha",
	"Akame ga Kill!",
	"Fate/Stay Night: Unlimited Blade Works",
	"Fate/Zero",
	"Vinland Saga",
	"The Seven Deadly Sins",
	"Soul Eater",
	"Seraph of the End",
	"Claymore",
	"D.Gray-man",
	"Trigun",
	"Berserk",
	"Noragami",
	"Mob Psycho 100",
	"One Punch Man",
	"Kill la Kill",
	"Magi: The Labyrinth of Magic",
	"Rurouni Kenshin",
	"Gurren Lagann",
	"Sword of the Stranger",
	"No Game No Life",
	"Log Horizon",
	"Tower of God",
	"Black Lagoon",
	"Psycho-Pass",
	"Re:Zero",
	"Dr. Stone",
	"Made in Abyss",
	"Goblin Slayer",
	"The Rising of the Shield Hero",
	"Bungo Stray Dogs",
}

func LlmChoices() (string, error) {
	rand.Seed(time.Now().UnixNano()) // Seed the random generator
	randomIndex := rand.Intn(len(animeList))
	randomAnime := animeList[randomIndex]
	fmt.Println("anime : ", randomAnime)
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		return "", err
	}

	config := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText(fmt.Sprintf(`
You are a JSON Game Story Generator.

You must ALWAYS respond in **valid JSON format only**. No explanations, no comments, no extra text. Only valid JSON.

Your task is to generate a complete adventure game story using the following EXACT JSON structure and naming conventions optimized for Go parsers with PascalCase keys:
Make the story inspired by the popular anime %s  
DIFFICULTY: EASY

⚖️ Game Balance Rule:
All stats and items must be internally balanced to ensure a fair and engaging EASY-difficulty experience.  
- Bosses should be beatable with reasonable use of healing, defense, and superPower.  
- The player's base stats must allow survival for several turns against bosses.  
- The superPower must be strong but limited or conditional to maintain balance.  
- Inventory items must meaningfully contribute to survival or offense, but not break challenge completely.

{
  "WorldIntro": "A short and engaging introduction to the game world.",
  "Player": {
	"HeroName": "The name of the main character.",
    "Health": Integer value (starting health),
    "Attack": Integer value (starting attack power),
    "Defense": Integer value (starting defense power),
	"superPower" :  {
		// ItemType 0 == heal, 1 == attack, 2 == defense
		// those item are inspired by the anime
		// this is to equilibrate the fight and make it possible to kill the boss
		// the stat of this object are OVERPOWERED TO EQUILIBRATE
      "ItemType": Integer value (starting defense power),
      "NameItem": "item are inspired by the anime",
      "Health": Integer value (starting defense power),
      "Attack": Integer value (starting defense power),
      "Defense": Integer value (starting defense power)
    }
	"Inventory": [
    {
		// ItemType 0 == heal, 1 == attack, 2 == defense
		// those item are inspired by the anime
		// those item are really powerfull for Health Attack Or Defense and are comparable to the boss ability 
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

`, randomAnime), genai.RoleUser),
	}

	result, _ := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text("generate it"),
		config,
	)

	formated := result.Text()[7 : len(result.Text())-3]
	return formated, nil
}
