package loggermock

import (
	loggermocks "github.com/rockholla/go-lib/mocks/logger"
	mock "github.com/stretchr/testify/mock"
)

var logMock *loggermocks.Interface

// GetLogMock will return a common mock complete with generic expected calls
// this mock will help just swallow all logging used in other testing scenarios
func GetLogMock() *loggermocks.Interface {
	if logMock != nil {
		return logMock
	}
	logMock = &loggermocks.Interface{}
	var expectedArgs []interface{}
	expectedArgs = append(expectedArgs, mock.Anything)
	for i := 0; i <= 50; i++ {
		logMock.On("Info", expectedArgs...).Return(nil)
		logMock.On("InfoPart", expectedArgs...).Return(nil)
		logMock.On("Log", expectedArgs...).Return(nil)
		logMock.On("Focused", expectedArgs...).Return(nil)
		logMock.On("Error", expectedArgs...).Return(nil)
		logMock.On("Errors", expectedArgs...).Return(nil)
		logMock.On("ListItem", expectedArgs...).Return(nil)
		logMock.On("LogPart", expectedArgs...).Return(nil)
		expectedArgs = append(expectedArgs, mock.Anything)
	}
	logMock.On("SpinnerStart", mock.AnythingOfType("string")).Return(nil)
	logMock.On("SpinnerStop").Return(nil)
	return logMock
}
