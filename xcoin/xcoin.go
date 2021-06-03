package xcoin

import "errors"

const CoinBase int64 = 10000

type Coin int64

func (c Coin) Set() int64 {
	return int64(c) * CoinBase
}

func (c Coin) Get() (int64, error) {
	if int64(c)%CoinBase != 0 {
		return 0, errors.New("coin is invalid")
	}
	return int64(c) / CoinBase, nil
}
