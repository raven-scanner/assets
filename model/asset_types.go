package model

import (
	"fmt"
	"net"
	"regexp"
	"errors"
	"context"
)

type AssetTypeValidator func(string) (string, error)


var AssetTypeValidators map[string]AssetTypeValidator = map[string]AssetTypeValidator{
	"ipv4": validate_type_ipv4,
	"ipv6": validate_type_ipv6,
	"netipv4": validate_type_netipv4,
	"netipv6": validate_type_netipv6,
	"hostname": validate_type_hostname,
	"domain": validate_type_domain,
	"webapp": validate_type_webapp,
	"awsaccount": validate_type_awsaccount,
	"gitrepo": validate_type_gitrepo,
}

func validate_type(_type string, asset string) (string, error) {
	return AssetTypeValidators[_type](asset)
}

func validate_type_ipv4(asset string) (string, error) {
	ipv4 := net.ParseIP(asset).To4()
	if (ipv4 == nil) {
		return asset, errors.New("Invalid IPv4")
	}
	return asset, nil
}


func validate_type_ipv6(asset string) (string, error) {
	ipv4 := net.ParseIP(asset).To4()
	ipv6 := net.ParseIP(asset).To16()
	if (ipv4 != nil || ipv6 == nil) {
		return asset, errors.New("Invalid IPv6")
	}
	return asset, nil
}

func validate_type_netipv4(asset string) (string, error) {
	ip, _, err := net.ParseCIDR(asset)
	ipv4 := ip.To4()
	fmt.Println(err, ip, ipv4)
	if (err != nil || ipv4 == nil) {
		return asset, errors.New("Invalid Network IPv4")
	}
	return asset, nil
}

func validate_type_netipv6(asset string) (string, error) {
	ip, _, err := net.ParseCIDR(asset)
	ipv4 := ip.To4()
	ipv6 := ip.To16()
	if (err != nil || ipv4 != nil || ipv6 == nil) {
		return asset, errors.New("Invalid Network IPv6")
	}
	return asset, nil
}

func validate_type_hostname(asset string) (string, error) {
	_, err := net.LookupHost(asset)
	if (err != nil) {
		return asset, errors.New("Invalid Hostname")
	}
	return asset, nil
}

func validate_type_domain(asset string) (string, error) {
	// Bypass kubernetes DNS that do not allow NS resolutions.
	r := &net.Resolver{
		Dial: func(ctx context.Context, _, _ string) (net.Conn, error) {
			return net.Dial("udp", "8.8.8.8:53")
    		},
	}
	_, err := r.LookupNS(context.Background(), asset)
	if (err != nil) {
		return asset, errors.New("Invalid Domain")
	}
	return asset, nil
}

func validate_type_webapp(asset string) (string, error) {
	fmt.Println("Validate Web Application")
	return asset, nil
}

func validate_type_awsaccount(asset string) (string, error) {
	re := regexp.MustCompile("^[0-9]{12}$")
	if !re.Match([]byte(asset)) {
		return asset, errors.New("Invalid AWS Account")
	}
	return asset, nil
}

func validate_type_gitrepo(asset string) (string, error) {
	fmt.Println("Validate Git Repo")
	return asset, nil
}
