package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const (
	registerURL = "http://127.0.0.1:8000/frontend/user/register"
	loginURL    = "http://127.0.0.1:8000/frontend/user/login"
	outputFile  = "tokens.txt"
)

// User holds the user information for registration
type User struct {
	Name         string `json:"name"`
	Password     string `json:"password"`
	Avatar       string `json:"avatar"`
	Sign         string `json:"sign"`
	Sex          string `json:"sex"`
	SecretAnswer string `json:"secret_answer"`
}

// LoginResponse holds the login response
type LoginResponse struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

func registerUser(name string) error {
	user := User{
		Name:         name,
		Password:     "123456",
		Avatar:       "default_avatar.png",
		Sign:         "Hello, I am " + name,
		Sex:          "unknown",
		SecretAnswer: "test_answer",
	}

	payload, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user data: %v", err)
	}

	resp, err := http.Post(registerURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to send register request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("registration failed: %s", body)
	}

	fmt.Printf("User %s registered successfully.\n", name)
	return nil
}

func loginUser(name, password string) (string, error) {
	data := map[string]string{
		"name":     name,
		"password": password,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal login data: %v", err)
	}

	resp, err := http.Post(loginURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return "", fmt.Errorf("failed to send login request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("login failed: %s", body)
	}

	var loginResp LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		return "", fmt.Errorf("failed to decode login response: %v", err)
	}

	return loginResp.Data.Token, nil
}

func main() {
	shouldRegister := true // Set to false if registration is not needed
	err := os.Remove(outputFile)
	if err != nil {
		return
	}
	tokenFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Failed to create output file: %v\n", err)
		return
	}
	defer tokenFile.Close()

	for i := 1; i <= 1000; i++ {
		userName := "test_user#" + strconv.Itoa(i)

		if shouldRegister {
			err := registerUser(userName)
			if err != nil {
				fmt.Printf("Error registering user %s: %v\n", userName, err)
				continue
			}
		}

		token, err := loginUser(userName, "123456")
		if err != nil {
			fmt.Printf("Error logging in user %s: %v\n", userName, err)
			continue
		}

		_, err = tokenFile.WriteString(fmt.Sprintf("%s\n", token))
		if err != nil {
			fmt.Printf("Error writing token for user %s: %v\n", userName, err)
		}
	}

	fmt.Println("Operation completed. Tokens saved to", outputFile)
}
