package project

type AWSRegion string

const (
	AWSRegionUsEast1 = "us-east-1"
	AWSRegionUsWest1 = "us-west-1"
)

type AWS struct {
	AccessKeyId     string    `json:"access_key_id"`
	SecretAccessKey string    `json:"secret_access_key"`
	Region          AWSRegion `json:"region"`
}

// ValidAWSRegion method
func ValidAWSRegion(r AWSRegion) bool {
	AWSRegions := []AWSRegion{
		AWSRegionUsEast1,
		AWSRegionUsWest1,
	}

	for _, awsRegion := range AWSRegions {
		if awsRegion == r {
			return true
		}
	}

	return false
}
