package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var nums []float64
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Println("Error in input: %v", err)
			return
		}

		if len(nums) > 2 && IsOutliner(n, nums) {
			n = int(CalculateAvg(nums))
		}

		nums = append(nums, float64(n))

		low, high := GetLimits(nums)

		fmt.Printf("%f %f\n", low, high)
	}
}
