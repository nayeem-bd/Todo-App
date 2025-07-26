package dto

type Event struct {
	Event  string `json:"event" validate:"required"`
	TodoID *int   `json:"todo_id,omitempty"`
}
