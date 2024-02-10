package strings

func IntToRoman(num int) string {
	lst := make([]Node, 0)
	var res string
	lst = append(lst, Node{Str: "I", Int: 1}, Node{Str: "IV", Int: 4}, Node{Str: "V", Int: 5}, Node{Str: "IX", Int: 9},
		Node{Str: "X", Int: 10}, Node{Str: "XL", Int: 40}, Node{Str: "L", Int: 50}, Node{Str: "XC", Int: 90}, Node{Str: "C", Int: 100},
		Node{Str: "CD", Int: 400}, Node{Str: "D", Int: 500}, Node{Str: "CM", Int: 900}, Node{Str: "M", Int: 1000})

	for i := len(lst) - 1; i >= 0; i-- {
		if num/lst[i].Int > 0 {
			count := num / lst[i].Int
			for j := 1; j <= count; j++ {
				res += lst[i].Str
			}
		}
		num = num % lst[i].Int
	}

	return res
}

type Node struct {
	Str string
	Int int
}
