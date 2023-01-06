package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"math/big"
	"bufio"
	crypto_rand "crypto/rand"
)

var Proxy    	  = []string{}

func GetInt(max int) (int, error) {
	if max <= 0 {
		return 0, fmt.Errorf("can't define input as <=0")
	}
	nbig, err := crypto_rand.Int(crypto_rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return max, err
	}
	n := int(nbig.Int64())

	return n, err
}
func Choice(j []string) (string, error) {
	i, err := GetInt(len(j))
	if err != nil {
		return "", err
	}

	return j[i], nil
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	f, err := os.Open("proxies.txt")
	if err != nil {
        fmt.Println(err)
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    for scanner.Scan(){
    	result := strings.Split((scanner.Text()),"] ")[1]
    	tot := strings.Replace(result, ">","",1)
    	Proxy = append(Proxy, tot)
    }
    if err := scanner.Err(); err != nil {
    	fmt.Println(err)
    }
    prox, _ := Choice(Proxy)
    proxx := strings.Split((prox),":")[0]
    Proxy = proxx
    fmt.Println(Proxy)
}