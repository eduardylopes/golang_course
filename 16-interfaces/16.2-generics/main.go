package main

import "fmt"

func generic(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	generic("String")
	generic(1)
	generic(false)

	fmt.Println(1, 2, "String", true)

	// Dangerous code, this code accept anything and "interface{}" can be replaced by "any"
	anything := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
		"String":     "String",
		true:         float64(12),
	}

	fmt.Println(anything)
}
