import (
	"bytes"
	"sort"
	"strconv"
)

type (
	stack struct {
		top    *node
		length int
	}
	node struct {
		val rune
		pre *node
	}
)

func newStack() *stack {
	return &stack{}
}

func (s *stack) push(v rune) {
	s.top = &node{val: v, pre: s.top}
	s.length++
}

func (s *stack) peek() rune {
	if !s.isEmpty() {
		return s.top.val
	}
	return '\x00'
}

func (s *stack) pop() rune {
	if !s.isEmpty() {
		t := s.top
		s.top = t.pre
		s.length--
		return t.val
	}
	return '\x00'
}

func (s *stack) size() int {
	return s.length
}
func (s *stack) isEmpty() bool {
	return s.length == 0
}

func isNum(r rune) bool {
	return r >= '0' && r <= '9'
}

func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func prepend(s []rune, r rune) []rune {
	return append([]rune{r}, s...)
}

func countOfAtomsHelper(s *stack, dict map[string]int) {
	//numbers
	multiplier := 0
	//record the digits of number
	order := 1
	//current rune
	var pivot rune
	//record atoms
	atoms := make([]rune, 0, 1)
	for !s.isEmpty() {
		//get current character
		pivot = s.pop()
		if isNum(pivot) {
			//accumulate the number
			multiplier += int(pivot-'0') * order
			order *= 10
		} else if isLower(pivot) {
			atoms = prepend(atoms, pivot)
		} else if isUpper(pivot) {
			//if no num, then it should be 1
			if multiplier == 0 {
				multiplier = 1
			}
			atoms = prepend(atoms, pivot)
			//put atom->number to dictionary
			dict[string(atoms)] += multiplier
			//reset variables
			atoms = make([]rune, 0, 1)
			multiplier = 0
			order = 1
		} else if pivot == ')' {
			if multiplier == 0 {
				multiplier = 1
			}
			//create new dictionary for recursion
			nested := make(map[string]int, 26)
			//recursion
			countOfAtomsHelper(s, nested)
			//merge result from the recursion
			for k, v := range nested {
				dict[k] += v * multiplier
			}
			multiplier = 0
			order = 1
		} else if pivot == '(' {
			return
		}
	}
}

func countOfAtoms(formula string) string {
	s := newStack()
	dict := make(map[string]int, 26)
	//push elements to stack
	for _, r := range formula {
		s.push(r)
	}
	//calculate and put result in map
	countOfAtomsHelper(s, dict)
	//sort key, which is atoms
	atoms := make([]string, 0, len(dict))
	for k := range dict {
		atoms = append(atoms, k)
	}
	sort.Strings(atoms)
	//output result
	var res bytes.Buffer
	for _, atom := range atoms {
		res.WriteString(atom)
		if dict[atom] != 1 {
			res.WriteString(strconv.Itoa(dict[atom]))
		}
	}
	return res.String()
}