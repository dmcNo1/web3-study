package main

import "fmt"

func main() {
	testMap := make(map[string]string)
	testMap["q"] = "天音波/回音击"

	testMap["q"] = "精准礼仪"

	delete(testMap, "q")

	if skill, ok := testMap["q"]; ok {
		fmt.Printf("skill-q: %v\n", skill)
	}
}
