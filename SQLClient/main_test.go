package main

import (
	"errors"
	"github.com/ECE356-Final-Project/SQLClient/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCLI(t *testing.T) {
	mockDB := &mocks.DBConnector{}
	cli := NewCLI(mockDB)

	mockDB.On("Query", "SELECT * FROM VotesByState").Return(nil, errors.New("anError"))

	_, err := cli.DoQuery("SELECT * FROM VotesByState")
	assert.EqualError(t, err, "anError")


}
