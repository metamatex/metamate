package persistence

import (
	sdk "github.com/metamatex/metamatemono/gen/v0/sdk"
	"github.com/metamatex/metamatemono/metamate/pkg/v0/types"
)

type MemoryLinkStore struct {
	relations map[string][]types.Link
}

func NewMemoryLinkStore() *MemoryLinkStore {
	return &MemoryLinkStore{
		relations: map[string][]types.Link{},
	}
}

func (s *MemoryLinkStore) GetLinks(relationName string, active bool, id sdk.ServiceId) (ids []sdk.ServiceId, err error) {
	if active {
		for _, l := range s.relations[relationName] {
			if *l.Active.ServiceName == *id.ServiceName &&
				*l.Active.Value == *id.Value {
				ids = append(ids, l.Passive)
			}
		}
	}

	if !active {
		for _, l := range s.relations[relationName] {
			if *l.Passive.ServiceName == *id.ServiceName &&
				*l.Passive.Value == *id.Value {
				ids = append(ids, l.Active)
			}
		}
	}

	return
}

func (s *MemoryLinkStore) PostLinks(relationName string, active bool, id sdk.ServiceId, ids []sdk.ServiceId) (err error) {
	if active {
		for i, _ := range ids {
			s.relations[relationName] = append(s.relations[relationName], types.Link{
				Active:  id,
				Passive: ids[i],
			})
		}
	}

	if !active {
		for i, _ := range ids {
			s.relations[relationName] = append(s.relations[relationName], types.Link{
				Active:  ids[i],
				Passive: id,
			})
		}
	}

	return
}

func (s *MemoryLinkStore) DeleteLinks(relationName string, active bool, id sdk.ServiceId, ids []sdk.ServiceId) (err error) {
	return
}
