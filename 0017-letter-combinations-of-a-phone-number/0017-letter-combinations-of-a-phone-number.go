func letterCombinations(digits string) []string {
  if digits == "" {
    return []string{}
  }
  return letterCombinationsR(digits, "")
}

var numToLetters = map[string]string {
  "2": "abc",
  "3": "def",
  "4": "ghi",
  "5": "jkl",
  "6": "mno",
  "7": "pqrs",
  "8": "tuv",
  "9": "wxyz",
}

func letterCombinationsR(digits, lettersSoFar string) []string {
  if digits == "" {
    return []string{lettersSoFar}
  }

  digit := string(digits[0])
  digits = digits[1:]
  letters := numToLetters[digit]
  combos := []string{}
  for _, r := range letters {
    rCombos := letterCombinationsR(digits, lettersSoFar + string(r))
    combos = append(combos, rCombos...)
  }
  
  return combos
}