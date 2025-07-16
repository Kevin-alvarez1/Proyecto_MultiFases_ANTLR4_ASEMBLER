package cst

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

type CSTresponse struct {
	SVGTree string `json:"svgtree"`
}

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
func CstReport(input string) string {
	_, filename, _, _ := runtime.Caller(0)
	path := filepath.Dir(filename)

	gramaticaPath := filepath.Join(path, "..", "..", "backend", "analizador", "parser", "gramatica.g4")

	absPath, _ := filepath.Abs(gramaticaPath)
	fmt.Println("Leyendo gramática en:", absPath)

	gramaticaContent, err := ReadFile(gramaticaPath)
	if err != nil {
		fmt.Println("Error leyendo la gramática:", err)
		return ""
	}

	gramaticaJSON, err := json.Marshal(gramaticaContent)
	if err != nil {
		fmt.Println("Error haciendo marshal de gramática:", err)
		return ""
	}

	jinput, err := json.Marshal(input)
	if err != nil {
		fmt.Println("Error haciendo marshal de input:", err)
		return ""
	}

	payload := []byte(fmt.Sprintf(
		`{
        "grammar": %s,
        "input": %s,
        "lexgrammar": %s,
        "start": "%s"
    }`,
		string(gramaticaJSON),
		string(jinput),
		string(gramaticaJSON),
		"init",
	))

	req, err := http.NewRequest("POST", "http://lab.antlr.org/parse/", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creando request:", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error enviando request:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error leyendo respuesta:", err)
		return ""
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
		return ""
	}

	resultRaw, ok := data["result"]
	if !ok {
		fmt.Println("La respuesta no contiene 'result'")
		return ""
	}

	result, ok := resultRaw.(map[string]interface{})
	if !ok {
		fmt.Println("'result' no es un objeto")
		return ""
	}

	svgtreeRaw, ok := result["svgtree"]
	if !ok {
		fmt.Println("No existe 'svgtree' en 'result'")
		return ""
	}

	svgtree, ok := svgtreeRaw.(string)
	if !ok {
		fmt.Println("'svgtree' no es un string")
		return ""
	}

	return svgtree
}
