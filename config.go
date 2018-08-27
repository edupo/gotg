package gotg

type Config struct {
	Address string
}

var DefaultConfig = Config{"127.0.0.1:4458"}
