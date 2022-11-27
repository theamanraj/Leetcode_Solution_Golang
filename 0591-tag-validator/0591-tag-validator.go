type tagStatus int

const (
	InvalidTag tagStatus = iota
	StartTag
	EndTag
)

func isValid(code string) bool {
	tags := make([]string, 0)
	// basic test, start with '<' and end with '>'
	if len(code) == 0 || code[0] != '<' || code[len(code)-1] != '>' {
		return false
	}
	for i := 0; i < len(code); i++ {
		if code[i] == '<' {
			if code[i+1] == '!' {
				// <![CDATA[
				// CDATA must in a tag
				if len(tags) == 0 {
					return false
				}
				if cdataLength, valid := cdata(code[i:]); valid {
					// loop will increment i by 1
					i += cdataLength-1
				} else {
					return false
				}
			} else {
				// TAG_NAME
				tagLength, valid, tagName := validTag(code[i:])
				nextIndex := i + tagLength
				switch valid {
				case EndTag:
					if len(tags) == 0 || tags[len(tags)-1] != tagName {
						return false
					}
					tags = tags[:len(tags)-1]
					// no tag, no more char in code
					if len(tags) == 0 {
						return nextIndex == len(code)
					}
				case InvalidTag:
					return false
				case StartTag:
					tags = append(tags, tagName)
				}
				// loop will increment i by 1
				i = nextIndex -1
			}
		}
		// TAG_CONTENT: no check is needed
	}
	// abnormal here
	return false
}

// cdata test for valid cdata and return (cdata length, valid or not)
func cdata(code string) (int, bool) {
	if !strings.HasPrefix(code, "<![CDATA[") {
		return -1, false
	}
	if index := strings.Index(code, "]]>"); index != -1 {
		return index+3, true
	}
	return -1, false
}

// validTag test for valid tag and return (tag length include </>, tagStatus, tag)
func validTag(code string) (int, tagStatus, string) {
	startIndex, endIndex := 1, strings.Index(code, ">")
	if endIndex == -1 {
		return -1, InvalidTag, ""
	}
	if strings.HasPrefix(code, "</") {
		startIndex = 2
	}
	tagNameLength := endIndex - startIndex
	if tagNameLength < 1 || tagNameLength > 9 {
		return -1, InvalidTag, ""
	}
	for i := startIndex; i < endIndex; i++ {
		if code[i] < 'A' || code[i] > 'Z' {
			return -1, InvalidTag, ""
		}
	}
	if startIndex == 1 {
		return endIndex+1, StartTag, code[startIndex:endIndex]
	}
	return endIndex+1, EndTag, code[startIndex:endIndex]
}