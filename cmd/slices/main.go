package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	fmt.Println("cap", cap(s))
	fmt.Println("len:", len(s))
	fmt.Println("set:", s)
	s[2] = "c"
	fmt.Println("get:", s[2])
	fmt.Println("cap", cap(s))
	fmt.Println("len:", len(s))

	sp := &s

	fmt.Println("len:", len(*sp))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	fmt.Println("cap", cap(s))
	fmt.Println("len:", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)
	fmt.Printf("sl1 p:%p\n", &l)

	l = s[2:]
	fmt.Println("sl1:", l)
	fmt.Printf("sl1 p:%p\n", &l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < len(twoD); i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d: ", twoD)

}
