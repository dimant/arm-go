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

	err, exists := a.ResGrpExists(cfgMap["resource-group"])
	if err != nil {
		panic(err)
	}
	fmt.Println("Resource group exists: ", exists)

	err, res := a.List(cfgMap["resource-group"])
	if err != nil {
		panic(err)
	}

	vnetType := "Microsoft.Network/virtualNetworks"
	err, vnet := arm.Find(res, "type", vnetType)
	if err != nil {
		panic(fmt.Errorf("Couldn't find type: %s", vnetType))
	}

	storageType := "Microsoft.Storage/storageAccounts"
	err, storage := arm.Find(res, "type", storageType)
	if err != nil {
		panic(fmt.Errorf("Couldn't find type: %s", storageType))
	}

	fmt.Println(vnet.Id)
	fmt.Println(storage.Id)
}
