package main

import (
	"fmt"
    "log"
    "io/ioutil"
    "sort"

	"github.com/ghodss/yaml"
)

type Status struct {
    Year int `json:"year"`
    Status string `json:"status"`
}

type Section struct {
    Name string `json:"name"`
    Version string `json:"version"`
    Years []Status `json:"years"`
}

type TechnologyRadar struct {
    TechnologyType string `json:"techType"`
    Version string `json:"version"`
    Sections []Section `json:"sections"`
}

func main() {
    yamlFile, err := ioutil.ReadFile("test.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    var radar TechnologyRadar
    err = yaml.Unmarshal(yamlFile, &radar)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }
    fmt.Println(radar)
    presentYears := getYearsPresent(radar)
    fmt.Println(presentYears)
    fmt.Printf("Number of Sections: %d\n", len(radar.Sections))
    fmt.Printf("Number of Years: %d\n", len(presentYears))
}

func getYearsPresent(radar TechnologyRadar) []int {
    years := map[int]bool{}

    for _, section := range radar.Sections {
        for _, year := range section.Years {
            years[year.Year] = true
        }
    }

    var keys []int
    for k := range years {
        keys = append(keys, k)
    }
    sort.Ints(keys)

    return keys
}
