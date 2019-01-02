package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// read pipeline.yml
	c := stepList{}
	err := getConf(&c)
	if err != nil {
		panic(err)
	}

	// execute pipeline steps
	fmt.Println(c.steps)

	for k, v := range c.steps {
		fmt.Printf("key[%d] value[%s]\n", k, v.Script)
	}

	// docker build current directory
	cmdName := "docker"
	cmdArgs := []string{"build", "."}

	cmd := exec.Command(cmdName, cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
		os.Exit(1)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
		os.Exit(1)
	}
}
