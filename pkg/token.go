package pkg

import (
	"fmt"
	"os"

	proxy "github.com/maxneuvians/go-copilot-proxy/pkg"
)

func authenticate(code string) string {
	var loginResponse = proxy.LoginResponse{
		VerificationURI: "",
		UserCode:        "",
		DeviceCode:      code,
	}

	token, err := proxy.Authenticate(loginResponse)

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return err.Error()
	}

	return token.AccessToken
}

func chat(token string, body []proxy.Message) (string, error) {

	var resp string

	err := proxy.Chat(token, body, false, func(completionResponse proxy.CompletionResponse) error {
		resp = completionResponse.Choices[0].Message.Content
		return nil
	})

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return "", err
	}

	return resp, nil
}

func getSessionToken() string {
	token := getToken()
	if token == "" {
		return ""
	}

	var err error

	sessionResponse, err := proxy.GetSessionToken(token)

	if err != nil {
		fmt.Println("Error: " + err.Error())
		return ""
	}

	return sessionResponse.Token
}

func getToken() string {
	if tokenFileExists() {
		token, _ := os.ReadFile(consts["tokenFile"])
		return string(token)
	}
	return ""
}

func login() (string, string, string) {
	token := getToken()
	if token != "" {
		return "", "", ""
	}

	loginResponse, err := proxy.Login()

	if err != nil {
		return "", "", ""
	}

	return loginResponse.VerificationURI, loginResponse.UserCode, loginResponse.DeviceCode
}

func saveToken(token string) {
	os.WriteFile(consts["tokenFile"], []byte(token), 0644)
}

func tokenFileExists() bool {
	_, err := os.Stat(consts["tokenFile"])
	return !os.IsNotExist(err)
}
