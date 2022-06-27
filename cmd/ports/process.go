package ports

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var (
	processCommand = &cli.Command{
		Name:    "process",
		Aliases: []string{"p"},
		Usage:   "start task processor",
		Action:  processAction,
	}
)

// Placeholder for command to launch separate task processor
// when tasks receiving using queue like RabbitMq or Kafka
func processAction(c *cli.Context) error {
	_, err := initConfig(c.String("config"), "pds")
	if err != nil {
		return fmt.Errorf("error initialising config: %w", err)
	}

	_, err = initLogger(zap.PanicLevel)
	if err != nil {
		return fmt.Errorf("error initialising logger: %w", err)
	}
	return nil
}
