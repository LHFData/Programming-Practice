package main

import "fmt"

func main() {
	grades := make(map[string]float32)
	grades["LHF"] = 42
	grades["WY"] = 92
	grades["LYL"] = 67

	fmt.Println(grades)

	Lhf_Grade := grades["LHF"]
	fmt.Println(Lhf_Grade)

	delete(grades, "LHF")
	fmt.Println(grades)

	for k, v := range grades {
		fmt.Println(k, "'s score is ", v)
	}

}
