package main

import (
	"os"
    "fmt"
    "github.com/tsyeyuanfeng/go-bsdp/bsdp"
    "gopkg.in/alecthomas/kingpin.v2"
)

var (
	gobsdp = kingpin.New("go-bsdp", "An binary diff&patch library based on bsdiff algorithm(v4.3).")
    diff = gobsdp.Command("diff", "Generates binary patch")
    patch = gobsdp.Command("patch", "Applies binary patch")
    doldfile = diff.Arg("oldfile", "The old file").Required().String()
    dnewfile = diff.Arg("newfile", "The new file").Required().String()
    dpatchfile = diff.Arg("patchfile", "The patch file").Required().String()
    poldfile = patch.Arg("oldfile", "The old file").Required().String()
    pnewfile = patch.Arg("newfile", "The new file").Required().String()
    ppatchfile = patch.Arg("patchfile", "The patch file").Required().String()
)

func main() {
    switch kingpin.MustParse(gobsdp.Parse(os.Args[1:])) {
        case diff.FullCommand():
            err := bsdp.Diff(*doldfile, *dnewfile, *dpatchfile)
            if err != nil {
                fmt.Printf("%v", err)
            }

        case patch.FullCommand():
            err := bsdp.Patch(*poldfile, *pnewfile, *ppatchfile)
            if err != nil {
                fmt.Printf("%v", err)
            }
    }
}