# CSV to markdown table
The tool csvtomarkdown is CLI tool.

## Install
You can install it simply with go command.
```
go get github.com/josjos7/csvtomarkdown
```

## Usage
After install, you execute this tool by calling 'csvtomarkdown'

```
  -file string
        file path
  -out string
        output file name
```

## Example
- Read csv and write markdown to stdin
```
$ ./csvtomarkdown -file temp/sample.csv
```

- Read csv and write markdown to specified file
```
$ ./csvtomarkdown -file temp/sample.csv -out out.md
```

## TODO
- markdown to csv
- tsv
- cli option