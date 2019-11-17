package magicalroleapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func Run(port int) {
	router := mux.NewRouter().StrictSlash(true)
	setupRoutes(router)
	run(router, port)
}

func setupRoutes(router *mux.Router) {
	router.HandleFunc("/", home)
	router.HandleFunc("/magicalroleapi/v1", getSubjectRoles)
}

func run(router *mux.Router, port int) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All sheep are welcome here, provided all they want to know is Who Has What Roles!")
}

func getSubjectRoles(w http.ResponseWriter, r *http.Request) {
	// Parse request into strings
	subjectFilter := r.Form.Get("subjectFilter")

	if subjectFilter == "" {
		// No filter provided, return the empty set.
		fmt.Fprintf(w, "Horse")
	}

	//subjects := splitSubjectParam(subjectFilter)

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	roles, err := clientset.RbacV1().RoleBindings("default").List(meta.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Roles: %v", roles)
}
