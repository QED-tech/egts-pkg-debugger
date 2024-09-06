package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"github.com/docopt/docopt-go"
	"github.com/kuznetsovin/egts-protocol/libs/egts"
)

func main() {
	usage := `EGTS Packages Debugger.

	Usage:
	  egts-packages-debugger decode <pkg>
	  egts-packages-debugger --version
	
	Options:
	  -h --help     Show this screen.
	  --version     Show version.`

	arguments, err := docopt.ParseDoc(usage)
	if err != nil {
		log.Fatalf("failed to parse doc arguments, err: %v", err)
	}

	pkg, err := arguments.String("<pkg>")
	if err != nil {
		log.Fatalf("failed to get egts pkg argument, err: %v", err)
	}

	bytes, err := hex.DecodeString(pkg)
	if err != nil {
		log.Fatalf("failed to decode hex string, err: %v", err)
	}

	egtsPackage := egts.Package{}
	_, err = egtsPackage.Decode(bytes)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.MarshalIndent(egtsPackage, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(b))
}
