package generics

type Transaction struct {
	From, To string
	Sum      float64
}

type Account struct {
	Name    string
	Balance float64
}

func NewTransaction(from, to Account, amount float64) Transaction {
	return Transaction{
		From: from.Name,
		To:   to.Name,
		Sum:  amount,
	}
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}

func applyTransaction(account Account, t Transaction) Account {
	if t.To == account.Name {
		account.Balance += t.Sum
	}

	if t.From == account.Name {
		account.Balance -= t.Sum
	}

	return account
}
