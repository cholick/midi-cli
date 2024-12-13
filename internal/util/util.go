package util

import "github.com/spf13/cobra"

// Sigh; I consider a package named "util" a code smell, but leaving these here for now...
// todo: Pull these out once a structure makes sense

// Annoying this is needed... https://github.com/spf13/cobra/issues/216
func ParentPreRun(cmd *cobra.Command, args []string) error {
	parent := cmd.Parent()
	if parent != nil {
		if parent.PersistentPreRun != nil {
			parent.PersistentPreRun(parent, args)
		}
		if parent.PersistentPreRunE != nil {
			err := parent.PersistentPreRunE(parent, args)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
