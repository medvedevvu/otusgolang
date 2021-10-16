package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type Rec struct {
	K string
	V int
}

type ByMx []Rec

func (a ByMx) Len() int      { return len(a) }
func (a ByMx) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByMx) Less(i, j int) bool {
	less := a[i].V > a[j].V
	if a[i].V == a[j].V {
		less = a[i].K < a[j].K
	}
	return less
}

func Top10(inStr string) []string {
	if len(inStr) == 0 {
		return nil
	}
	sliceOfStrings := strings.ReplaceAll(inStr, "\n\t", " ")
	submatchall := strings.Split(sliceOfStrings, " ")
	sliceOfWords := make([]string, 1, len(inStr))
	for _, el := range submatchall {
		sliceOfWords = append(sliceOfWords, strings.Trim(el, " "))
	}
	mapFreq := make(map[string]int)
	for _, v := range sliceOfWords {
		if v == "" {
			continue
		}
		mapFreq[v]++
	}
	// отсортируем значения
	valFreq := make([]int, 0, len(mapFreq))
	// var valFreq []int
	for _, val := range mapFreq {
		valFreq = append(valFreq, val)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(valFreq)))
	tenKeys := valFreq[:10]
	// взять первые 10
	var outString []string
	var sortedOutString []Rec
	for _, vx := range tenKeys {
		for k, v := range mapFreq {
			if vx == v && !containValue(outString, k) {
				outString = append(outString, k)
				sortedOutString = append(sortedOutString, Rec{K: k, V: v})
			}
		}
	}
	sort.Sort(ByMx(sortedOutString))
	fnlString := make([]string, 0, len(sortedOutString))
	// var fnlString []string
	for _, v := range sortedOutString {
		fnlString = append(fnlString, v.K)
	}
	return fnlString
}

func containValue(vstr []string, s string) bool {
	if len(vstr) == 0 {
		return false
	}
	for _, v := range vstr {
		if s == v {
			return true
		}
	}
	return false
}
