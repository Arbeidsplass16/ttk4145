package main

import(
	"strings"
	""
	""


func make_message(msg_id int, floor int, direction int, cost int)(string){
	message := string(msg_id)+","+string(floor)+","+string(direction)+","string(cost)
	return message
}



func read_message(message string){
	string_list := strings.Split(message,".")
	if string_list[0]{
		//delete order
	}
	else{
		//check_cost()
	}
}

