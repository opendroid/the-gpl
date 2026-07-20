// Package chapter9 covers Chapter 9 of The Go Programming Language:
// Concurrency with Shared Variables. It demonstrates sync.Mutex,
// sync.RWMutex, sync.Once, and race-safe patterns.
package chapter9

import "sync"

// SafeBank is a bank account protected by a mutex (book §9.2).
type SafeBank struct {
	mu      sync.Mutex
	balance int
}

// Deposit adds amount to the account balance.
func (b *SafeBank) Deposit(amount int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance += amount
}

// Withdraw removes amount from the balance and reports whether it succeeded.
// Returns false when funds are insufficient.
func (b *SafeBank) Withdraw(amount int) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.balance < amount {
		return false
	}
	b.balance -= amount
	return true
}

// Balance returns the current balance.
func (b *SafeBank) Balance() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.balance
}

// RWBank uses sync.RWMutex to allow concurrent reads (book §9.3).
type RWBank struct {
	mu      sync.RWMutex
	balance int
}

// Deposit adds amount to the account balance.
func (b *RWBank) Deposit(amount int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance += amount
}

// Balance returns the current balance; multiple goroutines may read concurrently.
func (b *RWBank) Balance() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.balance
}
