package main

import (
	"fmt"
	f "guess-it-1/functions"
	"os"
	"strconv"
	"strings"
)

func main() {
	// fmt.Println(len(os.Args))

	file := os.Args[1]
	// fmt.Println(file)
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Fprint(os.Stderr, "The file doesn't exist")
		os.Exit(1)
	}
	File_Content := strings.Split(strings.TrimSpace(string(data)), "\n")
	var Numbers []int
	temp := []string{}
	for _, str := range File_Content {
		check := false
		for _, char := range str {
			if char < '0' || char > '9' {
				break
			} else {
				check = true
			}
		}
		if check {
			temp = append(temp, str)
		}
	}
	File_Content = temp
	// fmt.Println(len(File_Content))
	for _, str := range File_Content {
		temp, _ := strconv.Atoi(str)
		// fmt.Println(temp)
		Guess_The_Range(Numbers)
		Numbers = append(Numbers, int(temp))
		// fmt.Println(Numbers)
	}

}

func Guess_The_Range(Numbers []int) {
	var low, hight float64
	var num int
	Av := f.Average(Numbers)

	Va, SD := f.Variance(Numbers, Av)
	fmt.Println(Av, SD, Va)
	if len(Numbers) == 0 {
		low = 0
		hight = 100
	} else if len(Numbers) < 1 {
		num = Numbers[len(Numbers)-1]
		low = Av
		hight = Av
	} else {
		num = Numbers[len(Numbers)-1]
		low = Av - SD
		hight = Av + SD
	}
	fmt.Println(low, "<", num, ">", hight)
}
