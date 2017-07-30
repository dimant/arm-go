package aad

import (
	"fmt"
	"net/url"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
)

func Auth(tenantId string, appId string, clientSecret string) string {
	aadUrl := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/token?api-version=1.0", tenantId)

	data := url.Values{}
	data.Set("grant_type","client_credentials")
	data.Set("resource","https://management.core.windows.net")
	data.Set("client_id", appId)
	data.Set("client_secret", clientSecret)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", aadUrl, bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if(err != nil) {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		tokenMap := map[string]string{}
		if err := json.Unmarshal(bodyBytes, &tokenMap); err != nil {
			panic(err)
		}
		return tokenMap["access_token"]
	}

	panic(fmt.Sprintf("AAD Authentication failed: %d", resp.StatusCode))
}