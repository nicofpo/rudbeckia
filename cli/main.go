package main

import (
	"fmt"
	"github.com/nicofpo/rudbeckia"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
	"time"
)

func main() {
	interval := 5 * time.Second

	app := cli.NewApp()
	app.Name = "rudbeckia"
	app.Usage = "niconico video new arriaval crawler"
	app.Flags = []cli.Flag{
		cli.DurationFlag{
			Name:        "interval, i",
			Usage:       "Fetch interval `DURATION`",
			EnvVar:      "RUDBECKIA_INTERVAL",
			Value:       interval,
			Destination: &interval,
		},
	}
	app.Action = func(c *cli.Context) error {
		crawler := rudbeckia.NewCrawler(rudbeckia.FEED_URL)
		crawler.Logger = logrus.StandardLogger()
		for {
			videos, err := crawler.Fetch()
			if err != nil {
				fmt.Printf("%s\n", err)
				continue
			}

			for _, video := range videos {
				fmt.Printf("%s\n", video.Title)
			}

			<-time.After(interval)
		}

		return nil
	}

	app.Run(os.Args)
}
