package workbench

func IsPalindrome(s string) bool {
	left := 0           // start at the beginning of the string
	right := len(s) - 1 // and at the end
	// run until the pointers meet each other
	for left < right {
		if s[left] != s[right] {
			return false
		}
		// otherwise move the pointers towards each other
		left++
		right--
	} // end of the loop is the end of the palindrome, that means we have a palindrome
	return true
}
