package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	sum := 0

	fmt.Println("nums:", nums)

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	keystore := map[string]string{"key1": "value1", "key2": "value2", "key3": "value3"}

	for key, value := range keystore {
		fmt.Println("key:", key, "value:", value)
	}

	for key := range keystore {
		fmt.Println("key:", key)
	}

	str := "golang"

	for index, char := range str {
		fmt.Println("index:", index, "char:", char)
	}
}
