package controllers

import (
	"fmt"
	"strings"

	"stockin/internal/setting"

	"github.com/gin-gonic/gin"
)

type QueryRequest struct {
	SQL    string `json:"sql"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}

// SelectFromDatabase runs a (read-only) SQL SELECT supplied in the request and returns
// the rows as an array of JSON objects. We enforce that only SELECT queries are allowed
// and wrap the query as a sub-query so we can safely apply LIMIT/OFFSET for pagination.
func SelectFromDatabase(c *gin.Context) {
	var req QueryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request payload", "details": err.Error()})
		return
	}

	sqlStr := strings.TrimSpace(req.SQL)
	sqlStr = strings.TrimRight(sqlStr, ";\n \t")
	if sqlStr == "" {
		c.JSON(400, gin.H{"error": "sql is required"})
		return
	}

	// Only allow read-only SELECT queries
	if !isReadOnlySQL(sqlStr) {
		c.JSON(400, gin.H{
			"error": "only read-only SELECT queries are allowed",
		})
		return
	}

	if req.Limit <= 0 {
		req.Limit = 10
	}
	if req.Offset < 0 {
		req.Offset = 0
	}

	db := setting.DB()

	composed := fmt.Sprintf("SELECT * FROM (%s) AS q LIMIT %d OFFSET %d", sqlStr, req.Limit, req.Offset)
	rows, err := db.Raw(composed).Rows()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to execute query", "details": err.Error()})
		return
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get columns", "details": err.Error()})
		return
	}

	results := []map[string]interface{}{}

	for rows.Next() {
		values := make([]interface{}, len(cols))
		valuePtrs := make([]interface{}, len(cols))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			c.JSON(500, gin.H{"error": "failed to scan row", "details": err.Error()})
			return
		}

		rowMap := make(map[string]interface{})
		for i, col := range cols {
			var v interface{}
			rawVal := values[i]
			// convert []byte -> string for readability
			if b, ok := rawVal.([]byte); ok {
				v = string(b)
			} else {
				v = rawVal
			}
			rowMap[col] = v
		}

		results = append(results, rowMap)
	}

	c.JSON(200, gin.H{"success": true, "count": len(results), "data": results})
}

func isReadOnlySQL(sqlStr string) bool {
	s := strings.ToLower(sqlStr)

	// forbid dangerous keywords anywhere
	forbidden := []string{
		"insert ", "update ", "delete ", "drop ", "alter ",
		"truncate ", "create ", "replace ", "grant ", "revoke ",
		"commit", "rollback",
	}

	for _, bad := range forbidden {
		if strings.Contains(s, bad) {
			return false
		}
	}

	// must contain at least one SELECT
	return strings.Contains(s, "select ")
}
