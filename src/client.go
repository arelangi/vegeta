package main

import (
	"fmt"
	"vegeta"
)

func main() {
	var client vegeta.Handler
	client.Token = "ZcV2zWsDtOmshRnJxtexxYM4Dyd6p1MFuIHjsnAPIijtfpuP3X"
	clientHeaders := make(map[string]string)
	clientHeaders["X-Mashape-Key"] = client.Token
	clientHeaders["Accept"] = "application/json"
	client.SetHeaders(clientHeaders)
	client.URL = "https://wordsapiv1.p.mashape.com/words/bamboozle?accessToken=" + client.Token
	if resp, ok := client.GetRequest(); ok {
		fmt.Println(string(resp))
	}
}
