package readerWriter

import (
	"fmt"
	"io"
)

// Методы Read и Write относятся к интерфейсам io.Read и io.Write

type PhoneReader string

func (ph PhoneReader) Read(p []byte) (int, error) {
	count := 0
	for i := 0; i < len(ph); i++ {
		if ph[i] >= '0' && ph[i] <= '9' {
			p[count] = ph[i]
			count++
		}
	}
	// Возвращаем количество считанных байт и если в потоке больше
	// Нет данных то возвращаем ошибку типа io.EOF
	return count, io.EOF
}

type PhoneWriter struct{}

// Write занимается копированием байтов из среза и выводом их куда-то
// В нашем случае мы выводим их в консоль
func (pw PhoneWriter) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	for i := 0; i < len(p); i++ {
		if p[i] >= '0' && p[i] <= '9' {
			fmt.Print(string(p[i]))
		}
	}
	fmt.Println()
	// Возвращаем длину слайса и ошибку
	return len(p), nil
}
