package fakes

import "github.com/cloudfoundry/bosh-bootloader/storage"

type GCPUp struct {
	Name        string
	ExecuteCall struct {
		CallCount int
		Receives  struct {
			State storage.State
		}
		Returns struct {
			Error error
		}
	}
}

func (u *GCPUp) Execute(state storage.State) error {
	u.ExecuteCall.CallCount++
	u.ExecuteCall.Receives.State = state
	return u.ExecuteCall.Returns.Error
}
