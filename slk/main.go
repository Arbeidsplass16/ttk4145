package main
//testing
import(
    "os"
    "bufio"
)

func main(){
    out_socket, in_socket := connect()
    go read(in_socket)
    
    reader  := bufio.NewReader(os.Stdin)
    for{
        input,_ := reader.ReadString('\n')
        output  := string(input)
        for i:=0; i<5; i++ {
            write(out_socket,output)
        }
    }
    
}
