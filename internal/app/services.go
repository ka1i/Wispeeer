package app

import (
	"github.com/wispeeer/wispeeer/internal/app/wispeeer"
	"github.com/wispeeer/wispeeer/pkg/info"
)

func init() {
	info.ShowBanner()
}

type service string

const (
	Wispeeer service = "wispeeer"
)

func (service service) Usage() string {
	switch service {
	case Wispeeer:
		return wispeeer.Usage()
	}
	return ""
}

func (service service) Flags() func() (bool, error) {
	switch service {
	case Wispeeer:
		return wispeeer.Flags
	}
	return func() (bool, error) {
		return true, nil
	}
}

func (service service) Service() func() error {
	switch service {
	case Wispeeer:
		return wispeeer.Service
	}
	return nil
}
