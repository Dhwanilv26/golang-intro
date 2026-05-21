package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "123",
		"blue":  "234",
		"white": "567",
	}
	printMap(colors)

}
func printMap(mp map[string]string) {

	for key, value := range mp {
		fmt.Println(key, value)
	}
}
