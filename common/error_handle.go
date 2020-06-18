package common

import "github.com/google/logger"

func Handle(err error) {
	if err != nil {
		logger.Error("An error occurred ", err)
	}
}
