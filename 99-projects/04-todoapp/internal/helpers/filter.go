package helpers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type QueryOptions struct {
	Page     int
	PageSize int
	Sort     map[string]string
	Filters  map[string][]interface{}
}

func BuildSQLQuery(tableName string, options *QueryOptions) (string, []interface{}) {
	// Initialize the SQL statement and arguments
	sql := "SELECT * FROM my_table"
	args := make([]interface{}, 0)

	// Add filters to the SQL statement and arguments
	if len(options.Filters) > 0 {
		sql += " WHERE "
		filters := make([]string, 0)
		for key, values := range options.Filters {
			inClause := strings.Trim(strings.Repeat("?, ", len(values)), ", ")
			filters = append(filters, fmt.Sprintf("%s IN (%s)", key, inClause))
			args = append(args, values...)
		}
		sql += strings.Join(filters, " AND ")
	}

	// Add sorting to the SQL statement
	if len(options.Sort) > 0 {
		sql += " ORDER BY "
		sorts := make([]string, 0)
		for key, direction := range options.Sort {
			sorts = append(sorts, fmt.Sprintf("%s %s", key, direction))
		}
		sql += strings.Join(sorts, ", ")
	}

	// Add paging to the SQL statement and arguments
	sql += " LIMIT ? OFFSET ?"
	args = append(args, options.PageSize, (options.Page-1)*options.PageSize)

	// Return the SQL statement and arguments
	return sql, args
}

func RequestToQueryOptions(req *http.Request) QueryOptions {
	options := QueryOptions{}

	// Parse the page parameter from the request URL
	pageParam := req.URL.Query().Get("page")
	if pageParam != "" {
		page, err := strconv.Atoi(pageParam)
		if err == nil {
			options.Page = page
		}
	}

	// Parse the pageSize parameter from the request URL
	pageSizeParam := req.URL.Query().Get("pageSize")
	if pageSizeParam != "" {
		pageSize, err := strconv.Atoi(pageSizeParam)
		if err == nil {
			options.PageSize = pageSize
		}
	}

	// Parse the sort parameter from the request URL
	sortParam := req.URL.Query().Get("sort")
	if sortParam != "" {
		options.Sort = parseSortParam(sortParam)
	}

	// Parse the filter parameters from the request URL
	filterParams := req.URL.Query()["filter"]
	if len(filterParams) > 0 {
		options.Filters = parseFilterParams(filterParams)
	}

	return options
}

func parseSortParam(sortParam string) map[string]string {
	sort := make(map[string]string)
	fields := strings.Split(sortParam, ",")
	for _, field := range fields {
		parts := strings.Split(field, ":")
		if len(parts) == 2 {
			sort[parts[0]] = parts[1]
		}
	}
	return sort
}

func parseFilterParams(filterParams []string) map[string][]interface{} {
	filters := make(map[string][]interface{})
	for _, filterParam := range filterParams {
		parts := strings.Split(filterParam, ":")
		if len(parts) == 2 {
			field := parts[0]
			values := strings.Split(parts[1], ",")
			var interfaces []interface{}
			for _, value := range values {
				interfaces = append(interfaces, value)
			}
			filters[field] = interfaces
		}
	}
	return filters
}
