package main

import (
	"flag"
	"fmt"
	"os"
)

// Reference: https://www.rapid7.com/blog/post/2016/08/04/build-a-simple-cli-tool-with-golang
func main() {
	// Subcommands
    createAssociationsCommand := flag.NewFlagSet("create-associations", flag.ExitOnError)
    cleanAnswersCommand := flag.NewFlagSet("clean-answers", flag.ExitOnError)

    // Verify that a subcommand has been provided
    // os.Arg[0] is the main command
    // os.Arg[1] will be the subcommand
    if len(os.Args) < 2 {
        fmt.Println("specify subcommand: create-associations | clean-answers")
        os.Exit(1)
    }

	// Switch on the subcommand
    // Parse the flags for appropriate FlagSet
    // FlagSet.Parse() requires a set of arguments to parse as input
    // os.Args[2:] will be all arguments starting after the subcommand at os.Args[1]
    switch os.Args[1] {
    case "create-associations":
        createAssociationsCommand.Parse(os.Args[2:])
    case "clean-answers":
        cleanAnswersCommand.Parse(os.Args[2:])
    default:
        flag.PrintDefaults()
        os.Exit(1)
    }

	var err error 
    // Check which subcommand was Parsed using the FlagSet.Parsed() function. Handle each case accordingly.
    // FlagSet.Parse() will evaluate to false if no flags were parsed (i.e. the user did not provide any flags)
    if createAssociationsCommand.Parsed() {
		err = createAssociations()
	}  
    if cleanAnswersCommand.Parsed() {
        err = cleanAnswers()
    }

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
}