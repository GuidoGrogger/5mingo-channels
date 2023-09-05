package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	const shoeSalesGoal = 300
	const shoeSaleAgents = 10
	const verboseOutput = true

	start := time.Now()
	soldShoeCountChan := make(chan int)

	// Sell shoes with 10 agents
	for i := 0; i < shoeSaleAgents; i++ {
		go func(agentNo int) {
			for {
				soldShoes := rand.Intn(10)
				waitTime := time.Duration(rand.Intn(10000)) // 10 sec
				time.Sleep(time.Duration(waitTime * time.Millisecond))
				if verboseOutput {
					fmt.Printf("agent_no %d sold %d shoes\n", agentNo, soldShoes)
				}
				soldShoeCountChan <- soldShoes
			}
		}(i)
	}

	allShoes := 0
	for elem := range soldShoeCountChan {
		allShoes += elem
		if verboseOutput {
			fmt.Printf("Total shoes sold: %d\n", allShoes)
		}
		if allShoes > shoeSalesGoal {
			fmt.Printf("Goal reached! - %d shoes sold - Time elapsed %s", allShoes, time.Since(start))
			os.Exit(0)
		}
	}
}
