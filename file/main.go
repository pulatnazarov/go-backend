package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Company struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Owner   Owner    `json:"owner"`
	Workers []Worker `json:"workers"`
}

type Owner struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Worker struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Salary    int    `json:"salary"`
	Level     string `json:"level"`
}

func main() {
	file, err := os.Open("test.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	companies := []Company{}

	if err = json.NewDecoder(file).Decode(&companies); err != nil {
		panic(err)
	}

	//workers := []Worker{}
	//
	//for _, company := range companies {
	//	for _, worker := range company.Workers {
	//		workers = append(workers, worker)
	//	}
	//}
	//
	//sort.Slice(workers, func(i, j int) bool {
	//	return workers[i].Salary > workers[j].Salary
	//})
	//
	//for i, worker := range workers {
	//	fmt.Println(worker)
	//	if i == 2 {
	//		break
	//	}
	//}

	sort.Slice(companies, func(i, j int) bool {
		sumI, sumJ := 0, 0
		for _, worker := range companies[i].Workers {
			sumI += worker.Salary
		}
		fmt.Println("sum i", sumI, companies[i].Name)

		for _, worker := range companies[j].Workers {
			sumJ += worker.Salary
		}
		fmt.Println("sum j", sumJ, companies[j].Name)

		return sumI > sumJ
	})

	for _, company := range companies {
		fmt.Println(company.Owner)
	}
}
