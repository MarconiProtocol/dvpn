package configs

var serverLocation = ""

func SetRegion(region string) {
	serverLocation = region
}

func GetMyRegion() string {
	return serverLocation
}
