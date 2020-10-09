package main

import (
	"fmt"
	"mediadup/mediainfo"
	"mediadup/movieid"
	"os"
)

func main() {
	mediainfo, err := mediainfo.GetMediaInfo(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	t,err := mediainfo.GetMovieTrackID()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	movie := t.Movie
	md := movieid.New()
	result, err := md.SearchMovie(movie)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%v", result)
}
