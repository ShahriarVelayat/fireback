package workspaces

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtractQueryDslFromGinContext(c *gin.Context) QueryDSL {
	workspaceId := c.GetHeader("workspace-id")
	internal_sql := c.GetString("internal_sql")
	id := c.Param("uniqueId")
	jsonQuery := c.Query("jsonQuery")
	sort := c.Query("sort")
	linkerId := c.Param("linkerId")
	queryString, _ := c.GetQuery("query")
	withPreloads, _ := c.GetQuery("withPreloads")
	isDeep, _ := c.GetQuery("deep")

	searchPhrase := c.Query("searchPhrase")

	o, _ := c.GetQuery("startIndex")
	startIndex, _ := strconv.Atoi(o)

	l, _ := c.GetQuery("itemsPerPage")
	itemsPerPage, _ := strconv.Atoi(l)

	if startIndex < 0 {
		startIndex = 0
	}

	switch {
	case itemsPerPage > 1000:
		itemsPerPage = 1000
	case itemsPerPage <= 0:
		itemsPerPage = 20
	}

	userHas := c.GetStringSlice("user_has")

	user, isUserSet := c.Get("user_id")
	var userId string

	if isUserSet {
		value, ok := user.(string)
		if ok {
			userId = value
		} else if value2, ok2 := user.(*string); ok2 {
			userId = *value2
		}
	}

	var f QueryDSL = QueryDSL{
		Query:         queryString,
		StartIndex:    startIndex,
		ItemsPerPage:  itemsPerPage,
		InternalQuery: internal_sql,
		UserHas:       userHas,
		Sort:          sort,
		JsonQuery:     jsonQuery,
		SearchPhrase:  searchPhrase,
		LinkerId:      linkerId,
		WorkspaceId:   workspaceId,
		Language:      "en",
		Region:        "us",
		UniqueId:      id,
		Authorization: c.GetHeader("Authorization"),
	}

	f.UserId = userId

	if len(withPreloads) > 0 {
		f.WithPreloads = strings.Split(strings.Trim(withPreloads, " "), ",")
	}

	deep := c.GetHeader("deep")

	if deep == "true" || deep == "yes" || deep == "1" || isDeep == "true" || isDeep == "yes" || isDeep == "1" {
		f.Deep = true
	}

	query := c.GetHeader("query")
	if query != "" && f.Query == "" {
		f.Query = query
	}

	acceptLang := c.GetHeader("accept-language")
	if acceptLang != "" && len(acceptLang) == 2 {
		f.Language = strings.ToLower(acceptLang)
	}

	// The language set in the header has higher priority
	language := c.Query("acceptLanguage")
	if language != "" {
		f.Language = language
	}

	return f
}
