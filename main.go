//go:build go1.16
// +build go1.16

/*
 *
 *
 */

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"ermites.io/pkisauce/config"
)

func jsonErrorLineChar(input string, offset int) {
}

func Usage(cmd string) {
	fmt.Printf("usage: %s [options] <pkidesc1.json pkidesc2.json ...>\n", cmd)
	fmt.Printf("where [options]:\n")
	fmt.Printf("  -userdata <data>\n")
	fmt.Printf("  -pkidebug <name>\n")
	fmt.Printf("  -convert <outputfile>\n")
}

// args:
// -pkidebug <name>: output to <name> stdin using test templates.
// -test	: output pkscd & pkscc code to debug the policy.
func main() {
	var err error
	//fmt.Printf("Flavors: %d %d %d %d %d\n", FlavorNone, FlavorGoTLS, FlavorGoGRPC, FlavorPEM, FlavorJavaGRPC)

	debugFlag := flag.String("pkidebug", "", "display the defined pki node to stdout")
	userFlag := flag.String("userdata", "", "add the following user data to the certificate in 'Locality'")
	verFlag := flag.Bool("version", false, "output version")
	//json2cfFlag := flag.Bool("j2c", false, "read a json conf stdin and stdout a conf")
	//cf2jsonFlag := flag.Bool("c2j", false, "read a conf stdin and stdout a json conf")
	flag.Parse()
	argv := flag.Args()

	if *verFlag {
		fmt.Printf("PKI sauce version %s\nsaucing TLS on your services\n", PkiSauceVersion)
		os.Exit(0)
	}

	// let's test our shit...
	ca, err := NewCA(182)
	if err != nil {
		panic(err)
	}

	// create the ephmeral PKI.
	// the pki will include some userdata
	// maybe later at node level inside the config directly..
	pki, err := NewPKI(ca, *userFlag)
	if err != nil {
		panic(err)
	}

	// if multiple files are given they represent ONE pki.
	// we build that one view.
	for _, fileArgv := range argv {
		configFile, err := filepath.Abs(fileArgv)
		if err != nil {
			panic(err)
		}

		fmt.Printf("parsing '%s'\n", configFile)
		//cfg, err := config.NewJSON(configFile)
		cfg, err := config.New(configFile)
		if err != nil {
			panic(err)
		}

		// we parse and verify nothing is missing to build a policy
		err = pki.IncludeConfig(cfg)
		if err != nil {
			panic(err)
		}
	}

	if len(*debugFlag) > 0 {
		fmt.Printf("debug mode\n")
		err = pki.Debug(*debugFlag)
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}

	// now let's generate a file for each
	// templating.
	for _, node := range pki.Nodes() {
		// create a node local list of all pki servers
		fmt.Printf("generating '%s' (srv:%s clt:%s) in '%s'\n", node.Name, node.ServerUUID, node.ClientUUID, node.Path)

		err := pki.BuildNodePolicy(node)
		if err != nil {
			panic(err)
		}

		err = node.generateFiles()
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("done\n")
	os.Exit(0)
}