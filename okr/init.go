package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const ApplicationFile string = "application.yml"
const ResourcePath string = "okr/resources/"

type Application struct {
	Profiles struct {
		Active string `yaml:"active"`
		Name   string
	}
}

type OkrConfig struct {
	Sql struct {
		Url            string
		Authentication string
	}
}

var application Application

var config OkrConfig

func init() {
	b, err := ioutil.ReadFile(ResourcePath + ApplicationFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	application = Application{}
	err = yaml.Unmarshal(b, &application)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Application.Profiles.Active =" + application.Profiles.Active)
	fmt.Println("Application.Profiles.Name =" + application.Profiles.Name)
	b, err = ioutil.ReadFile(ResourcePath + application.Profiles.Name)
	if err != nil {
		fmt.Println(err)
		return
	}
	config = OkrConfig{}
	err = yaml.Unmarshal(b, &config)
	if err != nil {
		fmt.Println(err)
		return
	}
}