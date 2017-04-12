package service_functions

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"strings"
)

//using while parsing sites
func RandomSleep(args ...int) {
	r := rand.Int63n(2) + 1
	for _, val := range args {
		r += int64(val)
	}

	time.Sleep(time.Duration(r) * time.Second)
}

//using while debuging
func SaveContentTofile(content []byte, filename string) {
	output, err := os.Create("./static/" + filename)
	defer output.Close()
	if err != nil {
		fmt.Println("Error while creating", err)
		return
	}
	_, err = output.Write(content)
	if err != nil {
		fmt.Println("Error while writing", err)
		return
	}
}

//while parsing some turkish web sites, turkish characters sometimes not valid
//checking at first with utf8.ValidString
func ReplaceTurkishNotValid(text string) string  {
	text = strings.Replace(text,"\xE7","ç",-1)
	text = strings.Replace(text,"\xC7","Ç",-1)

	text = strings.Replace(text,"\x131","ı",-1)
	text = strings.Replace(text,"\x30","İ",-1)

	text = strings.Replace(text,"\xD6","Ö",-1)
	text = strings.Replace(text,"\xF6","ö",-1)

	text = strings.Replace(text,"\xDC","Ü",-1)
	text = strings.Replace(text,"\xFC","ü",-1)

	text = strings.Replace(text,"\x11E","Ğ",-1)
	text = strings.Replace(text,"\x11F","ğ",-1)

	text = strings.Replace(text,"\x15E","Ş",-1)
	text = strings.Replace(text,"\x15F","ş",-1)

	text = strings.Replace(text,"\x20A4","TL",-1)

	return text
}
