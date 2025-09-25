package chapter10

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Car struct {
	TopSpeed   int      `yaml:"topspeed"`
	Name       string   `yaml:"name"`
	Brand      string   `yaml:"brand"`
	Passengers []string `yaml:"passengers"`
}

func YamlParse() {
	// c := Car{
	// 	TopSpeed:   231,
	// 	Name:       "SF-90",
	// 	Brand:      "Ferrari",
	// 	Passengers: []string{"skarekroe", "sanjana"},
	// }

	// out, err := yaml.Marshal(c)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(out))

	f, err := os.ReadFile("test.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var c Car
	if err := yaml.Unmarshal(f, &c); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", c)
}
