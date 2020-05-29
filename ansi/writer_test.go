// SPDX-License-Identifier: MIT

package ansi

import (
	"io"
	"math"
	"os"
	"testing"
)

var _ io.Writer = &Writer{}

func TestWriter(t *testing.T) {
	w := NewWriter(os.Stdout)

	for i := uint8(0); i < math.MaxUint8; i++ {
		w.Color256(i, 255-i)
		w.Printf("FColor(%d),BColor(%d)", i, 255-i)
		w.WriteESC(CSI(ResetCode))
		w.Println()
	}
}
