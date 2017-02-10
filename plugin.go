package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	respFormat = "URL: %s\n  RESPONSE STATUS: %s\n  RESPONSE BODY: %s\n"
)

type (
	Repo struct {
		Owner string `json:"owner"`
		Name  string `json:"name"`
	}

	Build struct {
		Tag     string `json:"tag"`
		Event   string `json:"event"`
		Number  int    `json:"number"`
		Commit  string `json:"commit"`
		Ref     string `json:"ref"`
		Branch  string `json:"branch"`
		Author  string `json:"author"`
		Status  string `json:"status"`
		Link    string `json:"link"`
		Started int64  `json:"started"`
		Created int64  `json:"created"`
	}

	Config struct {
		Webhook    string
		Token      string
		SkipVerify bool
	}

	Job struct {
		Started int64
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Config Config
		Job    Job
	}
)

func (p Plugin) Exec() error {
	var buf bytes.Buffer
	data := struct {
		Repo   Repo   `json:"repo"`
		Build  Build  `json:"build"`
	}{p.Repo, p.Build}

	if err := json.NewEncoder(&buf).Encode(&data); err != nil {
		return err
	}

	uri, err := url.Parse(p.Config.Webhook)

	if err != nil {
		return err
	}

	b := buf.Bytes()
	r := bytes.NewReader(b)

	req, err := http.NewRequest("POST", uri.String(), r)

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.Config.Token))

	client := http.DefaultClient
	if p.Config.SkipVerify {
		client = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}
	}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return err
		}
  	fmt.Printf(
	  	respFormat,
  		req.URL,
		  resp.Status,
		  string(body),
	  )
	}

	return nil
}
