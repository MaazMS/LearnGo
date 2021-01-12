package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//rand.Seed(time.Now().Unix())
	//fmt.Print(rand.Intn(200), ",")
	//take := add()

	no1 := randomno()
	fmt.Printf("%d \t\n", no1)
	s1 := mysql()
	fmt.Printf("%s \t\n", s1)

	no1s1 := join()
	fmt.Printf("%s \t \n", no1s1)

}
func randomno() int {
	rand.Seed(time.Now().Unix())
	number := rand.Intn(10)
	//fmt.Print(rand.Intn(200), ",")
	return number
}

func mysql() string {

	mysqlauth := "mysql-auth "
	return mysqlauth
}

func join() string {
	rand.Seed(time.Now().Unix())
	number := rand.Intn(10)
	mysqlauth := "mysql-auth "
	concatinate := fmt.Sprintf("%d%s", number, mysqlauth)

	return concatinate
}
