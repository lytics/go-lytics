package main

import (
	"errors"
)

func (c *Cli) getSchema(table string) (interface{}, error) {
	if table == "" {
		return nil, errors.New("Missing '-table' (e.g. user, content).")
	}

	schema, err := c.Client.GetSchema(table)
	if err != nil {
		return nil, err
	}

	return schema, nil
}
