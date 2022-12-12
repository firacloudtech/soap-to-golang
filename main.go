package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Request struct {
	XMLName     xml.Name `xml:"request"`
	PlateNumber string   `xml:"plateNumber"`
}

type Response struct {
	XMLName     xml.Name `xml:"finalmsgcustomer"`
	PlateNumber string   `xml:"plateNumber"`
	Location    int      `xml:"location"`
}

const (
	worldName = "World"
)

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)

}

// function to convert json data to xml. Yet to be implemented
func jsonToXML(jsonRes []byte) ([]byte, error) {
	// Unmarshal the JSON response into a map.
	var res map[string]interface{}
	if err := json.Unmarshal(jsonRes, &res); err != nil {
		return nil, err
	}

	xmlRes, err := xml.Marshal(res)

	if err != nil {
		return nil, err
	}
	return xmlRes, nil

}

// function to parse incoming xml body from incoming request
func parseXML(r *http.Request) (*Request, error) {

	var req Request
	if err := xml.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil

}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	request, _ := parseXML(r)

	res := Response{PlateNumber: request.PlateNumber, Location: 1111}

	// sent response in xml
	if err := xml.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
