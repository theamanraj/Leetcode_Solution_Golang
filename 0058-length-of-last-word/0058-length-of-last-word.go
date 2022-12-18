import ("strings")

func lengthOfLastWord(s string) int {
	if len(s) < 1 {
		return 0;
	}

    ns := strings.Trim(s, " ")
    sp := strings.Split(ns, " ")
    return len(sp[len(sp) - 1])
}