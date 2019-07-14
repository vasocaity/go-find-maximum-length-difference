package main

import (
	"fmt"
	"math"
)

func main() {
	var s1 = []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	var s2 = []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
	fmt.Println("= ", MxDifLg(s1, s2))
	s1 = []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
	s2 = []string{"bbbaaayddqbbrrrv"}
	fmt.Println("= ", MxDifLg(s1, s2))
}

func MxDifLg(a1 []string, a2 []string) int {
	max, tmp := 0, 0
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	for i := 0; i < len(a1); i++ {
		for y := 0; y < len(a2); y++ {
			tmp = int(math.Abs(float64(len(a1[i])) - float64(len(a2[y]))))
			if tmp > max {
				max = tmp
			}
		}
	}
	return max
}
