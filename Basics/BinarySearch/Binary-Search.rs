fn binary_search(arr: &[i32], target: i32) -> i32 {
    let mut low = 0;
    let mut high = arr.len() - 1;

    while low <= high {
        let mid = (low + high) / 2;

        if arr[mid] == target {
            return mid as i32;
        } else if arr[mid] < target {
            low = mid + 1;
        } else {
            high = mid - 1;
        }
    }

    -1
}
