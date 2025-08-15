package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mm := []string{"Miss", "Moneypenny", "I'm 008."}

	jbmm := [][]string{jb, mm}
	fmt.Println(jbmm)

	for i, v := range jbmm {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
		}
	}

	fmt.Println("--------------------")

	fmt.Println(jbmm[0])
	fmt.Println(jbmm[1])

	fmt.Println(len(jbmm))

	for i := 0; i < len(jbmm); i++ {
		fmt.Println(jbmm[i])
	}
}

/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional
slice:

"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "I'm 008."

Range over the records, then range over the data in each record.

fmt.Println(jbmm[0])
	fmt.Println(jbmm[1])

	fmt.Println(len(jbmm))

	for i := 0; i < len(jbmm); i++ {
		fmt.Println(jbmm[i])
	}

*/
