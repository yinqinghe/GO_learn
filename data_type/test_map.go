package main

import "fmt"

func main() {
	// Create a map with various data types
	data := map[string]interface{}{
		"name":      "John Doe",
		"age":       25,
		"isMale":    true,
		"interests": []string{"programming", "gaming", "reading"},
		"address": map[string]string{
			"street":  "123 Main St",
			"city":    "New York",
			"country": "USA",
		},
	}

	// Access and print values from the map
	fmt.Println("Name:", data["name"])
	fmt.Println("Age:", data["age"])
	fmt.Println("Is Male:", data["isMale"])
	fmt.Println("Interests:", data["interests"])
	fmt.Println("Address:", data["address"])
}
