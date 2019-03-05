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

func read_yaml(file string) Config {

    var config Config

    source, err := ioutil.ReadFile(file)
    if err != nil {
        panic(err)
    }

    err = yaml.Unmarshal(source, &config)
    if err != nil {
        panic(err)
    }

    return(config)

}

func write_yaml(file string, config Config) {

    return(1)

}
