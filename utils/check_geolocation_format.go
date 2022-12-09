package utils

func IsGeolocationStringInputAllowed(lat string, lng string) error {
	var err error

	if len(lat) != 10 {
		return &latError{}
	}

	if len(lng) != 11 {
		return &lngError{}
	}

	return err
}