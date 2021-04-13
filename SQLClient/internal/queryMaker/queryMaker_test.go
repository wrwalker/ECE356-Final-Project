package queryMaker

import (
	"errors"
	"github.com/ECE356-Final-Project/SQLClient/internal/dbConnector/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryMaker(t *testing.T) {
	mockDB := &mocks.DBConnector{}
	qm := NewQueryMaker(mockDB)

	queryString := "testIn"
	errMsg := "anErr"

	mockDB.On("Queryx", queryString).Return(nil, errors.New(errMsg))

	_, err := qm.DoQuery(queryString)
	assert.EqualError(t, err, errMsg)

}
