package main

import "fmt"

func main() {
	movies := movie{"avengers", "star-wars"}
	movies = append(movies, "ant-man")
	movies.printMovieNames() // any variable of type "movie" will have access to the printMovieNames method
}

// custom type
type movie []string

// receiver function syntax
func (m movie) printMovieNames() {

	for id, name := range m {
		fmt.Println(id, name)
	}
}
