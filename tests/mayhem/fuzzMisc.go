package fuzzMisc

import "strconv"
import "github.com/rancher/elemental/tests/e2e/helpers/misc"
import fuzz "github.com/AdaLogics/go-fuzz-headers"


func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) > 2 {
        num, _ = strconv.Atoi(string(bytes[0]))
        bytes = bytes[1:]

        switch num {
    
        case 0:
            fuzzConsumerString := fuzz.NewConsumer(bytes)
            var content string
            err := fuzzConsumerString.CreateSlice(&content)
            if err != nil {
                return 0
            }

            fuzzConsumerInt := fuzz.NewConsumer(bytes)
            var number int
            err = fuzzConsumerInt.CreateSlice(&number)
            if err != nil {
                return 0
            }

            misc.GetServerId(content, number)
            return 0

        case 1:
            fuzzConsumerString := fuzz.NewConsumer(bytes)
            var content1 string
            var content2 string
            err := fuzzConsumerString.CreateSlice(&content1)
            if err != nil {
                return 0
            }
            err = fuzzConsumerString.CreateSlice(&content2)
            if err != nil {
                return 0
            }

            misc.AddSelector(content1, content2)
            return 0

        case 2:
            fuzzConsumerString := fuzz.NewConsumer(bytes)
            var content1 string
            var content2 string
            err := fuzzConsumerString.CreateSlice(&content1)
            if err != nil {
                return 0
            }
            err = fuzzConsumerString.CreateSlice(&content2)
            if err != nil {
                return 0
            }

            misc.ConcateFiles(content1, content2, bytes)
            return 0

        case 3:
            fuzzConsumerString := fuzz.NewConsumer(bytes)
            var content1 string
            var content2 string
            err := fuzzConsumerString.CreateSlice(&content1)
            if err != nil {
                return 0
            }
            err = fuzzConsumerString.CreateSlice(&content2)
            if err != nil {
                return 0
            }

            misc.TrimStringFromChar(content1, content2)
            return 0

        default:
            fuzzConsumerString := fuzz.NewConsumer(bytes)
            var content1 string
            var content2 string
            err := fuzzConsumerString.CreateSlice(&content1)
            if err != nil {
                return 0
            }
            err = fuzzConsumerString.CreateSlice(&content2)
            if err != nil {
                return 0
            }

            fuzzConsumerInt := fuzz.NewConsumer(bytes)
            var number int
            err = fuzzConsumerInt.CreateSlice(&number)
            if err != nil {
                return 0
            }

            misc.AddNode(content1, content2, number)
            return 0

        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}