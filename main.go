package main

func main() {
	c := stepList{}

	// read pipeline.yml
	if err := getConf(&c); err != nil {
		panic(err)
	}

	// execute pipeline steps
	for _, v := range c.steps {
		run(v)
	}
}
