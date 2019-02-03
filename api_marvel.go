package main


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"crypto/md5"
)

type Response struct {
	Id int `json:"id"`
	Name    string    `json:"name"`
	//Hero []Hero `json:"name"`
}


type Hero struct {
	EntryNo int            `json:"id"`
	//Listado HeroListado `json:"name"`
	Name string `json:"name"`
}


type HeroListado struct {
	Name string `json:"name"`
}




func digestString(hash string) string {
    return fmt.Sprintf("%x", md5.Sum([]byte(hash)))
}

func main() {
	ts := "1"
	publickey := "86a6c35b1ae3e03447a201b560af1f3c"
	privatekey := "0f013d294001327e1a8fdcde3123f137c43877dc"
	
	hash := ts + privatekey + publickey

	//Verificacion del md5
	//fmt.Println(digestString(hash))
	
	link := "http://gateway.marvel.com/v1/public/characters?ts=" + ts + "&apikey=" + publickey + "&hash=" + digestString(hash)

	response, err := http.Get(link)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//Se imprime data traida de la api marvel
	fmt.Println(string(responseData))

	

	var responseObject HeroListado
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Name))

	
} 