package main

import (
	"errors"
	"github.com/ECE356-Final-Project/SQLClient/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryMaker(t *testing.T) {
	mockDB := &mocks.DBConnector{}
	qm := NewQueryMaker(mockDB)

	mockDB.On("Query", "SELECT * FROM VotesByState").Return(nil, errors.New("anError"))

	_, err := qm.DoQuery("SELECT * FROM VotesByState")
	assert.EqualError(t, err, "anError")

}
