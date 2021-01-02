package dcmtime

import (
	"testing"
	"time"
)

var fullDateTime = time.Date(
	1010,
	2,
	3,
	4,
	5,
	6,
	456789,
	time.FixedZone("", -3720),
)

func TestTimeToDA(t *testing.T) {
	testCases := []struct{
		TimeVal    time.Time
		Expected   string
		Truncation TruncationLimit
		NoOffset bool
	}{
		{
			TimeVal:    time.Date(
				1010,
				2,
				3,
				4,
				5,
				6,
				456789000,
				time.FixedZone("", -3720),
			),
			Expected:   "10100203040506.456789000-0102",
			Truncation: Truncation.None,
			NoOffset:   false,
		},
	}
	
	for _, thisCase := range testCases {
		t.Run(thisCase.Expected, func(t *testing.T) {
			
		})
	}
}
