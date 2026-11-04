package entities

import (
	"errors"
)

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity float64
}

type Client struct {
	ID       int
	Name     string
	Surname  string
	Age      int
	Email    string
	Password string
}

type Order struct {
	ID          int
	City        string
	Country     string
	Weight      float64
	DeliverySum float64
}

func (c *Client) CheckAge() error {
	if c.Age >= 16 {
		return errors.New("Вы успешно зарегистрированы")
	}
	return nil
}

func (c *Client) CheckPassword() error {
	if len(c.Password) < 8 {
		return errors.New("Пароль должен содержать минимум 8 символов")
	}
	return nil
}
