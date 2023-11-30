fn factorial(n: u64) -> u64 {
    if n == 0 || n == 1 {
        1
    } else {
        n * factorial(n - 1)
    }
}

// Example usage:
fn main() {
    let result = factorial(5);
    println!("{}", result);
}
