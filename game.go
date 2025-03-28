package main

import (
	"fmt"
	"os"
)

func play_game() {
	hidden_num := random()
	fmt.Println("Загаданно число от 1 до 100")
	attempts := check_num(hidden_num)
	fmt.Printf("\nВы использовали %d попыток", attempts)
	game_over()
}

func game_over() {
	var play_again int
	fmt.Println("\nХотите сыграть ещё раз: \n1 - Да \n2 - Нет")
	for {
		fmt.Scan(&play_again)
		if play_again == 1 {
			main()
		} else if play_again == 2 {
			os.Exit(0)
		} else {
			fmt.Printf("Вы ввели недопустимое число. Попробуйте ещё раз:")
		}
	}
}
