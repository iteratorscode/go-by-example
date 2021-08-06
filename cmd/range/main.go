package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	for index, num := range nums {
		if num == 3 {
			fmt.Println("index:", index)
		}
	}

	kvs := map[string]string{
		"a": "apple",
		"b": "banana",
	}

	for key, value := range kvs {
		fmt.Printf("%s -> %s\n", key, value)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
