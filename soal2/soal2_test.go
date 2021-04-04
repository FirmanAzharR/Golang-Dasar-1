package main

import (
	"fmt"
	"testing"
)

func TestMissingNumber(t *testing.T)  {
	input:="100110021003100410061007100810091010"
	//input2:="1234567891012"
	result:= findMissingNumber(input)
	if result!=1005{
		t.Error("missing number incorrect")
	}
	fmt.Println("TestMissingNumber Done")
}
