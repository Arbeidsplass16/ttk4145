package main

import(
    "net"
    "fmt"
)

func connect()(*net.UDPConn, *net.UDPConn){
    localport      := "129.241.187.255:20016"
    inAddr,_       := net.ResolveUDPAddr("udp4",localport)
    
    out_socket,err := net.DialUDP("udp4", nil, inAddr)
    CheckError(err)
    in_socket, err := net.ListenUDP("udp4", inAddr)
    CheckError(err)
    /*
    out_socket,err := net.DialUDP("udp4",&net.UDPAddr{IP: net.IPv4(129,241,187,147)
    
    Port: 20016,},&net.UDPAddr{IP: net.IPv4(129,241,187,255), Port: 20016,})
    CheckError(err)
    in_socket, err  := net.ListenUDP("udp4",inAddr)
    CheckError(err)
    */
    return out_socket, in_socket
}

func CheckError(err error){
    if err!=nil{
        fmt.Println(err.Error())
    }    
}

func write(out_socket *net.UDPConn, message string){
    out_socket.Write([]byte(message+"\x00"))
}

func read(in_socket *net.UDPConn){
    for{
        var buffer [1024]byte
        n, _,_    := in_socket.ReadFromUDP(buffer[0:]) 
        message      := string(buffer[0:n])
        fmt.Println("read got:", message)
   }
}
