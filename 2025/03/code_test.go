package main

import (
	"os"
	"testing"
)

func readExampleData() string {
	bytes, _ := os.ReadFile("input-example.txt")
	return string(bytes)
}

func TestPart1(t *testing.T) {
	want := 357
	part1 := run(false, readExampleData())

	if want != part1 {
		t.Errorf("got %d, wanted %d", part1, want)
	}
}

func TestPart2(t *testing.T) {
	want := 3121910778619
	part2 := run(true, readExampleData())

	if want != part2 {
		t.Errorf("got %d, wanted %d", part2, want)
	}
}
