fn merge_sort(arr: &mut [i32]) {
    let len = arr.len();

    if len <= 1 {
        return;
    }

    let mid = len / 2;
    let mut left = arr[0..mid].to_vec();
    let mut right = arr[mid..].to_vec();

    merge_sort(&mut left);
    merge_sort(&mut right);

    let (mut i, mut j, mut k) = (0, 0, 0);

    while i < left.len() && j < right.len() {
        if left[i] < right[j] {
            arr[k] = left[i];
            i += 1;
        } else {
            arr[k] = right[j];
            j += 1;
        }
        k += 1;
    }

    while i < left.len() {
        arr[k] = left[i];
        i += 1;
        k += 1;
    }

    while j < right.len() {
        arr[k] = right[j];
        j += 1;
        k += 1;
    }
}

// Example usage:
fn main() {
    let mut unsorted_array = vec![3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5];
    merge_sort(&mut unsorted_array);
    println!("{:?}", unsorted_array); // Output: [1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9]
}
