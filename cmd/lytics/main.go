package main

import (
	"encoding/json"
	"flag"
	"fmt"
	lytics "github.com/lytics/go-lytics"
	"io/ioutil"
	"os"
	"strings"
)

var (
	apikey        string
	dataapikey    string
	method        string
	id            string
	output        string
	file          string
	fields        string
	fieldsSlice   []string
	segments      string
	segmentsSlice []string
	entitytype    string
	fieldname     string
	fieldvalue    string
	table         string
	limit         int
)

type Cli struct {
	Client *lytics.Client
}

func init() {
	flag.Usage = func() {
		usageExit()
	}

	flag.StringVar(&apikey, "apikey", os.Getenv("LIOKEY"), "Lytics API Key - Or use env LIOKEY")
	flag.StringVar(&dataapikey, "dataapikey", os.Getenv("LIODATAKEY"), "Lytics Data API Key - Or use env LIODATAKEY")
	flag.StringVar(&id, "id", "", "Id of object")
	flag.StringVar(&segments, "segments", "", "Comma Separated Segments")
	flag.StringVar(&fields, "fields", "", "Comma Separated Fields")
	flag.StringVar(&fieldname, "fieldname", "", "Field Name")
	flag.StringVar(&fieldvalue, "fieldvalue", "", "Field Value")
	flag.StringVar(&entitytype, "entitytype", "", "Entity Type")
	flag.StringVar(&table, "table", "", "Schema Table")
	flag.StringVar(&file, "file", "", "Output File Name")
	flag.IntVar(&limit, "limit", -1, "Result Limit")
	flag.Parse()
}

func main() {
	if apikey == "" && dataapikey == "" {
		fmt.Println(`Missing -apikey and/or -method: use -help for assistance

    LIOKEY env variable will fullfill api key needs
    `)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println(`Missing command use -help for assistance.

    export LIOKEY="mykey"

    lytics schema
    `)
		os.Exit(1)
	}
	method = os.Args[1]

	// create lytics client with auth info
	c := Cli{
		Client: lytics.NewLytics(apikey, dataapikey, nil),
	}

	output, err := c.handleFunction(method)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %v", err))
		os.Exit(1)
	}

	if file != "" {
		err := writeToFile(file, output)
		if err != nil {
			fmt.Println("Failed to write data to file %s: %v", file, err)
			return
		}

		fmt.Println("Data written to %s successfully.", file)
		return
	}

	fmt.Println(output)
}

func writeToFile(file, data string) error {
	err := ioutil.WriteFile(file, []byte(data), 0644)
	return err
}

func appendToFile(file, data string) error {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.WriteString(data); err != nil {
		return err
	}

	return nil
}

func (c *Cli) handleFunction(method string) (string, error) {
	var (
		result interface{}
		err    error
	)

	if fields != "" {
		fieldsSlice = strings.Split(fields, ",")
	}

	if segments != "" {
		segmentsSlice = strings.Split(segments, ",")
	}

	switch method {
	case "account":
		result, err = c.getAccounts(id)

	case "auth":
		result, err = c.getAuths(id)

	case "schema":
		result, err = c.getSchema(table)

	case "entity":
		result, err = c.getEntity(entitytype, fieldname, fieldvalue, fieldsSlice)

	case "provider":
		result, err = c.getProviders(id)

	case "segment":
		result, err = c.getSegments("user", segmentsSlice)

	case "segmentsize":
		result, err = c.getSegmentSizes(segmentsSlice)

	case "segmentattribution":
		result, err = c.getSegmentAttributions(segmentsSlice, limit)

	case "work":
		result, err = c.getWorks(id)

	case "user":
		result, err = c.getUsers(id)

	case "query":
		result, err = c.getQueries(id)

	case "watch":
		c.watch()
	default:
		usageExit()
	}

	if err != nil {
		return "", err
	}

	return makeJSON(result), err
}

func validate() bool {
	return true
}

func makeJSON(blob interface{}) string {
	jsonOut, err := json.MarshalIndent(blob, "", "	")
	if err != nil {
		return fmt.Sprintf("Failed: %v", err)
	}

	return string(jsonOut)
}

func exitIfErr(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v: %s\n", err, msg)
		os.Exit(1)
	}
}

func errExit(err error, msg string) {
	fmt.Fprintf(os.Stderr, "%v: %s\n", err, msg)
	os.Exit(1)
}

func usageExit() {
	fmt.Printf(`
--------------------------------------------------------
**************  LYTICS COMMAND LINE HELP  **************
--------------------------------------------------------

GLOBAL PARAMS:
    <apikey>                     REQUIRED       string
    <dataapikey>                 OPTIONAL       string
    <limit>                      OPTIONAL       int

ENV Vars:
    export LIOKEY="your_api_key"
    export LIODATAKEY="your_api_key"

METHODS:
    [account]
         retrieves account information based upon api key.
         if no id is passed, all accounts returned.
         -------
         params:
         -------
              <id>               OPTIONAL       string
         -------
         example:
         -------
              lytics account
              lytics account -id=<id>

    [auth]
         retrieves auth information based upon api key.
         if no id is passed, all auths returned.
         -------
         params:
         -------
              <id>               OPTIONAL       string
         -------
         example:
         -------
              lytics auth
              lytics auth -id=<id>

    [schema]
         retrieves table schema based upon api key.
         -------
         params:
         -------
              <table>            REQUIRED       string
              <limit>            OPTIONAL       int
         -------
         example:
         -------
              lytics schema
              lytics schema -limit=<limit>

    [fieldinfo]
         retrieves detailed field info based upon api key.
         -------
         params:
         -------
              <table>            REQUIRED       string
              <fields>           OPTIONAL       string (comma separated list of fields)
              <limit>            OPTIONAL       int
         -------
         example:
         -------
              lytics fieldinfo -table=user
              lytics fieldinfo -table=user -fields=one,two -limit=2

    [entity]
         retrieves entity information based upon api key.
         -------
         params:
         -------
              <entitytype>       REQUIRED       string (user or content)
              <fieldname>        REQUIRED       string (name of field to search by, e.g. email)
              <fieldvalue>       REQUIRED       string (value of field to search by)
              <fields>           OPTIONAL       string (comma separated list of fields to filter by)
         -------
         example:
         -------
              lytics entity -entitytype=user -fieldname=email -fieldvalue="me@me.com"
              lytics entity -entitytype=user -fieldname=email -fieldvalue="me@me.com" -fields=email

    [provider]
         retrieves provider information based upon api key.
         if no id is passed, all providers returned.
         -------
         params:
         -------
              <id>               OPTIONAL       string
         -------
         example:
         -------
              lytics provider
              lytics provider -id=<id>

    [segment]
         retrieves segment information based upon api key.
         if no id is passed, all segments returned.
         -------
         params:
         -------
              <segments>         OPTIONAL       string (comma separated list of segment ids, max 1)
         -------
         example:
         -------
              lytics segment
              lytics segment -segments=one

    [segmentsize]
         retrieves segment sizes information based upon api key.
         if no id is passed, all segment sizes returned.
         -------
         params:
         -------
              <segments>         OPTIONAL       string (comma separated list of segment ids)
         -------
         example:
         -------
              lytics segmentsize
              lytics segmentsize -segmentes=one,two

    [segmentattribution]
         retrieves segment information based upon api key.
         if no id is passed, all segments returned.
         -------
         params:
         -------
              <segments>         OPTIONAL       string (comma separated list of segment ids)
              <limit>            OPTIONAL       int
         -------
         example:
         -------
              lytics segmentattribution
              lytics segmentattribution -segments=one,two -limit=5

    [user]
         retrieves user information based upon api key.
         if no id is passed, all users returned.
         -------
         params:
         -------
              <id>               OPTIONAL       string
         -------
         example:
         -------
              lytics user
              lytics user -id=<id>

    [work]
         retrieves work information based upon api key.
         if no id is passed, all works returned.
         -------
         params:
         -------
              <id>               OPTIONAL       string
         -------
         example:
         -------
              lytics work
              lytics work -id=<id>

    [query]
         retrieves query information
         -------
         params:
         -------
              <alias>               OPTIONAL       string
         -------
         example:
         -------
              lytics query
              lytics query -id=<alias>

    [watch]
         watch the current folder for .lql, .json files to evaluate
         the .lql query against the data in .json to preview output.

         .lql file name must match the json file name.

         For Example: 
            cd /tmp 
            ls *.lql       # assume a temp.lql
            cat temp.json  # file of data

         -------
         example:
         -------
              lytics watch

`, os.Args[0])
	os.Exit(1)
}
