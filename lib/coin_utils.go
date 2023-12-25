package lib

const changeFactor = float64(1000000000)

// converting PAC coin amount to satoshi amount.
func PACToSatoshiPAC(coin float64) int64 {
	return int64(coin * changeFactor)
}

// converting satoshi amount to PAC coin amount.
func SatoshiPACtoPAC(change int64) float64 {
	return float64(change) / changeFactor
}
