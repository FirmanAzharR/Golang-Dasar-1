package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func palindrome(min int, max int) {
	maxInput := int(math.Pow(10,9))
	if min>0 && max>0 && max<=maxInput && min>max {
		fmt.Println("input must be numb1>0; numb1 < numb2;")
	}else{
		count:=0
		for i:= min; i<=max; i++ {
			reverseNumb:=""
			strNumb := strconv.Itoa(int(i))
			for length := len(strNumb); length > 0; length-- {
				reverseNumb += string(strNumb[length-1])
			}
			convertIntNumb, error := strconv.Atoi(reverseNumb)
			if error != nil {
				fmt.Println("Failure to cast String to int")
			}
			if i==convertIntNumb {
				count+=1
			}
		}
		fmt.Println(count)
	}
}

func scanner() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Input number example '1 10' : ")
		scanner.Scan()
		input := scanner.Text()
		var numb1 int
		var numb2 int
		if strings.ToLower(input) == "exit" {
			fmt.Println("Goodbye")
			break
		}

		split := strings.Split(input, " ")
		if len(split) != 2 {
			fmt.Println("Input must be 2 number separated with space ex: 21 31")
		} else {
			numb1, _ = strconv.Atoi(split[0])
			numb2, _ = strconv.Atoi(split[1])
			if numb1>0 && numb2>0 {
				palindrome(numb1, numb2)
			}else{
				fmt.Println("Input must be 2 number separated with space ex: 21 31")
			}
		}
	}
}

func main(){
	scanner()
}