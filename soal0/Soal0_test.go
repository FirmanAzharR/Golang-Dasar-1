package main

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T)  {
	result:= palindrome(1,100)
	result2:= palindrome(100,30)
	if result!=18{
		t.Error("result palindrom 1 - 100 must 18")
	}
	if result2!=0{
		t.Error("an error has occurred in the input")
	}
	fmt.Println("TestPalindrome Done")
}
