package go_intro_tests

import (
	"testing"
	"strings"

	"github.com/drafailov/dimitar_rafailov_go_testing/gostring"
)

func TestStr(t *testing.T) {
	var StrTest = [] int{0, 1, 7, 21, 66, 183, 1093, 1037423}
	for _, e := range StrTest {
		received := gostring.Str(e)

		// testing if length is equal to entered
		if len(received) != e {
			t.Errorf("String lenght:(%v); expected:(%v)", len(received), e)
		}

		// testing for unique value if length is more then 2
		compare := gostring.Str(e)
		if strings.Join(received, "") == strings.Join(compare, "") && len(received) > 2 {
			t.Errorf("String from length(%v) is not unique %v", e, received)
		}
	}
}
