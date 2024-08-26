package main

import (
	"fmt"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"log"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	k := koanf.New(".")

	var err error

	if err = k.Load(file.Provider("config.yml"), yaml.Parser()); err != nil {
		log.Fatal(err)
	}

	dst := &Config{}

	if err = k.Unmarshal("", dst); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", dst)
}
