package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sortBookByCtg(listBook []string) []string{
	category:= [...]string{"6","7","0","9","4","8","1","2","5","3"} //declare category

	var bookSelf []string
	//looping for sorting by category
	for _,ctg:= range category {
		var temp []string
		count:=0
		for i:=0;i<len(listBook);i++ {
			if strings.HasPrefix(listBook[i],ctg)==true { //check listbook does have category
				count++
				temp = append(temp, listBook[i]) //adding listbook[i] with same category to temp variable
			}
		}
		if count>0{ //check if book have same code category from category value
			if len(temp)>0{ //check length temp > 0
				//looping for sorting book by number
				for k:=0;k<len(temp)-1;k++{
					for l := 0; l < len(temp)-k-1; l++ {
						num1,_:=strconv.Atoi(temp[l][2:4])
						num2,_:=strconv.Atoi(temp[l+1][2:4])//convert to string
						if num1 < num2 {
							temp[l], temp[l+1] = temp[l+1], temp[l]//sorting book by book size (num1,num2)
						}
					}
				}
			}
			bookSelf = append(bookSelf, temp...) //adding value temp to bookself
		}
	}
	return bookSelf
}

func Contains(a string) []string{
	//this function for removing duplicate book
	listBook := strings.Split(a, " ")
	identity:= getIdentity(listBook) //call function getIdentity
	var find string
	var index int
	for _, n := range identity { //looping for value identity book
		var duplicatBook []string
		for k:=0;k<len(listBook);k++ { //looping for value listbook
			if strings.Contains(listBook[k],n) { // condition for get the book with same identity
				duplicatBook = append(duplicatBook, listBook[k]) // add value listbook to duplicateBook
				if len(duplicatBook) > 2 { //condition if duplicatBook > 2
					for i:=0;i<len(duplicatBook)-1;i++{//sorting duplicatebook by size book
						for j:= 0; j < len(duplicatBook)-i-1; j++ {
							num1,_:=strconv.Atoi(duplicatBook[j][2:4])
							num2,_:=strconv.Atoi(duplicatBook[j+1][2:4])
							//sorting duplicate book
							if num1 < num2 {
								duplicatBook[j], duplicatBook[j+1] = duplicatBook[j+1], duplicatBook[j]
							}
						}
					}
					find = duplicatBook[len(duplicatBook)-1] //duplicatebook to be removed
				}
			}
		}

	}
	//looping for get index duplicate book
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
	//looping to store 2 digit book to var identity
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
	for _,n:=range books { //looping value input book
		tall,_:=strconv.Atoi(n[2:4]) //get 2 last digit as size book from book value
		if len(n)>4 { //checking length character book
			lenBooks++
			lenBookValue = append(lenBookValue, n) //if true add book with length>4 to lenBookValue
		}else{
			//if condition false
			if n[1:2]!="A"&&n[1:2]!="B"&&n[1:2]!="C"&&n[1:2]!="D"&&n[1:2]!="E"&&n[1:2]!="F"&&n[1:2]!="G"&&n[1:2]!="H"&&n[1:2]!="I"&&n[1:2]!="J"&&n[1:2]!="K"&&n[1:2]!="L"&&n[1:2]!="M"&&n[1:2]!="N"&&n[1:2]!="O"&&n[1:2]!="P"&&n[1:2]!="Q"&&n[1:2]!="R"&&n[1:2]!="S"&&n[1:2]!="T"&&n[1:2]!="U"&&n[1:2]!="V"&&n[1:2]!="W"&&n[1:2]!="X"&&n[1:2]!="Y"&&n[1:2]!="Z" {
				//check title must = A-Z
				cekTitle++
				cekTitleValue = append(cekTitleValue, n)  //if true add book with title not A-Z to cekTitleValue
			}
			if tall<13 || tall>24{
				//check tall book must >13 and < 24
				cekSize++
				cekSizeValue = append(cekSizeValue, n) //if tall <13 and >24 add book to cekSizeValue
			}
		}
	}
	if lenBooks>0 { // check if character book >4
		fmt.Println("Kode buku melebihi batas, Maksimal kode buku 4 digit: ",lenBookValue)
	}
	if cekTitle>0{ //check if title book not A-Z
		fmt.Println("Kode buku title harus hurus besar A-Z: ",cekTitleValue)
	}
	if cekSize>0{ //check if size book <13 || >24
		fmt.Println("Panjang buku(2 digit terakhir) minimal 14 maksimal 24: ",cekSizeValue)
	}
	if lenBooks==0 && cekTitle==0 && cekSize==0 {
		listBook := Contains(input)
		fmt.Println(listBook)
	}
}
