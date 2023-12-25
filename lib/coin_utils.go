package lib

const changeFactor = float64(1000000000)

func PACToSatoshiPAC(coin float64) int64 {
	return int64(coin * changeFactor)
}

func SatoshiPACtoPAC(change int64) float64 {
	return float64(change) / changeFactor
}
