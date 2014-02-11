package main

import(
    "net"
    "fmt"
)

func connect()(*net.UDPConn, *net.UDPConn){
    localport      := "129.241.187.255:20016"
	for{
    	inAddr,_       		:= net.ResolveUDPAddr("udp4",localport)
	    out_socket,err_out 	:= net.DialUDP("udp4", nil, inAddr)
	    in_socket, err_in 	:= net.ListenUDP("udp4", inAddr)

	    if CheckError(err_out) && CheckError(err_in){
			return out_socket, in_socket		
		}
	}    
}


func CheckError(err error)(bool){
    if err!=nil{
        fmt.Println(err.Error())
		return false
    }
	return true
}

func broadcast(out_socket *net.UDPConn, message string){
	out_socket.Write([]byte(message+"\x00"))
}



func recieve(in_socket *net.UDPConn){
    for{
        var buffer [1024]byte
        n, _,_		:= in_socket.ReadFromUDP(buffer[0:]) 
        message		:= string(buffer[0:n])
        fmt.Println("recieved:", message)
   }
}
