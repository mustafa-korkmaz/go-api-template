package secure_test

import (
	"testing"

	"github.com/mustafa-korkmaz/goapitemplate/pkg/utl/secure"

	"github.com/magiconair/properties/assert"
)

func TestIsPasswordSecure(t *testing.T) {

	cases := []struct {
		name string
		pass string
		want bool
	}{
		{
			name: "Insecure 1",
			pass: "notSec",
			want: false,
		},
		{
			name: "Insecure 2",
			pass: "MuteEsra1907",
			want: false,
		},
		{
			name: "Secure",
			pass: "MuteEsra_1907",
			want: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got := secure.IsPasswordSecure(tt.pass)
			assert.Equal(t, tt.want, got)
		})
	}
}
