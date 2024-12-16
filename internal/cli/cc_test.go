package cli_test

import (
	"bytes"
	"flag"
	"testing"

	"github.com/cholick/midi-cli/internal/cli"
	"github.com/cholick/midi-cli/internal/midifakes"
	"github.com/cholick/midi-cli/internal/ui"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestControlChange(t *testing.T) {
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}

	fakeOut := &midifakes.FakeOut{}
	fakeOpener := &midifakes.FakeOpener{}
	fakeOpener.NewOutForPortReturns(fakeOut, nil)
	out := ui.NewOutput(stdOut, stdErr)

	cmd := cli.NewCCCommand(fakeOpener, out)

	cmd.SetOut(stdOut)
	cmd.SetErr(stdErr)

	cmd.SetArgs([]string{
		"--port", "testPort", "--number", "2", "--value", "22", "--channel", "9"},
	)
	flag.Parse()

	output := &bytes.Buffer{}
	cmd.SetOut(output)
	cmd.SetErr(output)

	err := cmd.Execute()
	require.NoError(t, err)
	require.Equal(t, 1, fakeOut.ControlChangeCallCount())

	num, val, channel := fakeOut.ControlChangeArgsForCall(0)
	assert.Equal(t, 2, num)
	assert.Equal(t, 22, val)
	assert.Equal(t, 9, channel)
}

func TestControlChangeNumberValidation(t *testing.T) {
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}

	fakeOut := &midifakes.FakeOut{}
	fakeOpener := &midifakes.FakeOpener{}
	fakeOpener.NewOutForPortReturns(fakeOut, nil)
	out := ui.NewOutput(stdOut, stdErr)

	cmd := cli.NewCCCommand(fakeOpener, out)

	cmd.SetOut(stdOut)
	cmd.SetErr(stdErr)

	cmd.SetArgs([]string{
		"--port", "testPort", "--number", "130", "--value", "22", "--channel", "9"},
	)
	flag.Parse()

	output := &bytes.Buffer{}
	cmd.SetOut(output)
	cmd.SetErr(output)

	err := cmd.Execute()
	require.Error(t, err)
	require.ErrorContains(t, err, "number")
}

func TestControlChangeValueValidation(t *testing.T) {
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}

	fakeOut := &midifakes.FakeOut{}
	fakeOpener := &midifakes.FakeOpener{}
	fakeOpener.NewOutForPortReturns(fakeOut, nil)
	out := ui.NewOutput(stdOut, stdErr)

	cmd := cli.NewCCCommand(fakeOpener, out)

	cmd.SetOut(stdOut)
	cmd.SetErr(stdErr)

	cmd.SetArgs([]string{
		"--port", "testPort", "--number", "2", "--value", "2222", "--channel", "9"},
	)
	flag.Parse()

	output := &bytes.Buffer{}
	cmd.SetOut(output)
	cmd.SetErr(output)

	err := cmd.Execute()
	require.Error(t, err)
	require.ErrorContains(t, err, "value")
}
