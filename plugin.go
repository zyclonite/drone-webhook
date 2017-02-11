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
		Scm      string `json:"scm"`
		Owner    string `json:"owner"`
		Name     string `json:"name"`
		Link     string `json:"link"`
		Avatar   string `json:"avatar"`
		Branch   string `json:"branch"`
		Private  bool `json:"private"`
		Trusted  bool `json:"trusted"`
	}

	Build struct {
		Tag      string `json:"tag"`
		Number   int    `json:"number"`
		Event    string `json:"event"`
		Status   string `json:"status"`
		Link     string `json:"build_url"`
		Deploy   string `json:"deploy_to"`
		Created  int64  `json:"created_at"`
		Started  int64  `json:"started_at"`
		Finished int64  `json:"finished_at"`
		Url      string `json:"url"`
		Commit   string `json:"commit"`
		Ref      string `json:"ref"`
		Branch   string `json:"branch"`
		Clink    string `json:"link_url"`
		Message  string `json:"message"`
		Author   string `json:"author"`
		Email    string `json:"author_email"`
		Avatar   string `json:"author_avatar"`
	}

	Config struct {
		Webhook    string
		Token      string
		SkipVerify bool
	}

	Job struct {
		Number  int
		Status  string
		Error   string
		Code    int
		Started int64
		Finished int64
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
