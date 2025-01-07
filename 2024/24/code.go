package main

// NOTE: Got stuck really hard on part two. I quickly found the broken gates with
// my own code, but didn't really understand it until I visualized with GraphViz.

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type Gate int8

const (
	AND Gate = iota
	OR
	XOR
)

type Command struct {
	c1, c2, r string
	gate      Gate
}

func extractBinaryString(gates map[string]bool, prefix string) string {
	binaryString := ""
	for z := 0; z < len(gates); z++ {
		key := fmt.Sprintf("%s%02d", prefix, z)

		if value, exists := gates[key]; exists {
			if value {
				binaryString = "1" + binaryString
			} else {
				binaryString = "0" + binaryString
			}
		} else {
			break
		}
	}
	return binaryString
}

func processCommandList(command_list []Command, gates map[string]bool) {
	for len(command_list) > 0 {
		command := command_list[0]
		command_list = command_list[1:]

		// Only check gates that are actually used in the command
		c1, e1 := gates[command.c1]
		c2, e2 := gates[command.c2]

		if e1 && e2 {
			var r bool
			switch command.gate {
			case AND:
				r = c1 && c2
			case OR:
				r = c1 || c2
			case XOR:
				r = c1 != c2
			}
			gates[command.r] = r
		} else {
			// If the command is not ready, add it back to the end of the list to retry later
			command_list = append(command_list, command)
		}
	}
}

func run(part2 bool, input string) any {
	regex_command := regexp.MustCompile(`\w+`)

	gates := map[string]bool{}
	all_gates := map[string]bool{}
	command_list := []Command{}
	for _, line := range strings.Split(input, "\n") {
		if strings.Contains(line, ":") {
			gate := strings.Split(line, ": ")
			gates[gate[0]] = gate[1] == "1"
		} else if len(line) > 0 {
			commands := regex_command.FindAllStringSubmatch(line, -1)

			command := Command{commands[0][0], commands[2][0], commands[3][0], AND}
			if commands[1][0] == "OR" {
				command.gate = OR
			} else if commands[1][0] == "XOR" {
				command.gate = XOR
			}

			all_gates[command.c1] = true
			all_gates[command.c2] = true
			all_gates[command.r] = true
			command_list = append(command_list, command)
		}
	}

	if part2 {
		x, y, z := []string{}, []string{}, []string{}
		and, or, xor := []string{}, []string{}, []string{}

		for gate := range all_gates {
			if gate[0] == 'x' {
				x = append(x, gate)
			} else if gate[0] == 'y' {
				y = append(y, gate)
			} else if gate[0] == 'z' {
				z = append(z, gate)
			}
		}

		for _, command := range command_list {
			if command.gate == AND {
				and = append(and, command.r)
			} else if command.gate == OR {
				or = append(or, command.r)
			} else if command.gate == XOR {
				xor = append(xor, command.r)
			}
		}

		slices.Sort(x)
		slices.Sort(y)
		slices.Sort(z)

		// GraphViz output
		fmt.Printf(`
digraph G {
	subgraph {
		node [style=filled,color=lightgrey]
		%s
	}
	subgraph {
		node [style=filled,color=lightgrey]
		%s
	}
	subgraph {
		%s
	}
	subgraph {
		node [style=filled,color=lightgreen]
		%s
	}
	subgraph {
		node [style=filled,color=yellow]
		%s
	}
	subgraph {
		node [style=filled,color=lightskyblue]
		%s
	}
`,
			strings.Join(x, "->"),
			strings.Join(y, "->"),
			strings.Join(z, "->"),
			strings.Join(and, " "),
			strings.Join(or, " "),
			strings.Join(xor, " "),
		)

		for _, command := range command_list {
			fmt.Println("\t", command.c1, "->", command.r, ";", command.c2, "->", command.r, ";")
		}
		fmt.Println("}")

		// setGateValues := func(x_value, y_value int) map[string]bool {
		// 	gates := map[string]bool{}
		// 	x_string := reverseString(fmt.Sprintf("%045b", x_value))
		// 	y_string := reverseString(fmt.Sprintf("%045b", y_value))
		// 	for i := 0; i < 45; i++ {
		// 		x, y := fmt.Sprintf("x%02d", i), fmt.Sprintf("y%02d", i)
		// 		gates[x] = x_string[i] == '1'
		// 		gates[y] = y_string[i] == '1'
		// 	}
		// 	return gates
		// }

		// // Leaving this code to find the broken gates
		// for i, y := 0, 1; y <= 17592186044416; i, y = i+1, y*2 {
		// 	gates_copy := setGateValues(0, y)

		// 	processCommandList(command_list, gates_copy)
		// 	b := fmt.Sprintf("%046b", 0+y)
		// 	if extractBinaryString(gates_copy, "z") != b {
		// 		fmt.Println("Wrong in position", i, fmt.Sprintf("z%02d", i))
		// 	}
		// }

		// Looking through Graphviz via the above code we can consider the following:
		broken_gates := map[string]bool{}
		for _, command := range command_list {
			c1_0, c2_0, r_0 := command.c1[0], command.c2[0], command.r[0]

			// zvalue should always be an XOR value
			if r_0 == 'z' && command.gate != XOR && command.r != "z45" {
				broken_gates[command.r] = true
			}

			// The XOR gates should always have x and y as input and z as output
			if command.gate == XOR && r_0 != 'z' &&
				!(c1_0 == 'x' || c1_0 == 'y') && !(c2_0 == 'x' || c2_0 == 'y') {
				broken_gates[command.r] = true
			}

			// The AND gates should always give their result to an OR gate
			if command.gate == AND && command.c1 != "x00" && command.c2 != "x00" {
				for _, command2 := range command_list {
					if (command.r == command2.c1 || command.r == command2.c2) && command2.gate != OR {
						broken_gates[command.r] = true
					}
				}
			}

			// The XOR gates should never give result to an OR gate
			if command.gate == XOR {
				for _, command2 := range command_list {
					if (command.r == command2.c1 || command.r == command2.c2) && command2.gate == OR {
						broken_gates[command.r] = true
					}
				}
			}
		}

		result := []string{}
		for gate := range broken_gates {
			result = append(result, gate)
		}

		slices.Sort(result)
		return strings.Join(result, ",")
	} else {
		processCommandList(command_list, gates)

		result, _ := strconv.ParseInt(extractBinaryString(gates, "z"), 2, 64)
		return result
	}
}
