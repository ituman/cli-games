package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var GAMES_LIST = []string{
	"guess",
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Список доступных игр.",
	Long:  `list показывает список игр, в которые ты можешь поиграть`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Доступные игры:\n\n")
		for _, game := range GAMES_LIST {
			fmt.Println("\t - ", game)
		}
		fmt.Print("\nЧтобы начать играть введи команду\n\n\tcli-games <game name>")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
