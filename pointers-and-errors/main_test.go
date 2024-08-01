package pointers_and_errors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(30)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		want := Bitcoin(20)
		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, "oh no, insufficient funds")
		assertBalance(t, wallet, startingBalance)
	})
	// make sure the balance hasn't changed
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, err error, want string) {
	t.Helper()
	if err == nil {
		t.Error("wanted an error but didn't get one")
	} else if err.Error() != want {
		t.Errorf("got %q, want %q", err, want)
	}
}
