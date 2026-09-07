package createOpenFile

import (
	"fmt"
	"os"
)

func FileManipulation() {
	// С помощью пакета os можно взаимодействовать с файлами
	// Создание файла
	file, err := os.Create("hello.txt")

	if err == nil {
		fmt.Println(file)
		// Закрываем файл
		file.Close()
	}

	// Открытие файла

	file2, err2 := os.Open("hello.txt")

	if err2 == nil {
		fmt.Println(file2)
		// Закрываем файл
		file2.Close()
	}

	// Открываем файл, либо создаем его если его до этого не было
	// Передаем:
	// Название файла
	// Режим открытия файла (Только для чтения в данном случае)
	// Разрешения для доступа к файлу
	file3, err3 := os.OpenFile("hello2.txt", os.O_RDONLY|os.O_CREATE, 0666)
	if err3 == nil {
		fmt.Println(file3)
		fmt.Println(file3.Name(), "имя файла")
		// Закрываем файл
		file3.Close()
	}

}
