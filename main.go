package main

import (
	"encoding/json"
	"fmt"
	"mediadup/mediainfo"
	"os"
)

func main() {
	mediainfo, err := mediainfo.GetMediaInfo(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	info, err := json.Marshal(mediainfo)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(info))
}
