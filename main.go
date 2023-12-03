package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	redGamer := []int{rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10)}
	blueGamer := []int{rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10)}
	var b int
	fmt.Printf("CHOOSE THE FIRST INDEX FOR THE BLUE TEAM:\n")
	fmt.Scan(&b)
	time.Sleep(time.Second * 1)
	var r int
	fmt.Printf("CHOOSE THE FIRST INDEX FOR THE RED TEAM:\n")
	fmt.Scan(&r)
	time.Sleep(time.Second * 1)
	a := 0

	for k := 1; k < 6; k++ {
		fmt.Printf("THE %d CARD REVEALED BY THE BLUE TEAM: %d\n", k, blueGamer[b])
		fmt.Printf("THE %d CARD REVEALED BY THE RED TEAM: %d\n", k, redGamer[r])
		if blueGamer[b] == 0 && redGamer[r] == 0 {
			fmt.Printf("BOTH TEAMS HAVE WON.\n=================================================\nTHE CARDS ARE BEING REDISTRIBUTED.")
			break
		} else if blueGamer[b] == 0 {
			fmt.Printf("THE BLUE TEAM HAS WON.\n")
			break
		} else if redGamer[r] == 0 {
			fmt.Printf("THE RED TEAM HAS WON.\n")
			break
		} else {
			time.Sleep(time.Second * 3)
		}
		b = blueGamer[b]
		r = redGamer[r]
		fmt.Printf("\nA NEW CARD FOR THE NEW INDEX IS BEING REVEALED.\n")
		a += 1
	}
	time.Sleep(time.Second * 5)
	if a == 5 {
		fmt.Printf("\nTHE NUMBER OF MOVES HAS REACHED ITS MAXIMUM.\n=================================================\nA NEW GAME IS BEGINNING.\n")
	}
	fmt.Printf("\nBLUE TEAM CARDS:%v \n", blueGamer)
	fmt.Printf("RED TEAM CARDS:%v \n", redGamer)
}
