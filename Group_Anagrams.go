
func groupAnagrams(strs []string) [][]string {
	sortedmap := make(map[string][]string)
	for _, s := range strs {
		sortedstr := Sortstring(s)
		sortedmap[sortedstr] = append(sortedmap[sortedstr], s)

	}

	var group [][]string

	for _, value := range sortedmap {
		group = append(group, value)
	}
	return group
}

func Sortstring(s string) string { 
	chars := []rune(s)
	for i := 0; i < len(chars)-1; i++ {
		for j := i + 1; j < len(chars); j++ {
			if chars[i] > chars[j] {
				chars[i], chars[j] = chars[j], chars[i]
			}
		}
	}
	return string(chars)
}

