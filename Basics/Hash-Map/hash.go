package main

import "fmt"

type HashMap map[string]string

func (h HashMap) insert(key, value string) {
	h[key] = value
}

func (h HashMap) retrieve(key string) string {
	return h[key]
}

func main() {
	myHashMap := make(HashMap)
	myHashMap.insert("key1", "value1")
	myHashMap.insert("key2", "value2")

	fmt.Println(myHashMap.retrieve("key1")) // Output: value1
	fmt.Println(myHashMap.retrieve("key2")) // Output: value2
}
