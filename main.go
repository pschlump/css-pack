// Copyright 2015 Philip Schlump
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Pack css Remove extra blank lines.

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/pschlump/css/scanner"
)

var opts struct {
	CssFileName       string `short:"i" long:"input"      description:"Path to css file"                  default:""`
	SourceMapFileName string `short:"s" long:"sourcemap"  description:"Path to css file"                  default:""`
	OutFileName       string `short:"o" long:"output"     description:"Path packed output to css file"    default:"out.css"`
	Deps              string `short:"d" long:"deps"       description:"Path packed output to css file"    default:""`
	Debug             bool   `short:"D" long:"debug"      description:"debug flag"                        default:"false"`
}

func main() {

	_, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		fmt.Printf("Invalid Command Line: %s\n", err)
		os.Exit(1)
	}

	if opts.CssFileName == "" {
		fmt.Printf("Usage test1 -s File.css\n")
		os.Exit(1)
	}

	myCSS, err := ioutil.ReadFile(opts.CssFileName)
	if err != nil {
		fmt.Printf("Errror: %s\n", err)
		os.Exit(1)
	}

	o := ""
	d := ""
	atKeyword := 0
	atImport := 0

	s := scanner.New(string(myCSS))
	for {
		token := s.Next()
		if token.Type == scanner.TokenError {
			fmt.Printf("Syntax Error: %s\n", token)
			os.Exit(2)
		} else if token.Type == scanner.TokenEOF {
			break
		}

		if opts.Debug {
			fmt.Printf("%+v\n", token)
		}

		switch token.Type {

		case scanner.TokenComment:
		case scanner.TokenS:

		case scanner.TokenAtKeyword:
			o += token.Value
			atKeyword = 1
			if token.Value == "@import" {
				atImport = 1
			}
		case scanner.TokenString:
			o += token.Value
			if atImport == 1 {
				token.Type = scanner.TokenImportURI
				if opts.Deps != "" {
					d += "@import " + token.Value + "\n"
				}
			}
			atKeyword, atImport = 0, 0

		case scanner.TokenIdent:
			fallthrough
		case scanner.TokenHash:
			fallthrough
		case scanner.TokenNumber:
			fallthrough
		case scanner.TokenPercentage:
			fallthrough
		case scanner.TokenDimension:
			fallthrough
		case scanner.TokenUnicodeRange:
			fallthrough
		case scanner.TokenCDO:
			fallthrough
		case scanner.TokenCDC:
			fallthrough
		case scanner.TokenFunction:
			fallthrough
		case scanner.TokenIncludes:
			fallthrough
		case scanner.TokenDashMatch:
			fallthrough
		case scanner.TokenPrefixMatch:
			fallthrough
		case scanner.TokenSuffixMatch:
			fallthrough
		case scanner.TokenSubstringMatch:
			fallthrough
		case scanner.TokenBOM:
			o += token.Value
		case scanner.TokenChar:
			if token.Value == "}" {
				o += "}\n"
			} else {
				o += token.Value
			}

		case scanner.TokenURI:
			o += token.Value
			if opts.Deps != "" {
				d += token.Value + "\n"
			}
		}
	}
	// fmt.Printf("s=%s\n", o)
	ioutil.WriteFile(opts.OutFileName, []byte(o), 0644)
	if opts.Deps != "" {

		ioutil.WriteFile(opts.Deps, []byte(d), 0644)
	}

	_, _ = atKeyword, atImport
}
