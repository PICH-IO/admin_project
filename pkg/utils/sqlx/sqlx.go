package util_sqlx

import (
	"fmt"
	"strconv"
	"strings"

	pkg_models "github.com/PICH-IO/admin-api/pkg/models"
)

// create sorts function
func BuildSQLSort(sorts []pkg_models.Sort) string {
	if len(sorts) == 0 {
		return " ORDER BY id" // use it as default order
	}
	var orderClauses []string
	for _, sort := range sorts {
		orderClauses = append(orderClauses, fmt.Sprintf("%s %s", sort.Property, sort.Direction))
	}
	return " ORDER BY " + strings.Join(orderClauses, ", ")
}

// create filter function
func BuildSQLFilter(req []pkg_models.Filter) (string, []interface{}) {
	var sqlFilters []string
	var params []interface{}

	for i, filter := range req {
		paramPlaceholder := fmt.Sprintf("$%d", i+1) // Placeholder for parameterized query

		// Convert the filter value to the appropriate type
		switch v := filter.Value.(type) {
		case string:
			if intValue, err := strconv.Atoi(v); err == nil {
				filter.Value = intValue
			} else if boolValue, err := strconv.ParseBool(v); err == nil {
				filter.Value = boolValue
			} else {
				filter.Value = v
			}
		}

		// Handle the converted value
		switch v := filter.Value.(type) {
		case int:
			sqlFilters = append(sqlFilters, fmt.Sprintf("%s = %s", filter.Property, paramPlaceholder))
			params = append(params, v)
		case bool:
			sqlFilters = append(sqlFilters, fmt.Sprintf("%s = %s", filter.Property, paramPlaceholder))
			params = append(params, v)
		case string:
			if strings.Contains(v, "%") {
				// Handle cases with LIKE for wildcard searches
				sqlFilters = append(sqlFilters, fmt.Sprintf("%s LIKE %s", filter.Property, paramPlaceholder))
			} else {
				sqlFilters = append(sqlFilters, fmt.Sprintf("%s = %s", filter.Property, paramPlaceholder))
			}
			params = append(params, v)
		default:
			return "", nil // Handle unsupported types if necessary
		}
	}

	// Join the filters with " AND "
	filterClause := strings.Join(sqlFilters, " AND ")

	return filterClause, params
}
