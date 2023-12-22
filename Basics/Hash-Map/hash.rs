use std::collections::HashMap;

fn main() {
    let mut my_hash_map = HashMap::new();
    
    my_hash_map.insert("key1", "value1");
    my_hash_map.insert("key2", "value2");

    println!("{:?}", my_hash_map.get("key1")); // Output: Some("value1")
    println!("{:?}", my_hash_map.get("key2")); // Output: Some("value2")
}
