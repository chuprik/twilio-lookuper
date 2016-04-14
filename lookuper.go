package lookuper

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// URL contains address for lookuping number
const URL = "https://lookups.twilio.com/v1/PhoneNumbers/%s"

// Client struct for quering Twilio API
type Client struct {
	auth string
	http *http.Client
}

// New creates Client struct
func New(sid string, token string) (client *Client) {
	data := fmt.Sprintf("%s:%s", sid, token)

	return &Client{
		auth: fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(data))),
		http: &http.Client{},
	}
}

// DoLookup returns Twilio number data
func (client *Client) DoLookup(number string) (data map[string]interface{}, err error) {
	url := fmt.Sprintf(URL, number)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("authorization", client.auth)
	response, err := client.http.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
