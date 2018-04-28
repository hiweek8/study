package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/nsf/termbox-go"
	"strings"
)
func GetUUID() string {
	uid := uuid.Must(uuid.NewV4())
	return uid.String()
}

func init() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	termbox.SetCursor(0, 0)
	termbox.HideCursor()
}

func pause() {
	fmt.Println("请按任意键继续...")
Loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			break Loop
		}
	}
}

func main() {
/*	f, _ := os.OpenFile("Z://1.txt", os.O_RDONLY, 0)
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	a:=GetUUID()
*/
	for i := 0; i < 10; i++ {
		fmt.Println(strings.ToUpper(GetUUID()))
	}
	pause()
}
