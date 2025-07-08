package main

import (
	"awesomeProject1/internal/gamelogic"
	"awesomeProject1/internal/llm"
	"fmt"
	"time"
)

func main() {
	gameDataUnformated, err := llm.LlmChoices()

	if err != nil {
		fmt.Println(err)
		return
	}
	gameData, err := gamelogic.GameInitialization(gameDataUnformated)
	if err != nil {
		fmt.Println(err)
	}
	gameData.Players.HealingPotion = 5
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("intro", gameData.Dialogues[0].Content)
	time.Sleep(2 * time.Second)
	for i, boss := range gameData.Bosses {
		result := gamelogic.Combat(&gameData.Players, &boss, gameData.Dialogues[i+1].Content)
		if result == false {
			fmt.Println("you lost ")
			return

		}
		fmt.Println("end", gameData.Dialogues[len(gameData.Dialogues)-1].Content)
	}
}
