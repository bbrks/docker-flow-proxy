package logging

import (
	"fmt"
	"log"
	"log/syslog"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type LoggingTestSuite struct {
	suite.Suite
}

func TestLoggingUnitTestSuite(t *testing.T) {
	logPrintf = func(format string, v ...interface{}) {}
	suite.Run(t, new(LoggingTestSuite))
}

func (s *LoggingTestSuite) SetupTest() {}

// StartLogging

func (s LoggingTestSuite) Test_StartLogging_OutputsSyslogToStdOut() {
	/*
		FIXME: Disabled test that fails when run inside the Docker builder
				--- FAIL: TestLoggingUnitTestSuite (10.81s)
			    --- FAIL: TestLoggingUnitTestSuite/Test_StartLogging_OutputsSyslogToStdOut (10.81s)
			        logging_test.go:51:
			                Error Trace:    logging_test.go:51
			                Error:          Should be true
			                Test:           TestLoggingUnitTestSuite/Test_StartLogging_OutputsSyslogToStdOut
			FAIL
			coverage: 100.0% of statements
			FAIL    github.com/docker-flow/docker-flow-proxy/logging        10.833s
	*/
	s.T().Skip("test fails inside Docker builder")

	logPrintfOrig := logPrintf
	defer func() { logPrintf = logPrintfOrig }()
	actual := ""
	logPrintf = func(format string, v ...interface{}) {
		actual = fmt.Sprintf(format, v...)
	}

	go StartLogging()

	for i := 1; i <= 10; i++ {
		sysLog, err := syslog.Dial("udp", "127.0.0.1:1514", syslog.LOG_INFO, "testing")
		if err != nil {
			log.Fatal(err)
		}
		expected := fmt.Sprintf("This is a syslog message %d", i)
		go sysLog.Info(expected)
		logged := false
		for c := 0; c < 1000; c++ {
			if strings.Contains(actual, expected) {
				logged = true
				break
			}
			time.Sleep(10 * time.Millisecond)
		}
		s.True(logged)
	}
}
