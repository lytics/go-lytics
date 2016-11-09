## Command Line Tool


### Installation

```
go install github.com/lytics/go-lytics/cmd/lytics
```

### Usage

```

export LIOKEY="your_api_key"
lytics -help



# Watch A folder of lql files and .json files
# will output any lql syntax errors upon edit
# as well as the evaluated output of .lql + .json files


lytics watch
```

**Lytics Watch Usage**

1.  Create NAME.lql (any name) file in a folder.
2.  Create NAME.lql (any name, must match lql file name) in folder.
3.  run the `lytics watch` command from the folder with files.
4.  Edit .lql, or .json files, upon change the evaluated result of .lql, json will be output.

