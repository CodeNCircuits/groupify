package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	var participants []string
	scanner := bufio.NewScanner(os.Stdin)

	var teamSize int
	fmt.Println("Enter the desired team size (minimum 2):")
	_, err := fmt.Scanf("%d", &teamSize)
	if err != nil || teamSize < 2 {
		fmt.Println("Invalid input for team size. Please enter an integer greater than 1.")
		return
	}

	fmt.Println("Enter participant names (press Enter without typing a name to finish):")

	for {
		scanner.Scan()
		name := scanner.Text()
		name = strings.TrimSpace(name)

		if name == "" {
			break
		}

		if len(name) == 0 {
			fmt.Println("Invalid name entered. Please enter a non-empty name.")
			continue
		}

		participants = append(participants, name)
	}

	if len(participants) < 3 {
		fmt.Println("Not enough participants. Please enter at least 3 names.")
		return
	}

	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	randGen.Shuffle(len(participants), func(i, j int) {
		participants[i], participants[j] = participants[j], participants[i]
	})

	numTeams := len(participants) / teamSize
	remainder := len(participants) % teamSize

	fmt.Printf("\nTotal Participants: %d\n", len(participants))
	fmt.Printf("Forming %d teams of %d members each.\n", numTeams, teamSize)

	for i := 0; i < numTeams; i++ {
		team := participants[i*teamSize : (i+1)*teamSize]
		fmt.Printf("Team %d: %v\n", i+1, team)
	}

	if remainder > 0 {
		remainingTeam := participants[numTeams*teamSize:]
		fmt.Printf("Remaining Participants: %v\n", remainingTeam)
		fmt.Println("You can choose to redistribute the remaining members or create a smaller team.")
		redistributeParticipants(&participants, numTeams, remainder)
	}

	if len(participants)%teamSize != 0 {
		fmt.Println("\nNote: The last team might have fewer participants than the others.")
	}
}

func redistributeParticipants(participants *[]string, numTeams, remainder int) {
	if remainder > 0 {
		fmt.Println("\nRedistributing remaining participants into teams...")
		for i := 0; i < remainder; i++ {
			teamIndex := rand.Intn(numTeams)
			*participants = append((*participants)[:teamIndex*len(*participants)/numTeams], append([]string{(*participants)[numTeams*len(*participants)/numTeams]}, (*participants)[teamIndex*len(*participants)/numTeams:]...)...)
		}

		fmt.Println("Redistribution completed.")
	}
}
