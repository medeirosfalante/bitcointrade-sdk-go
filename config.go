package bitcointrade

type Config struct {
}

var (
	uriV1 = "https://api.bitcointrade.com.br/v1"
)

//GetBasePath api bitcointrade
func GetBasePath() string {
	return uriV1
}
