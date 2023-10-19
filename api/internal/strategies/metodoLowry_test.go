package strategies

//func TestLowrySystem_GetOperation(t *testing.T) {
//	strategyConfig := &domain.StrategyConfig{
//		Name: "lowry_system",
//		Parameters: []*domain.Parameter{
//			{
//				Name:  "lenght_4_sma",
//				Value: 4,
//				Min:   0,
//				Max:   0,
//			},
//			{
//				Name:  "lenght_18_sma",
//				Value: 18,
//				Min:   0,
//				Max:   0,
//			},
//			{
//				Name:  "lenght_40_sma",
//				Value: 40,
//				Min:   0,
//				Max:   0,
//			},
//		},
//	}
//	lowrySystem := NewLowrySystemFromConfig(strategyConfig)
//
//	candleStickCollection, err := utils.GetCandleStickOfMock("../mocks/historical_BTCUSDT_5m_phenomenon.json")
//	assert.Nil(t, err)
//
//	operation := lowrySystem.GetOperation(candleStickCollection)
//	assert.NotNil(t, operation)
//}
