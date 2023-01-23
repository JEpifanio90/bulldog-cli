package aws

import (
	"testing"

	"github.com/JEpifanio90/bulldog-cli/internal/models"
)

func TestFetchResourcesCDK(t *testing.T) {
	FetchResources(models.Command{CLI: false})
}

func TestFetchResourcesCli(t *testing.T) {
	FetchResources(models.Command{CLI: true})
}
