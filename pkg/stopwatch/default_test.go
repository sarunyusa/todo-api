/*
 * Copyright (c) 2020. Inception Technology Co., Ltd.
 *
 */

package stopwatch

import (
	"fmt"
	"testing"
	"time"
	"todo/pkg/logger"
)

func TestStopWatch_Stop(t *testing.T) {
	t.Run("with logger", func(t *testing.T) {
		t.Parallel()
		defer StartWithLogger(logger.New("TestStopWatch")).Stop()

		time.Sleep(1500 * time.Millisecond)
	})

	t.Run("with default logger", func(t *testing.T) {
		t.Parallel()
		defer StartWithDefaultLogger().Stop()

		time.Sleep(1500 * time.Millisecond)
	})

	t.Run("without logger", func(t *testing.T) {
		t.Parallel()
		defer func(stopper Stopper) {
			fmt.Println(stopper.Stop())
		}(Start())

		time.Sleep(1500 * time.Millisecond)
	})
}
