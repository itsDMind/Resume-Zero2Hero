pragma solidity ^0.8.0;

contract PalindromeChecker {
    function isPalindrome(string memory s) public pure returns (bool) {
        bytes memory b = bytes(s);
        for (uint i = 0; i < b.length / 2; i++) {
            if (b[i] != b[b.length - 1 - i]) {
                return false;
            }
        }
        return true;
    }
}
