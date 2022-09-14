package test

import (
	"daily_test_go/service"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := service.Wallet{}
	wallet.Deposit(10)

	want := "10 BTC"
	got := wallet.Balance().String()

	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
