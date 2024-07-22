package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var in *bufio.Reader
var out *bufio.Writer

func prettifyJSON(inputJSON interface{}) interface{} {
	switch v := inputJSON.(type) {
	case []interface{}:
		var result []interface{}
		for _, item := range v {
			if item != nil {
				prettifiedItem := prettifyJSON(item)
				if prettifiedItem != nil {
					result = append(result, prettifiedItem)
				}
			}
		}
		return result
	case map[string]interface{}:
		result := make(map[string]interface{})
		for key, value := range v {
			if value != nil && value != "null" { // Проверяем значение на "null"
				prettifiedValue := prettifyJSON(value)
				if prettifiedValue != nil {
					result[key] = prettifiedValue
				}
			}
		}
		// Удаляем пустые элементы
		if len(result) == 0 {
			return nil
		}
		return result
	default:
		return v
	}
}

func main() {
	// Буферизованный ввод и вывод
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// Чтение числа наборов входных данных
	var numSets int
	fmt.Fscan(in, &numSets)
	// Пропускаем символ новой строки после числа наборов
	in.ReadString('\n')

	// Проход по каждому набору входных данных
	for i := 0; i < numSets; i++ {
		// Чтение числа строк в JSON-объекте
		var numLines int
		fmt.Fscan(in, &numLines)
		// Пропускаем символ новой строки после числа строк
		in.ReadString('\n')

		// Считываем строки JSON-объекта
		var inputStrings []string
		for j := 0; j < numLines; j++ {
			inputString, _ := in.ReadString('\n')
			inputStrings = append(inputStrings, strings.TrimSpace(inputString))
		}

		// Преобразование входных данных в JSON
		var inputData interface{}
		err := json.Unmarshal([]byte(strings.Join(inputStrings, "\n")), &inputData)
		if err != nil {
			fmt.Fprintln(out, "Ошибка при разборе входных данных:", err)
			return
		}

		// Применение операции prettify
		prettifiedData := prettifyJSON(inputData)

		// Преобразование результата в JSON и вывод
		prettifiedJSON, err := json.Marshal(prettifiedData)
		if err != nil {
			fmt.Fprintln(out, "Ошибка при преобразовании результата в JSON:", err)
			return
		}
		prettifiedString := string(prettifiedJSON)
		prettifiedString = strings.ReplaceAll(prettifiedString, "[null", "[")
		prettifiedString = strings.ReplaceAll(prettifiedString, ",null", ",")
		prettifiedString = strings.ReplaceAll(prettifiedString, "null,", "")
		prettifiedString = strings.ReplaceAll(prettifiedString, "null]", "]")

		fmt.Fprint(out, prettifiedString)
	}
}
