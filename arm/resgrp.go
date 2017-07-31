package arm

import (
	"fmt"
	"net/http"
)

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
