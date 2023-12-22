def quicksort(arr):
    if len(arr) <= 1:
        return arr

    pivot = arr[0]
    left = [x for x in arr[1:] if x < pivot]
    right = [x for x in arr[1:] if x >= pivot]

    return quicksort(left) + [pivot] + quicksort(right)

# Example usage:
unsorted_array = [3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5]
sorted_array = quicksort(unsorted_array)
print(sorted_array)  # Output: [1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9]
