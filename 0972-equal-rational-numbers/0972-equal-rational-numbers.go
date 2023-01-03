func isRationalEqual(s string, t string) bool {
	if !strings.Contains(s, "(") && !strings.Contains(t, "(") {
		// RepeatingPart not exist in both s and t
		if strings.Contains(s, ".") {
			// remove invalid '0'
			s = strings.TrimRight(s, "0")
		}
		if strings.Contains(t, ".") {
			// remove invalid '0'
			t = strings.TrimRight(t, "0")
		}
		return strings.EqualFold(strings.TrimSuffix(s, "."), strings.TrimSuffix(t, "."))
	}
	if strings.Contains(s, "(") {
		// expand RepeatingPart, X.X(00) should be X.X
		repeatingPart := strings.TrimSuffix(strings.Split(s, "(")[1], ")")
		if strings.Trim(repeatingPart, "0") == "" {
			repeatingPart = ""
		}
		s = strings.Split(s, "(")[0]
		for repeatingPart != "" && len(s) < 32 {
			s += repeatingPart
		}
	}
	if strings.Contains(t, "(") {
		// expand RepeatingPart, X.X(00) should be X.X
		repeatingPart := strings.TrimSuffix(strings.Split(t, "(")[1], ")")
		if strings.Trim(repeatingPart, "0") == "" {
			repeatingPart = ""
		}
		t = strings.Split(t, "(")[0]
		for repeatingPart != "" && len(t) < 32 {
			t += repeatingPart
		}
	}
	if len(s) >= 32 && len(t) >= 32 {
		return s[:30] == t[:30]
	}
	if len(s) >= 32 {
		s = s[:13]
	}
	if len(t) >= 32 {
		t = t[:13]
	}
	s = strings.TrimSuffix(s, ".")
	t = strings.TrimSuffix(t, ".")
	sf, _ :=  strconv.ParseFloat(s, 64)
	tf, _ :=  strconv.ParseFloat(t, 64)
	return math.Abs(sf-tf) < 0.00000001
}