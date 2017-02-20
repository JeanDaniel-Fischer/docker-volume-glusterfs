package utils_test

import (
	"fmt"
	"testing"

	"github.com/Gimi/docker-volume-glusterfs/utils"
)

func TestGetGid(t *testing.T) {
	if id := utils.GetGid(); id < 0 {
		t.Error(fmt.Sprintf("expected utils.GetGid returns an integer no less than 0, but got %v", id))
	}
}
