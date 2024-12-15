package note_test

import (
	"bytes"
	"flag"
	"testing"

	"github.com/cholick/midi-cli/internal/cli/note"
	"github.com/cholick/midi-cli/internal/midifakes"
	"github.com/cholick/midi-cli/internal/ui"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateNote(t *testing.T) {
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}

	cmd := note.NewNoteCommand()
	cmd.SetOut(stdOut)
	cmd.SetErr(stdErr)

	opener := &midifakes.FakeOpener{}
	out := ui.NewOutput(stdOut, stdErr)

	cmd.AddCommand(note.NewOnCommand(opener, out))

	cmd.SetArgs([]string{
		"on", "--port", "testPort", "--note", "H4"},
	)
	flag.Parse()

	err := cmd.Execute()
	require.Error(t, err)

	assert.ErrorContains(t, err, "note")
	assert.ErrorContains(t, err, "H4")
}

func TestValidateVelocity(t *testing.T) {
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}

	cmd := note.NewNoteCommand()
	cmd.SetOut(stdOut)
	cmd.SetErr(stdErr)

	opener := &midifakes.FakeOpener{}
	out := ui.NewOutput(stdOut, stdErr)

	cmd.AddCommand(note.NewOnCommand(opener, out))

	cmd.SetArgs([]string{
		"on", "--port", "testPort", "--note", "C4", "-o", "-1"},
	)
	flag.Parse()

	err := cmd.Execute()
	require.Error(t, err)

	assert.ErrorContains(t, err, "velocity")
}

func TestValidateChannel(t *testing.T) {
	stdOut := &bytes.Buffer{}
	stdErr := &bytes.Buffer{}

	cmd := note.NewNoteCommand()
	cmd.SetOut(stdOut)
	cmd.SetErr(stdErr)

	opener := &midifakes.FakeOpener{}
	out := ui.NewOutput(stdOut, stdErr)

	cmd.AddCommand(note.NewOnCommand(opener, out))

	cmd.SetArgs([]string{
		"on", "--port", "testPort", "--note", "C4", "-c", "17"},
	)
	flag.Parse()

	err := cmd.Execute()
	require.Error(t, err)

	assert.ErrorContains(t, err, "channel")
}
