package grafana

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Datasource struct {
	Name           string `json:"name"`
	Type           string `json:"type"`
	URL            string `json:"url"`
	Access         string `json:"access"`
	OrgID          int    `json:"orgId"`
	BasicAuth      bool   `json:"basicAuth"`
	BasicAuthUser  string `json:"basicAuthUser"`
	SecureJSONData struct {
		BasicAuthPassword string `json:"basicAuthPassword"`
	} `json:"secureJsonData"`
}

func RequestAUTH(url, method string) []byte {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Errorf("Got error %s", err.Error())
	}
	// import grafana credentials from environment var
	user := os.Getenv("K8S_GRAFANA_BA_USER")
	pass := os.Getenv("K8S_GRAFANA_BA_PASSWORD")
	req.SetBasicAuth(user, pass)
	response, err := client.Do(req)
	if err != nil {
		fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Errorf("Got error %s", err.Error())
	}
	return data
}

func CreateDatasource() {}

func CreateOrganization() {}

func GetOrganizationID(orgname string) string {
	url := fmt.Sprintf("http://grafana/api/orgs/name/%v", orgname)
	data := RequestAUTH(url, "GET")
	var jsonData map[string]interface{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Errorf("Got error %s", err.Error())
	}

	id := fmt.Sprintf("%v", jsonData["id"])
	return id
}
