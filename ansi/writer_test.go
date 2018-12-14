// Copyright 2014 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package ansi

import (
	"io"
	"os"
	"testing"
)

var _ io.Writer = &Writer{}

func TestWriter(t *testing.T) {
	w := NewWriter(os.Stdout)

	for i := 0; i < 256; i += 10 {
		w.Color256(i, 255-i)
		w.Printf("FColor(%d),BColor(%d)", i, 255-i)
		w.WriteString(Reset)
		w.Println()
	}

	w.WriteString(Reset) //.Move(50, 100)
}
