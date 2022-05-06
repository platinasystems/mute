// Copyright © 2015-2022 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package printer

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/platinasystems/mute"
)

const Flags = log.Lshortfile

type Logger = log.Logger
type Printer = mute.Printer

var (
	Prog  = filepath.Base(os.Args[0])
	Alert = mute.Off(log.New(os.Stdout, Prog+":alert:", Flags))
	Error = mute.Off(log.New(os.Stdout, Prog+":error:", Flags))
	Warn  = mute.On(log.New(os.Stdout, Prog+":warning:", Flags))
	Info  = mute.On(log.New(os.Stdout, Prog+":info:", Flags))
	Debug = mute.On(log.New(os.Stdout, Prog+":debug:", Flags))
	Trace = mute.On(log.New(os.Stdout, Prog+":trace:", Flags))
)

func Name(p Printer) string {
	s := "*NOT LOGGER*"
	if l, ok := mute.Off(p).(*Logger); ok {
		s = strings.TrimSpace(l.Prefix())
		s = strings.TrimPrefix(s, Prog+":")
		s = strings.TrimSuffix(s, ":")
	}
	return s
}

func Redirect(p Printer, w io.Writer) {
	if l, ok := mute.Off(p).(*Logger); ok {
		l.SetOutput(w)
	}
}

func Reflag(p Printer, flags int) {
	if l, ok := mute.Off(p).(*Logger); ok {
		l.SetFlags(flags)
	}
}

func Rename(p Printer, name string) {
	if l, ok := mute.Off(p).(*Logger); ok {
		l.SetPrefix(name)
	}
}

func Restyle(p Printer, name string) {
	if l, ok := mute.Off(p).(*Logger); ok {
		l.SetOutput(os.Stdout)
		l.SetFlags(Flags)
		l.SetPrefix(fmt.Sprint(Prog, ":", name, ":"))
	}
}
