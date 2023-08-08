package netgear_nighthawk_m5_api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

type NetgearNighthawkM5Api struct {
	httpClient    http.Client
	routerBaseURL string
}

func NewNetgearNighthawkM5Api(routerBaseURL string) *NetgearNighthawkM5Api {
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	return &NetgearNighthawkM5Api{
		httpClient: http.Client{
			Jar:           cookieJar,
			Timeout:       5 * time.Second,
			CheckRedirect: checkRedirect,
		},
		routerBaseURL: routerBaseURL,
	}
}

// checkRedirect stops redirects to /index.html to allow getting the redirect on the request.
func checkRedirect(req *http.Request, via []*http.Request) error {
	//if req.URL.Path == "/success.json" || req.URL.Path == "/error.json" {
	//	return http.ErrUseLastResponse
	//}
	return nil
}

func (n *NetgearNighthawkM5Api) FormRequest(formName string, v url.Values) error {
	modelInfo, err := n.ModelInfo()
	if err != nil {
		return fmt.Errorf("failed to get model info: %w", err)
	}

	const failedLocation = "/error.json"
	const successLocation = "/success.json"
	values := url.Values{
		"token":        []string{modelInfo.Session.SecToken},
		"err_redirect": []string{failedLocation},
		"ok_redirect":  []string{successLocation},
	}
	for k, v := range v {
		values[k] = v
	}

	req, err := http.NewRequest(http.MethodPost, n.routerBaseURL+"/Forms/"+formName, bytes.NewBufferString(values.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create config request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := n.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send config request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("config failed: %d %s, expected (200 Ok)", resp.StatusCode, resp.Status)
	}

	if resp.Request.URL.Path != successLocation {
		return fmt.Errorf("config failed: %s, expected %s", resp.Request.URL.Path, successLocation)
	}

	return nil
}

// Login will log the session into the router.
func (n *NetgearNighthawkM5Api) Login(password string) error {
	if err := n.FormRequest("config", url.Values{
		"session.password": []string{password},
	}); err != nil {
		return fmt.Errorf("failed to login: %w", err)
	}

	return nil
}

func (n *NetgearNighthawkM5Api) Logout() error {
	if err := n.FormRequest("config", url.Values{
		"session.password": []string{""},
	}); err != nil {
		return fmt.Errorf("failed to logout: %w", err)
	}

	return nil
}
