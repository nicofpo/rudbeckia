package rudbeckia

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	extAPIPrefix = "http://ext.nicovideo.jp/api/getthumbinfo/"
)

var (
	ErrCanNotFetchVideoInfo = errors.New("Can't fetch video info")
)

type Video struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	ThumbnailURL  string    `json:"thumbnail_url"`
	Length        string    `json:"length"`
	FirstRetrieve time.Time `json:"first_retrieve"`
	Tags          []*Tag    `json:"tags"`
	User          *User     `json:"user"`
}

func NewVideo(id string) *Video {
	return &Video{
		ID: id,
	}
}

func (v *Video) Fetch() error {
	res, err := http.Get(extAPIPrefix + v.ID)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	return v.Parse(res.Body)
}

func (v *Video) Parse(r io.Reader) error {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return err
	}

	if doc.Find("nicovideo_thumb_response").AttrOr("status", "") != "ok" {
		return ErrCanNotFetchVideoInfo
	}

	v.ID = strip(doc.Find("nicovideo_thumb_response thumb video_id").Text())
	v.Title = strip(doc.Find("nicovideo_thumb_response thumb title").Text())
	v.Description = strip(doc.Find("nicovideo_thumb_response thumb description").Text())
	v.Length = strip(doc.Find("nicovideo_thumb_response thumb length").Text())
	v.ThumbnailURL = strip(doc.Find("nicovideo_thumb_response thumb thumbnail_url").Text())
	fr := doc.Find("nicovideo_thumb_response thumb first_retrieve").Text()

	if t, err := time.Parse(time.RFC3339, fr); err == nil {
		v.FirstRetrieve = t
	}

	tagNodes := doc.Find("nicovideo_thumb_response thumb tags tag")
	tags := make([]*Tag, tagNodes.Length())
	tagNodes.Each(func(i int, tagNode *goquery.Selection) {
		category := tagNode.AttrOr("category", "0") == "1"
		lock := tagNode.AttrOr("lock", "0") == "1"
		name := strip(tagNode.Text())

		tags[i] = &Tag{
			Name:     name,
			Category: category,
			Lock:     lock,
		}
	})
	v.Tags = tags

	uid := strip(doc.Find("nicovideo_thumb_response thumb user_id").Text())
	unickname := strip(doc.Find("nicovideo_thumb_response thumb user_nickname").Text())
	uiconurl := strip(doc.Find("nicovideo_thumb_response thumb user_icon_url").Text())
	v.User = &User{
		ID:       uid,
		Nickname: unickname,
		IconURL:  uiconurl,
	}

	return nil
}

func strip(s string) string {
	return strings.TrimSpace(s)
}
