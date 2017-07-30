package main

import (
	"io/ioutil"
	"encoding/json"
	"./aad"
	"fmt"
)

func main() {
	cfgBytes, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	cfgMap := map[string]string{}
	if err := json.Unmarshal(cfgBytes, &cfgMap); err != nil {
		panic(err)
	}

	token := aad.Auth(
		cfgMap["tenantId"],
		cfgMap["appId"],
		cfgMap["clientSecret"])

	fmt.Println(token)
}
