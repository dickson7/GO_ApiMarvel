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
	Code int `json:"code"`
	Data  Heroes  `json:"data"`
}


type Heroes struct {
	Total int            `json:"total"`
	//Listado HeroListado `json:"name"`
	HeroListado []HeroListado   `json:"results"`
	//Listado HeroListado
	//Name string `json:"name"`
}


type HeroListado struct {
	Name string `json:"name"`
	Description string `json:"description"`
}




func digestString(hash string) string {
    return fmt.Sprintf("%x", md5.Sum([]byte(hash)))
}

func listado () {
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
		//fmt.Println(string(responseData))


	// La cadena JSON devuelta en una nueva variable.
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	
	for i := 0; i < len(responseObject.Data.HeroListado); i++ {
		
		fmt.Print((i*1)+1)
		fmt.Print(".   Name: ")
		fmt.Println(responseObject.Data.HeroListado[i].Name)
		//fmt.Print("Description : ")
		//fmt.Println(responseObject.Data.HeroListado[i].Description)
	}

}

func main() {
	var opcion int
	opcion = 0

		fmt.Println("____________________________________________________________")
		fmt.Println("I                                                          I")
		fmt.Println("I      Este codigo consume la API de MARVEL con            I")	
		fmt.Println("I                       GOLANG                             I")	
		
	for opcion != 3 {

		fmt.Println("____________________________________________________________")
		fmt.Println("I                                                          I")
		fmt.Println("I  Digite la opcion 1 (uno) puedes obtener un listado      I")
		fmt.Println("I  de Superheroes                                          I")
		fmt.Println("I                                                          I")
		fmt.Println("I  Digite la opcion 2 (dos) para buscar Superheroes        I")
		fmt.Println("I                                                          I")
		fmt.Println("I  Digite la opcion 3 (tres) para salir                    I")
		fmt.Println("____________________________________________________________")
		fmt.Println("                                                           ")
		
		fmt.Scanf("%d\n",&opcion)
		
		if opcion == 1 {
			
		}else {
			if opcion == 2 {
				fmt.Println("____________________________________________________________")
				fmt.Println("I                                                          I")
				fmt.Println("I                Listado de Superheroes                    I")
				fmt.Println("____________________________________________________________")
		
				listado()
			} else {
				if opcion == 3 {
					fmt.Println("OUT")
				}
			}
		}
	}	
	
} 
