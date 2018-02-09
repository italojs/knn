package main

import (
	"fmt"
	"github.com/italojs/knn/algorithm"
	"os"
	"log"
	"encoding/csv"
)
func readFile(path string)(record [][]string){
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln("Error: ", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	record, err = reader.ReadAll()
	if err != nil {
		log.Fatalln("Error:", err)
		return
	}
	return
}

func main (){
	records := readFile("datas/wdbc.csv")
	
	train, test := knn.PrepareDataset(0.6, records)

	var hits int
	for i := range test{
		result := knn.Classify(train,test[i],10)
		columnIndex := len(test[i])-1
		fmt.Println("tumor: ",test[i][columnIndex]," classificado como: ",result)
		if result == test[i][columnIndex]{
			hits++
		}
	}

	fmt.Println("Total de dados: ", len(records))
	fmt.Println("Total de treinamento: ", len(train))
	fmt.Println("Total de testes: ", len(test))
	fmt.Println("Total de acertos: ", hits)
	fmt.Println("Porcentagem de acertos: ", (100 * hits / len(test)),"%")

}