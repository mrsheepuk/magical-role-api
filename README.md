# Brief

Design and implement a simple API that will enumerate Kubernetes RBAC Roles/ClusterRoles that are bound to a given subject name provided in the API request payload. Subject name(s) could be provided as a string, a regex pattern, or a list of strings and/or regex patterns. API response should be presented in a format specified in the request payload in either JSON or YAML format. Subjects and their respective Roles/ClusterRoles should be ordered alphabetically or by role name (string) length.

* The application should be written to run as a Kubernetes application 

* Include all sample resources for deployment of the project

* Previous experience will be taken into account as appropriate

* Implementation language choice left to you to decide, however we’d prefer that to be Golang.

# Implementation Notes

* The required API is naturally an HTTP GET (simply retrieving data, idempotent, doesn't affect the state of any resources it is touching) therefore query parameters / path elements make more sense than a request body (GET has no body, and URL parameters to identify the resource you're requesting is the RESTful way).
* Flexible response format (json/yaml) could be explicitly requested via a parameter/body attribute, or could be via an HTTP accepts header; depends on use case. Accepts header would be more standardised, but requires client of the API to correctly understand and set this header 
* No mention made of authentication/authorisation so assuming this is handled elsewhere/out of scope, but would look to verify this in a real use case.

# API design

## URL pattern
/magicalroleapi/v1/roles?subjectFilter=a&format=b&pageSize=0&page=1

1. subject - 