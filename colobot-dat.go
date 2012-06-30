package main

import (
	"flag"
	"fmt"
	"github.com/csi235/colobot-dat/cdat"
	"io"
	"os"
)

var format = flag.String("H", "colobot", "the container encoding to use")
var list = flag.Bool("l", false, "list files in a container")
var verbose = flag.Bool("v", false, "output extra information")

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("no file specified")
		flag.PrintDefaults()
		os.Exit(1)
	} else if flag.NArg() > 1 {
		fmt.Println("too many arguments")
		flag.PrintDefaults()
		os.Exit(1)
	}
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch {
	case *list:
		if err := listContainer(file); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		if err := extractContainer(file); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func extractContainer(file *os.File) error {
	c, err := cdat.New(file, getCodec(*format))
	if err != nil {
		return fmt.Errorf("extract %s: %v", file.Name(), err)
	}

	for _, v := range c.Files {
		if *verbose {
			fmt.Println(v.Name)
		}
		dst, err := os.OpenFile(v.Name, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
		if os.IsExist(err) {
			return fmt.Errorf("extract %s/%s: %v", file.Name(), v.Name, err)
		}
		fmt.Println(file.Name())
		if _, err := io.Copy(dst, v); err != nil {
			return fmt.Errorf("extract %s/%s: %v", file.Name(), v.Name, err)
		}
	}

	return nil
}

func getCodec(codec string) cdat.Codec {
	switch codec {
	case "ceebot":
		return cdat.CeebotCodec
	case "ceebot-demo":
		return cdat.CeebotDemoCodec
	case "colobot":
		return cdat.ColobotCodec
	case "colobot-demo":
		return cdat.ColobotDemoCodec
	case "none":
		return cdat.IdentityCodec
	}
	return nil
}

func listContainer(file *os.File) error {
	c, err := cdat.New(file, getCodec(*format))
	if err != nil {
		return fmt.Errorf("list %s: %v", file.Name(), err)
	}

	for _, v := range c.Files {
		if *verbose {
			fmt.Printf("%s %d-%d\n", v.Name, v.Offset, v.Offset+v.Length-1)
		} else {
			fmt.Println(v.Name)
		}
	}

	return nil
}
