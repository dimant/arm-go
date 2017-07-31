package main

import (
	"io/ioutil"
	"encoding/json"
	"./aad"
	"./arm"
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
		cfgMap["audience"],
		cfgMap["tenantId"],
		cfgMap["appId"],
		cfgMap["clientSecret"])

	armer := arm.Arm{
		cfgMap["api-version"],
	cfgMap["audience"],
	cfgMap["subscriptionId"],
		token}

	// XXX mental node to add timeouts to all http clients

	exists := armer.ResGrpExists("quilt")
	fmt.Println("Resource group 'quilt' exists: ", exists)

}
