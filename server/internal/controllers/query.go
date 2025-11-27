package controllers

import (
	"encoding/json"
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
        c.JSON(400, gin.H{"error": "invalid request", "details": err.Error()})
        return
    }

    sqlStr := strings.TrimSpace(req.SQL)
    sqlStr = strings.TrimRight(sqlStr, ";\n \t")

    if sqlStr == "" {
        c.JSON(400, gin.H{"error": "sql is required"})
        return
    }

    if !isReadOnlySQL(sqlStr) {
        c.JSON(400, gin.H{"error": "only SELECT queries allowed"})
        return
    }

    if req.Limit <= 0 {
        req.Limit = 10
    }
    if req.Offset < 0 {
        req.Offset = 0
    }

    db := setting.DB()

    composed := fmt.Sprintf("SELECT * FROM (%s) AS q LIMIT %d OFFSET %d", 
        sqlStr, req.Limit, req.Offset)

    rows, err := db.Raw(composed).Rows()
    if err != nil {
        c.JSON(500, gin.H{"error": "query failed", "details": err.Error()})
        return
    }
    defer rows.Close()

    cols, _ := rows.Columns()
    results := []string{} // each row as JSON string

    for rows.Next() {
        values := make([]interface{}, len(cols))
        ptrs := make([]interface{}, len(cols))

        for i := range values {
            ptrs[i] = &values[i]
        }

        rows.Scan(ptrs...)

        // Build JSON object manually to preserve order
        var b strings.Builder
        b.WriteString("{")

        for i, col := range cols {
            b.WriteString(`"` + col + `":`)

            val := values[i]

            // convert []byte -> string
            if bts, ok := val.([]byte); ok {
                val = string(bts)
            }

            // write json encoded value
            jsonVal, _ := json.Marshal(val)
            b.Write(jsonVal)

            if i < len(cols)-1 {
                b.WriteString(",")
            }
        }

        b.WriteString("}")
        results = append(results, b.String())
    }

    c.Data(200, "application/json", 
        []byte(`{"success":true,"count":` +
            fmt.Sprint(len(results)) +
            `,"data":[` + strings.Join(results, ",") + `]}`))
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
