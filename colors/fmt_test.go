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
