package problems

func BugfixLeaderSorted(a []int) int {

	si := len(a)
	count := si / 2

	tmp := make(map[int]int)
	for _, v := range a {
		if val, ok := tmp[v]; !ok {
			tmp[v] = 1
		} else {
			tmp[v] = val + 1
		}
	}

	var max int
	for _, v := range tmp {
		if max < v {
			max = v
		}
	}

	if max > count {
		return 1
	}

	return -1
}
