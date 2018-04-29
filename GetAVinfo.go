package main

import (
	"encoding/json"
	"go-ffprobe"
	"log"
	"time"
	"fmt"
)

func main() {
	path := "Z:/POV - Nekane - Nekane takes care of you/POV - Nekane - Nekane takes care of you.mp4"

	data, err := ffprobe.GetProbeData(path, 500 * time.Millisecond)
	if err != nil {
		log.Panicf("Error getting data: %v", err)
	}

	buf, err := json.MarshalIndent(data, "", "  ")
	fmt.Print(string(buf))

	buf, err = json.MarshalIndent(data.GetFirstVideoStream(), "", "  ")
	log.Print("视频信息：",string(buf))

	buf, err = json.MarshalIndent(data.GetFirstAudioStream(), "", "  ")
	log.Print("音频信息：",string(buf))

	log.Printf("总时长（Duration）: %v\n", data.Format.Duration())
	log.Printf("开始时间（StartTime）: %v\n", data.Format.StartTime())

/*	start := time.Now()
*	for i := 0; i < 100; i++ {
*		_, err = ffprobe.GetProbeData(path)
*		if err != nil {
*			log.Panicf("Error getting data: %v", err)
*		}
*	}
*	log.Printf("100 times get time: %v", time.Now().Sub(start))
*/
}
