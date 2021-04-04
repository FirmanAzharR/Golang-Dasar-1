package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func palindrome(min int, max int) int {
	maxInput := int(math.Pow(10,9))// maximal input number
	var result int
	//check input min and max>0 and min>max
	if min>0 && max>0 && max<=maxInput && min>max {
		fmt.Println("input must be numb1>0; numb1 < numb2;")
	}else{
		count:=0
		//looping for min to max from input
		for i:= min; i<=max; i++ {
			reverseNumb:=""
			strNumb := strconv.Itoa(int(i))//conversion int to string
			for length := len(strNumb); length > 0; length-- { //looping strNumb after convert to string
				reverseNumb += string(strNumb[length-1]) //reverse strNumb
			}
			convertIntNumb, error := strconv.Atoi(reverseNumb) //convert string to int
			if error != nil { //check error conversion
				fmt.Println("Failure to cast String to int")
			}
			if i==convertIntNumb { //check number (i) == convertIntNumb (reverse number)
				count+=1 // increment if condition true
			}
		}
		result=count // result total palindrom
	}
	return result
}

func scanner() {
	for {
		scanner := bufio.NewScanner(os.Stdin) //read input
		fmt.Printf("Input number example '1 10' : ")
		scanner.Scan()
		input := scanner.Text()
		var numb1 int
		var numb2 int
		//if input == exit close the program
		if strings.ToLower(input) == "exit" {
			fmt.Println("Goodbye")
			break
		}

		split := strings.Split(input, " ") //split string to array
		if len(split) != 2 { //check input
			fmt.Println("Input must be 2 number separated with space ex: 21 31")
		} else {
			numb1, _ = strconv.Atoi(split[0]) //convert to int
			numb2, _ = strconv.Atoi(split[1])
			if numb1>0 && numb2>0 { //check converted number > 0
				result:=palindrome(numb1, numb2) //call function palindrom
				fmt.Println(result) // print result
			}else{
				fmt.Println("Input must be 2 number separated with space ex: 21 31")
			}
		}
	}
}

func main(){
	scanner()
}