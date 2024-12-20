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

func TestBankSelect(t *testing.T) {
	fakeOut := &midifakes.FakeOut{}
	fakeOpener := &midifakes.FakeOpener{}
	fakeOpener.NewOutForPortReturns(fakeOut, nil)

	out := ui.NewOutputForTesting()
	cmd := cli.NewBSCommand(fakeOpener, out)

	cmd.SetOut(out.StdOut)
	cmd.SetErr(out.ErrOut)

	cmd.SetArgs([]string{
		"--port", "testPort", "--value", "13", "--channel", "2",
	})
	flag.Parse()

	err := cmd.Execute()
	require.NoError(t, err)
	require.Equal(t, 1, fakeOut.ControlChangeCallCount())

	num, val, channel := fakeOut.ControlChangeArgsForCall(0)
	assert.Equal(t, 0, num)
	assert.Equal(t, 13, val)
	assert.Equal(t, 2, channel)
}
