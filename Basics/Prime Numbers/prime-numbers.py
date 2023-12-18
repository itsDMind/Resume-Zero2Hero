def is_prime(n):
    if n < 2:
        return False
    for i in range(2, int(n**0.5) + 1):
        if n % i == 0:
            return False
    return True

def generate_primes_up_to_limit(limit):
    return [x for x in range(2, limit + 1) if is_prime(x)]

limit = 50
primes = generate_primes_up_to_limit(limit)
print(f"Prime numbers up to {limit}: {primes}")