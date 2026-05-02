package types

import (
	"strings"
	"testing"
)

// TestAccountDepositAndWithdraw は入金と出金で残高が更新されることを確認する。
func TestAccountDepositAndWithdraw(t *testing.T) {
	account := Account{Owner: "masa", Balance: 1000}

	if err := account.Deposit(500); err != nil {
		t.Fatalf("Deposit() error = %v", err)
	}
	if err := account.Withdraw(300); err != nil {
		t.Fatalf("Withdraw() error = %v", err)
	}
	if account.Balance != 1200 {
		t.Fatalf("Balance = %d, want 1200", account.Balance)
	}
}

// TestAccountRejectsInvalidWithdraw は残高不足の出金がエラーになることを確認する。
func TestAccountRejectsInvalidWithdraw(t *testing.T) {
	account := Account{Owner: "masa", Balance: 100}
	err := account.Withdraw(200)
	if err == nil || !strings.Contains(err.Error(), "insufficient balance") {
		t.Fatalf("Withdraw() error = %v, want insufficient balance", err)
	}
}

// TestLabelUsesInterface は Account が Describer として扱えることを確認する。
func TestLabelUsesInterface(t *testing.T) {
	got := Label(Account{Owner: "masa", Balance: 1000})
	want := "account: masa has 1000 yen"
	if got != want {
		t.Fatalf("Label() = %q, want %q", got, want)
	}
}
