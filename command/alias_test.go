package command

import (
	"testing"

	"github.com/cli/cli/test"
)

func TestAliasSet_existing_alias(t *testing.T) {
	initBlankContext("", "OWNER/REPO", "trunk")

	// TODO setup

	_, err := RunCommand("alias set co pr checkout")

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestAliasSet_arg_processing(t *testing.T) {
	initBlankContext("", "OWNER/REPO", "trunk")
	cases := []struct {
		Cmd          string
		ExpectedLine string
	}{
		{"alias set co pr checkout", "- Adding alias for co = pr checkout"},
		{`alias set il "issue list"`, "- Adding alias for il = issue list"},
	}

	for _, c := range cases {
		output, err := RunCommand(c.Cmd)
		if err != nil {
			t.Fatalf("got unexpected error running %s: %s", c.Cmd, err)
		}

		test.ExpectLines(t, output.String(), c.ExpectedLine)

	}
}
