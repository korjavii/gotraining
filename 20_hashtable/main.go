package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}
	sc := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	sc.Split(bufio.ScanWords)
	buckets := make([]int, 200)

	for sc.Scan() {
		n := HashBucket(sc.Text())
		buckets[n]++
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading output: ", err)
	}

	fmt.Println(buckets[65:122])
}

func HashBucket(word string) int {
	return int(word[0])
}
