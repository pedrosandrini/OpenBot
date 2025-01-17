package entity

import (
	"errors"
	"github.com/google/uuid"
	"github.com/pkoukk/tiktoken-go"
	"time"
)

type Message struct {
	ID        string
	Role      string
	Content   string
	Tokens    int
	Model     *Model
	CreatedAt time.Time
}

func NewMessage(role, content string, model *Model) (*Message, error) {
	totalTokens, err := CountTokens(content, model.GetModelName())
	if err != nil {
		return nil, errors.New("error on counting tokens")
	}
	msg := &Message{
		ID:        uuid.New().String(),
		Role:      role,
		Content:   content,
		Tokens:    totalTokens,
		Model:     model,
		CreatedAt: time.Now(),
	}
	if err := msg.Validate(); err != nil {
		return nil, err
	}
	return msg, nil
}

func (m *Message) Validate() error {
	if m.Role != "user" && m.Role != "system" && m.Role != "assistant" {
		return errors.New("invalid role")
	}
	if m.Content == "" {
		return errors.New("content is empty")
	}
	if m.CreatedAt.IsZero() {
		return errors.New("invalid created at")
	}
	return nil
}

func (m *Message) GetQtdTokens() int {
	return m.Tokens
}
func CountTokens(content string, encoding string) (int, error) {
	tkm, err := tiktoken.EncodingForModel(encoding)
	if err != nil {
		return 0, err
	}
	token := tkm.Encode(content, nil, nil)

	return len(token), nil
}
