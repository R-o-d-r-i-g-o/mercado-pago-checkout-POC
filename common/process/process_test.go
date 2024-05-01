package process

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type GracefulShutdownTestSuite struct {
	suite.Suite
	shutdownCallbackCalled chan os.Signal
}

func (suite *GracefulShutdownTestSuite) SetupTest() {
	suite.shutdownCallbackCalled = make(chan os.Signal)
	signal.Notify(suite.shutdownCallbackCalled, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
}

func (suite *GracefulShutdownTestSuite) TearDownTest() {
	signal.Reset(os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
}

func (suite *GracefulShutdownTestSuite) TestGracefulShutdown() {
	shutdownCallback := func(ctx context.Context) {
		close(suite.shutdownCallbackCalled)
	}

	go func() {
		GracefulShutdown(shutdownCallback, GracefulShutdownOptions{
			TimeoutSeconds: 1,
			Quit:           suite.shutdownCallbackCalled,
		})
	}()

	// Simulate receiving an interrupt signal
	suite.shutdownCallbackCalled <- os.Interrupt

	select {
	case <-suite.shutdownCallbackCalled:
		// Shutdown callback was called, the test is successful
	case <-time.After(2 * time.Second):
		suite.Fail("Timeout waiting for shutdown callback to be called")
	}
}

func TestGracefulShutdownSuite(t *testing.T) {
	suite.Run(t, new(GracefulShutdownTestSuite))
}
