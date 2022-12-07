package schemagen

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/whyrusleeping/gosky/lex/util"
	"github.com/whyrusleeping/gosky/xrpc"
)

// schema: app.bsky.feed.getPostThread

type FeedGetPostThread_NotFoundPost struct {
	Uri      string `json:"uri"`
	NotFound bool   `json:"notFound"`
}

func (t *FeedGetPostThread_NotFoundPost) MarshalJSON() ([]byte, error) {
	t.NotFound = true
	out := make(map[string]interface{})
	out["notFound"] = t.NotFound
	out["uri"] = t.Uri
	return json.Marshal(out)
}

type FeedGetPostThread_MyState struct {
	Repost   string `json:"repost"`
	Upvote   string `json:"upvote"`
	Downvote string `json:"downvote"`
}

func (t *FeedGetPostThread_MyState) MarshalJSON() ([]byte, error) {
	out := make(map[string]interface{})
	out["downvote"] = t.Downvote
	out["repost"] = t.Repost
	out["upvote"] = t.Upvote
	return json.Marshal(out)
}

type FeedGetPostThread_Output struct {
	Thread *FeedGetPostThread_Output_Thread `json:"thread"`
}

func (t *FeedGetPostThread_Output) MarshalJSON() ([]byte, error) {
	out := make(map[string]interface{})
	out["thread"] = t.Thread
	return json.Marshal(out)
}

type FeedGetPostThread_Output_Thread struct {
	FeedGetPostThread_Post         *FeedGetPostThread_Post
	FeedGetPostThread_NotFoundPost *FeedGetPostThread_NotFoundPost
}

func (t *FeedGetPostThread_Output_Thread) MarshalJSON() ([]byte, error) {
	if t.FeedGetPostThread_Post != nil {
		return json.Marshal(t.FeedGetPostThread_Post)
	}
	if t.FeedGetPostThread_NotFoundPost != nil {
		return json.Marshal(t.FeedGetPostThread_NotFoundPost)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *FeedGetPostThread_Output_Thread) UnmarshalJSON(b []byte) error {
	typ, err := util.EnumTypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.feed.getPostThread#post":
		t.FeedGetPostThread_Post = new(FeedGetPostThread_Post)
		return json.Unmarshal(b, t.FeedGetPostThread_Post)
	case "app.bsky.feed.getPostThread#notFoundPost":
		t.FeedGetPostThread_NotFoundPost = new(FeedGetPostThread_NotFoundPost)
		return json.Unmarshal(b, t.FeedGetPostThread_NotFoundPost)

	default:
		return nil
	}
}

type FeedGetPostThread_Post struct {
	ReplyCount    int64                                  `json:"replyCount"`
	MyState       *FeedGetPostThread_MyState             `json:"myState"`
	Uri           string                                 `json:"uri"`
	Parent        *FeedGetPostThread_Post_Parent         `json:"parent"`
	Record        any                                    `json:"record"`
	Embed         *FeedEmbed                             `json:"embed"`
	Replies       []*FeedGetPostThread_Post_Replies_Elem `json:"replies"`
	RepostCount   int64                                  `json:"repostCount"`
	UpvoteCount   int64                                  `json:"upvoteCount"`
	DownvoteCount int64                                  `json:"downvoteCount"`
	Cid           string                                 `json:"cid"`
	Author        *ActorRef_WithInfo                     `json:"author"`
	IndexedAt     string                                 `json:"indexedAt"`
}

func (t *FeedGetPostThread_Post) MarshalJSON() ([]byte, error) {
	out := make(map[string]interface{})
	out["author"] = t.Author
	out["cid"] = t.Cid
	out["downvoteCount"] = t.DownvoteCount
	out["embed"] = t.Embed
	out["indexedAt"] = t.IndexedAt
	out["myState"] = t.MyState
	out["parent"] = t.Parent
	out["record"] = t.Record
	out["replies"] = t.Replies
	out["replyCount"] = t.ReplyCount
	out["repostCount"] = t.RepostCount
	out["upvoteCount"] = t.UpvoteCount
	out["uri"] = t.Uri
	return json.Marshal(out)
}

type FeedGetPostThread_Post_Parent struct {
	FeedGetPostThread_Post         *FeedGetPostThread_Post
	FeedGetPostThread_NotFoundPost *FeedGetPostThread_NotFoundPost
}

func (t *FeedGetPostThread_Post_Parent) MarshalJSON() ([]byte, error) {
	if t.FeedGetPostThread_Post != nil {
		return json.Marshal(t.FeedGetPostThread_Post)
	}
	if t.FeedGetPostThread_NotFoundPost != nil {
		return json.Marshal(t.FeedGetPostThread_NotFoundPost)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *FeedGetPostThread_Post_Parent) UnmarshalJSON(b []byte) error {
	typ, err := util.EnumTypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.feed.getPostThread#post":
		t.FeedGetPostThread_Post = new(FeedGetPostThread_Post)
		return json.Unmarshal(b, t.FeedGetPostThread_Post)
	case "app.bsky.feed.getPostThread#notFoundPost":
		t.FeedGetPostThread_NotFoundPost = new(FeedGetPostThread_NotFoundPost)
		return json.Unmarshal(b, t.FeedGetPostThread_NotFoundPost)

	default:
		return nil
	}
}

type FeedGetPostThread_Post_Replies_Elem struct {
	FeedGetPostThread_Post         *FeedGetPostThread_Post
	FeedGetPostThread_NotFoundPost *FeedGetPostThread_NotFoundPost
}

func (t *FeedGetPostThread_Post_Replies_Elem) MarshalJSON() ([]byte, error) {
	if t.FeedGetPostThread_Post != nil {
		return json.Marshal(t.FeedGetPostThread_Post)
	}
	if t.FeedGetPostThread_NotFoundPost != nil {
		return json.Marshal(t.FeedGetPostThread_NotFoundPost)
	}
	return nil, fmt.Errorf("cannot marshal empty enum")
}
func (t *FeedGetPostThread_Post_Replies_Elem) UnmarshalJSON(b []byte) error {
	typ, err := util.EnumTypeExtract(b)
	if err != nil {
		return err
	}

	switch typ {
	case "app.bsky.feed.getPostThread#post":
		t.FeedGetPostThread_Post = new(FeedGetPostThread_Post)
		return json.Unmarshal(b, t.FeedGetPostThread_Post)
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