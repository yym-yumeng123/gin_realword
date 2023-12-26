package storage

import (
	"context"
	"gin_realword/utils"
	"testing"
)

func TestListTags(t *testing.T) {
	ctx := context.TODO()
	res, err := ListPopularTag(ctx)
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Logf("tags: %v\n", utils.JsonMarshal(res))
}
