package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
	
)
func main(){
	fbytes, err:=ioutil.ReadFile("azure.tf")
	if err!=nil{
		fmt.Println("error")
		os.Exit(1)
	}
	astfile, hclerr:=hcl.ParseBytes(fbytes)
	if hclerr!=nil{
		fmt.Println("error")
		os.Exit(2)
		
	}

	ast.Walk(astfile, func(n ast.Node)(ast.Node , bool){
		if n == nil {
			return n, false
		}

		fmt.Println(n)
		return n, true		
	})

}
