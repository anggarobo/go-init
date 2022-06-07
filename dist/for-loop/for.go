package forloop

import "fmt"

func LetsCount(num int) {
	for i := 0; i < num; i++ {
		fmt.Println(i)
	}
}

func WhichEven(num int) {
	if num < 1 {
		fmt.Println("0 value")
	}

	for num > 0 {
		if num % 2 == 0 {
			fmt.Println(num)
		}
		num--
	}
}

func LabelLoop(num int) {
	outerloop:
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			if i == 3 {
				break outerloop
			}
			fmt.Print("Matriks [", i, "] [", j, "]\n")
		}
	}
}