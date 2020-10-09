package main

import (
	"fmt"
	"github.com/kenshaw/imdb"
	"log"
	"regexp"
	"strings"
)

type movieid struct {
	cl *imdb.Client
}

func NewMovieId() *movieid {
	cl := imdb.New("9ff7bfda")
	return &movieid{cl}
}

func (md *movieid) SearchMovie(movieName string) (*imdb.MovieResult, error){
	y, err := extractYear(movieName)
	if err != nil {
		return nil, err
	}
	if len(y) < 1 {
		return nil, fmt.Errorf("year not found")
	}
	for _,year := range y {
		name, err := extractName(year, movieName)
		if err != nil {
			break
		}
		cl := imdb.New("9ff7bfda")
		log.Println("finding:" + name + " year:" + year )
		res, err := cl.MovieByTitle(name, year)
		if err != nil {
			break
		}
		return res, nil
	}
	return nil, fmt.Errorf("no match found")
}

func extractYear(name string) ([]string, error) {
	re,err := regexp.Compile("\\b19\\d{2}|20\\d{2}\\b")
	if err!=nil{
		return nil, err
	}
	allSubmatch := re.FindStringSubmatch(name)
	return allSubmatch,nil
}

func extractName(year string, title string) (string, error) {
	re,err := regexp.Compile("[^\\s.\\-_]*")
	if err!=nil{
		return "", err
	}
	allSubmatch := re.FindAllStringSubmatch(title,-1)
	name := ""
	for _,s := range allSubmatch {
		if s[0] == year {
			break
		}
		name += s[0]+" "
	}
	return strings.Trim(name," "), nil
}


