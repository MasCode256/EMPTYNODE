package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// Обработчик для запросов с параметрами
func queryHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры запроса
	query := r.URL.Query()
	response := ""

	// Перебираем параметры и добавляем их к ответу
	if query.Has("cmd") {
		response += executeCommand(query.Get("cmd"))
	}

	w.Write([]byte(response))
}

// Обработчик для запросов на получение файлов
func fileHandler(w http.ResponseWriter, r *http.Request) {
	// Укажите путь к файлу, который хотите вернуть
	filePath := "./html" + r.URL.Path // Например, если запрашивается /example.txt, файл будет в текущей директории

	http.ServeFile(w, r, filePath)
}

func main() {
	// Обрабатываем запросы на корневой путь и добавляем обработчик для файлов
	http.HandleFunc("/", fileHandler) // Обрабатываем все запросы на файлы
	http.HandleFunc("/query", queryHandler) // Обрабатываем запросы на /query

	address := "127.0.0.1:3223" // Указываем адрес и порт
	log.Printf("Сервер запущен на %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}
}



func executeCommand(fullCommand string) string {
	if fullCommand == "" {
		return ""
	}

	fmt.Printf("Executing command: %s\n", fullCommand)

	var cmd *exec.Cmd

	// Определяем команду в зависимости от операционной системы
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe", "/C", fullCommand)
	} else {
		cmd = exec.Command("bash", "-c", fullCommand)
	}

	var out bytes.Buffer
	var errOut bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errOut

	err := cmd.Run()
	if err != nil {
		// Выводим стандартный вывод и ошибки для диагностики
		fmt.Println("Ошибка выполнения команды:", err)
		if errOut.Len() > 0 {
			fmt.Println("Ошибка вывода:", errOut.String())
		}
		return "error:30"
	}

	return out.String()
}

func countFilesInDirectory(directory string) int {
	count := 0

	// Используем filepath.Walk для обхода файлов в директории
	err := filepath.Walk(directory, func(_ string, info os.FileInfo, _ error) error {
		if !info.IsDir() { // Если это не директория
			count++
		}
		return nil
	})

	if err != nil {
		fmt.Println("Ошибка при обходе директории:", err)
		return 0}

	return count
}

func int_to_str(num int) string {
	ret := strconv.Itoa(num)
	return ret
}

func str_to_int(str string) int {
	ret, _ := strconv.Atoi(str)
	return ret
}

func removeFile(filePath string) error {
	// Используем os.Remove для удаления файла.
	err := os.Remove(filePath)
	if err != nil {
		return err}
	return nil
}


func CopyFile(src, dst string) error {
    // Открываем исходный файл
    sourceFile, err := os.Open(src)
    if err != nil {
        return err }
    defer sourceFile.Close()

    if _, err := os.Stat(dst); err == nil {
        return fmt.Errorf("файл назначения уже существует")
    } else if !os.IsNotExist(err) {
        return err }

    // Создаем файл назначения
    destinationFile, err := os.Create(dst)
    if err != nil {
        return err }
    defer destinationFile.Close()

    // Копируем данные из исходного файла в файл назначения
    _, err = io.Copy(destinationFile, sourceFile)
    if err != nil {
        return err }

    // Опционально: копируем права доступа
    sourceInfo, err := os.Stat(src)
    if err != nil {
        return err
    }
    err = os.Chmod(dst, sourceInfo.Mode())
    if err != nil {
        return err }

    return nil
}


func before(str string, c byte) string {
	// Используем strings.IndexByte для поиска первого вхождения символа c
	index := strings.IndexByte(str, c)
	if index == -1 {
// Если символ не найден, возвращаем всю строку
		return str
	}
	// Возвращаем подстроку от начала до найденного индекса
	return str[:index]
}


func after(str string, c byte) string {
	// Используем strings.IndexByte для поиска первого вхождения символа c
	index := strings.IndexByte(str, c)
	if index == -1 {
		// Если символ не найден, возвращаем пустую строку
		return ""
	}
	// Возвращаем подстроку от символа, следующего за найденным индексом, до конца строки
	return str[index+1:]
}

