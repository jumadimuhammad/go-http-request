package model

type ArticleStoreInMemory struct {
	ArticleMap []Article
}

func NewArticleStoreInMemory() *ArticleStoreInMemory {
	return &ArticleStoreInMemory{
		ArticleMap: []Article{
			Article{ID:1, Title: "Membuat website", Body: "Hallo website"},
		},
	}
}

func (store *ArticleStoreInMemory) Save(article *Article) error{
	lastID := len(store.ArticleMap)

	article.ID = lastID + 1

	store.ArticleMap = append(store.ArticleMap, *article)

	return nil
}

func (store *ArticleStoreInMemory) Put(id int, title, body string) error {
	store.ArticleMap[id-1] = Article{ID: id, Title: title, Body: body}

	return nil
}

func (store *ArticleStoreInMemory) Delete(id int) error {
	store.ArticleMap = append(store.ArticleMap[:id-1], store.ArticleMap[id:]...)

	return nil
}