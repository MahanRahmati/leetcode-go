package main

func romanToInt(s string) int {
	length := len(s)

	if length == 1 {
		return getNumber(s[0])
	}

	sum := 0
	sum = sum + getNumber(s[length-1])
	for i := length - 1; i > 0; i-- {
		now := getNumber(s[i-1])
		before := getNumber(s[i])
		if now < before {
			sum = sum - now
		} else {
			sum = sum + now
		}
	}
	return sum
}

func getNumber(r byte) int {
	if r == 73 {
		return 1
	}
	if r == 86 {
		return 5
	}
	if r == 88 {
		return 10
	}
	if r == 76 {
		return 50
	}
	if r == 67 {
		return 100
	}
	if r == 68 {
		return 500
	}
	if r == 77 {
		return 1000
	}
	return 0
}
