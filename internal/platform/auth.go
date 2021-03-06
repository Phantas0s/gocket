package platform

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

const authURL = "https://getpocket.com/auth/authorize"

type RequestToken struct {
	Code string `json:"code"`
}

type Authorization struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
}

// OAuth2
// 1. Send the consumer key.
// 2. User needs to confirm on a webpage.
// 3. Get a token back - save it to a file.
// 4. Can retrieve pocket list thanks to token.
func Auth(consumerKey string) (Authorization, error) {
	auth := Authorization{}
	authPath := filepath.Join(configDir(), "auth.json")

	r, err := os.Open(authPath)
	defer r.Close()
	if err != nil {
		ch := make(chan struct{})
		ts := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				if req.URL.Path == "/favicon.ico" {
					http.Error(w, "Not Found", 404)
					return
				}

				w.Header().Set("Content-Type", "text/plain")
				fmt.Fprintln(w, "Authorized.")
				ch <- struct{}{}
			}))
		defer ts.Close()

		requestToken, err := ObtainRequestToken(consumerKey, ts.URL)
		if err != nil {
			panic(err)
		}

		url := GenerateAuthorizationURL(requestToken, ts.URL)

		err = openBrowser(url)
		if err != nil {
			panic(err)
		}

		<-ch

		accessToken, err := ObtainAccessToken(consumerKey, requestToken)
		if err != nil {
			panic(err)
		}
		auth.AccessToken = accessToken

		w, err := os.Create(authPath)
		if err != nil {
			return Authorization{}, err
		}
		defer w.Close()

		json.NewEncoder(w).Encode(&auth)
	} else {
		json.NewDecoder(r).Decode(&auth)
	}

	return auth, nil
}

func configDir() (configDir string) {
	configDir = filepath.Join(xdg.ConfigHome, "gocket")
	err := os.MkdirAll(configDir, 0755)
	if err != nil {
		// TODO return error
		panic(err)
	}

	return
}

func ObtainRequestToken(consumerKey, redirectURL string) (*RequestToken, error) {
	res := &RequestToken{}
	err := Post(
		"/v3/oauth/request",
		map[string]string{
			"consumer_key": consumerKey,
			"redirect_uri": redirectURL,
		},
		res,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func ObtainAccessToken(consumerKey string, requestToken *RequestToken) (string, error) {
	res := &Authorization{}
	err := Post(
		"/v3/oauth/authorize",
		map[string]string{
			"consumer_key": consumerKey,
			"code":         requestToken.Code,
		},
		res,
	)
	if err != nil {
		return "", err
	}

	return res.AccessToken, nil
}

func GenerateAuthorizationURL(requestToken *RequestToken, redirectURL string) string {
	values := url.Values{"request_token": {requestToken.Code}, "redirect_uri": {redirectURL}}
	return fmt.Sprintf("%s?%s", authURL, values.Encode())
}
