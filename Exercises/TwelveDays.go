package Exercises

import (
	"fmt"
	"io/ioutil"
)

/*
  Outputs the lyrics to 12 days of christmas
*/
func Output(lyrics string) string {
	bs, err := ioutil.ReadFile(lyrics)

	if err != nil {
		fmt.Println("Error")
	}

	str := string(bs)
	fmt.Println(str)

	return str

}
