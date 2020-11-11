package common

import (
	"fmt"
	"net/http"
	"regexp"
)

type Permission string

const (
	PermitAdmin Permission = "admin"
	PermitUser  Permission = "user"
	Public      Permission = "public"
)

func toRegex(expr string) *regexp.Regexp {
	return regexp.MustCompile(expr)
}

var Permissions = map[string]map[string][]Permission{
	"/users": {
		http.MethodGet: {PermitUser, PermitAdmin},
	},
	"/users/login": {
		http.MethodPost: {Public},
	},
	"/users/logout": {
		http.MethodDelete: {Public},
	},
	"/users/([0-9]+)": {
		http.MethodGet: {Public},
	},
	"/users/tags": {
		http.MethodGet: {PermitUser, PermitAdmin},
	},
	"/users/archives": {
		http.MethodGet: {PermitUser, PermitAdmin},
	},
	"/archives": {
		http.MethodPost: {PermitUser, PermitAdmin},
		http.MethodGet:  {PermitUser, PermitAdmin},
	},
	"archives/my": {
		http.MethodGet: {PermitUser, PermitAdmin},
	},
	"/archives/url": {
		http.MethodGet: {PermitUser, PermitAdmin},
	},
	"/archives/preview": {
		http.MethodGet: {PermitUser, PermitAdmin},
	},
}

var PathPatterns = makePatterns()

func makePatterns() map[string]*regexp.Regexp {
	patterns := map[string]*regexp.Regexp{}
	for key, _ := range Permissions {
		patterns[key] = regexp.MustCompile(fmt.Sprintf("^%s$", key))
	}
	return patterns
}
