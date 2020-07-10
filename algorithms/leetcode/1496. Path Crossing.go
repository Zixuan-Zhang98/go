package main

import "fmt"

func isPathCrossing(path string) bool {
	visited := make(map[string]bool)

	x, y := 0, 0
	visited["00"] = true
	for _, direction := range path {
		switch direction {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		case 'W':
			x--
		}
		position := fmt.Sprintf("%d%d", x, y)
		if _, ok := visited[position]; ok {
			return true
		}
		visited[position] = true
	}
	return false
}

func isPathCrossing2(path string) bool {
	set := make(map[[2]int]bool)

	x, y := 0, 0
	set[[2]int{x, y}] = true
	for _, direction := range path {
		switch direction {
		case 'N':
			y++
		case 'S':
			y--
		case 'E':
			x++
		case 'W':
			x--
		}
		if _, ok := set[[2]int{x, y}]; ok {
			return true
		}
		set[[2]int{x, y}] = true
	}
	return false
}
