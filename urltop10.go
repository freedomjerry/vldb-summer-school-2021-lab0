package main

import (
	"bytes"
	"fmt"
	//"sort"
	"strconv"
	"strings"
	//"sync/atomic"
)

// URLTop10 .

func URLTop10(nWorkers int) RoundsArgs {
	var args RoundsArgs
	args = append(args, RoundArgs{
		MapFunc   : MyURLCountMap,
		ReduceFunc: MyURLCountReduce,
		NReduce   : nWorkers,
	})

	args = append(args, RoundArgs{
		MapFunc   : MyURLTop10Map,
		ReduceFunc: MyURLTop10Reduce,
		NReduce   : 1,
	})
	return args
}

func MyURLCountMap(filename string, contents string) []KeyValue {
	urls := strings.Split(string(contents),"\n")
	Cnturl := make(map[string]int)
	for _, url := range urls {
		url = strings.TrimSpace(url)
		if len(url) == 0 {
			continue
		}
		Cnturl[url] ++
	}
	Kv := make([]KeyValue, 0, len(Cnturl))
	for url, count := range Cnturl{
		Kv = append(Kv, KeyValue{Key: url, Value: strconv.Itoa(count)})
	}
	return Kv//GET THE URL COUNT FOR ONE FILE(KEY , COUNT)
}

func MyURLCountReduce(Key string, Values []string) string {
		cnt := 0
		for _,value := range Values{
			value, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			cnt += value////SUM THE KEY COUNT OF ALL FILES
		}
		return fmt.Sprintf("%s %s\n", Key, strconv.Itoa(cnt))
}

func MyURLTop10Map(filename string, contents string) []KeyValue {
	//log.Println("filename: ", filename, "contents: ", contents)
	lines := strings.Split(contents, "\n")
	ucnt := make([]*urlCount, 0, len(lines))
	//log.Println("lines ", lines)
	for _, line := range lines {
		line := strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		words := strings.Split(line, " ")
		val, err := strconv.Atoi(words[1])
		if err != nil {
			panic(err)
		}
		ucnt = append(ucnt, &urlCount{words[0], val})
	}
	//log.Println("size of lines", len(lines)) {
		kvs := Top10heap(ucnt)
		//log.Println("kvs :  ", kvs)

	return kvs

}
func MyURLTop10Reduce(key string, values []string) string  {
	/**
	"" 【“url count”, "url count", ""】
	 */
	cnts := make([]*urlCount, len(values))
	for _, v := range values {
		v := strings.TrimSpace(v)
		if len(v) == 0 {
			continue
		}
		tmp := strings.Split(v, " ")
		n, err := strconv.Atoi(tmp[1])
		if err != nil {
			panic(err)
		}
		cnts = append(cnts, &urlCount{tmp[0], n})
	}
	foreresult := Top10heap(cnts)
	buf := new(bytes.Buffer)
	length := len(foreresult)
	for i := range foreresult {
		tmp :=strings.Split(foreresult[length-1-i].Value," ")
		fmt.Fprintf(buf, "%s: %s\n", tmp[0], tmp[1])
	}
	return buf.String()
}

