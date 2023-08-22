package main

func twoSum(nums []int, target int) []int {
	output := make([]int, 2)
out:
	for i, number := range nums {
		for j, secondNumber := range nums {
			if number+secondNumber == target {
				if i == j {
					break
				}
				output[0] = i
				output[1] = j
				break out
			}
		}
	}
	return output
}
