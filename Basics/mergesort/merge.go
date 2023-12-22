package main

import "fmt"

func mergeSort(arr []int) []int {
    if len(arr) <= 1 {
        return arr
    }

    middle := len(arr) / 2
    left := mergeSort(arr[:middle])
    right := mergeSort(arr[middle:])

    return merge(left, right)
}

func merge(left, right []int) []int {
    result := []int{}
    leftIndex, rightIndex := 0, 0

    for leftIndex < len(left) && rightIndex < len(right) {
        if left[leftIndex] < right[rightIndex] {
            result = append(result, left[leftIndex])
            leftIndex++
        } else {
            result = append(result, right[rightIndex])
            rightIndex++
        }
    }

    result = append(result, left[leftIndex:]...)
    result = append(result, right[rightIndex:]...)

    return result
}

func main() {
    unsortedArray := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
    sortedArray := mergeSort(unsortedArray)
    fmt.Println(sortedArray) // Output: [1 1 2 3 3 4 5 5 5 6 9]
}
