package arm

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"../util"
)

func Find(arr []interface{}, k string, v string) (error, *Resource) {
	res := Resource{}
	mapRes := util.Find(arr, func(e interface{}) bool {
		return e.(map[string]interface{})[k] == v
	}).(map[string]interface{})

	if err := res.FillStruct(mapRes); err != nil {
		return err, nil
	} else {
		return nil, &res
	}
}

func (ctx *Arm) List(name string) (error, []interface{}) {
	url := fmt.Sprintf("%ssubscriptions/%s/resourcegroups/%s/resources?api-version=%s",
		ctx.Audience,
		ctx.SubscriptionId,
		name,
		ctx.ApiVersion)


	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err, nil
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

		return nil, l
	} else {
		return fmt.Errorf("Response: %s",resp.StatusCode), nil
	}
}

func (ctx *Arm) ResGrpExists(name string) (error, bool) {
	url := fmt.Sprintf("%ssubscriptions/%s/resourcegroups/%s?api-version=%s",
		ctx.Audience,
		ctx.SubscriptionId,
		name,
		ctx.ApiVersion)

	client := http.Client{}

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return err, false
	}

	req.Header.Set("Authorization", ctx.AuthToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 204 {
		return nil, true
	} else if resp.StatusCode == 404 {
		return nil, false
	} else {
		panic(resp.Status)
	}
}
