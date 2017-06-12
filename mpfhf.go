package main

import (
    "flag"
    "fmt"
    "strings"
    "strconv"
)

func toBinArray(message string) []int {
    binString := ""
    var binArray []int

    for _, c := range message {
        binString = fmt.Sprintf("%s%.8b", binString, c)
    }

    for _, i := range strings.Split(binString, "") {
        j, err := strconv.Atoi(i)
        if err != nil {
            panic(err)
        }
        binArray = append(binArray, j)
    }

    return binArray
}

func flip(field []int, index int) {
    if field[index] == 0 {
        field[index] = 1
    } else {
        field[index] = 0
    }
}

func invert(field []int) {
    for i, _ := range field {
        flip(field, i)
    }
}

func screw(a []int, b []int, M_pos int, half bool) {
    count := 0
    if half {
        count = len(a)/2
    } else {
        count = len(a)
    }

    for i := 0; i<count; i++ {
        flip(b, ((i*M_pos)%len(b)))
    }
}

func main() {
    message := flag.String("message", "MP kicks ass!", "message")
    bits := flag.Int("bits", 64, "num of bits")
    flag.Parse()

    M := toBinArray(*message)
    fmt.Println(M)
    R := make([]int, *bits)
    S := make([]int, 1)

    M_pos := 0
    step := 0
    R_len := *bits

    for M_pos < len(M) {
        if M[M_pos] == 0 {
            S = append(S, 0)
            screw(S, R, M_pos, false)
            if R[M_pos%R_len] == 0 {
                flip(R, M_pos%R_len)
                if M_pos != 0 {
                    M_pos -= 1
                }
            } else {
                flip(R, M_pos%R_len)
                invert(S)
            }
        } else {
            screw(S, R, M_pos, true)
            if R[M_pos%R_len] == S[M_pos%len(S)] {
                S = append(S, 0)
                screw(R, S, M_pos, false)
            } else {
                flip(R, M_pos%R_len)
            }
        }
        M_pos += 1
        step += 1
    }

    for _, i := range R {
        fmt.Print(i)
    }
    fmt.Print("\n")
}
