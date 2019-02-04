package main

import (
	"fmt"
	"math"
)

func calcEntropy(input []float64, size int){
	var count int
	count=0
	var entropy float64
	entropy=0
	for count=0;count<size;count++{
		entropy += input[count] * math.Log2(float64(1)/input[count]) - input[count] * math.Log2(input[count])
	}
	fmt.Println("Entropy value:", entropy)
}

func main() {
	fmt.Println("Enter the total number of datapoints")
        var total_datapoints int
	fmt.Scanf("%d", &total_datapoints)

	fmt.Println("Enter the number of similar points")
	var similar_points int
	fmt.Scanf("%d", &similar_points)

        var different_points int
        different_points = total_datapoints - similar_points
	
	var total_diff float64
	var total_sim float64
	var count int

	total_diff=0.5
	total_sim=0.5
	count=0

	fl_array := make([]float64, total_datapoints)

	if different_points == 0 {
		for count=0;count<similar_points;count++{
			fl_array[count] = float64(float64(1)/float64(similar_points))
		}
		calcEntropy(fl_array, similar_points)
		return
	}

	for count=0;count<similar_points;count++{
		similar_val := float64(total_sim/float64(similar_points))
		fl_array[count] = similar_val
		fmt.Println(similar_val)
	}

	buffer:=count
	count=0
	for count=0;count<different_points-1;count++{
                diff_val := float64(total_diff/float64(2))
                fmt.Println(diff_val)
		fl_array[count + buffer] = diff_val
		total_diff -= diff_val
        }
	fl_array[count+buffer] = total_diff
	fmt.Println(total_diff)
	calcEntropy(fl_array, count+buffer+1)
}

