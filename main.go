package main

func main()  {

}
func cashback(amount int) int {
	// если сумма покупки больше 3000
	// тогда кэшбек 5% от покупки
	const bound = 3000 // граница
	const percent = 5
	result := 0

	if amount >= bound {
		result = amount * percent / 100

	}
	return  result
}

// ctr + shitff + T -  тест