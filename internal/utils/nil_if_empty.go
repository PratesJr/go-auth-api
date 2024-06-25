package utils

func NilIfIsEmpty(value string) *string {

	if len(value) == 0 {
		return nil
	}

	return &value

}
