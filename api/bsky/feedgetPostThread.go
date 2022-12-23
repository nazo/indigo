package schemagen

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/whyrusleeping/gosky/lex/util"
	"github.com/whyrusleeping/gosky/xrpc"
)

// schema: app.bsky.feed.getPostThread

func init() {
}

type FeedGetPostThread_NotFoundPost struct {
	NotFound bool   `json:"notFound" cborgen:"notFound"`
	Uri      string `json:"uri" cborgen:"uri"`
}

type FeedGetPostThread_Output struct {
	Thread *FeedGetPostThread_Output_Thread `json:"thread" cborgen:"thread"`
}

type FeedGetPostThread_Output_Thread struct {
	FeedGetPostThread_ThreadViewPost *FeedGetPostThread_ThreadViewPost
	FeedGetPostThread_NotFoundPost   *FeedGetPostThread_NotFoundPost
}

func (t *FeedGetPostThread_Output_Thread) MarshalJSON() ([]byte, error) {
	if t.FeedGetPostThread_ThreadViewPost != nil {
		return json.Marshal(t.FeedGetPostThread_ThreadViewPost)
	}
	if t.FeedGetPostThread_NotFoundPost != nil {
		return json.Marshal(t.FeedGetPostThread_NotFoundPost)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *FeedGetPostThread_Output_Thread) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.feed.getPostThread#threadViewPost":
		t.FeedGetPostThread_ThreadViewPost = new(FeedGetPostThread_ThreadViewPost)
		return json.Unmarshal(b, t.FeedGetPostThread_ThreadViewPost)
	case "app.bsky.feed.getPostThread#notFoundPost":
		t.FeedGetPostThread_NotFoundPost = new(FeedGetPostThread_NotFoundPost)
		return json.Unmarshal(b, t.FeedGetPostThread_NotFoundPost)

	default:
		return nil
	}
}

type FeedGetPostThread_ThreadViewPost struct {
	Parent  *FeedGetPostThread_ThreadViewPost_Parent         `json:"parent" cborgen:"parent"`
	Post    *FeedPost_View                                   `json:"post" cborgen:"post"`
	Replies []*FeedGetPostThread_ThreadViewPost_Replies_Elem `json:"replies" cborgen:"replies"`
}

type FeedGetPostThread_ThreadViewPost_Parent struct {
	FeedGetPostThread_ThreadViewPost *FeedGetPostThread_ThreadViewPost
	FeedGetPostThread_NotFoundPost   *FeedGetPostThread_NotFoundPost
}

func (t *FeedGetPostThread_ThreadViewPost_Parent) MarshalJSON() ([]byte, error) {
	if t.FeedGetPostThread_ThreadViewPost != nil {
		return json.Marshal(t.FeedGetPostThread_ThreadViewPost)
	}
	if t.FeedGetPostThread_NotFoundPost != nil {
		return json.Marshal(t.FeedGetPostThread_NotFoundPost)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *FeedGetPostThread_ThreadViewPost_Parent) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.feed.getPostThread#threadViewPost":
		t.FeedGetPostThread_ThreadViewPost = new(FeedGetPostThread_ThreadViewPost)
		return json.Unmarshal(b, t.FeedGetPostThread_ThreadViewPost)
	case "app.bsky.feed.getPostThread#notFoundPost":
		t.FeedGetPostThread_NotFoundPost = new(FeedGetPostThread_NotFoundPost)
		return json.Unmarshal(b, t.FeedGetPostThread_NotFoundPost)

	default:
		return nil
	}
}

type FeedGetPostThread_ThreadViewPost_Replies_Elem struct {
	FeedGetPostThread_ThreadViewPost *FeedGetPostThread_ThreadViewPost
	FeedGetPostThread_NotFoundPost   *FeedGetPostThread_NotFoundPost
}

func (t *FeedGetPostThread_ThreadViewPost_Replies_Elem) MarshalJSON() ([]byte, error) {
	if t.FeedGetPostThread_ThreadViewPost != nil {
		return json.Marshal(t.FeedGetPostThread_ThreadViewPost)
	}
	if t.FeedGetPostThread_NotFoundPost != nil {
		return json.Marshal(t.FeedGetPostThread_NotFoundPost)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *FeedGetPostThread_ThreadViewPost_Replies_Elem) UnmarshalJSON(b []byte) error {
	typ, err := util.TypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.feed.getPostThread#threadViewPost":
		t.FeedGetPostThread_ThreadViewPost = new(FeedGetPostThread_ThreadViewPost)
		return json.Unmarshal(b, t.FeedGetPostThread_ThreadViewPost)
	case "app.bsky.feed.getPostThread#notFoundPost":
		t.FeedGetPostThread_NotFoundPost = new(FeedGetPostThread_NotFoundPost)
		return json.Unmarshal(b, t.FeedGetPostThread_NotFoundPost)

	default:
		return nil
	}
}

func FeedGetPostThread(ctx context.Context, c *xrpc.Client, depth int64, uri string) (*FeedGetPostThread_Output, error) {
	var out FeedGetPostThread_Output

	params := map[string]interface{}{
		"depth": depth,
		"uri":   uri,
	}
	if err := c.Do(ctx, xrpc.Query, "", "app.bsky.feed.getPostThread", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
