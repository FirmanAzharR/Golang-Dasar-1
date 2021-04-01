package main

import (
	"fmt"
	"strings"
)

func sortBookByCtg(bookCode string) []string {
	category:= [...]string{"6","7","0","9","4","8","1","2","5","3"}
	listBook := strings.Split(bookCode, " ")
	var bookshelf []string
	for _,ctg:= range category {
		for i:=0;i<len(listBook);i++ {
			if strings.HasPrefix(listBook[i],ctg)==true {
				bookshelf = append(bookshelf, listBook[i])
			}
		}
	}
	return bookshelf
}

func removeDuplicateBook(listBook []string) []string  {
	var bookshelf []string
	for i:=0;i<len(listBook);i++ {
		book1:=strings.Split(listBook[i],"")
		identityBook:=strings.Join(book1[0:2],"")
		for j:=i+1;j<len(listBook);j++{
			if strings.HasPrefix(listBook[j],identityBook)==true {
				bookshelf = append(bookshelf, listBook[i])
				fmt.Println(listBook[j],j)
			}
		}
	}
	return  bookshelf
}

func main(){
	listCode := "3A13 3A15 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14"
	sortByCtg:=sortBookByCtg(listCode)
	fmt.Println(sortByCtg)
	removeDuplicate:=removeDuplicateBook(sortByCtg)
	fmt.Println(removeDuplicate)
}
