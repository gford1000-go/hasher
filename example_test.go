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

	h, _ := Hash(v)
	fmt.Println(h)
	// Output: [74 175 82 242 158 82 154 251 165 204 11 169 52 189 125 83 245 147 182 197 53 143 167 187 38 128 108 203 210 240 5 226]
}
