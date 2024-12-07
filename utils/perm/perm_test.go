package perm

import (
	"testing"

	"github.com/staffbase-robert/aoc2024/utils"
)

func TestEqual(t *testing.T) {
	perms := Equal(3, []string{"a", "b"})
	utils.MustLen(perms, 2*2*2)
}
