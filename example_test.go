// Copyright © 2021-2022 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package mute

import (
	"flag"
	"log"
	"os"
)

var (
	verbose   = flag.Bool("verbose", false, "unmute info")
	quiet     = flag.Bool("quiet", false, "mute errors")
	immutable = log.New(os.Stdout, "error ", 0)
	errata    = Off(immutable)
	info      = On(log.New(os.Stdout, "info ", 0))
)

func Example() {
	if *verbose {
		info = Off(info)
	} else if *quiet {
		errata = On(errata)
	}
	immutable.Print("immutable")   // always printed
	errata.Print("mutable")        // printed unless quiet
	info.Println("hello", "world") // skipped unless verbose
	Off(info).Print("done")        // always printed
}
