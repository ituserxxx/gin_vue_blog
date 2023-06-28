package utils

import "testing"

func TestSearch(t *testing.T) {
	Meili.Init( "http://localhost:7700")
	Meili.SearchByIndexAndQuery("test1","sd")
}
