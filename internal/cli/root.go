package cli

import (
	"os"

	"github.com/cholick/midi-cli/internal/cli/note"
	"github.com/cholick/midi-cli/internal/cli/port"
	"github.com/cholick/midi-cli/internal/ui"
	"github.com/spf13/cobra"
)

func NewRootCommand() (*cobra.Command, error) {
	var verbose bool
	con := ui.NewOutput(os.Stdout, os.Stderr)

	root := &cobra.Command{
		Use:   "midi-cli",
		Short: "Send MIDI messages",

		// Skip completions, not a feature I use
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
		// using error to indicate lots of problems, usage often doesn't make sense
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// need to do this here, due to flag parsing lifecycle
			con.SetDebug(verbose)
		},
	}

	root.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")

	portCmd := port.NewPortCommand()
	portCmd.AddCommand(port.NewListCommand(con))
	root.AddCommand(portCmd)

	noteCmd := note.NewNoteCommand()
	noteCmd.AddCommand(note.NewOnCommand(con))
	noteCmd.AddCommand(note.NewOffCommand(con))
	root.AddCommand(noteCmd)

	return root, nil
}