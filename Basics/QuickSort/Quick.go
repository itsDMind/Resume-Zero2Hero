package main

import "fmt"

func quicksort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    pivot := arr[0]
    var left, right []int

    for _, x := range arr[1:] {
        if x < pivot {
            left = append(left, x)
        } else {
            right = append(right, x)
        }
    }

    return append(append(quicksort(left), pivot), quicksort(right)...)
}

func main() {
    unsortedArray := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
    sortedArray := quicksort(unsortedArray)
    fmt.Println(sortedArray) // Output: [1 1 2 3 3 4 5 5 5 6 9]
}
