package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const VERSION = "0.1.0"

var (
	ipv4  = regexp.MustCompile(`\s?([\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}\.[\d]{1,3})\s?`)
	email = regexp.MustCompile(`(\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,3})`)

	collection []string

	opts struct {
		format string
		unique bool
		ignore string
	}
)

func addUnique(value string) {
	for _, item := range collection {
		if item == value {
			return
		}
	}

	collection = append(collection, value)
}

func scan(regex *regexp.Regexp, data []byte) {
	matches := regex.FindAllStringSubmatch(string(data), -1)

	for _, submatches := range matches {
		match := submatches[1]

		if opts.ignore != "" && strings.Contains(opts.ignore, match) {
			continue
		}

		if opts.unique {
			addUnique(match)
		} else {
			fmt.Println(match)
		}
	}
}

func init() {
	var version bool

	flag.StringVar(&opts.format, "f", "", "Extraction format: (ipv4, email)")
	flag.StringVar(&opts.ignore, "ignore", "", "List of ignore values")
	flag.BoolVar(&opts.unique, "uniq", false, "Return only unique matches")
	flag.BoolVar(&version, "v", false, "Print version")
	flag.Parse()

	if version {
		fmt.Printf("xtract v%s\n", VERSION)
		os.Exit(0)
	}

	if opts.format == "" {
		flag.Usage()
		os.Exit(0)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		switch opts.format {
		case "ipv4":
			scan(ipv4, scanner.Bytes())
		case "email":
			scan(email, scanner.Bytes())
		}
	}

	if opts.unique {
		for _, str := range collection {
			fmt.Println(str)
		}
	}
}
