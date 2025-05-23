// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package ozone

// schema: tools.ozone.team.listMembers

import (
	"context"

	"github.com/bluesky-social/indigo/xrpc"
)

// TeamListMembers_Output is the output of a tools.ozone.team.listMembers call.
type TeamListMembers_Output struct {
	Cursor  *string            `json:"cursor,omitempty" cborgen:"cursor,omitempty"`
	Members []*TeamDefs_Member `json:"members" cborgen:"members"`
}

// TeamListMembers calls the XRPC method "tools.ozone.team.listMembers".
func TeamListMembers(ctx context.Context, c *xrpc.Client, cursor string, disabled bool, limit int64, q string, roles []string) (*TeamListMembers_Output, error) {
	var out TeamListMembers_Output

	params := map[string]interface{}{
		"cursor":   cursor,
		"disabled": disabled,
		"limit":    limit,
		"q":        q,
		"roles":    roles,
	}
	if err := c.Do(ctx, xrpc.Query, "", "tools.ozone.team.listMembers", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
