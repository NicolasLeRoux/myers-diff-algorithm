package diff

import (
    "testing"
)

func TestSomething(t *testing.T) {
    s1, s2 := "ABCABBA", "CBABAC"
    expectedCommands := []Command{
        {Delete, "A"},
        {Add, "C"},
        {Skip, "B"},
        {Delete, "C"},
        {Skip, "A"},
        {Skip, "B"},
        {Delete, "B"},
        {Skip, "A"},
        {Add, "C"},
    }

    commands := Diff(s1, s2)

    if len(commands) == 0 {
        t.Error("Array cannot be empty !")
    }

    if len(commands) != len(expectedCommands) {
        t.Error("The length of the commands must be the same.")
    }

    for i, cmd := range commands {
        if cmd.Action != expectedCommands[i].Action {
            t.Errorf("The action at index %q is wrong.", i)
        }
    }
}
