package storage

import (
	"context"
	"encoding/json"
	"gin_realword/models"
	"testing"
	"time"
)

var data = `[
  {
    "slug": "Try-to-transmit-the-HTTP-card-maybe-it-will-override-the-multi-byte-hard-drive!-120863",
    "title": "Try to transmit the HTTP card, maybe it will override the multi-byte hard drive!",
    "description": "Assumenda molestiae laboriosam enim ipsum quaerat enim officia vel quo. Earum odit rem natus totam atque cumque. Sint dolorem facere non.",
    "body": "Sunt excepturi ut dolore fuga.\\nAutem eum maiores aut nihil magnam corporis consectetur sit. Voluptate et quasi optio eos et eveniet culpa et nobis.\\nSint aut sint sequi possimus reiciendis nisi.\\nRerum et omnis et sit doloribus corporis voluptas error.\\nIusto molestiae tenetur necessitatibus dolorem omnis. Libero sed ut architecto.\\nEx itaque et modi aut voluptatem alias quae.\\nModi dolor cupiditate sit.\\nDelectus consectetur nobis aliquid deserunt sint ut et voluptas.\\nCorrupti in labore laborum quod. Ipsa laudantium deserunt. Ut atque harum inventore natus facere sed molestiae.\\nQuia aliquid ut.\\nAnimi sunt rem et sit ullam dolorem ab consequatur modi. Cupiditate officia voluptatum.\\nTenetur facere eum distinctio animi qui laboriosam.\\nQuod sed voluptatem et cumque est eos.\\nSint id provident suscipit harum odio et. Et fuga repellendus magnam dignissimos eius aspernatur rerum. Quo perferendis nesciunt.\\nDolore dolorem porro omnis voluptatibus consequuntur et expedita suscipit et.\\nTempora facere ipsa.\\nDolore accusamus soluta officiis eligendi.\\nEum quaerat neque eum beatae odio. Ad voluptate vel.\\nAut aut dolor. Cupiditate officia voluptatum.\\nTenetur facere eum distinctio animi qui laboriosam.\\nQuod sed voluptatem et cumque est eos.\\nSint id provident suscipit harum odio et.",
    "tagList": [
      "voluptate",
      "rerum",
      "ducimus",
      "hic"
    ],
    "createdAt": "2022-12-09T13:46:24.264Z",
    "updatedAt": "2022-12-09T13:46:24.264Z",
    "favorited": false,
    "favoritesCount": 2202,
    "author": {
      "username": "Anah Benešová",
      "bio": null,
      "image": "https://api.realworld.io/images/demo-avatar.png",
      "following": false
    }
  },
]`

func TestCreateArticle(t *testing.T) {
	ctx := context.TODO()
	err := CreateArticle(ctx, &models.Article{
		AuthorUsername: "xx",
		Title:          "xxx",
		Slug:           "xxx",
		Body:           "111",
		Description:    "111",
		TagList:        []string{"111"},
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDataImport(t *testing.T) {
	ctx := context.TODO()
	var articles []map[string]interface{}
	err := json.Unmarshal([]byte(data), &articles)
	if err != nil {
		t.Fatal(err)
	}
	for _, article := range articles {
		var tagList []string
		for _, tag := range article["tagList"].([]interface{}) {
			tagList = append(tagList, tag.(string))
		}
		createdAt, err := time.Parse(time.RFC3339Nano, article["createdAt"].(string))
		if err != nil {
			t.Logf("parse time failed, err: %v, time: %v", err, article["createdAt"].(string))
			continue
		}
		updatedAt, err := time.Parse(time.RFC3339Nano, article["createdAt"].(string))
		if err != nil {
			t.Logf("parse time failed, err: %v, time: %v", err, article["updatedAt"].(string))
			continue
		}
		err = CreateArticle(ctx, &models.Article{
			AuthorUsername: article["author"].(map[string]interface{})["username"].(string),
			Title:          article["title"].(string),
			Slug:           article["slug"].(string),
			Body:           article["body"].(string),
			Description:    article["description"].(string),
			TagList:        tagList,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
		})
		if err != nil {
			t.Errorf("create article failed, err: %v", err)
			continue
		}
	}
}
