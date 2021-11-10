package main

import "fmt"

func main() {
	var basket1, basket2, basket3 int

	fmt.Println("Введите ёмкость корзин:")
	fmt.Scan(&basket1)
	fmt.Scan(&basket2)
	fmt.Scan(&basket3)

	var counter1, counter2, counter3 int
	for {
		if counter1 < basket1 {
			counter1++
			continue
		}

		if counter2 < basket2 {
			counter2++
			continue
		}

		if counter3 < basket3 {
			counter3++
			continue
		}

		break
	}
	fmt.Println(counter1)
	fmt.Println(counter2)
	fmt.Println(counter3)
}
