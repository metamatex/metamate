package types

import (
	"errors"
	"github.com/metamatex/metamate/gen/v0/mql"
	"strings"
)

type GetArgs struct {
	Instance   string
	TypePlural string
	Services   string
	Path       string
	Search     string
	Id         string
	Username   string
	Name       string
}

func (a GetArgs) Validate() (err error) {
	err = a.validateMode()
	if err != nil {
		return
	}

	err = a.validateId()
	if err != nil {
		return
	}

	if a.TypePlural == "" {
		return errors.New("no type set")
	}

	return
}

func (a GetArgs) validateMode() (err error) {
	hasId := a.HasId()

	if hasId && a.Search != "" {
		err = errors.New("--search doesn’t work with --id, --username, or --name")

		return
	}

	return
}

func (a GetArgs) GetSvcFilter() mql.ServiceFilter {
	var f *mql.ServiceIdFilter
	if a.Services != "" {
		f = &mql.ServiceIdFilter{
			Value: &mql.StringFilter{
				In: strings.Split(a.Services, ","),
			},
		}
	}

	return mql.ServiceFilter{
		Id: f,
	}
}

func (a GetArgs) validateId() (err error) {
	n := 0

	if a.Id != "" {
		n++

		s := strings.Split(a.Id, "/")
		if len(s) < 2 {
			err = errors.New("--id needs to be serviceName/value")

			return
		}
	}

	if a.Name != "" {
		n++
	}

	if a.Username != "" {
		n++
	}

	if n > 1 {
		err = errors.New("can only set one of --id, --username, or --name")
	}

	return
}

func (a GetArgs) GetId() (id mql.Id) {
	if a.Id != "" {
		s := strings.Split(a.Id, "/")
		serviceName := s[0]
		value := strings.Join(s[1:], "/")

		return mql.Id{
			Kind: &mql.IdKind.ServiceId,
			ServiceId: &mql.ServiceId{
				ServiceName: &serviceName,
				Value:       &value,
			},
		}
	}

	if a.Name != "" {
		return mql.Id{
			Kind: &mql.IdKind.Name,
			Name: &a.Name,
		}
	}

	if a.Username != "" {
		return mql.Id{
			Kind:     &mql.IdKind.Username,
			Username: &a.Username,
		}
	}

	panic("")
}

func (a GetArgs) HasId() bool {
	if a.Id != "" {
		return true
	}

	if a.Name != "" {
		return true
	}

	if a.Username != "" {
		return true
	}

	return false
}

func (a GetArgs) GetMode() (m mql.GetMode) {
	if a.Search != "" {
		return mql.GetMode{
			Kind: &mql.GetModeKind.Search,
			Search: &mql.SearchGetMode{
				Term: &a.Search,
			},
		}
	}

	if a.HasId() {
		id := a.GetId()
		return mql.GetMode{
			Kind: &mql.GetModeKind.Id,
			Id:   &id,
		}
	}

	return mql.GetMode{
		Kind:       &mql.GetModeKind.Collection,
		Collection: &mql.CollectionGetMode{},
	}
}
