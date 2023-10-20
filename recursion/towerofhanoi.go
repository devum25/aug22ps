package recursion

import "fmt"

func TOH(n int, src, dest, temp string) {
	if n == 0 {
		return
	}

	TOH(n-1, src, temp, dest)
	fmt.Printf("%s -> %s", src, dest)
	fmt.Println()
	TOH(n-1, temp, dest, src)
}
