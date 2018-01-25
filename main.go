package main

import (
	"fmt"
	"github.com/italojs/knn/knn"
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

	
	column := knn.GetCollum(&records,len(records[0])-1)
	classes := knn.Distinct(&column)
	var train [][]string
	var test [][]string
	p := float32(0.9)
	for i := range classes{
		values := knn.GetValuesByClass(&records, classes[i])
		trainRecords, testRecords := knn.DivideInPercent(&values,p)
		index := len(trainRecords)
		for index>0{
			train = append(train, trainRecords[index-1])
			index--
		}

		index = len(testRecords)
		for index>0{
			test = append(test, testRecords[index-1])
			index--
		}
	}

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