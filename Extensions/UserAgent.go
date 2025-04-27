package Extensions

import (
	"github.com/mssola/user_agent"
	"net/http"
)

type (
	UserAgent struct {
		httpRequest *http.Request
	}
	UserAgentData struct {
		Browser        string
		BrowserVersion string
		IsMobile       bool
		Localization   string
		OS             string
		Platform       string
		IP             string
	}
)

func NewUserAgent(httpRequest *http.Request) *UserAgent {
	return &UserAgent{httpRequest: httpRequest}
}

func (u UserAgent) Get() UserAgentData {
	var output UserAgentData

	userAgent := u.httpRequest.Header.Get("user-agent")
	ua := user_agent.New(userAgent)
	browser, browserVersion := ua.Browser()
	output = UserAgentData{
		IsMobile:       ua.Mobile(),
		OS:             ua.OS(),
		Platform:       ua.Platform(),
		Localization:   ua.Localization(),
		Browser:        browser,
		BrowserVersion: browserVersion,
	}

	forwarded := u.httpRequest.Header.Get("X-FORWARDED-FOR")

	if forwarded != "" {
		output.IP = forwarded
	} else {
		output.IP = u.httpRequest.RemoteAddr
	}

	return output
}
