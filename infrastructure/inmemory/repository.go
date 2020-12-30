package inmemory

import (
	"fmt"
	"github.com/reuben-baek/clean-go-application/domain"
	"sync"
)

type accountDto struct {
	id string
}

func fromAccount(a domain.Account) accountDto {
	return accountDto{id: a.Id()}
}

func (a *accountDto) to() domain.Account {
	return domain.NewAccount(a.id)
}

func NewAccountRepository() *accountRepository {
	storage := make(map[string]accountDto)
	return &accountRepository{storage: storage}
}

type accountRepository struct {
	rwMutex sync.RWMutex
	storage map[string]accountDto
}

func (r *accountRepository) FindOne(id string) (domain.Account, error) {
	r.rwMutex.RLock()
	defer r.rwMutex.RUnlock()
	if accountDto, ok := r.storage[id]; ok {
		return accountDto.to(), nil
	} else {
		return nil, domain.NewNotFoundError(fmt.Sprintf("cannot find %s in inmemory.accountRepository", id), nil)
	}
}

func (r *accountRepository) Save(account domain.Account) error {
	r.rwMutex.Lock()
	defer r.rwMutex.Unlock()

	accountDto := fromAccount(account)
	r.storage[accountDto.id] = accountDto
	return nil
}

func (r *accountRepository) Delete(account domain.Account) error {
	panic("implement me")
}
