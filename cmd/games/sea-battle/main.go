package seabattle

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SeaBattleCmd = &cobra.Command{
	Use:   "sea-battle",
	Short: "Старый добрый морской бой.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		seaBattle()
	},
}

func init() {
}

func seaBattle() {
	fmt.Println("BATTLE!")
	Game()
}
