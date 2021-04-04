package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sortBookByCtg(listBook []string) []string{
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
	return bookSelf
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
	sortedBook:=sortBookByCtg(listBook)
	return sortedBook
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
	input := "3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14"
	books:= strings.Split(input, " ")
	lenBooks:=0
	cekTitle:=0
	cekSize:=0
	var lenBookValue []string
	var cekTitleValue []string
	var cekSizeValue []string
	for _,n:=range books {
		tall,_:=strconv.Atoi(n[2:4])
		if len(n)>4 {
			lenBooks++
			lenBookValue = append(lenBookValue, n)
		}else{
			if n[1:2]!="A"&&n[1:2]!="B"&&n[1:2]!="C"&&n[1:2]!="D"&&n[1:2]!="E"&&n[1:2]!="F"&&n[1:2]!="G"&&n[1:2]!="H"&&n[1:2]!="I"&&n[1:2]!="J"&&n[1:2]!="K"&&n[1:2]!="L"&&n[1:2]!="M"&&n[1:2]!="N"&&n[1:2]!="O"&&n[1:2]!="P"&&n[1:2]!="Q"&&n[1:2]!="R"&&n[1:2]!="S"&&n[1:2]!="T"&&n[1:2]!="U"&&n[1:2]!="V"&&n[1:2]!="W"&&n[1:2]!="X"&&n[1:2]!="Y"&&n[1:2]!="Z" {
				cekTitle++
				cekTitleValue = append(cekTitleValue, n)
			}
			if tall<13 || tall>24{
				cekSize++
				cekSizeValue = append(cekSizeValue, n)
			}
		}
	}
	if lenBooks>0 {
		fmt.Println("Kode buku melebihi batas, Maksimal kode buku 4 digit: ",lenBookValue)
	}
	if cekTitle>0{
		fmt.Println("Kode buku title harus hurus besar A-Z: ",cekTitleValue)
	}
	if cekSize>0{
		fmt.Println("Panjang buku(2 digit terakhir) minimal 14 maksimal 24: ",cekSizeValue)
	}
	if lenBooks==0 && cekTitle==0 && cekSize==0 {
		listBook := Contains(input)
		fmt.Println(listBook)
	}
}
