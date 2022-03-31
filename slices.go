package main

import "fmt"

func main() {
	//Insert Code here
	weeklySpending := []float64{9.50, 8, 10.20, 7.40}

	fmt.Println(len(weeklySpending))
	fmt.Println(cap(weeklySpending))

	fmt.Println(weeklySpending[2])
	weeklySpending[2] = 9.80
	fmt.Println(weeklySpending)

	weeklySpending = append(weeklySpending, 8.40, 9.40, 7.20)
	fmt.Println(len(weeklySpending))
	fmt.Println(cap(weeklySpending))

	sub := weeklySpending[3:]
	fmt.Println(sub)
	fmt.Println(len(sub))
	fmt.Println(cap(sub))
}
