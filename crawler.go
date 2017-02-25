package rudbeckia

import (
	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"
)

const (
	watchURLPrefix = "http://www.nicovideo.jp/watch/"
)

type Crawler struct {
	Logger       logrus.FieldLogger
	feedURL      string
	parser       *gofeed.Parser
	latestVideID string
}

func NewCrawler(feedURL string) *Crawler {
	return &Crawler{
		feedURL: feedURL,
		parser:  gofeed.NewParser(),
	}
}

func (c *Crawler) Fetch() ([]*Video, error) {
	videos := []*Video{}

	c.Logger.WithFields(logrus.Fields{
		"url": c.feedURL,
	}).Info("Fetch newarrivals feed")

	feed, err := c.parser.ParseURL(c.feedURL)
	if err != nil {
		return videos, err
	}

	vIDs := make([]string, 0, len(feed.Items))
	for _, item := range feed.Items {
		vID := item.Link[len(watchURLPrefix):]
		if c.latestVideID == vID {
			break
		}
		vIDs = append(vIDs, vID)
	}

	if len(vIDs) == 0 {
		return videos, nil
	}
	c.latestVideID = vIDs[0]

	videos = make([]*Video, len(vIDs))
	for i, vID := range vIDs {
		video := NewVideo(vID)
		err := video.Fetch()
		if err != nil {
			return videos, err
		}
		videos[i] = video

		c.Logger.WithFields(logrus.Fields{
			"id":    video.ID,
			"video": video,
		}).Infof("Newarrival: %s", video.Title)
	}

	return videos, nil
}
