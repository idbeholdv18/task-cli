package task

func (tasks Tasks) Filter(status StatusState) Tasks {
	filtered := make(Tasks, 0, len(tasks))
	for _, t := range tasks {
		if t.Status == status {
			filtered = append(filtered, t)
		}
	}

	return filtered
}
