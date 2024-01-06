package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

func generateRandom() int {
	return rand.Intn(100)
}

func generate(w http.ResponseWriter, r *http.Request) {
	randomNumber := generateRandom()

	str := "random number is: " + strconv.Itoa(randomNumber)

	if _, err := w.Write([]byte(str)); err != nil {
		fmt.Println("Err")
		return
	}

	fmt.Println("done!")
}

func main() {
	// POST => body => save db => response => return successefully saved || id (uuid)
	// GET => no body => id => response => return data
	// PUT => body => update => response => return data || id
	// DELETE => no body => id => response => return successefully deleted

	http.HandleFunc("/register-user", func(w http.ResponseWriter, r *http.Request) {
		//logic, db request body kelgan body save
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("your data saved"))
	})

	// http.HandleFunc("/generate", generate)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("not found for example"))
	// })

	fmt.Println("program is running ...")
	http.ListenAndServe("localhost:8080", nil)

	/*var name1, name2, name3 string

	fmt.Print("enter name: ")
	fmt.Scan(&name1, &name2, &name3)

	client := http.Client{}

	// url := fmt.Sprintf("https://api.agify.io?name[]=%s&name[]=%s&name[]=%s", name1, name2, name3)
	url := "https://api.agify.io"

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("error while creating request", err.Error())
		return
	}

	query := request.URL.Query()
	query.Add("name[]", name1)
	query.Add("name[]", name2)
	query.Add("name[]", name3)
	request.URL.RawQuery = query.Encode()

	fmt.Println("request url", request.URL.String())

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("error while sending request", err.Error())
		return
	}
	fmt.Println("status code: ", response.StatusCode, response.Status)
	fmt.Println("resp: ", response)

	p := []Person{}

	b, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error while reading body", err.Error())
		return
	}

	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println("error while unmarshalling data", err.Error())
		return
	}

	for _, v := range p {
		fmt.Println(v.Age, v.Name, v.Count)
	}*/
}

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Count int    `json:"count"`
}
