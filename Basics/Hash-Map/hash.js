class HashMap {
    constructor() {
      this.map = {};
    }
  
    insert(key, value) {
      this.map[key] = value;
    }
  
    retrieve(key) {
      return this.map[key];
    }
  }
  
  // Example usage:
  const myHashMap = new HashMap();
  myHashMap.insert("key1", "value1");
  myHashMap.insert("key2", "value2");
  
  console.log(myHashMap.retrieve("key1")); // Output: value1
  console.log(myHashMap.retrieve("key2")); // Output: value2
  