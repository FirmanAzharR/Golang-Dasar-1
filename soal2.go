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
 for lengthNumber<len(input)/4 && !stat {
	 lengthNumber++
	 num1,_:=strconv.Atoi(string(input[0:lengthNumber]))
	 num2,_:=strconv.Atoi(string(input[lengthNumber:lengthNumber*2]))
	 num3,_:=strconv.Atoi(string(input[lengthNumber*2:lengthNumber*3]))
	 num4,_:=strconv.Atoi(string(input[lengthNumber*3:lengthNumber*4]))
	 result = num2-num1
	 result = result *(num3-num2)
	 result = result *(num4-num3)

	 if result==1||result==2 {
		 stat=true
	 }
 }

 if stat{
 	listNumber:=len(input)/lengthNumber
 	for j:=0;j<listNumber-1;j++{
 		number1,_:=strconv.Atoi(string(input[j*lengthNumber:(j+1)*lengthNumber]))
		number2,_:=strconv.Atoi(string(input[(j+1)*lengthNumber:(j+2)*lengthNumber]))
		if number2-number1>1 {
			missingNumber = number1+1
		}
	}
 }
 return missingNumber
}

func main(){
	input:="123567"
	find:=findMissingNumber(input)
	fmt.Println(find)
}
