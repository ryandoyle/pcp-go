package pmda

import (
	"testing"
	"os/exec"
	"github.com/stretchr/testify/assert"
	"os"
	"bytes"
	"regexp"
	"fmt"
)

func TestPmda_LoadsCorrectly(t *testing.T)  {
	dbpmda_output := runPmda(t, "status\n")

	assertPmdaLoaded(t, "gotestpmda", dbpmda_output)
}

func TestPmda_FetchesLong(t *testing.T) {
	dbpmda_output := runPmda(t, "status\n")

	assert.Equal(t, "asd", dbpmda_output)
}

func assertPmdaLoaded(t *testing.T, want string, output []byte)  {
	to_match := fmt.Sprintf("PMDA:.+%s", want)
	matched, err := regexp.Match(to_match, output)

	assert.NoError(t, err)
	assert.True(t, matched, fmt.Sprintf("tried to match \"%s\" from \"%s\" but couldn't find it", to_match, output))
}

func runPmda(t *testing.T, dbpmda_command string) []byte {
	buildPmda(t)
	command := exec.Command("dbpmda")
	buffer := bytes.NewBuffer([]byte("open pipe gotestpmda -d 50\n"))
	_, err := buffer.Write([]byte(dbpmda_command))
	assert.NoError(t, err)

	command.Stdin = buffer

	stdout, err := command.Output()
	assert.NoError(t, err)

	return stdout
}

func buildPmda(t *testing.T) {
	command := exec.Command("go", "build", "test/main/gotestpmda.go")
	command.Stderr = os.Stderr
	run := command.Run()

	assert.NoError(t, run)
}
