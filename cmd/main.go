package main

import (
	"awesomeProject1/internal/gamelogic"
	"awesomeProject1/internal/llm"
	"fmt"
)

func main() {
	gameData, err := llm.Llm_choices()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("intro", gameData.Dialogues[0].Content)
	fmt.Println("player", gameData.Players)

	for i, boss := range gameData.Bosses {
		gamelogic.Combat(&gameData.Players, &boss, gameData.Dialogues[i+1].Content)
	}
	fmt.Println("end", gameData.Dialogues[len(gameData.Dialogues)-1].Content)

}
