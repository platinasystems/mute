// Copyright © 2021-2022 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

package mute

import (
	"log"
	"strings"
	"testing"
)

func compare(t *testing.T, printer Printer, want string) {
	builder := Off(printer).(*log.Logger).Writer().(*strings.Builder)
	defer builder.Reset()
	t.Helper()
	if got := builder.String(); got != want {
		t.Errorf("got: %q; want %q", got, want)
	}
}

func Test(t *testing.T) {
	errata := Off(log.New(new(strings.Builder), "error ", 0))
	info := On(log.New(new(strings.Builder), "info ", 0))

	errata.Print("foo")
	compare(t, errata, "error foo\n")

	info.Print("bar")
	compare(t, info, "")

	info = Off(info)
	info.Print("foobar")
	compare(t, info, "info foobar\n")
}

func TestExample(t *testing.T) {
	if !testing.Verbose() {
		t.Skip()
	}
	Example()
}
