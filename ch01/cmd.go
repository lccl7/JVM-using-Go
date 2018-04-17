package main

import (
	"flag"
	"fmt"
	"os"
)

//define a struct for cmd to store command parameters
type Cmd struct {
	helpFlag bool
	versionFlag bool
	cpOption string
	class string
	args []string
}

//flag.Parse() is used to analysis parameters
//see more detail in flag package in go
func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage //Usage is a function
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
