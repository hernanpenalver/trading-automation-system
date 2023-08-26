package presenters

const (
	ConsolePresenterName = "console"
	MetricPresenterName  = "metric"
)

var presenterRepository = map[string]Presenter{
	ConsolePresenterName: NewConsolePresenter(),
	MetricPresenterName:  NewMetricPresenter(),
}

func GetPresenterByName(name string) Presenter {
	if presenter, ok := presenterRepository[name]; ok {
		return presenter
	}

	return nil
}
