package main

import (
	"fmt"
	"os"
	"log"
)
var path = "./Lista.txt"

func main() {
	CriaArquivo()
}

func CriaArquivo(){

// Checa se o arquivo já existe
    var _, err = os.Stat(path)

    // Cria caso não exista
    if os.IsNotExist(err) {
        var file, err = os.Create(path)
        if isError(err) {
            return
        }
        defer file.Close()
    }
	ColetaNomes()
}
	
func ColetaNomes(){
	file, err := os.Open(".")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	list, err := file.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range list {
		EscreveLista(f.Name())
	}
}

func EscreveLista(data string){

    var file, err = os.OpenFile(path, os.O_APPEND, 0644)
    if isError(err) {
        return
    }
    defer file.Close()

    // Escreve linha por linha
    _, err = file.WriteString(data + "\n")
    if isError(err) {
        return
    }

    // Salva alterações.
    err = file.Sync()
    if isError(err) {
        return
    }

}


func isError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }

    return (err != nil)
}