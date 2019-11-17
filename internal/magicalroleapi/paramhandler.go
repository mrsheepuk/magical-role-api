package magicalroleapi

import (
	"strings"
)

type subjectFilter struct {
	isRegex bool
	value   string
}

/*
SplitSubjectParam takes a string representing one or more subjects and
splits it into an array of subjects.
*/
func splitSubjectParam(in string) []subjectFilter {
	subjectParams := strings.Split(in, "||")
	var result []subjectFilter

	for i := 0; i < len(subjectParams); i++ {
		if strings.HasPrefix(subjectParams[i], "R:") {
			regex := ([]rune(subjectParams[i]))[2:]
			result = append(result, subjectFilter{isRegex: true, value: string(regex)})
		} else {
			result = append(result, subjectFilter{isRegex: false, value: subjectParams[i]})
		}
	}
	return result
}
