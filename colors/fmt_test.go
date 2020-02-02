// SPDX-License-Identifier: MIT

package colors

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/issue9/assert"
)

func TestPrint(t *testing.T) {
	a := assert.New(t)

	_, err := Printf(Red, Blue, "Print::foreground:%v;background:%v", Red, Blue)
	a.NotError(err)
	_, err = Printf(Blue, Red, "Print::foreground:%v;background:%v\n", Blue, Red)
	a.NotError(err)

	_, err = Println(Cyan, Default, "Println::foreground:Cyan;background:Default")
	a.NotError(err)

	_, err = Print(Red, Blue, "Print::foreground:Red;background:Blue\n\n")
	a.NotError(err)
}

func TestFprint(t *testing.T) {
	a := assert.New(t)
	path := "./fprint.test"

	f, err := os.Create(path)
	a.NotError(err).NotNil(f)
	_, err = Fprint(f, Red, Green, "abc")
	a.NotError(err)
	f.Close()

	data, err := ioutil.ReadFile(path)
	a.NotError(err).NotNil(data)
	a.Equal(string(data), "abc")
}
