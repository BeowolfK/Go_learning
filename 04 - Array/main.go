package main

import "fmt"

func main() {
	// Arrays
	grade := [3]int{15, 17, 18} // sized and typed array
	fmt.Printf("Grades: %v\n",grade)

	grade2 := [...]int{15, 17, 16, 19} // dynamically sized and typed array
	// grade2[4] = 20 // error => out of range
	fmt.Printf("Grades2: %v\n",grade2)
		

	grade3 := [3][2]int{
		{15, 17},
		{16, 19},
		{14, 16}} // 2D array
	fmt.Printf("Grades3: %v\n",grade3)

	arrayA := []int{1, 2, 3, 4, 5} 

	// Slices
	fmt.Println(arrayA[:]) // return entire array
	fmt.Println(arrayA[1:3]) // return index 1 to 2
	fmt.Println(arrayA[1:]) // return from index 1 to end
	fmt.Println(arrayA[:3]) // return up to index 2

	array := []int{} // un-sized and typed array
	fmt.Printf("%v => Length : %v, Capacity : %v\n", array, len(array), cap(array))
	array = append(array, 1)
	fmt.Printf("%v => Length : %v, Capacity : %v\n", array, len(array), cap(array))
	array = append(array, 2, 3, 4, 5)
	fmt.Printf("%v => Length : %v, Capacity : %v\n", array, len(array), cap(array))

	array2 := make([]int, 0, 5) // un-sized and typed array but with max capacity
	// make([]int, len(array), cap(array)) 
	fmt.Printf("%v => Length : %v, Capacity : %v\n", array2, len(array2), cap(array2))
	array2 = append(array2, 2)
	fmt.Printf("%v => Length : %v, Capacity : %v\n", array2, len(array2), cap(array2))
	array2 = append(array2, 2, 3, 4, 5)
	fmt.Printf("%v => Length : %v, Capacity : %v\n", array2, len(array2), cap(array2))
	array2 = append(array2, 2, 3, 4, 5)

	// Maps
	population := map[string]int{
		"Chine": 1.411e9,
		"Inde" : 1.372e9,
		"USA"  :  3.331e8,
	}
	fmt.Println(population)
	value, existence := population["Chine"] // Return value and boolean if exists
	fmt.Println(value, existence)
	value2, existence2 := population["Japon"]
	fmt.Println(value2, existence2)
	population["France"] = 67.02e6
	delete(population, "USA")
	fmt.Println(population)

}