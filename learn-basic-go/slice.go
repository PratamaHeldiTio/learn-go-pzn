package main

import "fmt"

func main() {
	var mont = [3]string{
		"jan",
		"feb",
		"mar",
	}

	var m = mont[0:2]

	fmt.Println(m)
	copySlice := make([]string, len(m), cap(m))
	copy(copySlice, m)
	fmt.Println(copySlice)
	copySlice[0] = "jii"
	copySlice = append(copySlice, "1")
	// fmt.Println(copySlice)
	copySlice = append(copySlice, "2")
	// fmt.Println(copySlice)
	copySlice = append(copySlice, "3")
	// fmt.Println(copySlice)
	copySlice = append(copySlice, "4")
	// fmt.Println(copySlice)
	copySlice = append(copySlice, "5")
	fmt.Println(copySlice)

	iniMap := make(map[int]string)
	iniMap[8] = "heldi"
	fmt.Println(iniMap[8])

	infit := []string{}
	fmt.Println(infit)
	fmt.Println(len(infit))
	fmt.Println(cap(infit))

	infit = append(infit, "1")
	fmt.Println(infit)
	fmt.Println(len(infit))
	fmt.Println(cap(infit))

	infit = append(infit, "1")
	fmt.Println(infit)
	fmt.Println(len(infit))
	fmt.Println(cap(infit))

	infit = append(infit, "1")
	fmt.Println(infit)
	fmt.Println(len(infit))
	fmt.Println(cap(infit))

	infit = append(infit, "1")
	fmt.Println(infit)
	fmt.Println(len(infit))
	fmt.Println(cap(infit))

	infit = append(infit, "1")
	fmt.Println(infit)
	fmt.Println(len(infit))
	fmt.Println(cap(infit))

}
