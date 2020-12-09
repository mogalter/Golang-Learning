package main

import "fmt"

func main() {
	fmt.Println("Maps are hash tables. They have keys and values")
	m := map[string]int{
		"kevin": 10,
		"meow":  70,
	} // indicates it's a string: int, k:v mapping
	m["james"] = 27
	m["rich"] = 28
	for k, v := range m {
		// will print key + value
		// in the case of range on a list, it iterations index, value at index
		fmt.Println(k, v)
	}
	// can we check for a value?
	//v, ok := m["kevin"]
	// if ok {
	// 	fmt.Println("Kevin has a value of", v)
	// } else {
	// 	fmt.Println("Kevin doesn't exist.")
	// }
	// alternatively... use init;condition format
	if v, ok := m["james"]; ok {
		fmt.Println("James has a value of", v, "--> We are now deleting James")
		// entry deletion: takes <map obj> key
		delete(m, "james")
	}
	fmt.Println("Post-delete of James")
	for k, v := range m {
		fmt.Println(k, v)
	}
}
