package main

import (
	"class-award-formatter/util"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func randShuffle(slice []string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

func main() {
	collection := util.LoadData("data.xlsx")
	ansStr := ""
	cnt := 0
	for _, student := range collection {
		award := student.Award
		randShuffle(award)
		ansStr += fmt.Sprintf("%s：%s\n\n", student.Name, strings.Join(student.Award, string('，')))
		cnt += len(student.Award)
	}
	ansStrByte := []byte(ansStr)
	err := ioutil.WriteFile("output.txt", ansStrByte, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully load:", cnt)
}
