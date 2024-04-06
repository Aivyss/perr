package perr

type preparedErrorGroup []PreparedError

func (g preparedErrorGroup) Contains(err error) bool {
	for _, pErr := range g {
		if pErr.Same(err) {
			return true
		}
	}

	return false
}

func ErrorGroup(err ...PreparedError) preparedErrorGroup {
	return err
}
