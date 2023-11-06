package darraygo

import "fmt"

func TopLeftAboveDiagonal(x [][]int) {

	for i := 0; i < len(x); i++ {
		for j := i + 1; j < len(x); j++ {
			fmt.Print(x[i][j])
		}
	}

}
