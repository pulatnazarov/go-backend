package main

import "fmt"

type Team struct {
	Name         string
	Coach        string
	PlayersCount int
	Captain      string
}

func main() {
	teams := []Team{
		{
			Name:         "Barca",
			Coach:        "Trainer 1",
			PlayersCount: 11,
			Captain:      "Captain 1",
		},

		{
			Name:         "Real",
			Coach:        "Trainer 2",
			PlayersCount: 13,
			Captain:      "Captain 2",
		},
		{
			Name:         "Bayern",
			Coach:        "Trainer 3",
			PlayersCount: 11,
			Captain:      "Captain 3",
		},
		{
			Name:         "Manchester",
			Coach:        "Trainer 4",
			PlayersCount: 5,
			Captain:      "Captain 4",
		},
		{
			Name:         "Liverpool",
			Coach:        "Trainer 1",
			PlayersCount: 10,
			Captain:      "Captain 1",
		},
	}

	names := []string{}

	for i := 0; i < 3; i++ {
		var name string
		fmt.Printf("name %d: ", i+1)
		fmt.Scan(&name)
		names = append(names, name)
	}

	newTeams := []Team{}

	for _, val := range teams {
		for _, name := range names {
			if name == val.Name {
				newTeams = append(newTeams, val)
			}
		}
	}

	for i := 0; i < len(newTeams); i++ {
		for j := i + 1; j < len(newTeams); j++ {
			if newTeams[i].PlayersCount < newTeams[j].PlayersCount {
				newTeams[i], newTeams[j] = newTeams[j], newTeams[i]
			}
		}
	}

	for _, val := range newTeams {
		fmt.Println(val)
	}
}
