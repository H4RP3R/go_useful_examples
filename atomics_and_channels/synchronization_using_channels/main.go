// Вот простейший пример синхронизации доступа к счётчику. Пусть имеется
// переменная, хранящая значение банковского счёта. Необходимо написать код,
// который бы позволил сделать имитацию снятия и пополнения денежных средств
// на него  — то есть такой, который имел бы механизм синхронизации доступа
// к памяти, где хранится сумма счета, но только с использованием каналов.

package main

import (
	"fmt"
	"sync"
)

// Количество транзакций, которые надо совершить над банковским счетом
const amountOfTransactions int = 1000

// Функция выполнения транзакции. Отправляет сумму транзакции в канал bank
// и ждет подтверждения выполнения транзакции через канал transactionDone.
func doTransaction(summ int, bank chan<- int, transactionDone <-chan int, wg *sync.WaitGroup) int {
	defer wg.Done()
	bank <- summ
	return <-transactionDone
}

// Функция создания банка
func createBank(bankAccount *int, wg *sync.WaitGroup) (bank chan int, transactionDone chan int) {
	bank = make(chan int) // для отправки сумм транзакций
	transactionDone = make(chan int)
	go func() {
		// Горутина обрабатывает транзакции
		defer wg.Done()
		for summ := range bank { // ожидает поступления сумм транзакций через канал bank
			// Логика, отвечающая за наращивание счета обязательно должна быть
			// в функции, которая работает только в одной горутине
			*bankAccount = *bankAccount + summ
			transactionDone <- *bankAccount
		}
	}()
	return
}

func main() {
	// Текущий баланс банковского счета.
	var bankAccount int = 0
	// Количество горутин, необходимых для выполнения всех транзакций
	var amountOfGoRoutines int = amountOfTransactions * 2
	var wg sync.WaitGroup
	wg.Add(amountOfGoRoutines)
	bank, transactionDone := createBank(&bankAccount, &wg)
	for i := 1; i <= amountOfTransactions; i++ {
		go doTransaction(1, bank, transactionDone, &wg)
		go doTransaction(-1, bank, transactionDone, &wg)
	}
	// Дожидаемся завершения всех транзакций
	wg.Wait()
	wg.Add(1)
	close(bank)
	// Дожидаемся закрытия банка
	wg.Wait()
	fmt.Println(bankAccount)
}
