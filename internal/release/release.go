package release

import (
	"fmt"
	"os"

	"github.com/afwilliams/gotfiles/internal/git"
)

func Run() (string, error) {
	if err := git.IsGitInstalled(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if err := git.HasGitRepo(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	return "v0.0.0", nil
}
