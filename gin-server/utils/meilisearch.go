package utils

import (
	"fmt"
	"github.com/meilisearch/meilisearch-go"
)

var Meili = &meili{}
type meili struct {
	cli *meilisearch.Client
}

func (ml *meili) Init(host string) {
	ml.cli = meilisearch.NewClient(meilisearch.ClientConfig{
		Host: host,
	})
}
func (ml *meili) Insert(index string, data map[string]interface{}) {
	v := []map[string]interface{}{
		data,
	}
	task, err := ml.cli.Index(index).AddDocuments(v)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println(task.TaskUID)
}
func (ml *meili) InsertBetch(index string, v []map[string]interface{}) {
	task, err := ml.cli.Index(index).UpdateDocuments(v)
	if err != nil {
		return
	}
	fmt.Println(task.TaskUID)
}
func (ml *meili) SearchByIndexAndQuery(index,query string) []interface{}{
	data, err := ml.cli.Index(index).Search(query, &meilisearch.SearchRequest{})
	if err != nil {
		return nil
	}
	if len(data.Hits) >0 {
		return data.Hits
	}
	return nil
}

func (ml *meili) DelDocById(index ,id string)  {
	ts, err := ml.cli.Index(index).DeleteDocument(id)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println(ts.Status)
}

func (ml *meili) DelDocByIndex(index string)  {
	ts, err := ml.cli.Index(index).DeleteAllDocuments()
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println(ts.Status)
}