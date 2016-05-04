package pulse

import "github.com/byuoitav/defibrillator/packages/elastic"

func Check() error {
	hostnames, err := elastic.GetHostnames()
	if err != nil {
		return err
	}

	for i := range hostnames.Hostnames {
		println(hostnames.Hostnames[i].Name)
	}

	return nil
}
