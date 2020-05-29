package profile

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const ApplicationFile string = "application.yml"

func active(path string, application interface{}) {
	b, err := ioutil.ReadFile(path + ApplicationFile)
	if err != nil {
		fmt.Println(err)
		return
	}
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
	fmt.Println("OkrConfig.Case =" + config.Case)
}
