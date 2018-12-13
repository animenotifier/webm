package main

import "flag"

func main() {
	flag.Parse()

	// The command we're going to perform
	command := flag.Arg(0)

	switch command {
	case "encode":
		println("encode")
	default:
		help()
	}
}
