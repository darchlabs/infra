package project

type DORegion string

const (
	DORegionNYC1 = "nyc1"
)

type DO struct {
	Token  string   `json:"token"`
	Region DORegion `json:"region"`
}

// ValidDORegion method
func ValidDORegion(r DORegion) bool {
	DORegions := []DORegion{
		DORegionNYC1,
	}

	for _, doRegion := range DORegions {
		if doRegion == r {
			return true
		}
	}

	return false
}
