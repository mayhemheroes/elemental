package fuzz

import "strconv"
import "github.com/rancher/elemental/tests/e2e/helpers/misc"


func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) > 2 {
        num, _ = strconv.Atoi(string(bytes[0]))
        bytes = bytes[1:]

        switch num {
    
        case 0:
            content := string(bytes)
            misc.GetServerId(content, num)
            return 0

        case 1:
            content := string(bytes)
            misc.AddSelector("mayhem", content)
            return 0

        case 2:
            content := string(bytes)
            misc.ConcateFiles("mayhem", content, bytes)
            return 0

        case 3:
            content := string(bytes)
            misc.TrimStringFromChar("mayhem", content)
            return 0

        default:
            content := string(bytes)
            misc.AddNode("mayhem", num, content)
            return 0

        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}