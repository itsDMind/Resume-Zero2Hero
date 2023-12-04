fn is_palindrome(s: &str) -> bool {
    let s: String = s.chars().filter(|c| c.is_alphanumeric()).map(|c| c.to_ascii_lowercase()).collect();
    let rev_s: String = s.chars().rev().collect();
    s == rev_s
}

fn main() {
    let string_to_check = "A man, a plan, a canal: Panama";
    let result = is_palindrome(string_to_check);
    println!("{}", result);
}
