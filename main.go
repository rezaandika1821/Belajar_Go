package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
var address, addresstemen string
var umur int  //deklarasi with var
var kelas int = 2  //deklarasi with var and value
alamat := "Sidoarjo" //deklarasi inference
no := 90  //deklarasi inference
var (
	nama string
	nim int
	ipk float32
)  //deklarasi with var majemuk
_ = "dasdsakdjaksdj" //deklarasi underscore

	fmt.Println(kelas)
	fmt.Println(alamat)
	fmt.Println(no)
	fmt.Println(nim)
	fmt.Println(ipk)
	fmt.Println(umur)
	
	fmt.Println("masukan nama anda : ")
	fmt.Scanln(&nama)
	fmt.Println(nama)

	input := bufio.NewScanner(os.Stdin) //inptan menggunakan spasi
	fmt.Print("masukan alamat anda :")
	input.Scan()
	address = input.Text()
	fmt.Println("ini address anda : ", address)

	fmt.Println("masukan alamat temen : ")
	input.Scan()
	addresstemen = input.Text()
	fmt.Println("ini address temen anda : ", addresstemen)
	//ini branch operator
}
