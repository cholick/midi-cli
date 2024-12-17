package cli_test

import (
	"flag"
	"testing"

	"github.com/cholick/midi-cli/internal/cli"
	"github.com/cholick/midi-cli/internal/midifakes"
	"github.com/cholick/midi-cli/internal/ui"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProgramChange(t *testing.T) {
	fakeOut := &midifakes.FakeOut{}
	fakeOpener := &midifakes.FakeOpener{}
	fakeOpener.NewOutForPortReturns(fakeOut, nil)

	out := ui.NewOutputForTesting()

	cmd := cli.NewPCCommand(fakeOpener, out)
	cmd.SetOut(out.StdOut)
	cmd.SetErr(out.ErrOut)

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
	fakeOut := &midifakes.FakeOut{}
	fakeOpener := &midifakes.FakeOpener{}
	fakeOpener.NewOutForPortReturns(fakeOut, nil)

	out := ui.NewOutputForTesting()

	cmd := cli.NewPCCommand(fakeOpener, out)
	cmd.SetOut(out.StdOut)
	cmd.SetErr(out.ErrOut)

	cmd.SetArgs([]string{
		"--port", "testPort", "--number", "-1"},
	)
	flag.Parse()

	err := cmd.Execute()
	require.Error(t, err)
	require.ErrorContains(t, err, "number")
}
