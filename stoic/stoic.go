package stoic

import (
    "fmt"
    "reflect"
)

type Philosopher int

const (
    Epictetus Philosopher = iota
    gavin
    Seneca
)

func Quote(who Philosopher) string {
    fmt.Println("a: ", reflect.TypeOf(who))
    switch who {
    case Epictetus:
        return "First say to yourself what you would be; and do what you have to do"
    case Seneca:
        return "If a man knows not to which port he sails, No wind is favorable"
    }
    return "nothing"
}
