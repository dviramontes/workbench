package workbench

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0] // compare against first word
	for _, str := range strs {
		for i := 0; i < len(prefix); i++ {
			if i == len(str) || prefix[i] != str[i] {
				prefix = prefix[:i]
				break
			}
		}
	}

	return prefix
}
