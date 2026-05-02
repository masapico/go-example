package types

import "fmt"

// Account は口座の所有者と残高を表す構造体です。
type Account struct {
	Owner   string
	Balance int
}

// Deposit は口座に正の金額を入金します。
func (a *Account) Deposit(amount int) error {
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be positive: %d", amount)
	}
	// ポインタレシーバなので、呼び出し元の Account の Balance が更新される。
	a.Balance += amount
	return nil
}

// Withdraw は口座から正の金額を引き出します。
func (a *Account) Withdraw(amount int) error {
	if amount <= 0 {
		return fmt.Errorf("withdraw amount must be positive: %d", amount)
	}
	if amount > a.Balance {
		return fmt.Errorf("insufficient balance: balance=%d amount=%d", a.Balance, amount)
	}
	// 検証を通過してから残高を減らす。
	a.Balance -= amount
	return nil
}

// Describer は説明文を返せる型が満たすインターフェースです。
type Describer interface {
	Describe() string
}

// Describe は Account を人が読める文字列に整形します。
func (a Account) Describe() string {
	return fmt.Sprintf("%s has %d yen", a.Owner, a.Balance)
}

// Label は Describer インターフェースを受け取り、共通の接頭辞を付けます。
func Label(d Describer) string {
	return "account: " + d.Describe()
}
