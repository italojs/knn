package main

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func readFile(path string) (record [][]string) {
	//Open the file
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln("Error: ", err)
		return
	}
	//Close the file when this method finish
	defer file.Close()

	//Read the file
	reader := csv.NewReader(file)
	record, err = reader.ReadAll()
	if err != nil {
		log.Fatalln("Error:", err)
		return
	}
	return
}
func distinct(elements []string) (result []string) {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}

	for i := range elements {
		if encountered[(elements)[i]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[(elements)[i]] = true
			// Append to result slice.
			result = append(result, (elements)[i])
		}
	}
	// Return the new slice.
	return result
}
func getCollum(elements [][]string, columnIndex int) (column []string) {
	//for each []line in dataset
	//get the the columIndex value
	//and add into the colum array
	for i := range elements {
		column = append(column, (elements)[i][columnIndex])
	}
	return
}
func getValuesByClass(records [][]string, class string) (newRecords [][]string) {
	//make a new matrix instance with 0 lines
	newRecords = make([][]string, 0)
	//for each line in your matrix(records)
	for l := range records {
		//if the value in the last index of this line(that is the class) is iguals
		//this method's parameter class
		if records[l][len(records[l])-1] == class {
			//add this array(this line) into "newRecords" matrix
			newRecords = append(newRecords, records[l])
		}
	}
	return
}
func divideInPercent(records [][]string, percent float32) (newRecords [][]string, residue [][]string) {
	//multiply the number of lines by percent(float number that is < 1.0)
	//"i" now is the lines quantity to get in the dataset (records)
	i := float32(len(records)) * percent
	//while i is bigger than 0
	for i > 0 {
		//so get the line index (i)
		//"-1" because in the "i" variable i had the records lenght and not the index values
		//obs: lenght doesnt start at 0 number, index start at 0 number
		index := int(i - 1)
		//add the line at "newRecords"
		newRecords = append(newRecords, records[index])
		i--
	}

	//getting the rest of records lines

	//get again the first percent of lines
	i = float32(len(records)) * percent
	//get records lenght
	j := 0
	//while the j(records lenght) is bigger then "i"
	//get the "j" line and put ir on "residue" matrix
	for float32(j) < i {
		residue = append(residue, records[j])
		j++
	}

	return
}

//"pi" is the line in your train dataset
//"qi" is your new value that you wanna know the distance of "pi"
func euclideanDist(pi []string, qi []string) (result float64) {
	//get the column quantity that "pi" have
	i := len(pi) - 1

	//sum all columns
	for i > 0 {
		//(pi - 1) because i dont wanna get the class column
		pif, _ := strconv.ParseFloat(pi[i-1], 32)
		qif, _ := strconv.ParseFloat(qi[i-1], 32)
		//get (pif - qif) squared
		result += math.Pow(pif-qif, 2)
		i--
	}
	//get square root of "result"
	result = math.Sqrt(result)

	return
}

//get only the keys of a map type (keyValue type)
func getMapKeys(m map[float64]string) (keys []float64) {
	for k := range m {
		keys = append(keys, k)
	}
	return
}

//get only the values of a map type (keyValue type)
func getMapValues(m map[float64]string) (values []string) {
	for _, v := range m {
		values = append(values, v)
	}
	return
}
func getKnn(list map[float64]string, k int) (sortedMap map[float64]string) {
	//get only keys of the (keyValue) distances
	//the keys contain the distance value(numeric)
	keys := getMapKeys(list)

	//sort this numbers
	sort.Float64s(keys)

	//new map[float64]string instance
	sortedMap = make(map[float64]string)

	//for each sorted distance in "keys"
	for i, key := range keys {
		if i < k {
			//add the keyValue into a new keyValue list
			sortedMap[key] = list[key]
		} else {
			break
		}
	}
	//in the end, this new keyValue list is be sorted
	return
}
func getPredominantClass(knn map[float64]string) (class string) {
	//distinct the classes
	classes := distinct(getMapValues(knn))
	var predominantClass int
	//for each class into the "knn" list
	//obs: it's not for each line, it's for each CLASS
	for c := range classes {
		//declare this variable here to every time that
		//the loop execute it again, this variable must be 0
		var countClass int
		//for each distance in "knn" list
		/*obs: knn have the list of K Nearest Neighbord. e.g. 10 smallest distances
		in "getKnn" method*/
		for i := range knn {
			//verify if the current distance is of the type of current class
			if knn[i] == classes[c] {
				//if true, cout +1 to this current class
				countClass++
			}
		}
		//if predominantClass count is biggest the current class count
		if predominantClass < countClass {
			//replace the old predominant class by current class
			predominantClass = countClass
			class = classes[c]
		}
	}
	return
}
func classify(train [][]string, valueToPredict []string, k int) (result string) {
	//new map[float64]string instace
	dists := make(map[float64]string)
	//get the line length
	i := len(train) - 1
	//get the line's euclidian distance to all other lines in train dataset
	for i >= 0 {
		class := train[i][len(train[i])-1]
		d := euclideanDist(train[i], valueToPredict)
		dists[d] = class
		i--
	}
	//to all distnce, get the k(e.g. 10) smallers distances
	knn := getKnn(dists, k)
	//get the predominant class of "knn"
	result = getPredominantClass(knn)

	return
}
func main() {

	//Read the dataset
	records := readFile("data.csv")
	//Get the classes
	classes := distinct(getCollum(records, len(records[0])-1))

	var train [][]string
	var test [][]string
	//Percent
	p := float32(0.9)
	//For each class, get 60% of lines
	for i := range classes {
		//Get all the values by current class
		values := getValuesByClass(records, classes[i])
		//Get the percent lines
		trainRecords, testRecords := divideInPercent(values, p)
		index := len(trainRecords)
		//While exist lines, add it into train array
		for index > 0 {
			train = append(train, trainRecords[index-1])
			index--
		}

		index = len(testRecords)
		//While exist lines, add it into test array
		for index > 0 {
			test = append(test, testRecords[index-1])
			index--
		}
	}

	var hits int
	//for each line in test lines
	for i := range test {
		//classify this line
		result := classify(train, test[i], 10)
		//get the correct class of this line
		columnIndex := len(test[i]) - 1
		//print the correct class and the predicted class by algorithm
		//fmt.Println("tumor: ", test[i][columnIndex], " classificado como: ", result)
		//if the predicted class is correct, add +1 to hits(count of correct class predicted by algorithm)
		if result == test[i][columnIndex] {
			hits++
		}
	}

	// fmt.Println("Total de dados: ", len(records))
	// fmt.Println("Total de treinamento: ", len(train))
	// fmt.Println("Total de testes: ", len(test))
	// fmt.Println("Total de acertos: ", hits)
	// fmt.Println("Porcentagem de acertos: ", (100 * hits / len(test)), "%")

}
