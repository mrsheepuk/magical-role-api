package magicalroleapi

import (
	rbac "k8s.io/api/rbac/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type roleGetter struct {
	clientset *kubernetes.Clientset
}

func (rg *roleGetter) Roles(namespace string) ([]rbac.RoleBinding, error) {
	roles, err := rg.clientset.RbacV1().RoleBindings(namespace).List(meta.ListOptions{})
	if err != nil {
		return nil, err
	}
	return roles.Items, nil
}

func (rg *roleGetter) ClusterRoles() ([]rbac.ClusterRoleBinding, error) {
	roles, err := rg.clientset.RbacV1().ClusterRoleBindings().List(meta.ListOptions{})
	if err != nil {
		return nil, err
	}
	return roles.Items, nil
}
