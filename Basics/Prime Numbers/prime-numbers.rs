fn generate_primes_up_to(limit: usize) -> Vec<usize> {
    let mut primes = Vec::new();
    let mut is_prime = vec![true; limit + 1];

    for num in 2..=limit {
        if is_prime[num] {
            primes.push(num);

            let mut multiple = 2 * num;
            while multiple <= limit {
                is_prime[multiple] = false;
                multiple += num;
            }
        }
    }

    primes
}

fn main() {
    let limit = 30; // Set your desired limit here
    let primes = generate_primes_up_to(limit);

    println!("Prime numbers up to {}: {:?}", limit, primes);
}