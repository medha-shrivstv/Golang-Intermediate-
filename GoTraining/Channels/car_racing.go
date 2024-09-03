package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Car function simulates a car racing to finish. It takes a name, a wait group, and a channel to signal the winner.
func Car(name string, wg *sync.WaitGroup, winnerCh chan string, stopCh chan struct{}) {
	defer wg.Done()

	// Simulating the race with random duration
	raceDuration := time.Duration(rand.Intn(10)+1) * time.Second

	select {
	case <-time.After(raceDuration):
		// The car finishes the race
		select {
		case winnerCh <- name:
			// If this car is the first to finish, it sends its name to the winner channel
		default:
			// If a winner is already declared, do nothing
		}
	case <-stopCh:
		// Stop the race if a winner is already declared
		fmt.Printf("%s stopped racing.\n", name)
		return
	}
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Wait group to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Channel to signal the winner
	winnerCh := make(chan string)

	// Channel to signal other cars to stop once a winner is declared
	stopCh := make(chan struct{})

	// List of car names
	cars := []string{"Car1", "Car2", "Car3"}

	// Start the race with each car running in its own goroutine
	for _, car := range cars {
		wg.Add(1)
		go Car(car, &wg, winnerCh, stopCh)
	}

	// Wait for a winner to be declared
	winner := <-winnerCh
	fmt.Printf("%s is the winner!\n", winner)

	// Close the stop channel to signal all other cars to stop
	close(stopCh)

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Race finished.")
}
