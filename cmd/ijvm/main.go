package main

import (
	"fmt"

	"github.com/gytsen/gojvm/pkg/binary"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

var (
	flagDebug bool
	flagInfo  bool
)

func main() {
	flag.BoolVarP(&flagInfo, "info", "i", false, "enable Info level logging")
	flag.BoolVarP(&flagDebug, "debug", "d", false, "enable Debug level logging")

	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("No binary selected to execute")
	} else if len(args) > 1 {
		log.Fatal("For now only single binary executions are supported")
	}

	if flagDebug {
		log.SetLevel(log.DebugLevel)
	} else if flagInfo {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}

	fileName := args[0]
	f, err := binary.Open(fileName)
	if err != nil {
		log.WithError(err).Fatal("failed to open binary")
	}

	for _, block := range f.Blocks {
		fmt.Printf("the size of this block is %d\n", block.Header.Size)
	}
}
