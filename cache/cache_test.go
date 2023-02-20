package cache

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/google/uuid"
)

func TestDiscovery(t *testing.T) {
	fmt.Println("tests discovered")
}

func TestNew(t *testing.T) {
	id := uuid.New()
    testPath := path.Join(os.TempDir(), id.String())
    defer os.RemoveAll(testPath)
    New(10, testPath)
}
