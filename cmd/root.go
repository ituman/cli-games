package cmd

import (
	"cli-games/cmd/games/guess"
	seabattle "cli-games/cmd/games/sea-battle"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "cli-games",
	Short: "cli-games - это набор игр, в которые можно поиграть в консоле)",
	Long: `
Можно поиграть в игры, пока выполняется миграция, 
выкатывается обновление на продакшн или ты просто устал чинить баги)
	
Чтобы поиграть - выбери игру, для этого введи команду
	cli-games list

Затем начни ее с помощью команды
	cli-games <game name> start
	`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Да начнется игра!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli-games.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(guess.GuessCmd)
	rootCmd.AddCommand(seabattle.SeaBattleCmd)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cli-games" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cli-games")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
