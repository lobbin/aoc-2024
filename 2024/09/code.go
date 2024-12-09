package main

import (
	"strconv"

	"github.com/jpillora/puzzler/harness/aoc"
)

type BlockType uint8

const (
	File BlockType = iota
	Free
)

type Block struct {
	blockType BlockType
	id        int
}

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	files := 0
	blocks := make([]Block, 0)

	// Create a list of blocks and of what kind they are
	for i, b := range input {
		size, _ := strconv.Atoi(string(b))

		var block BlockType
		if i%2 == 0 {
			block = File
			files++
		} else {
			block = Free
		}

		for j := 0; j < size; j++ {
			blocks = append(blocks, Block{block, files - 1})
		}
	}

	if part2 {
		// For part 2 we only move entire block of file if we can fit it in an empty
		// space. We start by going backwards from the highest file id.
		search := files - 1
		file_blocks := make([]int, 0)
		free_blocks := make([]int, 0)

		// Loop all blocks from the right
		for i := len(blocks) - 1; i > 0; i-- {
			if blocks[i].blockType == File && blocks[i].id == search {
				// If the block is a file and belongs to our file, we add the block to
				// the last
				file_blocks = append(file_blocks, i)
			} else {
				// If the block is empty, or another file, we start to process. We can
				// also end up here in case for empty blocks.
				if len(file_blocks) > 0 {
					// Start from left, trying to a find an empty block big enough to fit
					// our file
					for j := 0; j <= i; j++ {
						if blocks[j].blockType == Free {
							free_blocks = append(free_blocks, j)
							if len(free_blocks) == len(file_blocks) {
								// When we have enough free blocks, we simply swap them out with
								// our known file blocks
								for k := 0; k < len(file_blocks); k++ {
									blocks[file_blocks[k]], blocks[free_blocks[k]] = blocks[free_blocks[k]], blocks[file_blocks[k]]
								}
								break
							}
						} else {
							// If the block is not empty, we reset the list of empty blocks
							// since we don't want to spread them again (like we did in part1)
							free_blocks = nil
						}
					}

					// Reset everything for the loop. Decreasing the file id and increasing
					// the loop index since i is currently not on the same file block we
					// processed but one left.
					free_blocks = nil
					file_blocks = nil
					search--
					i++
				}
			}
		}
	} else {
		// For part 1 we move all blocks from the right to the last, no matter that
		// they belong to the same file or not.
		start := 0
		for i := len(blocks) - 1; i > (len(blocks) / 2); i-- {
			if blocks[i].blockType == File {
				for j := start; j < i; j++ {
					// For the first free block, we simply switch. We also store the
					// position of the last known empty block for fast access next time
					if blocks[j].blockType == Free {
						start = j
						// Swap values
						blocks[i], blocks[j] = blocks[j], blocks[i]
						break
					}
				}
			}
		}
	}

	// Calculate the checksum
	checksum := 0
	for i, block := range blocks {
		if block.blockType == File {
			checksum += i * block.id
		}
	}

	return checksum
}
