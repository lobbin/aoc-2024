package main

import (
	"os"
	"testing"
)

func readExampleData(e string) string {
	bytes, _ := os.ReadFile("input-example" + e + ".txt")
	return string(bytes)
}

func TestPart1(t *testing.T) {
	want := 161
	part1 := run(false, readExampleData(""))

	if want != part1 {
		t.Errorf("got %q, wanted %q", part1, want)
	}
}

func TestPart2(t *testing.T) {
	want := 48
	part2 := run(true, readExampleData("2"))

	if want != part2 {
		t.Errorf("got %q, wanted %q", part2, want)
	}
}
