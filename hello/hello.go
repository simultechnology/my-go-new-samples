package hello

import (
	quoteV3 "rsc.io/quote/v3"
)

func Hello() string {
	//return quote.Hello()
	return quoteV3.HelloV3()
}

func Proverb() string {
	return quoteV3.Concurrency()
}
