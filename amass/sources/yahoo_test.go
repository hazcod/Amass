package sources

import (
	"testing"

	"github.com/hazcod/amass/amass/core"
)

func TestYahoo(t *testing.T) {
	if *networkTest == false {
		return
	}

	config := setupConfig(domainTest)
	bus, out := setupEventBus(core.NewNameTopic)
	defer bus.Stop()

	srv := NewYahoo(config, bus)

	result := testService(srv, out)
	if result < expectedTest {
		t.Errorf("Found %d names, expected at least %d instead", result, expectedTest)
	}

}
