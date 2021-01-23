package top

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type House struct {
	ID    string `table:"ID"`
	Name  string `table:"接口"`
	Sigil string `table:"IP"`
	Motto int    `table:"请求次数"`
}

func Top() {

	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())

	//for {
	//	fmt.Println("Over")
	//	time.Sleep(time.Second * 1)
	//	cmd := exec.Command("cmd", "/c", "cls")
	//	cmd.Stdout = os.Stdout
	//	cmd.Run()
	//}

	//cmd := exec.Command("cmd", "/c", "cls")
	//cmd.Stdout = os.Stdout
	//cmd.Run()
	//
	//for i :=0;i!=10;i=i+1{
	//	fmt.Fprintf(os.Stdout,"result is %d\r",i)
	//	time.Sleep(time.Second*1)
	//}
	//fmt.Println("Over")

	//n := 0
	//for {
	//	s := []House{
	//		{"1", "api/login/login", "192.168.0.138", 3},
	//		{"2", "api/user/getUserInfo", "192.168.0.138", 2},
	//		{"3", "api/user/detail", "192.168.0.138", n},
	//	}
	//	t := table.Table(s)
	//	fmt.Fprintf(os.Stdout, "%s\r", t)
	//	n++
	//	time.Sleep(time.Second*1)
	//
	//	//app := "clear"
	//	//cmd := exec.Command(app)
	//	//stdout, err := cmd.Output()
	//	//
	//	//if err != nil {
	//	//	println(err.Error())
	//	//	return
	//	//}
	//	//print(string(stdout))
	//}


}
