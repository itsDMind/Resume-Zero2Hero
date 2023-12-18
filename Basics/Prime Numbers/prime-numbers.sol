// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract PrimeNumbers {
    function isPrime(uint256 n) internal pure returns (bool) {
        if (n < 2) {
            return false;
        }
        for (uint256 i = 2; i <= n / 2; i++) {
            if (n % i == 0) {
                return false;
            }
        }
        return true;
    }

    function generatePrimesUpToLimit(uint256 limit) external pure returns (uint256[] memory) {
        uint256[] memory primes;
        uint256 count = 0;
        for (uint256 i = 2; i <= limit; i++) {
            if (isPrime(i)) {
                assembly {
                    // Resize the dynamic array if needed
                    if lt(count, mload(primes)) {
                        mstore(primes, add(mload(primes), 1))
                    }
                    // Store the prime number in the array
                    sstore(add(primes, count), i)
                }
                count++;
            }
        }
        return primes;
    }
}
