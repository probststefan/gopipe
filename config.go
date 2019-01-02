package main

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
)

type step struct {
	Image  string
	Script string
}

type stepList struct {
	steps []*step
}

func getConf(sl *stepList) error {
	yamlFile, err := ioutil.ReadFile("pipeline.yml")

	if err != nil {
		return err
	}

	slice := yaml.MapSlice{}
	err = yaml.Unmarshal(yamlFile, &slice)

	if err != nil {
		return err
	}

	for _, sr := range slice {
		s := step{}

		out, err := yaml.Marshal(sr.Value)

		if err != nil {
			return err
		}

		err = yaml.Unmarshal(out, &s)

		if err != nil {
			return err
		}

		sl.steps = append(sl.steps, &s)
	}

	return err
}
