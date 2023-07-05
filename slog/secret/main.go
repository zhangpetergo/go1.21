// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log/slog"
	"os"
)

type People struct {
	name string
	Token
}

// A token is a secret value that grants permissions.
type Token string

// LogValue implements slog.LogValuer.
// It avoids revealing the token.
func (Token) LogValue() slog.Value {
	return slog.StringValue("REDACTED_TOKEN")
}

// This example demonstrates a Value that replaces itself
// with an alternative representation to avoid revealing secrets.
func main() {
	t := Token("shhhh!")
	p := People{
		name:  "peter",
		Token: t,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("permission granted", "user", "Perry", "token", t)
	logger.Info("permission granted", "user", "Perry", "people", p)

	// Output:
	// level=INFO msg="permission granted" user=Perry token=REDACTED_TOKEN
}
