package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"sponsortext/internal"
)

func TestParseSponsorTextVariables(t *testing.T) {
	tests := []struct {
		name  string
		given string
		want  map[string]string
	}{
		{
			name:  "parses variables with suffix",
			given: "$vars:discord=https://discord.gg/vx4AKRfj;provider=bf2hub.com$",
			want: map[string]string{
				"discord":  "https://discord.gg/vx4AKRfj",
				"provider": "bf2hub.com",
			},
		},
		{
			name:  "parses variables without suffix",
			given: "$vars:discord=https://discord.gg/vx4AKRfj;provider=bf2hub.com",
			want: map[string]string{
				"discord":  "https://discord.gg/vx4AKRfj",
				"provider": "bf2hub.com",
			},
		},
		{
			name:  "skips non-variables prefix",
			given: "Join our event this Sunday! $vars:discord=https://discord.gg/vx4AKRfj;provider=bf2hub.com$",
			want: map[string]string{
				"discord":  "https://discord.gg/vx4AKRfj",
				"provider": "bf2hub.com",
			},
		},
		{
			name:  "stops before non-variables suffix",
			given: "$vars:discord=https://discord.gg/vx4AKRfj;provider=bf2hub.com$ Apply to become an admin today!",
			want: map[string]string{
				"discord":  "https://discord.gg/vx4AKRfj",
				"provider": "bf2hub.com",
			},
		},
		{
			name:  "reads escaped syntax characters",
			given: "$vars:\\$trange\\=key=https://example.com?query\\=\\$start\\;end$",
			want: map[string]string{
				"$trange=key": "https://example.com?query=$start;end",
			},
		},
		{
			name:  "reads escaped whitespaces",
			given: "$vars:four-spaces=\\ \\ \\ \\ $",
			want: map[string]string{
				"four-spaces": "    ",
			},
		},
		{
			name:  "omits irrelevant whitespaces",
			given: "$vars: trimmed = one  two  three ; stri pped=four $",
			want: map[string]string{
				"trimmed":  "one two three",
				"stripped": "four",
			},
		},
		{
			name:  "ignores incomplete key",
			given: "$vars:discord=https://discord.gg/vx4AKRfj;website$",
			want: map[string]string{
				"discord": "https://discord.gg/vx4AKRfj",
			},
		},
		{
			name:  "includes key with no value",
			given: "$vars:discord=https://discord.gg/vx4AKRfj;teamspeak=$",
			want: map[string]string{
				"discord":   "https://discord.gg/vx4AKRfj",
				"teamspeak": "",
			},
		},
		{
			name:  "returns empty map for empty variables section",
			given: "$vars:$",
			want:  map[string]string{},
		},
		{
			name:  "returns nil for sponsor text without variables",
			given: "Our server is the best!",
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// WHEN
			variables := internal.ParseSponsorTextVariables(tt.given)

			// THEN
			assert.Equal(t, tt.want, variables)
		})
	}
}
