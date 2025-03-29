package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func VerifyJWT(jwt, server string) (string, error) {
	u, err := url.Parse(server + "/token")
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return "", err
	}

	params := url.Values{}
	params.Add("jwt", jwt)
	u.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return "", fmt.Errorf("incorrect JWT")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return "", err
	}
	return string(body)[1 : len(body)-1], nil
}
