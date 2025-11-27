package main

import "fmt"

func main() {

	nm := NewStringIntMap()

	nm.Add("a", 100)
	nm.Add("b", 1234)
	fmt.Println("myMap: ", nm.data)

	nm.Remove("a")
	fmt.Println("\n remove a: ", nm.data)
	var nmNew = nm.Copy()
	nmNew["c"] = 321
	fmt.Println("\n first: ", nm.data)
	fmt.Println("copy: ", nmNew)

	fmt.Printf("\n is equil b %v", nm.Exists("b"))
	v, ok := nm.Get("b")
	fmt.Printf("\n get b %v %v", v, ok)

}
