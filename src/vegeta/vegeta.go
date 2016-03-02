package vegeta

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Handler object handles the requests to the API
type Handler struct {
	Token    string
	User     string
	Password string
	URL      string
	Headers  map[string]string
	Request  http.Request
}

func (h *Handler) SetUsername(u string) {
	h.User = u
}

func (h *Handler) SetPassword(p string) {
	h.Password = p
}

//GetRequest is used to make requests using the get method
func (h *Handler) GetRequest() (output []byte, ok bool) {
	var client http.Client
	request, err := http.NewRequest("GET", h.URL, nil)
	if err != nil {
		log.Print(err)
		return
	}
	for key, value := range h.Headers {
		request.Header.Add(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Print(err)
		return
	}

	output, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print(err)
		return
	}
	ok = true
	return
}

//SetHeaders is used to set the HTTP headers
func (h *Handler) SetHeaders(headers map[string]string) {
	h.Headers = headers
}

//NewHandler is the constructor that initializes the map object
func NewHandler() *Handler {
	return &Handler{Headers: make(map[string]string)}
}
