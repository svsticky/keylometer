package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Config map[string][]string

func main() {
	log.SetOutput(os.Stderr)
	data, err := ioutil.ReadFile("/etc/keylometer.yml")
	var cfg Config
	if err != nil {
		log.Printf("Error reading config: %s", err)
		data, err = ioutil.ReadFile("./keylometer.yml")
		if err != nil {
			log.Panicf("No config found!")
		}
	}
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		log.Printf("Error parsing yaml: %s", err)
	}
	log.Printf("%#v", cfg)
	if len(os.Args) < 2 {
		log.Panicf("Missing user parameter!")
	}
	users := cfg[os.Args[1]]
	keys := make([]string, 0)
	for _, u := range users {
		if u == "" {
			continue
		}
		res, err := http.Get(fmt.Sprintf("https://github.com/%s.keys", u))
		if err != nil {
			log.Printf("Could not get keys for user %s: %s", u, err)
			continue
		}
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Printf("Could not get keys for user %s: %s", u, err)
			continue
		}
		for _, key := range strings.Split(string(b), "\n") {
			if key != "" {
				keys = append(keys, fmt.Sprintf("%s", key))
			}
		}
	}
	for _, k := range keys {
		fmt.Println(k)
	}

}
