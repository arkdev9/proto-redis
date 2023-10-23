package cli_test

import (
	"os"
	"strings"
	"testing"

	"github.com/arkdev9/bad-redis/cli"
)

func TestCli(t *testing.T) {
	// Read the test file
	data, err := os.ReadFile("../data/test.txt")
	if err != nil {
		t.Errorf("Could not read test data file")
	}

	testString := string(data)
	splitStrings := strings.Split(testString, "\n")

	for i := 0; i < len(splitStrings); i += 2 {
		cmd := splitStrings[i]
		output := splitStrings[i+1]
		computedOutput := cli.Cli(cmd)
		if computedOutput != output {
			t.Errorf(
				"Invalid testcase\nCMD: %s\nExpected: %s\nComputed: %s\n",
				cmd,
				output,
				computedOutput)
		}
	}
}
