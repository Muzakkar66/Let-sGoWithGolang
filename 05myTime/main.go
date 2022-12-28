package main

import (
	"fmt"
	"time"
)

func main()  {
	mycurrentTime := time.Now()

	fmt.Println(mycurrentTime)

	fmt.Println(mycurrentTime.Format("2006-01-02"))

	fmt.Println(mycurrentTime.Format(time.UnixDate))

}