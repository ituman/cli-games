package guess

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var GuessCmd = &cobra.Command{
	Use:   "guess",
	Short: "Игра \"угадай число\".",
	Long: `В данной игре тебе предстоит угадать число, которое я загадал.

Ты можешь делать предположение, какое это число, а я буду говорить, верное ли оно.

Если число, которое ты предположил не верное, то я буду давать тебе подсказку - 
говорить, больше или меньше твое число, чем мое.

Удачи!
	`,
	Run: func(cmd *cobra.Command, args []string) {
		flevel, _ := cmd.Flags().GetString("level")
		level, err := strconv.Atoi(flevel)
		if err != nil {
			level = 0
			fmt.Println("Ты даже сложность указать не можешь, пусть она будет равна 0 для тебя.")
		}
		guessGame(level)
	},
}

func init() {
	var level string
	GuessCmd.Flags().StringVarP(&level, "level", "l", "5", "Сложность игры, значения от 0 до 10.")
}

func guessGame(level int) {
	fmt.Printf("level: %v\n", level)

	rand.Seed(time.Now().UnixNano())
	myNum := rand.Intn(20 * level)
	if level > 5 {
		myNum -= 20 * level / 2
	}

	reader := bufio.NewReader(os.Stdin)
	inp, err := reader.ReadString('\n')
	guessNum, err := strconv.Atoi(strings.Trim(inp, " \n"))
	if err != nil {
		fmt.Println(err)
	}

	for myNum != guessNum {
		if myNum > guessNum {
			fmt.Println("Не угадал! Мое число больше)")
		} else {
			fmt.Println("Не угадал! Мое число меньше)")
		}
		inp, err = reader.ReadString('\n')
		guessNum, err = strconv.Atoi(strings.Trim(inp, " \n"))
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Printf("Ты угадал, это число: %v", myNum)
}
