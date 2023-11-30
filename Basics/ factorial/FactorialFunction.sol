// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Factorial {
    function calculateFactorial(uint256 n) public pure returns (uint256) {
        if (n == 0 || n == 1) {
            return 1;
        } else {
            return n * calculateFactorial(n - 1);
        }
    }
}
