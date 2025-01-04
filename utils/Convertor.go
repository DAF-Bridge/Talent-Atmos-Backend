package utils

import "github.com/google/uuid"

func ListStringToListUuid(ids []string) []uuid.UUID {
	// convert string to uuid
	var uuids []uuid.UUID
	for _, id := range ids {
		uid, err := uuid.Parse(id)
		if err != nil {
			continue
		}
		uuids = append(uuids, uid)
	}
	return uuids
}
