package domain

import (
	"dialogService/domain/helper"
	"errors"
	"github.com/google/uuid"
)

type CreateMessageDTO struct {
	From    string
	To      string
	Message string
}

type RegisterUserService func(dto *CreateMessageDTO) (*helper.ValidationResult, string, error)

func registerUserValidateDto(dto *CreateMessageDTO) *helper.ValidationResult {
	result := helper.NewValidationResult()
	if dto.From == "" {
		result.AddError("From", "From is empty")
	}
	if dto.To == "" {
		result.AddError("To", "To is empty")
	}
	if dto.From != "" && dto.To != "" && dto.From == dto.To {
		// TODO: Tricks with uuid?
		result.AddError("To", "To and from are same")
	}

	if dto.From != "" {
		_, err := uuid.Parse(dto.From)
		if err != nil {
			result.AddError("From", "From must be correct uuid")
		}
	}
	if dto.To != "" {
		_, err := uuid.Parse(dto.From)
		if err != nil {
			result.AddError("From", "From must be correct uuid")
		}
	}
	if dto.Message == "" {
		result.AddError("Message", "Message is empty")
	}

	return result
}

func CreateMessageService(messageModel MessageRepository) RegisterUserService {
	return func(dto *CreateMessageDTO) (*helper.ValidationResult, string, error) {
		validationResult := registerUserValidateDto(dto)
		if !validationResult.IsValid() {
			return validationResult, "", nil
		}
		id, err := uuid.NewUUID()
		if err != nil {
			return nil, "", errors.New("failed to create user, error:" + err.Error())
		}
		from, _ := uuid.Parse(dto.From)
		to, _ := uuid.Parse(dto.To)

		err = messageModel.Create(id, from, to, dto.Message)
		if err != nil {
			return nil, "", errors.New("failed to create user, error:" + err.Error())
		}

		return nil, id.String(), nil
	}
}
