package main


import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	
)


func readjson(){

	raw, err := ioutil.ReadFile("./azure.json")
	if err!=nil{
		log.Fatal("cannot open file")
	}
	var result  interface{}

	json.Unmarshal(raw, &result)

	resmap:= result.(map[string]interface{})
		
	for k, v := range resmap{
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}

