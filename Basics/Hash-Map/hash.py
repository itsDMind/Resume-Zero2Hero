class HashMap:
    def __init__(self):
        self.map = {}

    def insert(self, key, value):
        self.map[key] = value

    def retrieve(self, key):
        return self.map.get(key)

# Example usage:
my_hash_map = HashMap()
my_hash_map.insert("key1", "value1")
my_hash_map.insert("key2", "value2")

print(my_hash_map.retrieve("key1"))  # Output: value1
print(my_hash_map.retrieve("key2"))  # Output: value2
