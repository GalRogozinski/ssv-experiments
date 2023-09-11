package full_flow

import (
	"github.com/stretchr/testify/require"
	"ssv-experiments/new_arch_2/tests"
	"testing"
	"ssv-experiments/new_arch_2/spec/asgard/types"
	"ssv-experiments/new_arch_2/tests/spec/asgard/fixtures"
	"ssv-experiments/new_arch_2/tests/spec/asgard/ssv"
)

func TestFullFlow(t *testing.T) {
	tst, err := tests.NewSpecTest[*ssv.ProcessMessageTest](FullFlow())
	require.NoError(t, err)
	tst.Run(t, fixtures.Share)
}

var AllTests = []tests.TestObject{FullFlow()}
