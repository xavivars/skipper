package kubernetes

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHostsToRegex(t *testing.T) {
	for _, ti := range []struct {
		msg   string
		host  string
		regex string
	}{
		{
			msg:   "simple",
			host:  "simple.example.org",
			regex: "^(simple[.]example[.]org[.]?(:[0-9]+)?)$",
		},
		{
			msg:   "wildcard",
			host:  "*.example.org",
			regex: "^(*[.]example[.]org[.]?(:[0-9]+)?)$",
		},
	} {
		t.Run(ti.msg, func(t *testing.T) {
			regex := createHostRx(ti.host)
			require.Equal(t, ti.regex, regex)
			// maybe we should validate the generated regex and maybe add
			_, err := regexp.Compile(regex)
			require.NoError(t, err)
		})
	}
}
