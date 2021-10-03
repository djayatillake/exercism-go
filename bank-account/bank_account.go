package account

import "sync"

// declaring type account
type Account struct {
	balance int
	open    bool
	mu      sync.RWMutex
}

// Close method sets account to initial state
func (a *Account) Close() (bal int, ok bool) {
	if a == nil {
		return 0, false
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.open {
		return 0, false
	}
	a.open = false
	return a.balance, true
}

// Balance method returns the balance of the account
func (a *Account) Balance() (int, bool) {
	if a == nil || !a.open {
		return 0, false
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.balance, true
}

// Deposit method returns the balance of the account
func (a *Account) Deposit(amount int) (int, bool) {
	if a == nil || !a.open {
		return 0, false
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}

// Open opens a new bank account
func Open(amount int) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount, open: true, mu: sync.RWMutex{}}
}
