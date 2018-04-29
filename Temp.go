package main

import (
	"encoding/json"
	"fmt"
	"github.com/dreamCodeMan/go-mediainfo"
)

func main() {
	mediainfo, _ := mediainfo.GetMediaInfo("D:\\Jolin.ts")
	info, _ := json.Marshal(mediainfo)
	fmt.Println(string(info))
}
