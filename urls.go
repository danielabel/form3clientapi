package form3clientapi

import "fmt"

const defaultDomain = "localhost"
const defaultPort = "8080"
var userDomain = ""
var userPort = ""
func setDomain(domain string, port string) {
	userDomain = domain
	userPort = port
}

func getbaseUrl() string {
	domain := defaultDomain
	if userDomain != "" {
		domain = userDomain
	}
	port := defaultPort
	if userPort != "" {
		port = userPort
	}

	return fmt.Sprintf("http://%s:%s/v1", domain, port)
}
