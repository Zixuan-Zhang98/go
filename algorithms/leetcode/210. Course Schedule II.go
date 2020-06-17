package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	var sortedOrder []int

	graph := make(map[int][]int, numCourses)
	inDegree := make([]int, numCourses)

	for _, pair := range prerequisites {
		child, parent := pair[0], pair[1]
		// if _, ok := graph[parent]; !ok {
		//     graph[parent] = []int{}
		// }
		graph[parent] = append(graph[parent], child)
		inDegree[child]++
	}

	var sources []int
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			sources = append(sources, i)
		}
	}

	for len(sources) > 0 {
		course := sources[0]
		sortedOrder = append(sortedOrder, course)
		sources = sources[1:]
		for _, child := range graph[course] {
			inDegree[child] -= 1
			if inDegree[child] == 0 {
				sources = append(sources, child)
			}
		}
	}

	if len(sortedOrder) == numCourses {
		return sortedOrder
	} else {
		return []int{}
	}
}
