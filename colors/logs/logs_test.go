// Copyright 2018 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package logs

import "testing"

func TestLogs(t *testing.T) {
	Error.Println("ERROR")
	Warn.Println("WARN")
	Info.Println("INFO")
	Success.Println("SUCCESS")
}
