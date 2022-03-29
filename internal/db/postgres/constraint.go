package postgres

import (
	"context"
	"errors"
)

const (
	SQL = `SELECT kcu1.constraint_name
		   FROM information_schema.key_column_usage as kcu1
           WHERE kcu1.table_name = $1 AND kcu1.constraint_name IN (
				  SELECT DISTINCT kcu.constraint_name
				  FROM information_schema.key_column_usage as kcu
				  WHERE kcu.table_name = $1 AND kcu.column_name = ANY ($2)
			)
		   GROUP BY kcu1.constraint_name
		   HAVING COUNT(kcu1.constraint_name) = $3`
)

func (c *connection) GetConstraintName(ctx context.Context, tableName string, columnNames []string) (constraintName *string, err error) {
	if len(columnNames) < 1 {
		return nil, errors.New("missing column name's")
	}

	conn, err := c.GetMasterConn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	constraintName = nil
	err = conn.QueryRow(ctx, SQL, tableName, columnNames, len(columnNames)).Scan(&constraintName)

	return constraintName, err
}
