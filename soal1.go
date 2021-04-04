package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sortBookByCtg(listBook []string) {
	category:= [...]string{"6","7","0","9","4","8","1","2","5","3"}

	var bookSelf []string
	for _,ctg:= range category {
		var temp []string
		count:=0
		for i:=0;i<len(listBook);i++ {
			if strings.HasPrefix(listBook[i],ctg)==true {
				count++
				temp = append(temp, listBook[i])
			}
		}
		if count>0{
			if len(temp)>0{
				for k:=0;k<len(temp)-1;k++{
					for l := 0; l < len(temp)-k-1; l++ {
						num1,_:=strconv.Atoi(temp[l][2:4])
						num2,_:=strconv.Atoi(temp[l+1][2:4])
						if num1 < num2 {
							temp[l], temp[l+1] = temp[l+1], temp[l]
						}
					}
				}
			}
			bookSelf = append(bookSelf, temp...)
		}
	}
	fmt.Println(bookSelf)
}

func Contains(a string) []string{
	listBook := strings.Split(a, " ")
	identity:= getIdentity(listBook)
	var find string
	var index int
	for _, n := range identity {
		var duplicatBook []string
		for k:=0;k<len(listBook);k++ {
			if strings.Contains(listBook[k],n) {
				duplicatBook = append(duplicatBook, listBook[k])
				if len(duplicatBook) > 2 {
					for i:=0;i<len(duplicatBook)-1;i++{
						for j:= 0; j < len(duplicatBook)-i-1; j++ {
							num1,_:=strconv.Atoi(duplicatBook[j][2:4])
							num2,_:=strconv.Atoi(duplicatBook[j+1][2:4])
							if num1 < num2 {
								duplicatBook[j], duplicatBook[j+1] = duplicatBook[j+1], duplicatBook[j]
							}
						}
					}
					find = duplicatBook[len(duplicatBook)-1]
				}
			}
		}
	}

	for n, s := range listBook {
		if find == s {
			index = n
		}
	}
	listBook[index] = listBook[len(listBook)-1] // Copy last element to index i.
	listBook[len(listBook)-1] = ""   // Erase last element (write zero value).
	listBook = listBook[:len(listBook)-1]   // Truncate slice.

	return listBook
}

func getIdentity(book []string) []string{
	var identity []string
	for _,x:=range book {
		identity = append(identity, x[0:2])
	}

	check := make(map[string]int)
	res := make([]string,0)
	for _, val := range identity {
		check[val] = 1
	}
	for letter, _ := range check {
		res = append(res,letter)
	}
	return res
}

func main(){
	book := "3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14"
	listBook := Contains(book)
    sortBookByCtg(listBook)
}
