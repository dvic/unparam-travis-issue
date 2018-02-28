package unparamtest

//go:generate mockgen -destination mocks_test.go -package main github.com/dvic/unparam-travis-issue Processor

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestSomething(t *testing.T) {
	ctrl := gomock.NewController(t)
	ts := NewMockProcessor(ctrl)

	ts.EXPECT().Process(gomock.Any(), gomock.Any()).DoAndReturn(
		func(res []string) ([]string, error) {
			num, _ := strconv.Atoi(res[0])
			return []string{fmt.Sprintf("something %d", num)}, nil
		},
	)
}
