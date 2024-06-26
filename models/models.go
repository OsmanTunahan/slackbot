package models

type SlackRequest struct {
	StatusCode int    `json:"statusCode"`
	Content    string `json:"body"`
}

type SlackResponse struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

type Email struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

type Sms struct {
	From    string `json:"from"`
	Number  string `json:"number"`
	Message string `json:"message"`
}

type Url struct {
	Filename string `json:"filename"`
	Type     string `json:"type"`
	Duration string `json:"duration"`
}

const (
	EmailsCommand = "/emails"
	SmsCommand    = "/sms"
	UrlCommand    = "/url"
)