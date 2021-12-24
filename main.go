package main

import (
	"bufio"
	"encoding/json"
	"fmt"

	//"io/ioutil"
	"os"
	"strings"
)

type Elements struct {
	Value []string
}

type Holder struct {
	Value map[string]string
}

type ToJson struct {
	Value []Holder
}

func main() {
	var st string
	fmt.Println("Write a filename")
	fmt.Scan(&st)
	f, err := os.Open(st)
	wf, err := os.Create("result.json")

	if err != nil {
		fmt.Println("File does not exists")
	}
	defer f.Close()
	defer wf.Close()
	scanner := bufio.NewScanner(f)
	i := 0
	temp := 0

	elements := Elements{}
	holder := Holder{}
	holder.Value = map[string]string{}
	toJson := ToJson{}
	toJson.Value = []Holder{}
	var keys []string
	var all = "[\r\n"
	for scanner.Scan() {
		if i == 0 {
			keys = strings.Fields(scanner.Text())
			elements.Value = keys
			temp = len(keys)
		} else {
			values := strings.Fields(scanner.Text())
			for j := 0; j < temp; j++ {
				holder.Value[keys[j]] = values[j]
			}
			jsonst, err := json.Marshal(holder.Value)
			if err != nil {
				fmt.Println("Json Encoding Failed")
			}
			fmt.Println(string(jsonst))
			content := string(jsonst)
			if len(scanner.Bytes()) != 0 {
				content += ",\r\n"
			}
			//_, err = wf.Write([]byte(content))
			all += string(content)
			if err != nil {
				fmt.Println("Failed Writing")
			}
			//toJson.Value = append(toJson.Value, holder)
		}
		i++
	}
	all += "]"
	strings.NewReplacer("},\r\n]", "}\r\n]").Replace(all)
	wf.Write([]byte(all))
}
