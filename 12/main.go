package main

import "fmt"

func main() {
	stringArr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(setOfStrings(stringArr))
}

func setOfStrings(stringArr []string) []string {
	set := make(map[string]struct{})
	for _, s := range stringArr {
		if _, ok := set[s]; !ok {
			set[s] = struct{}{}
		}
	}
	result := make([]string, 0)
	for s, _ := range set {
		result = append(result, s)
	}
	return result
}
