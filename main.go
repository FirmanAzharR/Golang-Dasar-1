package main

import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type Result struct {
	InputNumber string
	TotalPalindrome int
}

func palindrome(input string) int {
	var min,max int
	split:=strings.Split(input, " ")
	min, _ = strconv.Atoi(split[0]) //convert to int
	max, _ = strconv.Atoi(split[1])
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

func main(){
	input:="1 10"
	callFunction:=palindrome(input)
	results:=Result{input,callFunction}
	templates:= template.Must(template.ParseFiles("templates/web-template.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if input := r.FormValue("input"); input != "" {
			results.TotalPalindrome = callFunction
		}
		if err:=templates.ExecuteTemplate(w, "web-template.html",results);err!=nil {
			http.Error(w, err.Error(),http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
