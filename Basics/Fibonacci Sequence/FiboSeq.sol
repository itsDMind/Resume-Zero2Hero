// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Fibonacci {
    function fibonacci(uint n) public pure returns (uint) {
        if (n <= 1) {
            return n;
        }
        return fibonacci(n - 1) + fibonacci(n - 2);
    }
}

// Example usage:
// Deploy the contract and call the `fibonacci` function with the desired parameter.
