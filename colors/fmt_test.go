// SPDX-License-Identifier: MIT

package colors

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/issue9/assert"
)

func TestFprint(t *testing.T) {
	a := assert.New(t)
	path := "./fprint.test"

	f, err := os.Create(path)
	a.NotError(err).NotNil(f)
	_, err = Fprint(f, Bold, Red, Green, "abc")
	a.NotError(err)
	f.Close()

	data, err := ioutil.ReadFile(path)
	a.NotError(err).NotNil(data)
	a.Equal(string(data), "abc")
}

func TestFprintf(t *testing.T) {
	a := assert.New(t)

	a.Panic(func() {
		Fprintf(os.Stderr, -100, Red, Green, "test")
	})

	Fprintf(os.Stderr, Normal, Red, Green, "test")

	Fprintf(os.Stderr, Blink, Red, Green, "test")
}

func TestFprintln(t *testing.T) {
	a := assert.New(t)

	a.Panic(func() {
		Fprintln(os.Stderr, -100, Red, Green, "test")
	})

	Fprintln(os.Stderr, Normal, Red, Green, "test")

	Fprintln(os.Stderr, Blink, Red, Green, "test")
}
