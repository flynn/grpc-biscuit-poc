package antireplay

import (
	"bytes"
	"errors"
	"fmt"
	"time"
)

var (
	ErrReplay   = errors.New("authorization: replay attempt")
	ErrNonceOOB = errors.New("authorization: nonce out of time window")
)

type Checker interface {
	Check(nonce Nonce) error
}

type checker struct {
	store       Store
	nonceWindow time.Duration
	nonceMaxAge time.Duration
}

func NewChecker(store Store, nonceWindow, nonceMaxAge time.Duration) Checker {
	return &checker{
		store:       store,
		nonceWindow: nonceWindow,
		nonceMaxAge: nonceMaxAge,
	}
}

func (c *checker) Check(nonce Nonce) error {
	now := time.Now()

	// reject if nonce is out of window
	if nonce.CreatedAt.Before(now.Add(-c.nonceWindow)) {
		return ErrNonceOOB
	}
	if nonce.CreatedAt.After(now.Add(c.nonceWindow)) {
		return ErrNonceOOB
	}

	existingNonces, err := c.store.Get(nonce.UserEmail)
	if err != nil {
		return err
	}

	for _, n := range existingNonces {
		// skip too old nonces
		if n.CreatedAt.Before(now.Add(-c.nonceMaxAge)) {
			continue
		}
		fmt.Printf("Check replay for %s: %x | %x: %v\n", nonce.UserEmail, nonce.Value, n.Value, bytes.Equal(n.Value, nonce.Value))
		if bytes.Equal(n.Value, nonce.Value) {
			return ErrReplay
		}
	}

	return c.store.Insert(nonce)
}