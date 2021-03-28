package encoderReplacer

import (
	"strings"
)

func Replace(input string, replacer *strings.Replacer) string {
	return replacer.Replace(input)
}

func GenerateReplacer(key int, dictionaries ...[]rune) *strings.Replacer {
	_len := 0
	for _, dict := range dictionaries {
		_len += len(dict)
	}
	replacerSlice := make([]string, _len*2, _len*2)
	i := 0
	for _, dict := range dictionaries {
		for j := range dict {
			replacerSlice[i] = string(dict[j])
			replaceNum := j + key
			for replaceNum >= len(dict) {
				replaceNum = replaceNum - len(dict)
			}
			if replaceNum < 0 {
				replaceNum = len(dict) - (replaceNum * -1)
			}
			replacerSlice[i+1] = string(dict[replaceNum])
			i += 2
		}
	}

	return strings.NewReplacer(replacerSlice...)
}
