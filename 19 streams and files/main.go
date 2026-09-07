package main

import (
	"fmt"
	"streams-files/createOpenFile"
	"streams-files/fscan"
	"streams-files/ioStream"
	"streams-files/readFromConsole"
	"streams-files/readWriteFile"
	"streams-files/readerWriter"
)

func main() {
	phone1 := readerWriter.PhoneReader("+1(234)567 9010")
	phone2 := readerWriter.PhoneReader("+2-345-678-12-35")

	buffer1 := make([]byte, len(phone1))
	// Передаем срез байтов в метод Read типа PhoneReader
	phone1.Read(buffer1)
	// В этот срез записываются числа из номера телефона
	fmt.Println(string(buffer1))

	buffer2 := make([]byte, len(phone2))
	phone2.Read(buffer2)
	fmt.Println(string(buffer2))

	phone3 := []byte("+1(234)567 9010")
	phone4 := []byte("+2-345-678-12-35")

	pw := readerWriter.PhoneWriter{}
	pw.Write(phone3)
	pw.Write(phone4)
	// --------------------------------------
	fmt.Println()
	fmt.Println()
	createOpenFile.FileManipulation()

	// --------------------------------------
	fmt.Println()
	fmt.Println()
	readWriteFile.ReadWriteFile()

	// --------------------------------------
	fmt.Println()
	fmt.Println()
	ioStream.IoStream()

	// --------------------------------------
	fmt.Println()
	fmt.Println()
	fscan.WriteData()
	fscan.ReadData()

	// --------------------------------------
	fmt.Println()
	fmt.Println()
	readFromConsole.ReadFromConsole()
}
