package main

import "fmt"

func main() {
	m := make(map[string]int)

	fmt.Println("emp:", m)

	m["key1"] = 1
	m["key2"] = 2
	m["key3"] = 3

	fmt.Println("map:", m)
	fmt.Println("len:", len(m))

	v1 := m["key1"]
	fmt.Println("v1: ", v1)

	delete(m, "key2")

	fmt.Println("after delete key 2", m)
	fmt.Println("len:", len(m))
	fmt.Println("missing value:", m["key2"])

	_, prs := m["key2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
