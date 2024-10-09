// Пусть это будет программа, в которой есть пять горутин, слушающих ввод с
// клавиатуры: при вводе сообщения с последующим нажатием клавиши Enter они
// реагируют на это событие, и каждая из них выводит обратно в консоль
// введённую строку.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"unicode/utf8"

	wg "wgroup"
)

const amountOfGoroutines int = 5
const quitCommand string = "quit"

func main() {
	var message string
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	// Для ожидания запуска слушающих горутин
	var startWg *wg.WaitGroup = wg.NewWaitGroup()
	// Для ожидания завершения работы всех запущенных слушающих горутин
	var closeWg *wg.WaitGroup = wg.NewWaitGroup()
	startWg.Add(amountOfGoroutines)
	closeWg.Add(amountOfGoroutines)
	c := sync.NewCond(&sync.Mutex{})
	for i := 0; i < amountOfGoroutines; i++ {
		go func(id int) {
			defer closeWg.Done()
			var outPutMessage string
			// Уведомляем главную горутину, что очередная слушающая
			// горутина запущена
			startWg.Done()
			for {
				c.L.Lock()
				c.Wait()
				c.L.Unlock()
				// Сравниваем введённую строку с командой выхода, не
				// учитывая регистр символов
				if strings.EqualFold(message, quitCommand) {
					return
				}
				if utf8.RuneCountInString(message) == 0 {
					outPutMessage = fmt.Sprintf(
						"Горутина №%d обработала событие: Осуществлен ввод пустой строки", id)
				} else {
					outPutMessage = fmt.Sprintf(
						"Горутина №%d обработала событие: Осуществлен ввод строки: \"%s\"", id, message)
				}
				fmt.Println(outPutMessage)
			}
		}(i)
	}
	// Ожидаем запуска всех слушающих горутин
	startWg.Wait()

	for {
		// В программе использовался метод Scan  объекта Scanner пакета bufio,
		// так как он позволяет выполнять блокирующее чтение с указанного источника.
		scanner.Scan()
		message = scanner.Text()
		fmt.Println("------------------------------------------------------------")
		c.Broadcast()
		// Сравниваем введённую строку с командой выхода, не учитывая
		// регистр символов
		if strings.EqualFold(message, quitCommand) {
			// Ждём завершение работы всех запущенных горутин
			closeWg.Wait()
			fmt.Println("Выход из программы")
			return
		}
	}
}
