def is_palindrome(s):
    s = ''.join(c.lower() for c in s if c.isalnum())
    return s == s[::-1]

# Example usage:
string_to_check = "A man, a plan, a canal: Panama"
result = is_palindrome(string_to_check)
print(result)
