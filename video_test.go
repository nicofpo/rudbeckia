package rudbeckia

import (
	"strings"
	"testing"
)

var ExampleVideoInfo = `<nicovideo_thumb_response status="ok">
<thumb>
<video_id>sm30712528</video_id>
<title>【三国志大戦４】武力なんて飾りです　その３</title>
<description>※この動画は低品、低スキル動画です。
兵力王異と遠弓蒋欽がほしい・・・</description>
<thumbnail_url>http://tn-skr1.smilevideo.jp/smile?i=30712528</thumbnail_url>
<first_retrieve>2017-02-26T01:51:30+09:00</first_retrieve>
<length>7:49</length>
<movie_type>mp4</movie_type>
<size_high>64034726</size_high>
<size_low>25154652</size_low>
<view_counter>11</view_counter>
<comment_num>0</comment_num>
<mylist_counter>0</mylist_counter>
<last_res_body/>
<watch_url>http://www.nicovideo.jp/watch/sm30712528</watch_url>
<thumb_type>video</thumb_type>
<embeddable>1</embeddable>
<no_live_play>0</no_live_play>
<tags domain="jp">
<tag category="1" lock="1">ゲーム</tag>
<tag lock="1">三国志大戦</tag>
<tag lock="1">三国志大戦４</tag>
</tags>
<user_id>36412383</user_id>
<user_nickname>保証書</user_nickname>
<user_icon_url>
https://secure-dcdn.cdn.nimg.jp/nicoaccount/usericon/defaults/blank_s.jpg
</user_icon_url>
</thumb>
</nicovideo_thumb_response>
`

func TestVideo(t *testing.T) {
	buf := strings.NewReader(ExampleVideoInfo)
	video := NewVideo("sm30712528")

	err := video.Parse(buf)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%#v", video)
	t.Logf("%s", video.Description)
	t.Logf("%#v", video.Tags)
	t.Logf("%#v", video.User)
}
