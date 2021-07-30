package main

import "fmt"

func main() {

	// for i := 1; i <= 100; i++ {
	// 	if i%3 == 0 && i%5 == 0 {
	// 		fmt.Println("FizzBuzz")
	// 		continue
	// 	}
	// 	if i%3 == 0 {
	// 		fmt.Println("Fizz")
	// 		continue
	// 	}
	// 	if i%5 == 0 {
	// 		fmt.Println("Buzz")
	// 		continue
	// 	}

	// 	fmt.Println("ULTIMO::: ", i)
	// }

	// mySlice()
	iteratingOverStrings()

}

func mySlice() {

	evenVals := []int{2, 4, 6, 7, 10, 12}

	for _, v := range evenVals {
		fmt.Println(v)
	}
}

func uniqueNames() {
	names := map[string]bool{"Fred": true, "Raul": true, "Douglas": false}

	//k: retorna o nome da chave
	//v: retorna o valor da chave.

	for k, v := range names {
		fmt.Println(k, v)
	}
}

func iteratingOverMap() {
	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
}

func iteratingOverStrings() {
	samples := []string{"hello", "apple"}

	//comportamento especial para o for;
	// no geral, strings são feitas de bytes.
	// no for-range, as strings são runas.
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}

		fmt.Println()
	}
}

func mySetFunctionToMaps() {
	//garantir o único valor da chave
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(intSet)
}
