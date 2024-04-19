package pkg 

import (
	"os"

	proxy "github.com/maxneuvians/go-copilot-proxy/pkg"
)

func getToken() string {
	if tokenFileExists() {
		token, _ := os.ReadFile(consts["tokenFile"])
		return string(token)
	}
	return ""
}

func login() (string, string) {
	token := getToken()
	if token != "" {
		return "", ""
	}

	loginResponse, err := proxy.Login()

	if err != nil {
		return "", ""
	}

	return loginResponse.VerificationURI, loginResponse.UserCode
}

func tokenFileExists() bool {
	_, err := os.Stat(consts["tokenFile"])
	return !os.IsNotExist(err)
}