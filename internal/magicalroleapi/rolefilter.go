package magicalroleapi

import (
	"regexp"
	"sort"

	rbac "k8s.io/api/rbac/v1"
)

type subjectRoles struct {
	Subject string
	Roles   []string
}

// filterRoles filters the provided list of roles and clusterRoles for those which apply
// to the subject or subjects listed.
func filterRoles(roles []rbac.RoleBinding, clusterRoles []rbac.ClusterRoleBinding, subjects []subjectFilter) ([]subjectRoles, error) {
	// Create empty slice to hold the result
	resultIndex := make(map[string][]string)

	// TODO(mrsheepuk) The below multiple-nested for loops implementation looks awful, I'm sure
	// there's something more efficient but I don't know enough Go yet to work out what the
	// right way of approaching this is. The whole of the below would collapse down to a couple
	// of lines of C#...

	// Iterate role bindings, check if any of the subjects on that binding match and add to
	// result
	for _, role := range roles {
		checkRoleBindingSubjects(resultIndex, role.Subjects, subjects, role.RoleRef.Name)
	}

	// Iterate cluster role bindings and do the same as above.
	for _, role := range clusterRoles {
		checkRoleBindingSubjects(resultIndex, role.Subjects, subjects, role.RoleRef.Name)
	}

	// TODO(mrsheepuk) The below looks potentially inefficient, research better ways to
	// implement this in Go.
	var matchedSubjects []string
	for key := range resultIndex {
		matchedSubjects = append(matchedSubjects, key)
	}
	sort.Strings(matchedSubjects)

	var result []subjectRoles
	for _, key := range matchedSubjects {
		value := resultIndex[key]
		sort.Strings(value)
		result = append(result, subjectRoles{
			Subject: key,
			Roles:   value,
		})
	}

	return result, nil
}

func checkRoleBindingSubjects(resultIndex map[string][]string, subjects []rbac.Subject, subjectFilters []subjectFilter, roleName string) error {
	for _, subject := range subjects {
		for _, subjectFilter := range subjectFilters {
			if subjectFilter.isRegex {
				match, err := regexp.MatchString(subjectFilter.value, subject.Name)
				if err != nil {
					return err
				}
				if match {
					resultIndex[subject.Name] = append(resultIndex[subject.Name], roleName)
				}
			} else {
				if subject.Name == subjectFilter.value {
					resultIndex[subject.Name] = append(resultIndex[subject.Name], roleName)
				}
			}
		}
	}
	return nil
}
