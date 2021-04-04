package main

import (
	"fmt"
	"strconv"
)

func findMissingNumber(input string) int{
 lengthNumber:=0
 stat:=false
 var result int
 missingNumber:=0
 for lengthNumber<len(input)/4 && !stat { //looping to get length number
	 lengthNumber++
	 num1,_:=strconv.Atoi(string(input[0:lengthNumber])) // convert string to int, get first value from input
	 num2,_:=strconv.Atoi(string(input[lengthNumber:lengthNumber*2])) // get 2nd value from input
	 num3,_:=strconv.Atoi(string(input[lengthNumber*2:lengthNumber*3])) //get 3rd value from input
	 num4,_:=strconv.Atoi(string(input[lengthNumber*3:lengthNumber*4])) //get 4th value from input
	 result = num2-num1
	 result = result *(num3-num2)
	 result = result *(num4-num3)

	 if result==1||result==2 { //check if result ==1 || result ==2
		 stat=true //change state to true
	 }
 }

 if stat{ //if stat !=false
 	listNumber:=len(input)/lengthNumber
 	for j:=0;j<listNumber-1;j++{ //looping to get missing number
		// get digit number from input and convert to int
 		number1,_:=strconv.Atoi(string(input[j*lengthNumber:(j+1)*lengthNumber]))
		number2,_:=strconv.Atoi(string(input[(j+1)*lengthNumber:(j+2)*lengthNumber]))
		//number2-number1 must be 1 if > 1 is missing number
		if number2-number1>1 {
			missingNumber = number1+1
		}
	}
 }
 return missingNumber
}

func main(){
	input:="12356789"
	find:=findMissingNumber(input)
	fmt.Println(find)
}
