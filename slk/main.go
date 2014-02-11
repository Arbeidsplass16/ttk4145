package main

import(
    "os"
    "bufio"
	"fmt"
)

func main(){
    out_socket, in_socket := connect()
    go recieve(in_socket)
    
    reader  := bufio.NewReader(os.Stdin)
    for{
        input,_ := reader.ReadString('\n')
        output  := string(input)
        for i:=0; i<5; i++ {
            broadcast(out_socket,output)
        }
    }
    
}
