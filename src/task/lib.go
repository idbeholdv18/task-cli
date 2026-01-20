package task

import (
	"task-cli/src/ids"
)

func (task Task) ToDto() TaskDto {
	return TaskDto{
		Id:          task.id,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.createdAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

func TaskFromDto(dto TaskDto) Task {
	return Task{
		id:          dto.Id,
		Description: dto.Description,
		Status:      dto.Status,
		createdAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
	}
}

// FindMaxId the maximum id among all the Task elements.
//
// If tasks doesn't contain any elements returns false and zero value Id.
//
// You should check if FindMaxId found any Id or not via first return value!
//
// Although in this instance, it is possible to retrieve the ID of the last element,
// the function still iterates over the whole slice, as a change in behaviour or data
// structure could result in the maximum ID not being the last element.
func (tasks Tasks) FindMaxId() (ids.Id, bool) {
	if len(tasks) == 0 {
		return 0, false
	}
	max := tasks[0].id
	for _, task := range tasks[1:] {
		if task.id > max {
			max = task.id
		}
	}
	return max, true
}
