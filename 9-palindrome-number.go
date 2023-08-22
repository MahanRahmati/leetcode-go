package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	string := strconv.Itoa(x)
	length := len(string)
	for i := range string {
		if string[i] != string[length-i-1] {
			return false
		}
	}
	return true
}
