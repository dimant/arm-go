package arm

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"../util"
)

func GetType(arr []interface{}, t string) map[string]interface{} {
	return util.Find(arr, func(e interface{}) bool {
		return e.(map[string]interface{})["type"] == t
	}).(map[string]interface{})
}

func (ctx *Arm) List(name string) []interface{} {
	url := fmt.Sprintf("%ssubscriptions/%s/resourcegroups/%s/resources?api-version=%s",
		ctx.Audience,
		ctx.SubscriptionId,
		name,
		ctx.ApiVersion)


	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", ctx.AuthToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var r interface{}

		if err := json.Unmarshal(bytes, &r); err != nil {
			panic(err)
		}

		a := r.(map[string]interface{})

		l := a["value"].([]interface{})

		return l
	} else {
		panic(resp.StatusCode)
	}
}

func (ctx *Arm) ResGrpExists(name string) bool {
	url := fmt.Sprintf("%ssubscriptions/%s/resourcegroups/%s?api-version=%s",
		ctx.Audience,
		ctx.SubscriptionId,
		name,
		ctx.ApiVersion)

	client := http.Client{}

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", ctx.AuthToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 204 {
		return true
	} else if resp.StatusCode == 404 {
		return false
	} else {
		panic(resp.Status)
	}
}
