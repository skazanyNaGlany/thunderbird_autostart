package main

import (
	"os"
	"strings"
)

type State struct {
	thunderbirdExePathname string
}

func (s *State) SaveStateToFile(pathname string) error {
	return os.WriteFile(pathname, []byte(s.thunderbirdExePathname), 0777)
}

func (s *State) LoadStateFromFile(pathname string) error {
	bytes, err := os.ReadFile(pathname)

	if err != nil {
		return err
	}

	s.thunderbirdExePathname = strings.TrimSpace(string(bytes))

	return err
}

func (s *State) GetThunderbirdExePathname() string {
	return s.thunderbirdExePathname
}

func (s *State) SetThunderbirdExePathname(thunderbirdExePathname string) *State {
	s.thunderbirdExePathname = thunderbirdExePathname
	return s
}
