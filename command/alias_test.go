package command

import (
	"bytes"
	"testing"

	"github.com/cli/cli/internal/config"
	"github.com/cli/cli/test"
)

func TestAliasSet_existing_alias(t *testing.T) {
	cfg := `---
hosts:
  github.com:
    user: OWNER
    oauth_token: token123
aliases:
  co: pr checkout
`
	initBlankContext(cfg, "OWNER/REPO", "trunk")

	_, err := RunCommand("alias set co pr checkout")

	if err == nil {
		t.Fatal("expected error")
	}

	eq(t, err.Error(), "alias co already exists")
}

func TestAliasSet_arg_processing(t *testing.T) {
	initBlankContext("", "OWNER/REPO", "trunk")
	buf := bytes.NewBufferString("")
	defer config.StubWriteConfig(buf)()
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

func TestAliasSet_init_alias_cfg(t *testing.T) {
	cfg := `---
hosts:
  github.com:
    user: OWNER
    oauth_token: token123
`
	initBlankContext(cfg, "OWNER/REPO", "trunk")

	buf := bytes.NewBufferString("")
	defer config.StubWriteConfig(buf)()

	output, err := RunCommand("alias set cool story")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	expected := `hosts:
    github.com:
        user: OWNER
        oauth_token: token123
aliases:
    cool: story
`

	test.ExpectLines(t, output.String(), "Adding alias for cool = story", "Added alias.")
	eq(t, buf.String(), expected)
}

func TestAliasSet_existing_aliases(t *testing.T) {
	cfg := `---
hosts:
  github.com:
    user: OWNER
    oauth_token: token123
aliases:
    foo: bar
`
	initBlankContext(cfg, "OWNER/REPO", "trunk")

	buf := bytes.NewBufferString("")
	defer config.StubWriteConfig(buf)()

	output, err := RunCommand("alias set cool story")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	expected := `hosts:
    github.com:
        user: OWNER
        oauth_token: token123
aliases:
    foo: bar
    cool: story
`

	test.ExpectLines(t, output.String(), "Adding alias for cool = story", "Added alias.")
	eq(t, buf.String(), expected)

}
