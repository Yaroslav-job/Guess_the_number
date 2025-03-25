package main

import (
	"fmt"
)

func play_game() {
	hidden_num := random()
	fmt.Println("Число загаданно")
	attempts := check_num(hidden_num)
	fmt.Printf("\nВы использовали %d попыток", attempts)
	play := play_again()
	if play == 1 {
		main()
	} else {
		fmt.Println("Спасибо за игру")
	}
}

func play_again() int {
	var play_again int
	fmt.Println("\nХотите сыграть ещё раз: \n1 - Да \n2 - Нет")
	for {
		fmt.Scan(&play_again)
		if play_again == 1 || play_again == 2 {
			break
		} else {
			fmt.Println("Вы ввели недопустимое число. Попробуйте ещё раз:")
		}
	}
	return play_again
}
