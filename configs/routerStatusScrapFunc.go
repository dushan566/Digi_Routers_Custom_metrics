package configs

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"router/api"
)

func scrap(rID string, simID string) (routerID int, err error) {
	api.GetEnvVars()
	drmUserName := os.Getenv("DRMUSER")
	drmPassword := os.Getenv("DRMPASS")
	drmStreamEndpoint := os.Getenv("STREAM_ENDPOINT")

	drmApi := "https://" + drmUserName + ":" + drmPassword + "@" + drmStreamEndpoint + rID + "/metrics/cellular/1/" + simID + "/connection_status"
	response, err := http.Get(drmApi)

	content, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var router Router
	json.Unmarshal(content, &router)

	if router.Value == "connected" {
		routerID = 1
	} else if response.StatusCode == 400 {
		routerID = 404
	} else {
		routerID = 0
	}
	return
}
