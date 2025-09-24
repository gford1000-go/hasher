package hasher

import "fmt"

func Example() {

	type Inner struct {
		B []int
	}
	type Outer struct {
		A int
		B *Inner
		C string
	}

	v := &Outer{
		A: 42,
		B: &Inner{B: []int{101, 102, 103}},
		C: "Hello",
	}

	h, _ := Hash(v) // Illustrates hashing arbitrary objects
	fmt.Println(h)
	// Output: [121 166 34 156 67 166 135 131 184 8 158 234 134 25 173 164 219 114 142 83 69 18 62 75 40 13 148 54 234 209 191 132]
}
