function isPalindrome(s) {
    s = s.replace(/[^a-zA-Z0-9]/g, '').toLowerCase();
    console.log(s)
    return s === s.split('').reverse().join('');
}

// Example usage:
let stringToCheck = "A man, a plan, a canal: Panama";
let result = isPalindrome(stringToCheck);
console.log(result);
