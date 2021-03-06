package structs

// HTTPMethodsResponse represents the response
// that will be sent back to clients requesting
// any of the http_methods endpoint.
type Response struct {
	URL       string                 `json:"url,omitempty"`
	Args      map[string]interface{} `json:"args,omitempty"`
	Headers   map[string]interface{} `json:"headers,omitempty"`
	Origin    string                 `json:"origin,omitempty"`
	Form      map[string]interface{} `json:"form,omitempty"`
	JSON      interface{}            `json:"json,omitempty"`
	Data      interface{}            `json:"data,omitempty"`
	Method    string                 `json:"method,omitempty"`
	IP        string                 `json:"ip,omitempty"`
	UserAgent string                 `json:"user-agent,omitempty"`
	Cookies   map[string]interface{} `json:"cookies,omitempty"`
}

type AuthResponse struct {
	Authenticated bool   `json:"authenticated"`
	User          string `json:"user,omitempty"`
	Token         string `json:"token,omitempty"`
}
