fn quicksort(arr: &mut [i32]) {
    if arr.len() <= 1 {
        return;
    }

    let pivot_index = partition(arr);
    quicksort(&mut arr[0..pivot_index]);
    quicksort(&mut arr[pivot_index + 1..]);
}

fn partition(arr: &mut [i32]) -> usize {
    let pivot_index = arr.len() / 2;
    arr.swap(pivot_index, arr.len() - 1);

    let mut i = 0;
    for j in 0..arr.len() - 1 {
        if arr[j] <= arr[arr.len() - 1] {
            arr.swap(i, j);
            i += 1;
        }
    }

    arr.swap(i, arr.len() - 1);
    i
}

// Example usage:
fn main() {
    let mut unsorted_array = vec![3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5];
    quicksort(&mut unsorted_array);
    println!("{:?}", unsorted_array); // Output: [1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9]
}
