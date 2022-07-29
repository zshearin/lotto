package main

import (
        "fmt"
        "math/rand"
        "sort"
        "time"
)

func main() {

        var numList []int
        for i := 0; i < 5; i++ {
                newNum := getRandomNumber(70)
                numList = append(numList, newNum)
        //      fmt.Printf("%d ", newNum)
        }

        sort.Ints(numList)

        for _, val := range numList {
                fmt.Printf("%d ", val)
        }
        mult := getRandomNumber(25)
        fmt.Printf("+ %d\n", mult)

        fmt.Printf("Odds of hitting the jackpot: 1 in %d\n", (70 * 69 * 68 * 67 * 66 * 25)/(5*4*3*2))
}

//Get number from [1, num]
func getRandomNumber(num int) int {
        rand.Seed(time.Now().UnixNano())
        return rand.Intn(num + 1)
}
