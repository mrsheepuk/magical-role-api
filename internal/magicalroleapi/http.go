package magicalroleapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

// roleAPI parses parameters from HTTP requests, routes them to the relevant underlying logic,
// then formats the responses back onto the HTTP response.
type roleAPI struct {
	roleGetter *roleGetter
}

func (ra *roleAPI) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All sheep are welcome here, provided all they want to know is Who Has What Roles!")
}

func (ra *roleAPI) getSubjectRoles(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	subjectFilter := r.Form.Get("subjectFilter")
	responseFormat := r.Form.Get("format")

	if subjectFilter == "" {
		// No filter provided, return the empty set.
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Please provide a subjectFilter parameter")
		return
	}
	if responseFormat != "" && responseFormat != "json" && responseFormat != "yaml" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "responseFormat must be blank, json or yaml")
		return
	}

	subjects := splitSubjectParam(subjectFilter)

	log.Println("Getting roles list")
	roles, err := ra.roleGetter.Roles("default")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error loading roles: %v", err)
		return
	}

	log.Println("Getting cluster roles list")
	clusterRoles, err := ra.roleGetter.ClusterRoles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error loading roles: %v", err)
		return
	}

	log.Println("Filtering role lists to produce result for subjects", subjects)
	filteredRoles, err := filterRoles(roles, clusterRoles, subjects)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error filtering roles: %v", err)
		return
	}

	// TODO: Format response optionally as yaml
	result := struct {
		Result []subjectRoles
	}{
		Result: filteredRoles,
	}
	if responseFormat == "" || responseFormat == "json" {
		log.Println("Encoding result to JSON")
		js, err := json.MarshalIndent(&result, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error encoding JSON: %v", err)
			return
		}

		log.Println("Sending JSON")
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(js)
		if err != nil {
			log.Printf("Failed to write the response body: %v", err)
			return
		}
	}
	if responseFormat == "yaml" {
		log.Println("Encoding result to YAML")
		yaml, err := yaml.Marshal(&result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error encoding YAML: %v", err)
			return
		}

		log.Println("Sending YAML")
		w.Header().Set("Content-Type", "application/x-yaml")
		_, err = w.Write(yaml)
		if err != nil {
			log.Printf("Failed to write the response body: %v", err)
			return
		}
	}
}
