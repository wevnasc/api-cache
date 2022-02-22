package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GithubProfile struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar_url"`
}

func githubUrl(username string) string {
	return fmt.Sprintf("https://api.github.com/users/%s", username)
}

func requestGithubProfile(username string) (*GithubProfile, error) {
	res, err := http.Get(githubUrl(username))

	if err != nil {
		return nil, fmt.Errorf("error to load github profile: %v", err)
	}

	gp := &GithubProfile{}
	json.NewDecoder(res.Body).Decode(&gp)
	return gp, nil
}
