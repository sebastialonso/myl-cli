package deck

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DeckTestSuite struct {
	suite.Suite
	deck Deck
}

func (s *DeckTestSuite) SetupTest() {
	deck := newBaseDeck(DefaultDeckCfg)
	s.deck = deck
}

func TestDeckTestSuite(t *testing.T) {
	suite.Run(t, new(DeckTestSuite))
}

func (s *DeckTestSuite) TestnnewBaseDeck(t *testing.T) {
	assert.NotNil(t, s.deck)
}

func (s *DeckTestSuite) TestString(t *testing.T) {
	deckStruct, ok := s.deck.(deck)
	assert.True(t, ok)
	assert.Equal(t, deckStruct.String(), fmt.Sprintf("<Deck ID:%s>", &deckStruct.UUID))
}