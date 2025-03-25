package main

import (
	"fmt"
)

func check_num(hidden_num int) int {
	var count = 0
	for {
		num := get_num()
		count += 1
		if num == hidden_num {
			fmt.Println("Вы угадали число")
			break
		} else if num > hidden_num {
			fmt.Println("Загаданное число меньше")
		} else {
			fmt.Println("Загаданное число больше")
		}
	}
	return count
}

func get_num() int {
	var num int
	fmt.Println("Попробуйте отгадать число:")
	for {
		fmt.Scan(&num)
		if num < 1 || num > 100 {
			fmt.Println("Вы ввели недопустимое число. Введитечисло от 1 до 100:")
		} else {
			break
		}
	}
	return num
}
