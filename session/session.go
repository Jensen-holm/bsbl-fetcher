package session

import (
	"fmt"
	"github.com/Jensen-holm/bsbl-api/rf"
	"github.com/Jensen-holm/bsbl-api/user"

	"github.com/gofiber/fiber/v2"
)

type Session interface {
	Main() error
	Results() string
	ValidRequest(r *map[string]string)
}

func WebPage(wp string, usr *user.User) (Session, error) {
	var PageMap = map[string]Session{
		"Baseball-Reference": rf.NewRef(usr),
	}

	if s, isIn := PageMap[wp]; isIn {
		return s, nil
	}

	return nil, fiber.NewError(
		fiber.StatusBadRequest,
		fmt.Sprintf("Invalid Source \"%s\"", wp),
	)
}
