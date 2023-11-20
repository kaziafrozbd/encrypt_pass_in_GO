package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	// "fmt"
)

func main(){

	// pass := "asd@123"
	//password encryt
	bs,err := bcrypt.GenerateFromPassword([]byte("asd@123"), 14)
	if err!=nil{
		fmt.Println(err.Error())
	}
	// fmt.Println(bs)
	fmt.Println(string(bs))

	//pass compare
	err = bcrypt.CompareHashAndPassword(bs,[]byte("asd@123"))
	if err!=nil{
		fmt.Println(err)
	}else{
		fmt.Println("Same password")
	}

}