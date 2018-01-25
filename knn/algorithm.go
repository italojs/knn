package knn

import (
	"sort"
	"math"
	"strconv"
)

func Distinct(elements *[]string) (result []string) {
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
func GetCollum(elements *[][]string, ColumnIndex int) (column []string) {
	for i := range *elements {
		column = append(column, (*elements)[i][ColumnIndex])
	}
	return
}
func GetValuesByClass(records *[][]string, class string) (newRecords [][]string) {
	newRecords = make([][]string, 0)
	for l := range *records {
		if (*records)[l][len((*records)[l]) - 1] == class{
			newRecords = append(newRecords, (*records)[l])
		}
	}
	return
}
func DivideInPercent(records *[][]string, percent float32)(newRecords [][]string, residue [][]string){
	i := float32(len(*records)) * percent
	for i > 0{
		index := int(i-1)
		newRecords = append(newRecords,(*records)[index])
		i--
	}

	i = float32(len(*records)) * percent
	j := float32(len(*records))
	for j>i{
		index := int(j-1)
		residue = append(residue,(*records)[index])
		j--
	}
	return
}
func EuclideanDist(pi *[]string, qi *[]string)(result float64){
	i := len(*pi) -1

	for i>0 {
		pif,_ := strconv.ParseFloat((*pi)[i-1], 32)
		qif,_ := strconv.ParseFloat((*qi)[i-1], 32) 
		result += math.Pow(pif - qif , 2)
		i--
	}
	result = math.Sqrt(result)
	
	return 
}
func GetMapKeys(m *map[float64]string)(keys []float64){
	for k := range *m {
	   keys = append(keys, k)
	}
	return
}
func GetMapValues(m *map[float64]string)(values []string){
	for _,v := range *m {
	   values = append(values, v)
	}
	return
}
func GetKnn(list *map[float64]string, k int)(sortedMap map[float64]string){
	 keys := GetMapKeys(&(*list))

	 sort.Float64s(keys)
 
	 sortedMap = make(map[float64]string)
	 
	 for i, key := range keys {
		 if i < k{
			sortedMap[key] = (*list)[key]
		 }else{
			 break
		 }
	 }
	 return
}
func GetPredominantClass(knn *map[float64]string)(class string){
	mapValues := GetMapValues(&(*knn))
	classes := Distinct(&mapValues)
	var predominantClass int
	for c := range classes{
		var countClass int
		for i := range *knn{
			if (*knn)[i] == classes[c]{
				countClass++
			}
		}
		if predominantClass < countClass{
			predominantClass = countClass
			class = classes[c]
		}
	}
	return 
}
func Classify(train [][]string, valueToPredict []string, k int)(result string){
	dists := make(map[float64]string)
	i := len(train) - 1
	for i>=0 {
		class := train[i][len(train[i]) - 1]
		d := EuclideanDist(&train[i], &valueToPredict)
		dists[d] = class
		i--
	}
	knn := GetKnn(&dists, k)
	result = GetPredominantClass(&knn)

	return
}