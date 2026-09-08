package sqlitelogger

import (
	"encoding/json"
	"time"

	"coding-guidelines/common/pkg/appfault"
	"coding-guidelines/common/pkg/errtype"
)

// TaskLogger provides task-scoped logging bound to an isolated task database.
type TaskLogger struct {
	taskId string
	mgr    *SplitDBManager
}

// NewTaskLogger constructs a TaskLogger for the specified task ID.
func NewTaskLogger(taskId string, mgr *SplitDBManager) (*TaskLogger, *appfault.AppError) {
	if taskId == "" {
		return nil, appfault.New(errtype.Validation, "taskId cannot be empty")
	}

	if mgr == nil {
		return nil, appfault.New(errtype.Validation, "manager cannot be nil")
	}

	return &TaskLogger{taskId: taskId, mgr: mgr}, nil
}

// TaskId returns the bound task identifier.
func (tl *TaskLogger) TaskId() string {
	return tl.taskId
}

// Log records a basic entry into the task's database.
func (tl *TaskLogger) Log(level string, message string) *appfault.AppError {
	entry := TaskLogEntry{
		TaskId:    tl.taskId,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Level:     level,
		Message:   message,
		Status:    "OK",
	}

	return tl.mgr.WriteTask(tl.taskId, entry)
}

// LogWithFields records an entry with extra contextual fields into the task's database.
func (tl *TaskLogger) LogWithFields(
	level string,
	message string,
	fields map[string]any,
) *appfault.AppError {
	fieldsBytes, err := json.Marshal(fields)
	if err != nil {
		return appfault.Wrap(errtype.Serialization, err, "failed to marshal fields")
	}

	entry := TaskLogEntry{
		TaskId:     tl.taskId,
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
		Level:      level,
		Message:    message,
		FieldsJson: string(fieldsBytes),
		Status:     "OK",
	}

	return tl.mgr.WriteTask(tl.taskId, entry)
}

// LogExecution records command/task completion status and elapsed duration.
func (tl *TaskLogger) LogExecution(
	message string,
	durationMs int64,
	status string,
) *appfault.AppError {
	entry := TaskLogEntry{
		TaskId:     tl.taskId,
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
		Level:      "INFO",
		Message:    message,
		DurationMs: durationMs,
		Status:     status,
	}

	return tl.mgr.WriteTask(tl.taskId, entry)
}
