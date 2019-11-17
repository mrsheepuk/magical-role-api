package magicalroleapi

import (
	"strings"
)

/*
SplitSubjectParam takes a string representing one or more subjects and
splits it into an array of subjects.
*/
func splitSubjectParam(in string) []string {
	result := strings.Split(in, "||")
	return result
}
