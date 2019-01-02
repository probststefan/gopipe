package main

import "fmt"

func main() {
	c := stepList{}

	// read pipeline.yml
	err := getConf(&c)

	if err != nil {
		panic(err)
	}

	// execute pipeline steps
	fmt.Println(c.steps)

	for k, v := range c.steps {
		fmt.Printf("key[%d] value[%s]\n", k, v.Script)
	}
}
