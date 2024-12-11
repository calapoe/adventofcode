package main

import (
	"fmt"
	"os"
)

func main() {
	day := 11
	in, err := os.ReadFile(fmt.Sprintf("2024/inputs/input%d.txt", day))

	if err != nil {
		fmt.Print(err)
	}

	input := string(in)
	//dayOne1(input)
	//dayOne2(input)
	//dayTwo2(input)
	//dayThree1(input)
	//dayThree2(input)
	//dayFour1(input)
	//dayFour2(input)
	//dayFive1(input)
	//dayFive2(input)
	//daySix1(input)
	//daySix2(input)
	//daySeven1(input)
	//daySeven2(input)
	//dayEight1(input)
	//dayEight2(input)
	//dayNine1(input)
	//dayNine2(input)
	//dayTen1(input)
	//dayTen2(input)
	dayEleven1(input)
	dayEleven2(input)
}
