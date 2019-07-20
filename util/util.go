package util

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"net"
)

// ResponseJSON response data
func ResponseJSON(w http.ResponseWriter, code int, data interface{}) error {
	_, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	err = json.NewEncoder(w).Encode(data)
	return err
}

// RequestJSON read data
func RequestJSON(r *http.Request, jsonObject interface{}) error {
	rawData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(rawData, jsonObject)
}

func GetIpAddr() string{
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && ! ipnet.IP.IsLoopback() {
			if ipnet.IP.To4()!=nil {
				return ipnet.IP.String()
			}
		}
	}
	panic("Cannot find the local ip")
}
