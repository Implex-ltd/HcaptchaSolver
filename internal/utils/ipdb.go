package utils

import (
	"net"

	"github.com/oschwald/maxminddb-golang"
)

func Lookup(Address string) (*IpInfos, error) {
	db, err := maxminddb.Open("../../assets/GeoLite2-City.mmdb")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	ip := net.ParseIP(Address)

	var data Data
	err = db.Lookup(ip, &data)
	if err != nil {
		return nil, err
	}

	return &IpInfos{
		IsoCode:  data.Country.IsoCode,
		TimeZone: data.Location.TimeZone,
	}, nil
}
