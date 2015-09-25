package main

import (
	"errors"
)

func (c *Cli) getSchema(table string, limit int) (interface{}, error) {
	if table == "" {
		return nil, errors.New("Missing '-table' (e.g. user, content).")
	}

	schema, err := c.Client.GetSchema(table, limit)
	if err != nil {
		return nil, err
	}

	return schema, nil
}
