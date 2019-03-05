package main

import (
    "fmt"
    "io/ioutil"
    "regexp"
    "math/rand"
    "time"

    "gopkg.in/yaml.v2"
)

type Config struct {
    Phrases map[string][]string
}
func main() {
	fmt.Printf(check_message("kelly"));
}

func check_message(text string) string {
    var config Config
    source, err := ioutil.ReadFile("phrases.yaml")
    if err != nil {
        panic(err)
    }
    err = yaml.Unmarshal(source, &config)
    if err != nil {
        panic(err)
    }
    //fmt.Printf("Value: %#v\n", config.Phrases["maru"])
    //fmt.Printf("%+v\n", config)

    for k, v := range config.Phrases {
        //fmt.Printf("key[%s] value[%s]\n", k, v)
	s := fmt.Sprintf(".*%s.*", k)
        matched,err := regexp.MatchString(s, text)
	if err != nil {
            panic(err)
        }
	if(matched == true) {
            rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
            return(v[rand.Intn(len(v))])
	}
    }
    return("")
}
