package main

import (
	"fmt"
	"os"
)

func main() {
	{
		timeMap := Constructor()
		timeMap.Set("alice", "happy", 1)
		expect("#1 > #1", timeMap.Get("alice", 1), "happy")
		expect("#1 > #2", timeMap.Get("alice", 2), "happy")
		timeMap.Set("alice", "sad", 3)
		expect("#1 > #3", timeMap.Get("alice", 3), "sad")
	}
	{
		timeMap := Constructor()
		timeMap.Set("key1", "value1", 10)
		expect("#2 > #1", timeMap.Get("key1", 1), "")
		expect("#2 > #2", timeMap.Get("key1", 10), "value1")
		expect("#2 > #3", timeMap.Get("key1", 11), "value1")
	}
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

type TimeMap struct {
	Entries map[string][]TimeMapEntryValue
}

type TimeMapEntryValue struct {
	Value     string
	Timestamp int
}

func Constructor() TimeMap {
	return TimeMap{make(map[string][]TimeMapEntryValue)}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if this.Entries[key] == nil {
		this.Entries[key] = make([]TimeMapEntryValue, 0)
	}
	this.Entries[key] = append(this.Entries[key], TimeMapEntryValue{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	values := this.Entries[key]
	n := len(values)
	if values == nil || n == 0 {
		return ""
	}
	i, j := 0, n-1
	for i <= j {
		mid := (i + j) / 2
		if values[mid].Timestamp <= timestamp {
			if mid == n-1 || values[mid+1].Timestamp > timestamp {
				return values[mid].Value
			}
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return ""
}
