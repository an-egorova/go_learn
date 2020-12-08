package main

import(
	"fmt"
	"os"
)

func main() {
	var first float64
	var second float64
	var oper float64
	var resultat float64
	fmt.Print("Введите первое число: ")
	fmt.Fscan(os.Stdin, &first)
	fmt.Print("Введите второе число: ")
	fmt.Fscan(os.Stdin, &second)
	fmt.Println("Выберите операцию:\n\t(1) Сложение\n\t(2) Вычитание\n\t(3) Умножение\n\t(4) Деление ")
	fmt.Fscan(os.Stdin, &oper)
	switch oper {
	case 1:
		resultat=first+second
		fmt.Print("Результат: ")
		fmt.Println(resultat)
		break
	case 2:
		resultat=first-second
		fmt.Print("Результат: ")
		fmt.Println(resultat)
		break
	case 3:
		resultat=first*second
		fmt.Print("Результат: ")
		fmt.Println(resultat)
		break
	case 4:
		resultat=first/second
		fmt.Print("Результат: ")
		fmt.Println(resultat)
		break
	default:
		break
	}
}
