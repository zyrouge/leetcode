package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", minWindow("OUZODYXAZV", "XYZ"), "YXAZ")
	expect("#2", minWindow("xyz", "xyz"), "xyz")
	expect("#3", minWindow("x", "xy"), "")
	expect("#4", minWindow("aaaaaaaaaaaabbbbbcdd", "abcdd"), "abbbbbcdd")
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func minWindow(s string, t string) string {
	sn, tn := len(s), len(t)
	if tn > sn {
		return ""
	}
	result := ""
	missing, matched := make(map[byte]int), 0
	for _, c := range t {
		missing[byte(c)]++
	}
	i := 0
	for j := range s {
		if count, ok := missing[s[j]]; ok {
			missing[s[j]]--
			if count > 0 {
				matched++
			}
		}
		for i < j {
			count, ok := missing[s[i]]
			if ok && count >= 0 {
				break
			}
			if ok {
				missing[s[i]]++
			}
			i++
		}
		if matched == tn && (result == "" || j-i < len(result)) {
			result = s[i : j+1]
		}
	}
	return result
}
