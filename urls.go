package form3clientapi

import "fmt"

const defaultDomain = "localhost"
const defaultPort = "8080"

var userDomain = defaultDomain
var userPort = defaultPort

func resetDomain() {
	userDomain = defaultDomain
	userPort = defaultPort
}

func setDomain(domain string, port string) {
	if domain != "" {
		userDomain = domain
	}
	if port != "" {
		userPort = port
	}
}

func getBaseUrl() string {
	return fmt.Sprintf("http://%s:%s/v1", userDomain, userPort)
}
