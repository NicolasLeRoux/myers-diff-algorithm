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
    length := len(s1) + len(s2)
    mapV := make(map[int]int)
    mapV[1] = 0

    for d := 0; d <= length; d++ {
        for k := -d; k <= d; k += 2 {
            // down or right?
            down := k == -d || (k != d && mapV[k - 1] < mapV[k + 1])

            var kPrev int
            if down {
                kPrev = k + 1
            } else {
                kPrev = k - 1
            }

            // start point
            var xStart int = mapV[ kPrev ]
            var yStart int = xStart - kPrev

            // mid point
            var xMid int
            if down {
                xMid = xStart
            } else {
                xMid = xStart + 1
            }
            var yMid int = xMid - k

            // end point
            var xEnd int = xMid
            var yEnd int = yMid

            // follow diagonal
            var snake int = 0
            for xEnd < N && yEnd < M && s1[xEnd] == s2[yEnd] {
                xEnd++
                yEnd++
                snake++
            }

            // save end point
            mapV[k] = xEnd

            // check for solution
            if ( xEnd >= N && yEnd >= M ) {}/* solution has been found */
        }
    }

    return make([]Command, 0)
}
