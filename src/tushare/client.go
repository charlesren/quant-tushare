package tushare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
)

var TushareConfig struct {
	Endpoint string `json:"endpoint"`
	Token    string `json:"token"`
}

func init() {
	config, err := ioutil.ReadFile("./tushare.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(config, &TushareConfig); err != nil {
		fmt.Println(err)
		return
	}
}

// TuShare instance
type TuShare struct {
	token  string
	client *http.Client
}

// New TuShare default client
func New(token string) *TuShare {
	return NewWithClient(token, http.DefaultClient)
}

// NewWithClient TuShare client with arguments
func NewWithClient(token string, httpClient *http.Client) *TuShare {
	return &TuShare{
		token:  token,
		client: httpClient,
	}
}

func (api *TuShare) request(method, path string, body interface{}) (*http.Request, error) {
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, path, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (api *TuShare) doRequest(req *http.Request) (*APIResponse, error) {
	// Set http content type
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	resp, err := api.client.Do(req)
	//Handle network error
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("oops! Network error")
	}

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var jsonData *APIResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonData)

	// Check mime type of response
	mimeType, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, err
	}
	if mimeType != "application/json" {
		return nil, fmt.Errorf("Could not execute request (%s)", fmt.Sprintf("Response Content-Type is '%s', but should be 'application/json'.", mimeType))
	}

	// @TODO: handle API exception
	// Argument required
	if jsonData.Code == -2001 {
		return jsonData, fmt.Errorf("Argument error: %s", jsonData.Msg)
	}

	// Permission deny
	if jsonData.Code == -2002 {
		return jsonData, fmt.Errorf("Your point is not enough to use this api")
	}

	return jsonData, nil
}

func (api *TuShare) postData(body map[string]interface{}) (*APIResponse, error) {
	req, err := api.request("POST", TushareConfig.Endpoint, body)
	if err != nil {
		return nil, err
	}
	resp, err := api.doRequest(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// ParsingData  parsing tushare response to gorm form data
func (resp *APIResponse) ParsingData() []Daily {
	items := resp.Data.Items
	fields := resp.Data.Fields
	var dbdata []Daily
	for _, value := range items {
		iterData := Daily{}
		for i := 0; i < len(value); i++ {
			/*
				switch v := value[i].(type) {
				case string, float64:
					iterData.(fmt.Println(a)) := v
				case nil:
					fmt.Println("data is empty")
				default:
					fmt.Println("data type is not string nor float64")
				}
			*/
			switch fields[i] {
			case "ts_code":
				switch v := value[i].(type) {
				case string:
					iterData.TsCode = v
				//case  float64:
				//	iterData.TsCode = value[i]
				case nil:
					fmt.Println("data is empty")
				default:
					fmt.Println("data type is not string nor float64")
				}
				/*
					case trade_date:
						iterData.TradeDate = value[i]
					case open:
						iterData.Open = value[i]
					case high:
						iterData.High = value[i]
					case low:
						iterData.Low = value[i]
					case close:
						iterData.Rlose = value[i]
					case pre_close:
						iterData.PreClose = value[i]
					case change:
						iterData.Change = value[i]
					case pct_chg:
						iterData.PctChg = value[i]
					case vol:
						iterData.Vol = value[i]
					case amount:
						iterData.Amount = value[i]
				*/
			}
		}
		dbdata = append(dbdata, iterData)
	}
	return dbdata
}
