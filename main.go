//// Тестовое задание для претендентов ПС Sendy
//
//// Описание:
//// Есть функция, которая принимает от всех модулей дробное число
//// Функция должна умножить входящее значение на 100 и вернуть целое число
//// Ожидаемое максимальное количество знаков после запятой: 2
//// Входящее значение может быть любым, но ожидаемый диапазон от 0 до 99.99
//// Возвращаемое корректное значение должно быть в диапазоне от 0 до 9999
//
//// Входящий тип данных string, исходящий тип данных uint64
//// Например для значения 12.34 функция должна вернуть значение 1234
//// Работа с функцией не должна требовать каких-либо проверок перед отправкой в неё значения, т.е. все проверки входящего значения должны быть на стороне функции
//// Результат работы функции не должен требовать перепроверки со стороны внешних модулей, т.е. возвращаемое значение функции всегда считается корректным
//
//// Разработчик реализовал функцию проверки корректности ввода и преобразования вводимого значения
//// От тех, кто использует функцию поступают жалобы на то, что вводимое ими число преобразовывается неверно
//
//// Задание для разработчика:
//// Исправьте функцию проверки вводимого значения и дальнейшего преобразования итогового значения.
//
//// Задание для тестировщика:
//// Проверьте корректность работы фукнции изменяя входящее значение по вашему усмотрению.
//// В случае выявления ошибок опишите порядок действий, которые привели к ошибке.
//
//// Задание для системного аналитика:
//// Напишите постановку данной задачи для разработчика с максимальной детализацией.
//
//// Готовое решение нужно предоставить в виде ссылки по кнопке "Share"
//
//package main
//
//// Импортируем библиотеки вывода и конвертации
//import (
//	"fmt"
//	"strconv"
//	"strings"
//)
//
//// Функция умножения дробного числа
//func num_x_100(number string) uint64 {
//	// Конвертируем присланное значение из строки в число с запятой во временную переменную
//	tempValue, err := strconv.ParseFloat(number, 64)
//	// Если конвертация не получилась
//	if err != nil {
//		// Скажем об этом
//		fmt.Println("Введено не число:", number)
//		// И выйдем из функции
//		return 0
//	}
//	// Если конвертация удалась
//	// Умножим значение на 100, тем самым сдвинем запятую
//	tempValue = tempValue * 100
//	// Преобразуем значение в нужный тип данных
//	result := uint64(tempValue)
//	// Возвращаем корректное значение
//	return result
//}
//
//// Ввиду ограничений входящее значение всегда в string
//// Пеменная input будет эмулировать входящее значение
//func main() {
//	// Проверяем маленькое значение
//	input := "0.01"
//	fmt.Println("Первое значение:", input)
//	fmt.Println("Преобразованное значение:", num_x_100(input))
//
//	// Проверяем большое значение
//	input = "99.99"
//	fmt.Println("\nВторое значение:", input)
//	fmt.Println("Преобразованное значение:", num_x_100(input))
//
//	// Проверяем если ввели не число
//	input = "неЧисло"
//	fmt.Println("\nСтроковое значение:", input)
//	fmt.Println("Преобразованное значение:", num_x_100(input))
//}

package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

const (
	msgRangeErr = "не верный диапозон мин 0 макс 99.99"
	msgErr      = "не верное число"
)

var (
	inputs = []string{
		"1.-1",
		"-1.2",
		"1",
		"15",
		"100",
		"1,5",
		"Text",
		" 1,8 8",
		"9999",
		"0",
		"1%.",
		"1,,5",
		"12..3",
		"-0.1",
		"-0",
		"99.999",
	}
	wg sync.WaitGroup
)

func num_x_100(number string) (uint64, error) {
	if len(number) > 5 {
		return 0, fmt.Errorf(msgRangeErr)
	}
	if strings.HasPrefix(number, "-") {
		return 0, fmt.Errorf(msgRangeErr)
	}

	parts := strings.FieldsFunc(number, func(r rune) bool {
		return r == '.' || r == ','
	})

	var integerPart, decimalPart string

	switch len(parts) {
	case 1:
		integerPart = parts[0]
		decimalPart = "0"
	case 2:
		integerPart = parts[0]
		decimalPart = parts[1]
	default:
		return 0, fmt.Errorf(msgErr)
	}

	if !isValidInteger(integerPart) || !isValidDecimal(decimalPart) || isNegativeDecimal(decimalPart) || hasMultipleDecimalSeparators(number) {
		return 0, fmt.Errorf(msgErr)
	}

	decimalValue, err := strconv.Atoi(decimalPart)
	if err != nil {
		return 0, err
	}

	if len(decimalPart) < 2 {
		decimalValue *= 10
	}

	integer, err := strconv.Atoi(integerPart)
	if err != nil {
		return 0, err
	}

	value := uint64(integer*100) + uint64(decimalValue)

	if value > 9999 {
		return 0, fmt.Errorf(msgRangeErr)
	}

	return value, nil
}

func isValidInteger(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isValidDecimal(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil && len(s) <= 2
}

func isNegativeDecimal(s string) bool {
	return strings.HasPrefix(s, "-")
}

func hasMultipleDecimalSeparators(s string) bool {
	dotCount := strings.Count(s, ".")
	commaCount := strings.Count(s, ",")
	return dotCount+commaCount > 1
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	for _, input := range inputs {
		wg.Add(1)
		go func(input string) {
			defer wg.Done()
			fmt.Println("\nЗначение:", input)
			value, err := num_x_100(input)
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Преобразованное значение:", value)
			}
		}(input)
	}

	wg.Wait()
	return nil
}
