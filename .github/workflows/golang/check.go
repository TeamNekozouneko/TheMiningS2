package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"bufio"
)

func main(){
	log("main", "Runtime Results Checking...")

	//実行結果の取得
	const LOG_FILE_PATH = "logs/latest.log"
	data, err := ioutil.ReadFile(LOG_FILE_PATH)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logContent := string(data)

	//ランタイム実行の結果処理
	if !compileCheck(logContent) { os.Exit(1) }
	log("compileCheck", "passed.")
	if !runtimeCheck(logContent) { os.Exit(1) }
	log("runtimeCheck", "passed.")

	//終了処理
	log("main", "Runtime Results are all passed successfully.")

}

func log(logName string, logContent string){
	fmt.Println("[TheMiningS2] "+logName+": "+logContent)
}

//コンパイルテスト
func compileCheck(logContent string) bool{
	const COMPILE_ERROR_DETECTION_STRING = "[Skript] Line"
	if strings.Contains(logContent, COMPILE_ERROR_DETECTION_STRING) { return false }
	return true
}

//ランタイムテスト
func runtimeCheck(logContent string) bool{
	
	//ランタイムエラー検出
	const RUNTIME_ERROR_DETECTION_STRING = "THEMININGS2_TEST_FAILED"
	if strings.Contains(logContent, RUNTIME_ERROR_DETECTION_STRING) { return false }

	//1-LogQuantityModuloCheck
	const SYSTEM_LOG_FILE_PATH = "plugins/Skript/logs/themining.log"
	const AMOUNT_OF_LOG_TYPES = 3
	
	file, err := os.Open(SYSTEM_LOG_FILE_PATH)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	logQuantity := 0
	for scanner.Scan() {
		logQuantity += strings.Count(scanner.Text(), "RUNTIME_TEST_LOG")
	}
	if logQuantity % AMOUNT_OF_LOG_TYPES != 0 {
		log("1-LogQuantityModuloCheck", "failed.")
		os.Exit(1)
	}

	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}

	return true
}
