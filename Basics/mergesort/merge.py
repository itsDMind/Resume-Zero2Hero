def merge_sort(arr):
    if len(arr) <= 1:
        return arr

    middle = len(arr) // 2
    left = arr[:middle]
    right = arr[middle:]

    return merge(merge_sort(left), merge_sort(right))

def merge(left, right):
    result = []
    left_index, right_index = 0, 0

    while left_index < len(left) and right_index < len(right):
        if left[left_index] < right[right_index]:
            result.append(left[left_index])
            left_index += 1
        else:
            result.append(right[right_index])
            right_index += 1

    return result + left[left_index:] + right[right_index:]

# Example usage:
unsorted_array = [3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5]
sorted_array = merge_sort(unsorted_array)
print(sorted_array)  # Output: [1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9]
