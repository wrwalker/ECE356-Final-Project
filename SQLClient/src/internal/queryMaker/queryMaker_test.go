package queryMaker

import (
	"errors"
	"github.com/ECE356-Final-Project/SQLClient/src/internal/dbConnector/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

const RunNonDeterministicTests = true

func TestQueryMaker(t *testing.T) {
	t.Run("doQuery", func(t *testing.T) {
		t.Run("error handling", func(t *testing.T) {
			mockDB := &mocks.DBConnector{}
			qm := NewQueryMaker(mockDB)

			queryString := "testIn"
			errMsg := "anErr"

			mockDB.On("Queryx", queryString).Return(nil, errors.New(errMsg))

			_, err := qm.doQuery(queryString)
			assert.EqualError(t, err, errMsg)
		})
	})

	t.Run("GetVotesForCandidate", func(t *testing.T) {
		t.Run("non deterministic tests", func(t *testing.T) {
			if !RunNonDeterministicTests {
				t.Skip()
			}
			qm := NewQueryMaker()
			defer qm.Db.Close()
			t.Run("no states, no county", func(t *testing.T) {
				numVotes, _, _ := qm.GetVotesForCandidate("Joe Biden", "", []string{})
				assert.Equal(t, 82046434, numVotes)
			})
			t.Run("no states, county", func(t *testing.T) {
				numVotes, _, _ := qm.GetVotesForCandidate("Joe Biden", "Autauga County", []string{})
				assert.Equal(t, 7503, numVotes)
			})
			t.Run("1 state, county", func(t *testing.T) {
				numVotes, _, _ := qm.GetVotesForCandidate("Joe Biden", "Autauga County", []string{"Alabama"})
				assert.Equal(t, 7503, numVotes)
			})
			t.Run("multiple state, county", func(t *testing.T) {
				numVotes, _, _ := qm.GetVotesForCandidate("Joe Biden", "Autauga County", []string{"Alabama", "Wyoming"})
				assert.Equal(t, 7503, numVotes)
			})
			t.Run("multiple state, no county", func(t *testing.T) {
				numVotes, _, _ := qm.GetVotesForCandidate("Joe Biden", "", []string{"Alabama", "Wyoming"})
				assert.Equal(t, 923139, numVotes)
			})
			t.Run("invalid candidates", func(t *testing.T) {
				_, _, err := qm.GetVotesForCandidate("fakeName", "", []string{"Alabama", "Wyoming"})
				assert.EqualError(t, err, "could not find any matches")
			})
			t.Run("empty", func(t *testing.T) {
				_, _, err := qm.GetVotesForCandidate("fakeName", "", []string{""})
				assert.EqualError(t, err, "could not find any matches")
			})
		})

	})

}

func TestGetStringForGetVotesForCandidate(t *testing.T) {
	t.Run("no states, no county", func(t *testing.T) {
		expected := "select sum(votes) from VotesByCountyCandidate where candidate = \"Joe Biden\""
		assert.Equal(t, expected, getStringForGetVotesForCandidate("Joe Biden", "", []string{}))
	})
	t.Run("no states, county", func(t *testing.T) {
		expected := "select sum(votes) from VotesByCountyCandidate where candidate = \"Joe Biden\" and county=\"Autauga County\""
		assert.Equal(t, expected, getStringForGetVotesForCandidate("Joe Biden", "Autauga County", []string{}))
	})
	t.Run("1 state, county", func(t *testing.T) {
		expected := "select sum(votes) from VotesByCountyCandidate where candidate = \"Joe Biden\" and county=\"Autauga County\" and (state=\"Alabama\")"
		assert.Equal(t, expected, getStringForGetVotesForCandidate("Joe Biden", "Autauga County", []string{"Alabama"}))
	})
	t.Run("multiple state, county", func(t *testing.T) {
		expected := "select sum(votes) from VotesByCountyCandidate where candidate = \"Joe Biden\" and county=\"Autauga County\" and (state=\"Alabama\" or state=\"Wyoming\")"
		assert.Equal(t, expected, getStringForGetVotesForCandidate("Joe Biden", "Autauga County", []string{"Alabama", "Wyoming"}))
	})
	t.Run("multiple state, no county", func(t *testing.T) {
		expected := "select sum(votes) from VotesByCountyCandidate where candidate = \"Joe Biden\" and (state=\"Alabama\" or state=\"Wyoming\")"
		assert.Equal(t, expected, getStringForGetVotesForCandidate("Joe Biden", "", []string{"Alabama", "Wyoming"}))
	})
}
