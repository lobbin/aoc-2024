package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// Get value based on combo operand
func getValue(registers map[string]int, operand int) int {
	if operand >= 0 && operand <= 3 {
		return operand
	} else if operand == 4 {
		return registers["A"]
	} else if operand == 5 {
		return registers["B"]
	} else if operand == 6 {
		return registers["C"]
	} else {
		return -1
	}
}

func calcB(a int) int {
	b := (a % 8) ^ 1
	return (b ^ (a >> b)) ^ 4
}

func run(part2 bool, input string) any {
	registers := make(map[string]int)
	outputs := make([]string, 0)

	// Loop the file
	for _, line := range strings.Split(input, "\n") {
		// Take care of the different registers
		if strings.Contains(line, "Register") {
			register_string := strings.Split(line, " ")
			register, _ := strconv.Atoi(strings.TrimSpace(register_string[2]))
			registers[string(register_string[1][0])] = register
		} else if strings.Contains(line, "Program") {
			// Start the program
			opcodes := strings.Split(strings.Split(line, " ")[1], ",")

			if part2 {
				start, _ := strconv.Atoi(opcodes[len(opcodes)-1])
				possibles := make([]int, 0)

				// Find the last magic number based on modulo 8
				for a := 1; a < 8; a++ {
					// This is a short representation of my opcodes
					b := calcB(a)
					if b%8 == start {
						possibles = append(possibles, a)
						break
					}
				}

				// We loop the result, calculating backwards
				for i := len(opcodes) - 2; i >= 0; i-- {
					// This is what we're trying to find
					b, _ := strconv.Atoi(opcodes[i])
					next_possibles := make([]int, 0)

					for _, a := range possibles {
						// For each possible range in mod 8 we check the next step
						for j := 0; j < 8; j++ {
							potential_a := a*8 + j // *8 is the last entry in my loop (0,3)
							potential_b := calcB(potential_a)

							if potential_b%8 == b {
								next_possibles = append(next_possibles, potential_a)
							}
						}
					}

					possibles = next_possibles
				}

				// Doesn't work on the example
				if len(possibles) > 0 {
					// Pick the lowest value and return
					slices.Sort(possibles)
					return possibles[0]
				}
			} else {
				// Loop and handle each opcode
				for i := 0; i < len(opcodes); i += 2 {
					opcode := opcodes[i]
					combo, _ := strconv.Atoi(opcodes[i+1])
					operand_value := getValue(registers, combo)

					switch opcode {
					case "0": // adv
						registers["A"] = int(float64(registers["A"]) / math.Pow(2, float64(operand_value)))
					case "1": // bxl
						registers["B"] = registers["B"] ^ combo
					case "2": // bst
						registers["B"] = operand_value % 8
					case "3": // jnz
						if registers["A"] != 0 {
							i = combo - 2
						}
					case "4": // bxc
						registers["B"] = registers["B"] ^ registers["C"]
					case "5": // out
						string_values := strings.Split(strconv.Itoa(operand_value%8), "")
						outputs = append(outputs, string_values...)
					case "6": // bdv
						registers["B"] = int(float64(registers["A"]) / math.Pow(2, float64(operand_value)))
					case "7": // cdv
						registers["C"] = int(float64(registers["A"]) / math.Pow(2, float64(operand_value)))
					default:
						fmt.Println("Unknown opcode", opcode)
					}
				}
			}
		}
	}

	return strings.Join(outputs, ",")
}
