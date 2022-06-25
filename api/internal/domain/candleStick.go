package domain

type CandleStick struct {
	CloseTime    int64
	OpenTime     int64
	Close        float64
	Open         float64
	Max          float64
	Min          float64
	OpenDateTime string
}
