package magicalroleapi

import (
	"testing"

	rbac "k8s.io/api/rbac/v1"
)

func TestRoleFiltering(t *testing.T) {
	cases := []struct {
		roles        []rbac.RoleBinding
		clusterRoles []rbac.ClusterRoleBinding
		subjects     []subjectFilter
		want         []subjectRoles
	}{
		{
			roles: []rbac.RoleBinding{
				rbac.RoleBinding{
					Subjects: []rbac.Subject{
						{Name: "mark"},
					},
					RoleRef: rbac.RoleRef{
						Name: "test",
					},
				},
			},
			clusterRoles: []rbac.ClusterRoleBinding{},
			subjects: []subjectFilter{
				subjectFilter{
					isRegex: false,
					value:   "mark",
				},
			},
			want: []subjectRoles{
				subjectRoles{
					Subject: "mark",
					Roles: []string{
						"test",
					},
				},
			},
		},
		{
			roles: []rbac.RoleBinding{
				rbac.RoleBinding{
					Subjects: []rbac.Subject{
						{Name: "mark"},
					},
					RoleRef: rbac.RoleRef{
						Name: "test",
					},
				},
				rbac.RoleBinding{
					Subjects: []rbac.Subject{
						{Name: "markzzz"},
					},
					RoleRef: rbac.RoleRef{
						Name: "testqq",
					},
				},
			},
			clusterRoles: []rbac.ClusterRoleBinding{
				rbac.ClusterRoleBinding{
					Subjects: []rbac.Subject{
						{Name: "marksheep"},
					},
					RoleRef: rbac.RoleRef{
						Name: "testcluster",
					},
				},
				rbac.ClusterRoleBinding{
					Subjects: []rbac.Subject{
						{Name: "marksheep"},
					},
					RoleRef: rbac.RoleRef{
						Name: "alphamonkey",
					},
				},
			},
			subjects: []subjectFilter{
				subjectFilter{
					isRegex: true,
					value:   "mark.*",
				},
			},
			want: []subjectRoles{
				subjectRoles{
					Subject: "mark",
					Roles: []string{
						"test",
					},
				},
				subjectRoles{
					Subject: "marksheep",
					Roles: []string{
						"alphamonkey",
						"testcluster",
					},
				},
				subjectRoles{
					Subject: "markzzz",
					Roles: []string{
						"testqq",
					},
				},
			},
		},
	}

	for _, c := range cases {
		got, _ := filterRoles(c.roles, c.clusterRoles, c.subjects)
		if len(c.want) != len(got) {
			t.Errorf("filterRoles(%v,%v,%v) = %v, want %v", c.roles, c.clusterRoles, c.subjects, got, c.want)
			continue
		}
		for i := 0; i < len(c.want); i++ {
			if got[i].Subject != c.want[i].Subject {
				t.Errorf("filterRoles(%v,%v,%v) = %v, want %v", c.roles, c.clusterRoles, c.subjects, got, c.want)
			} else {
				for j := 0; j < len(c.want[i].Roles); j++ {
					if got[i].Roles[j] != c.want[i].Roles[j] {
						t.Errorf("filterRoles(%v,%v,%v) = %v, want %v", c.roles, c.clusterRoles, c.subjects, got, c.want)
					}
				}
			}
		}
	}
}
