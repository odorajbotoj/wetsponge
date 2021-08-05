package WetSpongeCore

import "github.com/satori/go.uuid"

type UUID struct {
	uuid string
}

func (u UUID) String() string {
	if u.uuid == "" {
		uid4, err := uuid.NewV4()
		if err != nil {
			return ""
		}
		ret := uid4.String()
		return ret
	} else {
		return u.uuid
	}
}

func (u *UUID) Set(setUuid uuid.UUID, err error) error {
	if err != nil {
		u.uuid = setUuid.String()
		return nil
	} else {
		return err
	}
}

type RequestId = UUID
