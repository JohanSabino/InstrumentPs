package ctx

import (
	"github.com/sirupsen/logrus"
)

type questionRepositoryInterface interface {
	questionsRepository()
}

// Handler main lambda's structure
type Handler struct {
	questionsRepository questionRepositoryInterface
	logger              *logrus.Logger
}
