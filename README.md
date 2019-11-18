# Brief

Design and implement a simple API that will enumerate Kubernetes RBAC Roles/ClusterRoles that are bound to a given subject name provided in the API request payload. Subject name(s) could be provided as a string, a regex pattern, or a list of strings and/or regex patterns. API response should be presented in a format specified in the request payload in either JSON or YAML format. Subjects and their respective Roles/ClusterRoles should be ordered alphabetically or by role name (string) length.

* The application should be written to run as a Kubernetes application 

* Include all sample resources for deployment of the project

* Previous experience will be taken into account as appropriate

* Implementation language choice left to you to decide, however we’d prefer that to be Golang.

# Assumptions

* Authorisation / authentication of the API is out of scope.

* Only the "default" namespace is considered at this time. It would be trivial to add this
  as an extra parameter to the API if required.

* The implementation assumes it is executing "in-cluster" on the Kubernetes cluster it is 
  inspecting, so does not expose configuration options for connecting to a remote cluster.

# Implementation Choices / Discussion

* The required API is naturally an HTTP GET (simply retrieving data, idempotent, doesn't affect the state of any resources it is touching) therefore query parameters / path elements make more sense than a request body (GET has no body, and URL parameters to identify the resource you're requesting is the RESTful way).
* Flexible response format (json/yaml) could be explicitly requested via a parameter/body attribute, or could be via an HTTP accepts header; depends on use case. Accepts header would be more standardised, but requires client of the API to correctly understand and set this header. For the purposes of this implementation, a parameter has been chosen.

# API design

/magicalroleapi/v1?subjectFilter=a&format=b

1. subjectFilter - one or more subject names in the format NAME1||NAME2||NAME3
   Supports regexes using format R:regex1||R:regex2
   Regexes and full matches can be mixed, e.g. mrsheep||R:sys.*

2. format - optional, if specified should be either "json" or "yaml". Defaults to json.

# Installation

To run the latest built version from docker hub:

1. Run kubectl apply -f ./deployment.yaml
2. Run kubectl apply -f ./service.yaml
3. Browse to port 8080 on the cluster load balancer IP assigned by applying the service.