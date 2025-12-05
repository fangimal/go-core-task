package main

func intEntrySlice(a, b []int) (bool, []int) {
	set := make(map[int]bool, len(b))
	for _, v := range b {
		set[v] = true
	}

	result := make([]int, 0, min(len(a), len(b)))
	for _, v := range a {
		if set[v] {
			result = append(result, v)
		}
	}

	return len(result) > 0, result
}
