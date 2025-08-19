package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//root handler
func root (w http.ResponseWriter, r *http.Request) {
	//check if method is not GET then throw an error
	if  r.Method != http.MethodGet {
		http.Error(w,"Bad Request type",400)
	}
	//add cors
	w.Header().Set("Access-Control--Allow-Origin","*")
	//send as json format
	w.Header().Set("Content-Type", "application/json")
	//create a encoder
	encoder := json.NewEncoder(w)
	//encode the struct => json
	encoder.Encode(&RootResponse{
		Success: true,
		Message: "Server is operational",
	})
}

func main () {
	//router
	mux := http.NewServeMux()

	//root route
	mux.HandleFunc("/",root)

	//listen on 8080 port if catch any error then store in err
	err := http.ListenAndServe(":8080",mux)

	//if err is not nil then throw an err
	if err != nil {
		fmt.Println("Something went wrong on starting server",err)
		//if nil then print server status
	}else {
		fmt.Println("Server running at port 8080")
	}
}
