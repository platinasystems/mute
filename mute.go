// Copyright © 2021-2022 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// Package mute provides an encapsulation with no-op Print methods.
package mute

type Printer interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

// Returns a no-op'd Printer.
func On(printer Printer) Printer { return Mute{Off(printer)} }

// Returns the encapsulated Printer or itself if not currently no-op.
func Off(printer Printer) Printer {
	if m, ok := printer.(Mute); ok {
		return m.Printer
	}
	return printer
}

// true if muted
func Status(printer Printer) bool {
	_, ok := printer.(Mute)
	return ok
}

type Mute struct{ Printer }

func (Mute) Print(v ...interface{}) {}

func (Mute) Printf(format string, v ...interface{}) {}

func (Mute) Println(v ...interface{}) {}
