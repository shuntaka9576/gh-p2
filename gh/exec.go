package gh

import (
	"fmt"
	"strings"

	"github.com/cli/go-gh/v2"
)

func execGh(args ...string) ([]byte, error) {
	stdout, stderr, err := gh.Exec(args...)
	if err != nil {
		msg := strings.TrimSpace(stderr.String())
		if msg == "" {
			return nil, fmt.Errorf("gh %s: %w", strings.Join(args, " "), err)
		}
		return nil, fmt.Errorf("gh %s: %w: %s", strings.Join(args, " "), err, msg)
	}
	return stdout.Bytes(), nil
}
