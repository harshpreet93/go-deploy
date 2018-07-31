package deploy

import (
	"testing"
	"os"
)

func TestSpinUp(*testing.T) {
	spinUpNewDroplet("test-droplet", "sfo2", true, 1, 1,
		os.Getenv("DO_TOKEN"))
}
