package dbif

import (
	"fmt"

	"github.com/newton-miku/nali/pkg/cdn"
	"github.com/newton-miku/nali/pkg/geoip"
	"github.com/newton-miku/nali/pkg/ip2location"
	"github.com/newton-miku/nali/pkg/ip2region"
	"github.com/newton-miku/nali/pkg/ipip"
	"github.com/newton-miku/nali/pkg/qqwry"
	"github.com/newton-miku/nali/pkg/zxipv6wry"
)

type QueryType uint

const (
	TypeIPv4 = iota
	TypeIPv6
	TypeDomain
)

type DB interface {
	Find(query string, params ...string) (result fmt.Stringer, err error)
	Name() string
}

var (
	_ DB = &qqwry.QQwry{}
	_ DB = &zxipv6wry.ZXwry{}
	_ DB = &ipip.IPIPFree{}
	_ DB = &geoip.GeoIP{}
	_ DB = &ip2region.Ip2Region{}
	_ DB = &ip2location.IP2Location{}
	_ DB = &cdn.CDN{}
)
