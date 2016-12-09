## Lytics Command Line tool & Developers Aid

The goal of this tool is to provide CLI access to Lytics api, as well as a developers aid to enable writing and testing LQL as easily as possible.   

Would love any feature requests or ideas that would make this useful to you.

### Installation

Download a binary from https://github.com/lytics/go-lytics/releases and rename to lytics.
Or install from source.

```
# Or if you have go installed
go get github.com/lytics/go-lytics/cmd/lytics

```

### Usage

```

export LIOKEY="your_api_key"
lytics --help


```

**Lytics Watch Usage**

1.  Create NAME.lql (any name) file in a folder.
2.  Create NAME.json (any name, must match lql file name) in folder.
3.  run the `lytics watch` command from the folder with files.
4.  Edit .lql, or .json files, upon change the evaluated result of .lql, json will be output.


### Example

```

# get your api key from web app account settings
export LIOKEY="your_api_key"

cd /tmp

# start watching in background
lytics watch &


# create an lql file
echo '
SELECT 
   user_id,
   name,
   todate(ts),
   match("user.") AS user_attributes,
   map(event, todate(ts))   as event_times   KIND map[string]time  MERGEOP LATEST

FROM data
INTO USER
BY user_id
ALIAS hello
' > hello.lql

# Create a Json data of events to feed into lql query
echo '[
    {"user_id":"dump123","name":"Down With","company":"Trump", "event":"project.create", "ts":"2016-11-09"},
    {"user_id":"another234","name":"No More","company":"Trump", "event":"project.signup","user.city":"Portland","user.state":"Or", "ts":"2016-11-09"}
]' > hello.json



```
