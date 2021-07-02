package model

var (
	APPID      = "share-service"
	TimeLayout = "2006-01-02 15:04:05"

	_HShareIdPoolKey   = "Hshare:Seq"
	_HShareIdPoolField = "seq"
)

// Kratos hello kratos.
type Kratos struct {
	Hello string
}

func GetHShareIdPoolKey() string {
	return _HShareIdPoolKey
}
