package main

import (
	"math"
	"strings"
)

type FileName string
type Word string

type TermFreq map[Word]float64
type InverseDocFreq map[Word]float64
type TfIdf map[Word]map[FileName]float64

func Parse(files map[FileName]string) TfIdf {
	tfs := parallelCalcTfs(files)
	nominator := 1 + len(files)

	idf := InverseDocFreq{}
	for _, tf := range tfs {
		for word := range tf {
			idf[word] = math.Log10(float64(nominator)/(float64(1+howManyDocsHaveTheWord(word, tfs))))
		}
	}
	
	res := TfIdf{}
	for fileName, tf := range tfs {
		for word, tfRes := range tf {
			v, ok := res[word]
			if !ok {
				v = map[FileName]float64{}
			}
			v[fileName] = tfRes * idf[word]
			res[word] = v
		}
	}
	return res
}

func parallelCalcTfs(files map[FileName]string) map[FileName]TermFreq {
	type result struct {
		f FileName
		r TermFreq
	}

	c := make(chan result)
	for f,content := range files {
		go func(name FileName, content string) {
			c <- result{
				f: name, 
				r: tf(content),
			}
		}(f, content)
	}

	tfs := map[FileName]TermFreq{}
	for i := 0; i < len(files); i++ {
		res := <- c
		tfs[res.f] = res.r
	}

	return tfs
}

func howManyDocsHaveTheWord(w Word, tfs map[FileName]TermFreq) (out int) {
	for _, tf := range tfs {
		if _, ok := tf[w]; ok {
			out += 1
		}
	}
	return out
}

func tf(data string) TermFreq {
	words := strings.Fields(data)
	adjustWord := func(w string) string {
		w = strings.ToLower(w)
		w = strings.ReplaceAll(w, "-", "")
		w = strings.ReplaceAll(w, ",", "")
		w = strings.ReplaceAll(w, ".", "")
		w = strings.ReplaceAll(w, "/", "")
		return w
	}

	tf := TermFreq{}
	for _, w := range words {
		tf[Word(adjustWord(w))] += 1.0
	}

	for k,v := range tf {
		tf[k] = float64(v)/float64(len(words))
	}
	return tf
}