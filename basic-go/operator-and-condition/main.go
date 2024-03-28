package main

import "fmt"

func main() {
	value := 2 + 3*2
	isCheck := value >= 5

	if isCheck {
		fmt.Println(value, "lebih besar dari 5")
	} else {
		fmt.Println(value, "lebih kecil dari 5")

	}

	value2 := 50

	if checkMod := value2 % 2; checkMod == 0 {
		fmt.Println("is True modulus")
	}

	var point = 6

	switch point {
	case 8:
		fmt.Println("perfect")
	// case (point < 8) && (point > 6):
	// 	fmt.Println("grateful bang")
	// 	fallthrough
	case 6, 5, 4:
		fmt.Println("awesome")

	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better!")
		}
	}
}
