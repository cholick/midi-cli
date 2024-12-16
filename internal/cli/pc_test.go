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

func TestProgramChange(t *testing.T) {
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}

	fakeOut := &midifakes.FakeOut{}
	fakeOpener := &midifakes.FakeOpener{}
	fakeOpener.NewOutForPortReturns(fakeOut, nil)
	out := ui.NewOutput(stdOut, stdErr)

	cmd := cli.NewPCCommand(fakeOpener, out)
	cmd.SetOut(stdOut)
	cmd.SetErr(stdErr)

	cmd.SetArgs([]string{
		"--port", "testPort", "--number", "4", "--channel", "9"},
	)
	flag.Parse()

	err := cmd.Execute()
	require.NoError(t, err)
	require.Equal(t, 1, fakeOut.ProgramChangeCallCount())

	num, channel := fakeOut.ProgramChangeArgsForCall(0)
	assert.Equal(t, 4, num)
	assert.Equal(t, 9, channel)
}

func TestProgramChangeNumberValidation(t *testing.T) {
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}

	fakeOut := &midifakes.FakeOut{}
	fakeOpener := &midifakes.FakeOpener{}
	fakeOpener.NewOutForPortReturns(fakeOut, nil)
	out := ui.NewOutput(stdOut, stdErr)

	cmd := cli.NewPCCommand(fakeOpener, out)
	cmd.SetOut(stdOut)
	cmd.SetErr(stdErr)

	cmd.SetArgs([]string{
		"--port", "testPort", "--number", "-1"},
	)
	flag.Parse()

	err := cmd.Execute()
	require.Error(t, err)
	require.ErrorContains(t, err, "number")
}
