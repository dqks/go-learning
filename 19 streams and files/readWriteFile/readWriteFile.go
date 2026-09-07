package readWriteFile

import (
	"fmt"
	"io"
	"os"
)

func ReadWriteFile() {
	file, err := os.OpenFile("hello.txt", os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println("Ошибка открытия файла")
		return
	}
	defer file.Close()

	// Для записи текстовой информации применяем метод WriteString()
	file.WriteString("Hello, World!")
	fmt.Println("Done.")

	// Для записи бинарнйо информации используем метод Write()
	file.Write([]byte("\nВсем привет"))

	// Для чтения из файла используем метод Read()
	file2, err2 := os.Open("hello.txt")

	if err2 != nil {
		fmt.Println("Ошибка открытия файла")
		return
	}

	defer file2.Close()

	// 1 способ считывания файла, тут он в консоль
	// data := make([]byte, 64)
	// for {
	// 	n, err := file2.Read(data)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	// Преобразуем срез байтов в строку
	// 	fmt.Print(string(data[:n]))
	// }

	// 2 способ считывания файла, тут он также в консоль
	// Но можно указать разное место вывода
	io.Copy(os.Stdout, file2)
}
