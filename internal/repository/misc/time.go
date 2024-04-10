package misc

import "time"

func (mr *MiscRepositoryImpl) GetCurrentTimestamp() time.Time {
	return time.Now()
}
