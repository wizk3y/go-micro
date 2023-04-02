package postgresql

import "github.com/lib/pq"

// IsDuplicateError --
func IsDuplicateError(err error) bool {
	me, ok := err.(*pq.Error)

	if ok && me.Code.Name() == "unique_violation" {
		return true
	}

	return false
}
