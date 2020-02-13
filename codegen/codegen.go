package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	defaultURL = "https://releases.rancher.com/kontainer-driver-metadata/dev-v2.4/data.json"
	dataFile   = "data/data.json"
	envName    = "KDM_DATA_URL"
)

// Codegen fetch data.json from https://releases.rancher.com/kontainer-driver-metadata/dev-v2.4/data.json and generates bindata
func main() {
	u := os.Getenv(envName)
	if u == "" {
		u = defaultURL
	}
	data, err := http.Get(u)
	if err != nil {
		panic(fmt.Errorf("failed to fetch data.json from kontainer-driver-metadata repository"))
	}
	defer data.Body.Close()

	b, err := ioutil.ReadAll(data.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Writing data")
	if err := ioutil.WriteFile(dataFile, b, 0755); err != nil {
		return
	}
}
