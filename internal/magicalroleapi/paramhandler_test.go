package magicalroleapi

import "testing"

func TestSplitSubjectParam(t *testing.T) {
	cases := []struct {
		in   string
		want []subjectFilter
	}{
		{"singlesubject", []subjectFilter{
			subjectFilter{isRegex: false, value: "singlesubject"},
		}},
		{"singlesubject||secondsubject", []subjectFilter{
			subjectFilter{isRegex: false, value: "singlesubject"},
			subjectFilter{isRegex: false, value: "secondsubject"},
		}},
		{"R:regexSubject||R:regexSubject2", []subjectFilter{
			subjectFilter{isRegex: true, value: "regexSubject"},
			subjectFilter{isRegex: true, value: "regexSubject2"},
		}},
	}

	for _, c := range cases {
		got := splitSubjectParam(c.in)
		if len(c.want) != len(got) {
			t.Errorf("SplitSubjectParam(%v) = %v, want %v", c.in, got, c.want)
			continue
		}
		for i := 0; i < len(c.want); i++ {
			if got[i] != c.want[i] {
				t.Errorf("SplitSubjectParam(%v) = %v, want %v", c.in, got, c.want)
			}
		}
	}

}
