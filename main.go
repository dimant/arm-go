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
		cfgMap["tenant-id"],
		cfgMap["app-id"],
		cfgMap["client-secret"])

	a := arm.Arm{
		cfgMap["api-version"],
	cfgMap["audience"],
	cfgMap["subscription-id"],
		token}

	// XXX mental note to add timeouts to all http clients

	exists := a.ResGrpExists(cfgMap["resource-group"])
	fmt.Println("Resource group 'quilt' exists: ", exists)

	res := a.List(cfgMap["resource-group"])
	fmt.Println(res)
}
