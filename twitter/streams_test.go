package twitter

import (
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStream_MessageJSONError(t *testing.T) {
	badJSON := []byte(`{`)
	msg := getMessage(badJSON)
	assert.EqualError(t, msg.(error), "unexpected end of JSON input")
}

func TestStream_GetMessageTweet(t *testing.T) {
	msgJSON := []byte(`{"id": 20, "text": "just setting up my twttr", "retweet_count": "68535"}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &Tweet{}, msg)
}

func TestStream_GetMessageTweetWithEntities(t *testing.T) {
	message := []byte(`
	{
		"created_at": "Mon Sep 26 11:35:44 +0000 2016",
		"id": 780370264324509696,
		"id_str": "780370264324509696",
		"text": "RT @CITE_CCOOCat: Precarietat i irregularitat defineixen la pres\u00e8ncia de la poblaci\u00f3 estrangera al mercat de treball catal\u00e0 https:\/\/t.co\/8S\u2026",
		"source": "\u003ca href=\"http:\/\/twitter.com\" rel=\"nofollow\"\u003eTwitter Web Client\u003c\/a\u003e",
		"truncated": false,
		"in_reply_to_status_id": null,
		"in_reply_to_status_id_str": null,
		"in_reply_to_user_id": null,
		"in_reply_to_user_id_str": null,
		"in_reply_to_screen_name": null,
		"user": {
			"id": 91987403,
			"id_str": "91987403",
			"name": "CCOO de Catalunya",
			"screen_name": "ccoocatalunya",
			"location": "Catalunya",
			"url": "http:\/\/www.ccoo.cat",
			"description": "CCOO \u00e9s un sindicat d\u2019homes i dones que ens afiliem per defensar els nostres interessos i per aconseguir una societat m\u00e9s justa, democr\u00e0tica i participativa",
			"protected": false,
			"verified": true,
			"followers_count": 11606,
			"friends_count": 3756,
			"listed_count": 311,
			"favourites_count": 5566,
			"statuses_count": 29295,
			"created_at": "Mon Nov 23 11:04:35 +0000 2009",
			"utc_offset": 7200,
			"time_zone": "Madrid",
			"geo_enabled": false,
			"lang": "ca",
			"contributors_enabled": false,
			"is_translator": false,
			"profile_background_color": "FAFAFA",
			"profile_background_image_url": "http:\/\/pbs.twimg.com\/profile_background_images\/81724022\/twitter_ccoo_catalunya.jpg",
			"profile_background_image_url_https": "https:\/\/pbs.twimg.com\/profile_background_images\/81724022\/twitter_ccoo_catalunya.jpg",
			"profile_background_tile": false,
			"profile_link_color": "E82828",
			"profile_sidebar_border_color": "C1C3C7",
			"profile_sidebar_fill_color": "F2F2F2",
			"profile_text_color": "333333",
			"profile_use_background_image": true,
			"profile_image_url": "http:\/\/pbs.twimg.com\/profile_images\/2024868969\/ccoocat_twitter_normal.png",
			"profile_image_url_https": "https:\/\/pbs.twimg.com\/profile_images\/2024868969\/ccoocat_twitter_normal.png",
			"profile_banner_url": "https:\/\/pbs.twimg.com\/profile_banners\/91987403\/1398235157",
			"default_profile": false,
			"default_profile_image": false,
			"following": null,
			"follow_request_sent": null,
			"notifications": null
		},
		"geo": null,
		"coordinates": null,
		"place": null,
		"contributors": null,
		"retweeted_status": {
			"created_at": "Mon Sep 26 11:33:16 +0000 2016",
			"id": 780369646075740160,
			"id_str": "780369646075740160",
			"text": "Precarietat i irregularitat defineixen la pres\u00e8ncia de la poblaci\u00f3 estrangera al mercat de treball catal\u00e0 https:\/\/t.co\/8SBiTzTSYy",
			"source": "\u003ca href=\"https:\/\/about.twitter.com\/products\/tweetdeck\" rel=\"nofollow\"\u003eTweetDeck\u003c\/a\u003e",
			"truncated": false,
			"in_reply_to_status_id": null,
			"in_reply_to_status_id_str": null,
			"in_reply_to_user_id": null,
			"in_reply_to_user_id_str": null,
			"in_reply_to_screen_name": null,
			"user": {
				"id": 778161012478533632,
				"id_str": "778161012478533632",
				"name": "CITE",
				"screen_name": "CITE_CCOOCat",
				"location": "Catalunya",
				"url": "http:\/\/www.ccoo.cat\/aspnet\/immigracio.aspx",
				"description": "El Centre d'Informaci\u00f3 per a Treballadors Estrangers (CITE) \u00e9s una associaci\u00f3 sense afany de lucre que assessora gratu\u00eftament sobre la Llei d'estrangeria",
				"protected": false,
				"verified": false,
				"followers_count": 35,
				"friends_count": 60,
				"listed_count": 0,
				"favourites_count": 6,
				"statuses_count": 12,
				"created_at": "Tue Sep 20 09:16:57 +0000 2016",
				"utc_offset": null,
				"time_zone": null,
				"geo_enabled": false,
				"lang": "ca",
				"contributors_enabled": false,
				"is_translator": false,
				"profile_background_color": "000000",
				"profile_background_image_url": "http:\/\/abs.twimg.com\/images\/themes\/theme1\/bg.png",
				"profile_background_image_url_https": "https:\/\/abs.twimg.com\/images\/themes\/theme1\/bg.png",
				"profile_background_tile": false,
				"profile_link_color": "1B95E0",
				"profile_sidebar_border_color": "000000",
				"profile_sidebar_fill_color": "000000",
				"profile_text_color": "000000",
				"profile_use_background_image": false,
				"profile_image_url": "http:\/\/pbs.twimg.com\/profile_images\/778163234251603969\/YYkHnnSp_normal.jpg",
				"profile_image_url_https": "https:\/\/pbs.twimg.com\/profile_images\/778163234251603969\/YYkHnnSp_normal.jpg",
				"profile_banner_url": "https:\/\/pbs.twimg.com\/profile_banners\/778161012478533632\/1474363418",
				"default_profile": false,
				"default_profile_image": false,
				"following": null,
				"follow_request_sent": null,
				"notifications": null
			},
			"geo": null,
			"coordinates": null,
			"place": null,
			"contributors": null,
			"quoted_status_id": 780367987761086464,
			"quoted_status_id_str": "780367987761086464",
			"quoted_status": {
				"created_at": "Mon Sep 26 11:26:41 +0000 2016",
				"id": 780367987761086464,
				"id_str": "780367987761086464",
				"text": "Informe 2016 de #ccoocatalunya sobre la situaci\u00f3 laboral de la poblaci\u00f3 estrangera: https:\/\/t.co\/sRZpz4eJTi #CCOO",
				"source": "\u003ca href=\"https:\/\/about.twitter.com\/products\/tweetdeck\" rel=\"nofollow\"\u003eTweetDeck\u003c\/a\u003e",
				"truncated": false,
				"in_reply_to_status_id": null,
				"in_reply_to_status_id_str": null,
				"in_reply_to_user_id": null,
				"in_reply_to_user_id_str": null,
				"in_reply_to_screen_name": null,
				"user": {
					"id": 91987403,
					"id_str": "91987403",
					"name": "CCOO de Catalunya",
					"screen_name": "ccoocatalunya",
					"location": "Catalunya",
					"url": "http:\/\/www.ccoo.cat",
					"description": "CCOO \u00e9s un sindicat d\u2019homes i dones que ens afiliem per defensar els nostres interessos i per aconseguir una societat m\u00e9s justa, democr\u00e0tica i participativa",
					"protected": false,
					"verified": true,
					"followers_count": 11606,
					"friends_count": 3756,
					"listed_count": 311,
					"favourites_count": 5566,
					"statuses_count": 29294,
					"created_at": "Mon Nov 23 11:04:35 +0000 2009",
					"utc_offset": 7200,
					"time_zone": "Madrid",
					"geo_enabled": false,
					"lang": "ca",
					"contributors_enabled": false,
					"is_translator": false,
					"profile_background_color": "FAFAFA",
					"profile_background_image_url": "http:\/\/pbs.twimg.com\/profile_background_images\/81724022\/twitter_ccoo_catalunya.jpg",
					"profile_background_image_url_https": "https:\/\/pbs.twimg.com\/profile_background_images\/81724022\/twitter_ccoo_catalunya.jpg",
					"profile_background_tile": false,
					"profile_link_color": "E82828",
					"profile_sidebar_border_color": "C1C3C7",
					"profile_sidebar_fill_color": "F2F2F2",
					"profile_text_color": "333333",
					"profile_use_background_image": true,
					"profile_image_url": "http:\/\/pbs.twimg.com\/profile_images\/2024868969\/ccoocat_twitter_normal.png",
					"profile_image_url_https": "https:\/\/pbs.twimg.com\/profile_images\/2024868969\/ccoocat_twitter_normal.png",
					"profile_banner_url": "https:\/\/pbs.twimg.com\/profile_banners\/91987403\/1398235157",
					"default_profile": false,
					"default_profile_image": false,
					"following": null,
					"follow_request_sent": null,
					"notifications": null
				},
				"geo": null,
				"coordinates": null,
				"place": null,
				"contributors": [
				3162711857
				],
				"is_quote_status": false,
				"retweet_count": 2,
				"favorite_count": 0,
				"entities": {
					"hashtags": [
					{
						"text": "ccoocatalunya",
						"indices": [
						16,
						30
						]
					},
					{
						"text": "CCOO",
						"indices": [
						108,
						113
						]
					}
					],
					"urls": [
					{
						"url": "https:\/\/t.co\/sRZpz4eJTi",
						"expanded_url": "http:\/\/www.ccoo.cat\/noticia\/204005\/informe-2016-sobre-la-situacio-laboral-de-la-poblacio-estrangera-a-catalunya",
						"display_url": "ccoo.cat\/noticia\/204005\u2026",
						"indices": [
						84,
						107
						]
					}
					],
					"user_mentions": [],
					"symbols": []
				},
				"favorited": false,
				"retweeted": false,
				"possibly_sensitive": false,
				"filter_level": "low",
				"lang": "und"
			},
			"is_quote_status": true,
			"retweet_count": 1,
			"favorite_count": 0,
			"entities": {
				"hashtags": [],
				"urls": [
				{
					"url": "https:\/\/t.co\/8SBiTzTSYy",
					"expanded_url": "https:\/\/twitter.com\/ccoocatalunya\/status\/780367987761086464",
					"display_url": "twitter.com\/ccoocatalunya\/\u2026",
					"indices": [
					106,
					129
					]
				}
				],
				"user_mentions": [],
				"symbols": []
			},
			"favorited": false,
			"retweeted": false,
			"possibly_sensitive": false,
			"filter_level": "low",
			"lang": "und"
		},
		"quoted_status_id": 780367987761086464,
		"quoted_status_id_str": "780367987761086464",
		"quoted_status": {
			"created_at": "Mon Sep 26 11:26:41 +0000 2016",
			"id": 780367987761086464,
			"id_str": "780367987761086464",
			"text": "Informe 2016 de #ccoocatalunya sobre la situaci\u00f3 laboral de la poblaci\u00f3 estrangera: https:\/\/t.co\/sRZpz4eJTi #CCOO",
			"source": "\u003ca href=\"https:\/\/about.twitter.com\/products\/tweetdeck\" rel=\"nofollow\"\u003eTweetDeck\u003c\/a\u003e",
			"truncated": false,
			"in_reply_to_status_id": null,
			"in_reply_to_status_id_str": null,
			"in_reply_to_user_id": null,
			"in_reply_to_user_id_str": null,
			"in_reply_to_screen_name": null,
			"user": {
				"id": 91987403,
				"id_str": "91987403",
				"name": "CCOO de Catalunya",
				"screen_name": "ccoocatalunya",
				"location": "Catalunya",
				"url": "http:\/\/www.ccoo.cat",
				"description": "CCOO \u00e9s un sindicat d\u2019homes i dones que ens afiliem per defensar els nostres interessos i per aconseguir una societat m\u00e9s justa, democr\u00e0tica i participativa",
				"protected": false,
				"verified": true,
				"followers_count": 11606,
				"friends_count": 3756,
				"listed_count": 311,
				"favourites_count": 5566,
				"statuses_count": 29294,
				"created_at": "Mon Nov 23 11:04:35 +0000 2009",
				"utc_offset": 7200,
				"time_zone": "Madrid",
				"geo_enabled": false,
				"lang": "ca",
				"contributors_enabled": false,
				"is_translator": false,
				"profile_background_color": "FAFAFA",
				"profile_background_image_url": "http:\/\/pbs.twimg.com\/profile_background_images\/81724022\/twitter_ccoo_catalunya.jpg",
				"profile_background_image_url_https": "https:\/\/pbs.twimg.com\/profile_background_images\/81724022\/twitter_ccoo_catalunya.jpg",
				"profile_background_tile": false,
				"profile_link_color": "E82828",
				"profile_sidebar_border_color": "C1C3C7",
				"profile_sidebar_fill_color": "F2F2F2",
				"profile_text_color": "333333",
				"profile_use_background_image": true,
				"profile_image_url": "http:\/\/pbs.twimg.com\/profile_images\/2024868969\/ccoocat_twitter_normal.png",
				"profile_image_url_https": "https:\/\/pbs.twimg.com\/profile_images\/2024868969\/ccoocat_twitter_normal.png",
				"profile_banner_url": "https:\/\/pbs.twimg.com\/profile_banners\/91987403\/1398235157",
				"default_profile": false,
				"default_profile_image": false,
				"following": null,
				"follow_request_sent": null,
				"notifications": null
			},
			"geo": null,
			"coordinates": null,
			"place": null,
			"contributors": [
			3162711857
			],
			"is_quote_status": false,
			"retweet_count": 2,
			"favorite_count": 0,
			"entities": {
				"hashtags": [
				{
					"text": "ccoocatalunya",
					"indices": [
					16,
					30
					]
				},
				{
					"text": "CCOO",
					"indices": [
					108,
					113
					]
				}
				],
				"urls": [
				{
					"url": "https:\/\/t.co\/sRZpz4eJTi",
					"expanded_url": "http:\/\/www.ccoo.cat\/noticia\/204005\/informe-2016-sobre-la-situacio-laboral-de-la-poblacio-estrangera-a-catalunya",
					"display_url": "ccoo.cat\/noticia\/204005\u2026",
					"indices": [
					84,
					107
					]
				}
				],
				"user_mentions": [],
				"symbols": []
			},
			"favorited": false,
			"retweeted": false,
			"possibly_sensitive": false,
			"filter_level": "low",
			"lang": "und"
		},
		"is_quote_status": true,
		"retweet_count": 0,
		"favorite_count": 0,
		"entities": {
			"hashtags": [],
			"urls": [],
			"user_mentions": [
			{
				"screen_name": "CITE_CCOOCat",
				"name": "CITE",
				"id": 778161012478533632,
				"id_str": "778161012478533632",
				"indices": [
				3,
				16
				]
			}
			],
			"symbols": []
		},
		"favorited": false,
		"retweeted": false,
		"filter_level": "low",
		"lang": "und",
		"timestamp_ms": "1474889744250"
	}
	`)
	msg := getMessage(message)
	assert.IsType(t, &Tweet{}, msg)
	tweet := msg.(*Tweet)
	assert.NotNil(t, tweet.Entities)
}

func TestStream_GetMessageDirectMessage(t *testing.T) {
	msgJSON := []byte(`{"direct_message": {"id": 666024290140217347}}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &DirectMessage{}, msg)
}

func TestStream_GetMessageDelete(t *testing.T) {
	msgJSON := []byte(`{"delete": { "id": 20}}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &StatusDeletion{}, msg)
}

func TestStream_GetMessageLocationDeletion(t *testing.T) {
	msgJSON := []byte(`{"scrub_geo": { "up_to_status_id": 20}}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &LocationDeletion{}, msg)
}

func TestStream_GetMessageStreamLimit(t *testing.T) {
	msgJSON := []byte(`{"limit": { "track": 10 }}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &StreamLimit{}, msg)
}

func TestStream_StatusWithheld(t *testing.T) {
	msgJSON := []byte(`{"status_withheld": { "id": 20, "user_id": 12, "withheld_in_countries":["USA", "China"] }}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &StatusWithheld{}, msg)
}

func TestStream_UserWithheld(t *testing.T) {
	msgJSON := []byte(`{"user_withheld": { "id": 12, "withheld_in_countries":["USA", "China"] }}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &UserWithheld{}, msg)
}

func TestStream_StreamDisconnect(t *testing.T) {
	msgJSON := []byte(`{"disconnect": { "code": "420", "stream_name": "streaming stuff", "reason": "too many connections" }}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &StreamDisconnect{}, msg)
}

func TestStream_StallWarning(t *testing.T) {
	msgJSON := []byte(`{"warning": { "code": "420", "percent_full": 90, "message": "a lot of messages" }}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &StallWarning{}, msg)
}

func TestStream_FriendsList(t *testing.T) {
	msgJSON := []byte(`{"friends": [666024290140217347, 666024290140217349, 666024290140217342]}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &FriendsList{}, msg)
}

func TestStream_Event(t *testing.T) {
	msgJSON := []byte(`{"event": "block", "target": {"name": "XKCD Comic", "favourites_count": 2}, "source": {"name": "XKCD Comic2", "favourites_count": 3}, "created_at": "Sat Sep 4 16:10:54 +0000 2010"}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, &Event{}, msg)
}

func TestStream_Unknown(t *testing.T) {
	msgJSON := []byte(`{"unknown_data": {"new_twitter_type":"unexpected"}}`)
	msg := getMessage(msgJSON)
	assert.IsType(t, map[string]interface{}{}, msg)
}

func TestStream_Filter(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	reqCount := 0
	mux.HandleFunc("/1.1/statuses/filter.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertQuery(t, map[string]string{"track": "gophercon,golang"}, r)
		switch reqCount {
		case 0:
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Transfer-Encoding", "chunked")
			fmt.Fprintf(w,
				`{"text": "Gophercon talks!"}`+"\r\n"+
					`{"text": "Gophercon super talks!"}`+"\r\n",
			)
		default:
			// Only allow first request
			http.Error(w, "Stream API not available!", 130)
		}
		reqCount++
	})

	counts := &counter{}
	demux := newCounterDemux(counts)
	client := NewClient(httpClient)
	streamFilterParams := &StreamFilterParams{
		Track: []string{"gophercon", "golang"},
	}
	stream, err := client.Streams.Filter(streamFilterParams)
	// assert that the expected messages are received
	assert.NoError(t, err)
	defer stream.Stop()
	for message := range stream.Messages {
		demux.Handle(message)
	}
	expectedCounts := &counter{all: 2, other: 2}
	assert.Equal(t, expectedCounts, counts)
}

func TestStream_Sample(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	reqCount := 0
	mux.HandleFunc("/1.1/statuses/sample.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"stall_warnings": "true"}, r)
		switch reqCount {
		case 0:
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Transfer-Encoding", "chunked")
			fmt.Fprintf(w,
				`{"text": "Gophercon talks!"}`+"\r\n"+
					`{"text": "Gophercon super talks!"}`+"\r\n",
			)
		default:
			// Only allow first request
			http.Error(w, "Stream API not available!", 130)
		}
		reqCount++
	})

	counts := &counter{}
	demux := newCounterDemux(counts)
	client := NewClient(httpClient)
	streamSampleParams := &StreamSampleParams{
		StallWarnings: Bool(true),
	}
	stream, err := client.Streams.Sample(streamSampleParams)
	// assert that the expected messages are received
	assert.NoError(t, err)
	defer stream.Stop()
	for message := range stream.Messages {
		demux.Handle(message)
	}
	expectedCounts := &counter{all: 2, other: 2}
	assert.Equal(t, expectedCounts, counts)
}

func TestStream_User(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	reqCount := 0
	mux.HandleFunc("/1.1/user.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"stall_warnings": "true", "with": "followings"}, r)
		switch reqCount {
		case 0:
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Transfer-Encoding", "chunked")
			fmt.Fprintf(w, `{"friends": [666024290140217347, 666024290140217349, 666024290140217342]}`+"\r\n"+"\r\n")
		default:
			// Only allow first request
			http.Error(w, "Stream API not available!", 130)
		}
		reqCount++
	})

	counts := &counter{}
	demux := newCounterDemux(counts)
	client := NewClient(httpClient)
	streamUserParams := &StreamUserParams{
		StallWarnings: Bool(true),
		With:          "followings",
	}
	stream, err := client.Streams.User(streamUserParams)
	// assert that the expected messages are received
	assert.NoError(t, err)
	defer stream.Stop()
	for message := range stream.Messages {
		demux.Handle(message)
	}
	expectedCounts := &counter{all: 1, friendsList: 1}
	assert.Equal(t, expectedCounts, counts)
}

func TestStream_Site(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	reqCount := 0
	mux.HandleFunc("/1.1/site.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"follow": "666024290140217347,666024290140217349"}, r)
		switch reqCount {
		case 0:
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Transfer-Encoding", "chunked")
			fmt.Fprintf(w,
				`{"text": "Gophercon talks!"}`+"\r\n"+
					`{"text": "Gophercon super talks!"}`+"\r\n",
			)
		default:
			// Only allow first request
			http.Error(w, "Stream API not available!", 130)
		}
		reqCount++
	})

	counts := &counter{}
	demux := newCounterDemux(counts)
	client := NewClient(httpClient)
	streamSiteParams := &StreamSiteParams{
		Follow: []string{"666024290140217347", "666024290140217349"},
	}
	stream, err := client.Streams.Site(streamSiteParams)
	// assert that the expected messages are received
	assert.NoError(t, err)
	defer stream.Stop()
	for message := range stream.Messages {
		demux.Handle(message)
	}
	expectedCounts := &counter{all: 2, other: 2}
	assert.Equal(t, expectedCounts, counts)
}

func TestStream_PublicFirehose(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	reqCount := 0
	mux.HandleFunc("/1.1/statuses/firehose.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"count": "100"}, r)
		switch reqCount {
		case 0:
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Transfer-Encoding", "chunked")
			fmt.Fprintf(w,
				`{"text": "Gophercon talks!"}`+"\r\n"+
					`{"text": "Gophercon super talks!"}`+"\r\n",
			)
		default:
			// Only allow first request
			http.Error(w, "Stream API not available!", 130)
		}
		reqCount++
	})

	counts := &counter{}
	demux := newCounterDemux(counts)
	client := NewClient(httpClient)
	streamFirehoseParams := &StreamFirehoseParams{
		Count: 100,
	}
	stream, err := client.Streams.Firehose(streamFirehoseParams)
	// assert that the expected messages are received
	assert.NoError(t, err)
	defer stream.Stop()
	for message := range stream.Messages {
		demux.Handle(message)
	}
	expectedCounts := &counter{all: 2, other: 2}
	assert.Equal(t, expectedCounts, counts)
}

func TestStreamRetry_ExponentialBackoff(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	reqCount := 0
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch reqCount {
		case 0:
			http.Error(w, "Service Unavailable", 503)
		default:
			// Only allow first request
			http.Error(w, "Stream API not available!", 130)
		}
		reqCount++
	})
	stream := &Stream{
		client:   httpClient,
		Messages: make(chan interface{}),
		done:     make(chan struct{}),
		group:    &sync.WaitGroup{},
	}
	stream.group.Add(1)
	req, _ := http.NewRequest("GET", "http://example.com/", nil)
	expBackoff := &BackOffRecorder{}
	// receive messages and throw them away
	go NewSwitchDemux().HandleChan(stream.Messages)
	stream.retry(req, expBackoff, nil)
	defer stream.Stop()
	// assert exponential backoff in response to 503
	assert.Equal(t, 1, expBackoff.Count)
}

func TestStreamRetry_AggressiveBackoff(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	reqCount := 0
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch reqCount {
		case 0:
			http.Error(w, "Enhance Your Calm", 420)
		case 1:
			http.Error(w, "Too Many Requests", 429)
		default:
			// Only allow first request
			http.Error(w, "Stream API not available!", 130)
		}
		reqCount++
	})
	stream := &Stream{
		client:   httpClient,
		Messages: make(chan interface{}),
		done:     make(chan struct{}),
		group:    &sync.WaitGroup{},
	}
	stream.group.Add(1)
	req, _ := http.NewRequest("GET", "http://example.com/", nil)
	aggExpBackoff := &BackOffRecorder{}
	// receive messages and throw them away
	go NewSwitchDemux().HandleChan(stream.Messages)
	stream.retry(req, nil, aggExpBackoff)
	defer stream.Stop()
	// assert aggressive exponential backoff in response to 420 and 429
	assert.Equal(t, 2, aggExpBackoff.Count)
}
