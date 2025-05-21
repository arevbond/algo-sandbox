type ATM struct {
	banknotes [5]int
	cash [5]int
}

func Constructor() ATM {
	return ATM{banknotes: [5]int{20, 50, 100, 200, 500},
			cash: [5]int{}}
}

func (m *ATM) Deposit(banknotesCount []int) {
	for i := 0; i < len(m.cash); i++ {
		m.cash[i] += banknotesCount[i]
	}
}

func (m *ATM) Withdraw(amount int) []int {
	tmpCash := [5]int{}
	for i := 0; i < len(m.cash); i++ {
		tmpCash[i] = m.cash[i]
	}

	result := make([]int, 5)
	for i := len(m.banknotes) - 1; i >= 0; i-- {
		denomination := m.banknotes[i]
		count := tmpCash[i]

		need := amount / denomination

		exists := min(need, count)


		result[i] = exists
		tmpCash[i] -= exists

		amount -= exists * denomination
	}

	if amount > 0 {
		return []int{-1}
	}
	m.cash = tmpCash
	return result
}
