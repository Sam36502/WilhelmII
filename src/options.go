package main

import (
	"github.com/magiconair/properties"
)

var Options *properties.Properties

func LoadOptions(optionsFile string) {
	Options = properties.MustLoadFile(optionsFile, properties.UTF8)
}
