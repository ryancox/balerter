package elastic

import (
	"context"

	ec "github.com/olivere/elastic/v7"
	lua "github.com/yuin/gopher-lua"
	"go.uber.org/zap"
)

func (m *Elastic) query(luaState *lua.LState) int {
	index := luaState.Get(1).String()
	queryString := luaState.Get(2).String()

	m.logger.Debug("call elastic query string", zap.String("index", index), zap.String("queryStrin", queryString))

	q := ec.NewQueryStringQuery(queryString)

	ecResults, err := m.client.Search().
		Index(index).
		Query(q).
		Do(context.TODO())

	if err != nil {
		m.logger.Error("error elastic query string", zap.String("index", index), zap.String("queryStrin", queryString))
		luaState.Push(lua.LNil)
		luaState.Push(lua.LString(err.Error()))
		return 2
	}

	results := &lua.LTable{}
	for _, hit := range ecResults.Hits.Hits {
		row := &lua.LTable{}
		for name := range hit.Fields {
			s, _ := hit.Fields[name].(string) // TODO: may revisit making all return values strings
			row.RawSet(lua.LString(name), lua.LString(s))
		}
		results.Append(row)
	}
	luaState.Push(results)
	luaState.Push(lua.LNil)
	return 2
}
