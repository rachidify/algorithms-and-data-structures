package linkedlist_test

import (
	"testing"

	"github.com/rasheedhamdawi/algorithms-and-data-structures/data-structures/linkedlist"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SinglyLinkedListTestSuite struct {
	suite.Suite
	list linkedlist.LinkedList
}

func (suite *SinglyLinkedListTestSuite) SetupTest() {
	suite.list = &linkedlist.SinglyLinkedList{}
}

func (suite *SinglyLinkedListTestSuite) TestPushFront() {
	suite.list.PushFront("A")

	assert.Equal(suite.T(), "A", suite.list.TopFront())
}

func (suite *SinglyLinkedListTestSuite) TestPopFront() {
	assert.EqualError(suite.T(), suite.list.PopFront(), "empty list")

	suite.list.PushFront("A")
	suite.list.PushFront("B")
	suite.list.PopFront()

	assert.Equal(suite.T(), "A", suite.list.TopFront())
}

func (suite *SinglyLinkedListTestSuite) TestPushBack() {
	suite.list.PushFront("A")
	suite.list.PushBack("B")
	assert.Equal(suite.T(), "B", suite.list.TopBack())
}

func (suite *SinglyLinkedListTestSuite) TestPopBack() {
	assert.EqualError(suite.T(), suite.list.PopFront(), "empty list")

	suite.list.PushBack("A")
	suite.list.PushBack("B")
	suite.list.PushBack("C")
	suite.list.PushBack("D")
	suite.list.PushBack("E")
	suite.list.PopBack()
	suite.list.PopBack()

	assert.Equal(suite.T(), "C", suite.list.TopBack())
}

func (suite *SinglyLinkedListTestSuite) TestEmpty() {
	assert.True(suite.T(), suite.list.Empty())

	suite.list.PushBack("A")
	assert.False(suite.T(), suite.list.Empty())
}

func (suite *SinglyLinkedListTestSuite) TestFind() {
	assert.False(suite.T(), suite.list.Find("A"))

	items := []string{"A", "B", "C", "D", "E", "F"}

	for _, i := range items {
		suite.list.PushBack(i)
	}

	for _, i := range items {
		assert.True(suite.T(), suite.list.Find(i))
	}
}

func (suite *SinglyLinkedListTestSuite) TestErase() {
	assert.EqualError(suite.T(), suite.list.PopFront(), "empty list")

	items := []string{"A", "B", "C", "D", "E", "F"}

	for _, i := range items {
		suite.list.PushBack(i)
	}

	for _, i := range items {
		suite.list.Erase(i)
	}

	for _, i := range items {
		assert.False(suite.T(), suite.list.Find(i))
	}
}

func TestSinglyLinkedListTestSuite(t *testing.T) {
	suite.Run(t, new(SinglyLinkedListTestSuite))
}
