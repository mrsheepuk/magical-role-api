package magicalroleapi

import "testing"

func TestSplitSubjectParam(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{"singlesubject", []string{"singlesubject"}},
		{"singlesubject||secondsubject", []string{"singlesubject", "secondsubject"}},
	}

	for _, c := range cases {
		got := splitSubjectParam(c.in)
		if len(c.want) != len(got) {
			t.Errorf("SplitSubjectParam(%q) = %q, want %q", c.in, got, c.want)
			continue
		}
		for i := 0; i < len(c.want); i++ {
			if got[i] != c.want[i] {
				t.Errorf("SplitSubjectParam(%q) = %q, want %q", c.in, got, c.want)
			}
		}
	}

}
