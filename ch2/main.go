package ch2

import (
    "log"
    "math"
    "math/rand"
)

func Main() {
    //var q []string
    //q = append(q, "hello")
    //q = append(q, "ewq")
    //q = append(q, "qwe")
    //q = append(q, "weq")
    //res := Shuffle(q)
    //log.Println(res)

    log.Printf("Exponentiation result = %d \n", Exponentiation(7, 2))
}

func Shuffle(a []string) []string {
    maxA := len(a)
    for i := 0; i < maxA; i++ {
        j := rand.Intn(i + 1)
        a[j], a[i] = a[i], a[j]
    }

    return a
}

func Exponentiation(a float64, n int) int {
    return int(math.Pow(float64(n), a))
}
