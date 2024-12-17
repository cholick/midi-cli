package note_test

import (
	"flag"
	"testing"

	"github.com/cholick/midi-cli/internal/cli/note"
	"github.com/cholick/midi-cli/internal/midifakes"
	"github.com/cholick/midi-cli/internal/ui"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNotOff(t *testing.T) {
	cmd := note.NewNoteCommand()

	out := ui.NewOutputForTesting()
	cmd.SetOut(out.StdOut)
	cmd.SetErr(out.ErrOut)

	fakeOut := &midifakes.FakeOut{}
	fakeOpener := &midifakes.FakeOpener{}
	fakeOpener.NewOutForPortReturns(fakeOut, nil)

	cmd.AddCommand(note.NewOffCommand(fakeOpener, out))

	cmd.SetArgs([]string{
		"off", "--port", "testPort", "--note", "E4", "-o", "99", "-c", "14",
	})
	flag.Parse()

	err := cmd.Execute()
	require.NoError(t, err)
	require.Equal(t, 1, fakeOut.NoteOffCallCount())

	noteName, velocity, channel := fakeOut.NoteOffArgsForCall(0)
	assert.Equal(t, "E4", noteName)
	assert.Equal(t, 99, velocity)
	assert.Equal(t, 14, channel)
}
