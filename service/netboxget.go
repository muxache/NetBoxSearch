package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"../model"
)

//GetToNetBox send get request to netbox
func GetToNetBox(url, token string) model.NetBoxJSON {
	var (
		limit  int
		newURL string
	)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error when sending request to the server")
	}
	defer resp.Body.Close()
	var nb model.NetBoxJSON
	json.NewDecoder(resp.Body).Decode(&nb)
	if len(nb.Next) != 0 {
		limit, _ = strconv.Atoi(URLParse(nb.Next)["limit"][0])
		newURL = nb.Next
		for i := limit; i <= nb.Count; i += limit {

			reqnext, _ := http.NewRequest("GET", newURL, nil)
			reqnext.Header.Add("accept", "application/json")
			reqnext.Header.Add("Authorization", token)
			respnext, err1 := client.Do(reqnext)
			if err1 != nil {
				fmt.Println("Error when sending request to the server")
			}
			defer respnext.Body.Close()
			var pn model.NetBoxJSON
			json.NewDecoder(respnext.Body).Decode(&pn)
			newURL = pn.Next
			nb.Results = append(nb.Results, pn.Results...)
		}
	}
	return nb
}

//URLParse parses url string
func URLParse(urlField string) url.Values {
	u, _ := url.Parse(urlField)
	m, _ := url.ParseQuery(u.RawQuery)
	return m
}

//URLSet makes new url string
func URLSet(urlField, newLimit, offset string) string {
	u, _ := url.Parse(urlField)
	q := u.Query()
	q.Set("limit", newLimit)
	q.Set("offset", offset)
	u.RawQuery = q.Encode()
	return u.String()
}
