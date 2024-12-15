package ui_test

import (
	"bytes"
	"testing"

	"github.com/cholick/midi-cli/internal/ui"
	"github.com/stretchr/testify/assert"
)

func TestDebug(t *testing.T) {
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}

	o := ui.NewOutput(stdOut, stdErr)
	o.SetDebug(true)

	o.Debug("foo", "bar")
	o.Debugf("%s: %s", "hello", "🍕")

	assert.Equal(t, "foo bar\nhello: 🍕\n", stdOut.String())

	assert.Empty(t, stdErr.String())
}

func TestDebugOff(t *testing.T) {
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}

	o := ui.NewOutput(stdOut, stdErr)
	o.SetDebug(false)

	o.Debug("foo", "bar")

	assert.Len(t, stdOut.String(), 0)
}

func TestInfo(t *testing.T) {
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}

	o := ui.NewOutput(stdOut, stdErr)

	o.Info("baz", "qux")
	o.Infof("%s: %s", "goodbye", "🍔")

	assert.Equal(t, "baz qux\ngoodbye: 🍔\n", stdOut.String())

	assert.Empty(t, stdErr.String())
}

func TestError(t *testing.T) {
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}

	o := ui.NewOutput(stdOut, stdErr)

	o.Error("quux", "corge")
	o.Errorf("%s: %s", "salud", "🥂")

	assert.Empty(t, stdOut.String())

	assert.Equal(t, "quux corge\nsalud: 🥂\n", stdErr.String())
}
