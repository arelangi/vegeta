package vegeta

import (
	"fmt"
	"testing"
)

func TestGetRequestSimple(t *testing.T) {
	var resp []byte
	var err error
	handler := Handler{}
	url := "https://api.forecast.io/forecast/8ee11f0e046c284683899336b23c0a76/37.8267,-122.423"
	if resp, err = handler.GetRequest(url); err != nil {
		fmt.Errorf("Expected ", nil, " Received ", err)
	}
	if len(resp) < 1 {
		fmt.Errorf("Expected length to be greater than 0 got ", len(resp))
	}
}

func TestGetRequestWithHeaders(t *testing.T) {
	var resp []byte
	var err error
	handler := NewHandler()
	handler.Token = "ZcV2zWsDtOmshRnJxtexxYM4Dyd6p1MFuIHjsnAPIijtfpuP3X"
	handler.Headers["X-Mashape-Key"] = handler.Token
	handler.Headers["Accept"] = "application/json"
	url := "https://wordsapiv1.p.mashape.com/words/bamboozle?accessToken=" + handler.Token
	if resp, err = handler.GetRequest(url); err != nil {
		fmt.Errorf("Expected:", nil, " Received: ", err)
	}
	if len(resp) < 1 {
		fmt.Errorf("Expected length to be greater than 0 got ", len(resp))
	}
}

func TestNewHandler(t *testing.T) {
	h := NewHandler()
	if _, ok := h.Headers["fakeheader"]; ok {
		fmt.Errorf("Expected: false Received:", ok)
	}
}
