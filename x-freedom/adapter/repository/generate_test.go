package repository

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDescPager(t *testing.T) {
	pager := NewDescPager("created", "updated", "id")
	fmt.Printf("%v\n", pager)
	assert.Equal(t, pager.Order(), "desc created, desc updated, desc id")

}
