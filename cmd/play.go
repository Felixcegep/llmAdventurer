package cmd

import (
	"awesomeProject1/internal/gamelogic"
	"awesomeProject1/internal/llm"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"time"
)

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Start your anime-themed adventure!",
	Run: func(cmd *cobra.Command, args []string) {

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

		color.Yellow("Intro: %s", gameData.Dialogues[0].Content)
		time.Sleep(2 * time.Second)
		for i, boss := range gameData.Bosses {
			result := gamelogic.Combat(&gameData.Players, &boss, gameData.Dialogues[i+1].Content)
			if result == false {
				color.Red("YOU LOST")
				return

			}
			fmt.Println("end", gameData.Dialogues[len(gameData.Dialogues)-1].Content)
			color.Green("You won the battle!")
		}
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}
