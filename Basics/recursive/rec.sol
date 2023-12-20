pragma solidity ^0.8.0;

contract Fibonacci {
    function fibonacci(uint256 n) public pure returns (uint256) {
        if (n <= 1) {
            return n;
        } else {
            return fibonacci(n - 1) + fibonacci(n - 2);
        }
    }
}
