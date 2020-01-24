package gen_data

import (
	"github.com/google/uuid"
	"strings"
)

func GenUUID() (id string, err error) {
	var uuidObj uuid.UUID
	if uuidObj, err = uuid.NewUUID(); err != nil {
		return
	}
	id = strings.ReplaceAll(uuidObj.String(), "-", "")
	return
}
