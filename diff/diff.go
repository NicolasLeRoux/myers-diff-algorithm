package diff

type Action int
const (
    Add Action = 0
    Delete Action = 1
    Skip Action = 2
)

type Command struct {
    // Add, delete and skip
    Action Action
    Char string
}

func Diff(s1, s2 string) []Command {
    return make([]Command, 0)
}
