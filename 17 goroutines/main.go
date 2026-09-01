package main

import (
	"fmt"
	"sync"
	"time"
)

func sum(n int) {
	result := 0
	for i := 0; i < n; i++ {
		result += i
	}
	//fmt.Printf("%d - %d\n", n, result)
}

func work(id int, wg *sync.WaitGroup) {
	fmt.Printf("Начало горутины %d \n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Окончание горутины %d \n", id)
	wg.Done()
}

func main() {
	// main - это уже горутина
	// При достижении конца main все горутины закрываются
	// Обычно горутины ставят для какой-то фукнции
	for i := 0; i < 7; i++ {
		go sum(i)
	}
	// Можно использовать запрос на ввод, что наверняка позволит всем горутинам выполниться
	//fmt.Scanln()
	// либо использовать time.Sleep(), что позволит остановить выполнение программы на 200 мс
	//time.Sleep(200 * time.Millisecond)
	// Вывод в консоль чаще всего хаотичный, поскольку горутины заканчиваются в разном порядке

	// -------------------------------------------------------------------
	// Также можно добавлять WaitGroup - счетчик, который может блокировать выполнение фукнции
	// Пока счетчик не заполнится
	//var wg sync.WaitGroup
	//// Добавляем в ожидание 2 горутины
	//wg.Add(2)
	//
	//go work(1, &wg)
	//go work(2, &wg)
	//
	//// Блокируем выполнение фукнции, пока 2 горутины не выполнятся
	//wg.Wait()
	//
	//fmt.Println("Конец горутин для wg")

	// -------------------------------------------------------------------
	// Канал нужен для обмена данными между горутинами
	// make резервирует место в хранилище и создает базовый заголовок для канала
	//var channel chan int = make(chan int)
	//
	//go receiveNumber(channel)
	//
	//// После объявления горутины выше
	//// main продолжает свою работу, но дальше требуется чтение канала
	//// И main блокируется, пока в канал не попадет значение
	//fmt.Println(<-channel)
	//fmt.Println("The end")

	// -------------------------------------------------------------------
	// Одна горутина должна отправлять данные, а другая получать
	// Одна горутина не может быть отправителем и получателем данных через канал
	fmt.Println()
	fmt.Println()

	channel2 := make(chan int)
	// 1. Вызываем горутину
	go square(channel2)
	// 3. По каналу передается 5
	channel2 <- 5
	//4. Горутина блокируется, ей нужно значение
	// 6. Горутина получает значение
	fmt.Println("result", <-channel2)
	fmt.Println("The end")

	// -------------------------------------------------------------------
	// Чтобы сделать буферезированный канал нужно указать емкость канала при его инициализации
	fmt.Println()
	fmt.Println()

	buffedChannel := make(chan int, 3)
	// При буферезированном канале отправителем и получаетелем может быть сама горутина
	// 1. Отправка значения по каналу
	buffedChannel <- 10
	// 3. Отправка значения по каналу
	buffedChannel <- 20
	// 5. Отправка значения по каналу
	buffedChannel <- 30
	// 2. Получение значения по каналу
	fmt.Println(<-buffedChannel)
	// 4. Получение значения по каналу
	fmt.Println(<-buffedChannel)
	// 6. Получение значения по каналу
	fmt.Println(<-buffedChannel)

	// При этом если одновременное в буфер будет отправлено больше значений, чем емкость
	// То будет блокировка горутины

	// -------------------------------------------------------------------
	// Канал можно закрыть
	fmt.Println()
	fmt.Println()
	channelToClose := make(chan int)
	go getNumber(channelToClose)
	fmt.Println(<-channelToClose)
	// При повторной отправке данных мы получим значение типа по умолчанию
	fmt.Println(<-channelToClose)

	// При чтении данных можно получать не только значение, но и то
	// Закрытый ли канал или нет
	channelToClose2 := make(chan int)
	go getNumber(channelToClose2)
	val, ok := <-channelToClose2
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("Канал закрыт")
	}
	val1, ok1 := <-channelToClose2
	if ok1 {
		fmt.Println(val1)
	} else {
		fmt.Println("Канал закрыт")
	}

	// -------------------------------------------------------------------
	// Есть каналы также только для чтения или только для отправки

	// Канал только для чтения
	//var channelToSend <-chan int = make(<-chan int)
	// Канал только для отправки
	//var channelToRead chan<- int = make(chan<- int)

	// Двунаправленный канал можно использовать в функциях как однонаправленный канал (либо для чтения либо для отправки)
	fmt.Println()
	// Ниже будем применять передавать двунаправленный канал как канал только для отправки
	dualChannel := make(chan int)
	go getNumberNoClose(dualChannel)
	fmt.Println(<-dualChannel, "получение данных из канала только для отправки")

	// Ниже отправим канал только для чтения
	fmt.Println()
	go printChannelNumber(dualChannel)
	dualChannel <- 10000

	fmt.Println()

	// -------------------------------------------------------------------
	// Канал можно использовать для синхронизации данных
	// Результат передаваемый в канал необязательно должен использоваться в горутинах
	fmt.Println()
	fmt.Println()
	results := make(map[int]int)
	ch := make(chan struct{})
	// Передаем канал на закрытие а также мап для записи в него результат
	go getHashTable(10, ch, results)
	// Блокируем горутину, ожидая результат канала
	<-ch

	fmt.Println(results)
	// Хотя в примере выше лучше использовать WaitGroup
	// Поскольку именно для этого он и создан

	// -------------------------------------------------------------------
	// select нужен для выполнения операций над готовым каналов из списка каналов (который стал готовым быстрее)
	// он выполняется одномоментно - 1 select = 1 горутин, даже если в select несколько case
	// default добавляют на случай, если не нужна блокировка фукнционала
	// default не будет блокировать выполнение горутины, в которой он объявлен
	// Если ни одна горутина не готова - выполняется default
	doubleCh := make(chan int)
	squareCh := make(chan int)
	cubeCh := make(chan int)

	go square(squareCh)
	go double(doubleCh)
	go cube(cubeCh)

	squareCh <- 10
	doubleCh <- 2
	cubeCh <- 3

	for i := 0; i <= 2; i++ {
		select {
		case squareVal := <-squareCh:
			fmt.Println(squareVal, "squareVal")
		case doubleVal := <-doubleCh:
			fmt.Println(doubleVal, "doubleVal")
		case cubeVal := <-cubeCh:
			fmt.Println(cubeVal, "cubeVal")
			//default:
			//	fmt.Println("default")
		}
	}

	// -------------------------------------------------------------------
	// Данные из канала можно читать через цикл
	fmt.Println()
	fmt.Println()
	factorialChannel := make(chan int)
	go factorial(10, factorialChannel)

	// Можно написать так
	//for {
	//	val, ok := <-factorialChannel
	//	if !ok {
	//		break
	//	}
	//	fmt.Println(val)
	//}

	// Но так лучше
	for val := range factorialChannel {
		fmt.Println(val)
	}

	// -------------------------------------------------------------------
	// Гонка данных
	// Происходит когда несколько горутин пытаются могут изменить
	// Один и тот же участок памяти, от которого зависит выполнение самой горутины
	// Для профилактики горутины используеся флаг -race при сборке или запуске программы
}

func factorial(n int, ch chan int) {
	defer close(ch)
	result := 1
	for i := 1; i < n; i++ {
		result *= i
		ch <- result
	}
}

//func receiveNumber(channel chan int) {
//	fmt.Println("Начало горутины")
//	// Передача 13 по каналу
//	// И функция main получает это значение и продолжает свою работу
//	channel <- 13
//	fmt.Println("Конец горутины")
//}

func square(channel chan int) {
	// 2. Горутина блокируется на этом моменте, ей нужно значение
	num := <-channel
	//fmt.Println(num)
	// 5. По каналу передатеся 5*5
	channel <- num * num
}

func getNumber(number chan<- int) {
	number <- 10
	// Закрываем канал
	close(number)
	// В закрытый канал мы больше не можем отправлять данные
	// number <- 20
}

func getNumberNoClose(number chan<- int) {
	number <- 10
}

func printChannelNumber(number <-chan int) {
	fmt.Println(<-number)
}

func getHashTable(n int, ch chan struct{}, results map[int]int) {
	for i := 1; i < n; i++ {
		results[i] = i * i
	}
	// Закрытием горутины мы открываем канал
	close(ch)
}

func double(n chan int) {
	val := <-n
	n <- val * 2
}

func cube(n chan int) {
	val := <-n
	n <- val * val * val
}
