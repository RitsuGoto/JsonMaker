package function

import (
	"bufio"
	"encoding/json"
	"fmt"
	"jsonmaker/model"
	"os"
	"strings"
)

func GetFileText(st string) []string {
	result := []string{}
	f, err := os.Open(st)
	if err != nil {
		fmt.Println("File does not exist")
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}

func TextToJson(fvalue []string) string {
	all := "[\r\n"
	var keys []string
	temp := 0
	elements := model.Elements{}
	holder := model.Holder{}
	for i, st := range fvalue {
		if i == 0 {
			keys = strings.Fields(st)
			elements.Value = keys
			temp = len(keys)
		} else {
			values := strings.Fields(st)
			holder.Value = map[string]string{}
			for j := 0; j < temp && j < len(values); j++ {
				holder.Value[keys[j]] = values[j]
			}
			jsonst, err := json.Marshal(holder.Value)
			if err != nil {
				fmt.Println("Json Encoding Failed")
			}
			fmt.Println(string(jsonst))
			content := string(jsonst)
			all += string(content)
			if err != nil {
				fmt.Println("Failed Writing")
				panic(err)
			}
		}
		i++
	}
	all = strings.Replace(all, "}{", "},\r\n{", -1)
	all += "\r\n]"
	return all
}

func CreateJsonFile(value string) {
	f, err := os.Create("res.json")
	if err != nil {
		fmt.Println("File does not exist")
		panic(err)
	}
	f.WriteString(value)
}

func CreateJson(st string) {
	f, err := os.Open(st)
	if err != nil {
		fmt.Println("File does not exist")
		panic(err)
	}
	defer f.Close()

	wf, err := os.Create("result.json")
	if err != nil {
		fmt.Println("Can't make file")
		panic(err)
	}
	defer wf.Close()
	scanner := bufio.NewScanner(f)
	i := 0
	temp := 0

	elements := model.Elements{}
	holder := model.Holder{}

	var keys []string
	var all = "[\r\n"
	for scanner.Scan() {
		if i == 0 {
			keys = strings.Fields(scanner.Text())
			elements.Value = keys
			temp = len(keys)
		} else {
			values := strings.Fields(scanner.Text())
			holder.Value = map[string]string{}
			for j := 0; j < temp && j < len(values); j++ {
				holder.Value[keys[j]] = values[j]
			}
			jsonst, err := json.Marshal(holder.Value)
			if err != nil {
				fmt.Println("Json Encoding Failed")
			}
			fmt.Println(string(jsonst))
			content := string(jsonst)
			all += string(content)
			if err != nil {
				fmt.Println("Failed Writing")
				panic(err)
			}
		}
		i++
	}
	all = strings.Replace(all, "}{", "},\r\n{", -1)
	all += "\r\n]"
	wf.WriteString(all)
}
