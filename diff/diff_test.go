package diff

import (
    "testing"
)

func TestSomething(t *testing.T) {
    s1, s2 := "ABCABBA", "CBABAC"

    result := Diff(s1, s2)

    if len(result) == 0 {
        t.Error("Array cannot be empty !")
    }
}
