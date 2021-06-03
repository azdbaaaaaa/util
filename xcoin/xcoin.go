package xcoin

const CoinBase int64 = 10000

type Coin int64

func (c Coin) Set() int64 {
	return int64(c) * CoinBase
}

func (c Coin) Get() float32 {
	return float32(int64(c) / CoinBase)
}
