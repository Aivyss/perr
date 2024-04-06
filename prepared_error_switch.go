package perr

type preparedErrorSwitch struct {
	target   error
	finished bool
}

func (s *preparedErrorSwitch) Case(err PreparedError, runnable func()) *preparedErrorSwitch {
	if !s.finished && err.Same(s.target) {
		runnable()
		s.finished = true
	}

	return s
}

func (s *preparedErrorSwitch) CaseGroup(group preparedErrorGroup, runnable func()) *preparedErrorSwitch {
	if !s.finished && group.Contains(s.target) {
		runnable()
		s.finished = true
	}

	return s
}

func Switch(err error) *preparedErrorSwitch {
	return &preparedErrorSwitch{
		target:   err,
		finished: false,
	}
}
