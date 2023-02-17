type WordFilter struct {
	dictionay map[key]*metadata
}

type metadata struct {
	wordLen int
	offset  int
}
type key struct {
	prefix string
	suffix string
}

func Constructor(words []string) WordFilter {
	return WordFilter{
		dictionay: genDictionay(words),
	}
}

func genDictionay(words []string) map[key]*metadata {
	dictionary := make(map[key]*metadata)
	for index, word := range words {
		wordLen := len(word)
		for i := 0; i < wordLen; i++ {
			for j := wordLen - 1; j > -1; j-- {
				k := key{
					prefix: word[:i+1],
					suffix: word[j:wordLen],
				}
				value, ok := dictionary[k]
				if !ok || value.offset < index {
					dictionary[k] = &metadata{
						offset:  index,
						wordLen: wordLen,
					}
				}
			}
		}
	}
	return dictionary
}

func (this *WordFilter) F(prefix string, suffix string) int {
	k := key{
		prefix: prefix,
		suffix: suffix,
	}
	value, ok := this.dictionay[k]
	if !ok {
		return -1
	}
	return value.offset
}
