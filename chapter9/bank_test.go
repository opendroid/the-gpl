package chapter9

import (
	"sync"
	"testing"
)

func TestSafeBankDeposit(t *testing.T) {
	b := &SafeBank{}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			b.Deposit(1)
		}()
	}
	wg.Wait()
	if got := b.Balance(); got != 100 {
		t.Errorf("Balance() = %d, want 100", got)
	}
}

func TestSafeBankWithdraw(t *testing.T) {
	b := &SafeBank{}
	b.Deposit(50)
	if !b.Withdraw(30) {
		t.Fatal("Withdraw(30) from 50 should succeed")
	}
	if b.Withdraw(30) {
		t.Fatal("Withdraw(30) from 20 should fail")
	}
	if got := b.Balance(); got != 20 {
		t.Errorf("Balance() = %d, want 20", got)
	}
}

func TestRWBankConcurrentReads(t *testing.T) {
	b := &RWBank{}
	b.Deposit(100)
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if got := b.Balance(); got != 100 {
				t.Errorf("Balance() = %d, want 100", got)
			}
		}()
	}
	wg.Wait()
}
