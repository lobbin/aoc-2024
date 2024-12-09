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
	want := 1928
	part1 := run(false, readExampleData())

	if want != part1 {
		t.Errorf("got %q, wanted %q", part1, want)
	}
}

func TestPart2(t *testing.T) {
	want := 2858
	part2 := run(true, readExampleData())

	if want != part2 {
		t.Errorf("got %q, wanted %q", part2, want)
	}
}
