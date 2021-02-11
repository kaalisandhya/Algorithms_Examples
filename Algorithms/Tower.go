package main

import "fmt"

func hanoi(disk int, start string, end string, middle string) {
    if disk > 0 {
        hanoi(disk -1, start, middle, end)
        fmt.Printf("Move dist%d from %s to %s \n", disk, start, end)
        hanoi(disk - 1, middle, end, start)
    }
}

func main() {
    s := 3
    hanoi(s, "START", "END", "MIDDLE")
}