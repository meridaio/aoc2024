package main

import (
	"log"
	"strconv"
)

const FreeBlockId = -1

type Block struct {
	FileId int
}

type File struct {
	Block
	Space int
}

func checksum(blocks []Block) int {
	sum := 0
	for i, b := range blocks {
		if b.FileId != FreeBlockId {
			sum = sum + (i * b.FileId)
		}
	}

	return sum
}

func getLastNonEmpty(blocks []Block) int {
	for i := len(blocks) - 1; i > 0; i-- {
		if blocks[i].FileId != FreeBlockId {
			return i
		}
	}

	log.Panic("something horrible has happened")
	return -1
}

func sortBlocks(blocks []Block) {
	for i := range blocks {
		if blocks[i].FileId == FreeBlockId {
			lastBlockPos := getLastNonEmpty(blocks)
			if lastBlockPos <= i {
				return
			}

			blocks[i], blocks[lastBlockPos] = blocks[lastBlockPos], blocks[i]
		}
	}
}

func parseBlocks(files []File) []Block {
	blocks := make([]Block, 0)
	for _, f := range files {
		for range f.Space {
			blocks = append(blocks, Block{
				FileId: f.FileId,
			})
		}
	}

	return blocks
}

func parseFiles(line string) []File {
	files := make([]File, 0)
	fileId := 0

	for i, v := range line {
		num, err := strconv.Atoi(string(v))
		if err != nil {
			log.Panic(err)
		}

		isFile := i%2 == 0
		id := FreeBlockId
		if isFile {
			id = fileId
			fileId = fileId + 1
		}

		files = append(files, File{
			Space: num,
			Block: Block{FileId: id},
		})
	}

	return files
}

func getFirstSpaceFor(files []File, capacity int) int {
	for i, v := range files {
		if v.FileId == FreeBlockId && v.Space >= capacity {
			return i
		}
	}

	return -1
}

func sortFiles(files []File) []File {
	sorted := make([]File, len(files))
	copy(sorted, files)

	for i := len(sorted) - 1; i > 0; i-- {
		if sorted[i].FileId != FreeBlockId {
			firstFree := getFirstSpaceFor(sorted, sorted[i].Space)

			if firstFree > 0 && firstFree < i {
				freeSpace := sorted[firstFree]
				diff := freeSpace.Space - sorted[i].Space
				if diff > 0 {
					spaceLeftBehind := File{
						Block: Block{FileId: FreeBlockId},
						Space: sorted[i].Space,
					}
					remainingSpace := File{
						Block: Block{FileId: FreeBlockId},
						Space: diff,
					}
					pre := append(sorted[:firstFree], sorted[i], remainingSpace)
					sorted[i] = spaceLeftBehind
					sorted = append(pre, sorted[firstFree+1:]...)
				} else {
					sorted[firstFree], sorted[i] = sorted[i], sorted[firstFree]
				}
			}
		}
	}
	return sorted
}

func Day9() (int, int) {
	lines := getFileLines("./day9.txt")
	files := parseFiles(lines[0])
	log.Printf("%v", files[:10])
	blocks := parseBlocks(files)
	sortBlocks(blocks)

	sortedFiles := sortFiles(files)
	log.Printf("%v", sortedFiles[:10])
	fileBlocks := parseBlocks(sortedFiles)

	return checksum(blocks), checksum(fileBlocks)
}
