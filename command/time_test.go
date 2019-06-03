package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestTimeCommand_implement(t *testing.T) {
	var _ cli.Command = &TimeCommand{}
}
