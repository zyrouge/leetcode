package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", canAttendMeetings([]Interval{{0, 30}, {5, 10}, {15, 20}}), false)
	expect("#2", canAttendMeetings([]Interval{{5, 8}, {9, 15}}), true)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

type Interval struct {
	start int
	end   int
}

func canAttendMeetings(intervals []Interval) bool {
	n := len(intervals)
	if n <= 1 {
		return true
	}
	for i := range n {
		for j := i + 1; j < n; j++ {
			if intervals[j].start < intervals[i].start {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}
	currentEnd := intervals[0].end
	for i := 1; i < n; i++ {
		if intervals[i].start < currentEnd {
			return false
		}
		if intervals[i].end > currentEnd {
			currentEnd = intervals[i].end
		}
	}
	return true
}
