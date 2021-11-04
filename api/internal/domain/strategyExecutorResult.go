package domain

type StrategyExecutorResult struct {
	ClosedOperations []*Operation
	OpenedOperations []*Operation
}
