package main

import "fmt"

func main() {
	fmt.Println(StringShortener("aab"))       // output: b
	fmt.Println(StringShortener("aaabba"))    // output:
	fmt.Println(StringShortener("aabb"))      // output:
	fmt.Println(StringShortener("aaa"))       // output: a
	fmt.Println(StringShortener("aaabccddd")) // output: abd
}

func StringShortener(s string) string {
	res := make([]byte, len(s))

	for i := range s {
		if res[len(res)-1] == s[i] {
			res = res[:len(res)-1]
			continue
		}
		res = append(res, s[i])
	}

	return string(res)
}
