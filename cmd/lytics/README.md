## Command Line Tool


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
   match("user.") AS user_attributes

FROM data
INTO USER
BY user_id
ALIAS hello
' > hello.lql

# Create a Json data of events to feed into lql query
echo '[
    {"user_id":"dump123","name":"Down With","company":"Trump", "event":"project.create"},
    {"user_id":"another234","name":"No More","company":"Trump", "event":"project.signup","user.city":"Portland","user.state":"Or"}
]' > hello.json



```