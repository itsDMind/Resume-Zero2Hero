function isPrime(n) {
    if (n < 2) {
        return false;
    }
    for (let i = 2; i <= Math.sqrt(n); i++) {
        if (n % i === 0) {
            return false;
        }
    }
    return true;
}

function generatePrimesUpToLimit(limit) {
    const primes = [];
    for (let i = 2; i <= limit; i++) {
        if (isPrime(i)) {
            primes.push(i);
        }
    }
    return primes;
}

const limit = 50;
const primes = generatePrimesUpToLimit(limit);
console.log(`Prime numbers up to ${limit}: ${primes}`);
