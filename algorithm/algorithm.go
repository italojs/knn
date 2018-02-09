package knn

import (
	"math"
	"sort"
	"strconv"
)

var columClassIndex int

func distinct(elements *[]string) (result []string) {
	encountered := map[string]bool{}

	for i := range *elements {
		if encountered[(*elements)[i]] == true {
		} else {
			encountered[(*elements)[i]] = true
			result = append(result, (*elements)[i])
		}
	}
	return result
}
func getCollum(elements *[][]string, ColumnIndex int) (column []string) {
	for i := range *elements {
		column = append(column, (*elements)[i][ColumnIndex])
	}
	return
}
func getValuesByClass(records *[][]string, class string) (newRecords [][]string) {
	newRecords = make([][]string, 0)
	for l := range *records {
		if (*records)[l][len((*records)[l])-1] == class {
			newRecords = append(newRecords, (*records)[l])
		}
	}
	return
}
func divideInPercent(records *[][]string, percent float32) (newRecords [][]string, residue [][]string) {
	i := float32(len(*records)) * percent
	for i >= 0 {
		index := int(i - 1)
		newRecords = append(newRecords, (*records)[index])
		i--
	}

	i = float32(len(*records)) * percent
	j := float32(len(*records))
	for j > i {
		index := int(j - 1)
		residue = append(residue, (*records)[index])
		j--
	}
	return
}
func euclideanDist(pi *[]string, qi *[]string) (result float64) {
	i := len(*pi) - 1

	for i >= 0 {
		pif, _ := strconv.ParseFloat((*pi)[i], 32)
		qif, _ := strconv.ParseFloat((*qi)[i], 32)
		result += math.Pow(pif-qif, 2)
		i--
	}
	result = math.Sqrt(result)

	return
}
func getMapValues(m *map[float64]string) (values []string) {
	for _, v := range *m {
		values = append(values, v)
	}
	return
}
func getMapKeys(m *map[float64]string) (keys []float64) {
	for k := range *m {
		keys = append(keys, k)
	}
	return
}
func getKnn(list *map[float64]string, k int) (sortedMap map[float64]string) {
	keys := getMapKeys(&(*list))

	sort.Float64s(keys)

	sortedMap = make(map[float64]string)

	for i, key := range keys {
		if i < k {
			sortedMap[key] = (*list)[key]
		} else {
			break
		}
	}
	return
}
func getPredominantClass(knn *map[float64]string) (class string) {
	mapValues := getMapValues(&(*knn))
	classes := distinct(&mapValues)
	var predominantClass int
	for c := range classes {
		var countClass int
		for i := range *knn {
			if (*knn)[i] == classes[c] {
				countClass++
			}
		}
		if predominantClass < countClass {
			predominantClass = countClass
			class = classes[c]
		}
	}
	return
}
//PrepareDataset divide a dataset in x percet to test and the rest get to test
func PrepareDataset(percent float32, records [][]string) (train [][]string, test [][]string) {
	column := getCollum(&records, len(records[0])-1)
	classes := distinct(&column)
	for i := range classes {
		values := getValuesByClass(&records, classes[i])
		trainRecords, testRecords := divideInPercent(&values, percent)
		train = append(train, trainRecords[0:]...)
		test = append(test, testRecords[0:]...)
	}
	return
}
//Classify is a method that run all knn algorithm  to predict the claass of a new data
func Classify(train [][]string, dataToPredict []string, k int) (result string) {
	dists := make(map[float64]string)
	for i := range train {
		class := train[i][len(train[i])-1]
		d := euclideanDist(&train[i], &dataToPredict)
		dists[d] = class
	}
	knn := getKnn(&dists, k)
	result = getPredominantClass(&knn)

	return
}
