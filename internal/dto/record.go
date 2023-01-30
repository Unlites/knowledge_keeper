package dto

import validation "github.com/go-ozzo/ozzo-validation"

type RecordDTORequest struct {
	Topic   string `json:"topic"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (r RecordDTORequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Topic, validation.Required, validation.Length(1, 100)),
		validation.Field(&r.Title, validation.Required, validation.Length(1, 255)),
		validation.Field(&r.Content, validation.Required, validation.Length(1, 3000)),
	)
}

type RecordDTOResponse struct {
	Id      uint   `json:"id"`
	Topic   string `json:"topic"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
