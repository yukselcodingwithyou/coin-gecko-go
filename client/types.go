package client

type FromClient int64
type ImplementedMethodName string

func NewFromClient() *FromClient {
	return new(FromClient)
}

func (fc *FromClient) IsFromCli(clientType FromClient) bool {
	return clientType == CLI
}

const (
	CLI  FromClient = 1
	HTTP            = 2
)

const (
	Ping      ImplementedMethodName = "ping"
	ListCoins                       = "list-coins"
)
