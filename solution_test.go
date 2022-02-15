package solution

import (
	"testing"
)

func Test_GetMessage(t *testing.T) {
	res := GetMessage()
	want := "Hello \U0001f5fa\ufe0f !"
	if res != want {
		t.Error("Sprint ", res, want)
	}
}
