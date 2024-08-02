package nasConvert_test

import (
	"testing"

	"github.com/free5gc/nas/nasConvert"
	"github.com/free5gc/nas/nasType"
	"github.com/free5gc/openapi/models"
	"github.com/stretchr/testify/require"
)

func TestSnssaiToModels(t *testing.T) {
	testCase := []struct {
		Name         string
		nasSnssai    nasType.SNSSAI
		expectSnssai models.Snssai
	}{
		{
			Name: "Default",
			nasSnssai: nasType.SNSSAI{
				Iei:   uint8(1),
				Len:   uint8(4),
				Octet: [8]uint8{1, 1, 2, 3},
			},
			expectSnssai: models.Snssai{
				Sst: int32(1),
				Sd:  "010203",
			},
		},
		{
			Name: "Empty SD",
			nasSnssai: nasType.SNSSAI{
				Iei:   uint8(1),
				Len:   uint8(1),
				Octet: [8]uint8{1},
			},
			expectSnssai: models.Snssai{
				Sst: int32(1),
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.Name, func(t *testing.T) {
			result := nasConvert.SnssaiToModels(&tc.nasSnssai)
			require.Equal(t, tc.expectSnssai, result)
		})
	}
}
