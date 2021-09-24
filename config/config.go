//go:build go1.16
// +build go1.16

package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// [name]:Options
type Config map[string]Options

func (c Config) Add(name string, o Options) (err error) {
	_, ok := c[name]
	if ok {
		err = fmt.Errorf("already a value")
		return
	}
	c[name] = o
	return
}

func NewConfig() (cf Config) {
	cf = make(Config)
	return
}

func New(filename string) (cf Config, err error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	ok := json.Valid(file)
	switch ok {
	case true:
		cf, err = parseJSON(file)
	case false:
		cf, err = parseConf(file)
	}
	return
}

func parseConf(payload []byte) (cf Config, err error) {
	// Setup the input
	is := antlr.NewInputStream(string(payload))

	// Create the Lexer
	lexer := NewpkscLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := NewpkscParser(stream)
	//parser := NewParser(hmackey)
	parser := NewParser()

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(parser, p.Config())

	cf = parser.Config()
	return
}

func parseJSON(payload []byte) (cf Config, err error) {
	cf = make(Config)
	err = json.Unmarshal(payload, &cf)
	return
}

func NewFile(filename string) (cf Config, err error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	cf, err = parseConf(file)
	return
}

func NewJSON(filename string) (cf Config, err error) {
	cf = make(Config)

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	cf, err = parseJSON(file)
	return
}
