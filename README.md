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
Usage of csvtomarkdown:
  -f string
        alias of -file
  -file string
        CSV file path
  -header
        Use first line as headers
  -o string
        ailias of -out
  -out string
        Output file name
```

## Example command
- Read csv and write markdown to stdin
```
$ csvtomarkdown -file temp/sample.csv
```

- Read csv and write markdown to specified file
```
$ csvtomarkdown -file temp/sample.csv -out out.md
```

## Example input
```
First,Last,"Comma, ""double quate"""
Taro,Test,"Taro,Test"
```

## Example output
$ csvtomarkdown -f sample.csv
```
|First|Last|Comma, "double quate"|
|Taro|Test|Taro,Test|
```

$ csvtomarkdown -f sample.csv -header
```
|First|Last|Comma, "double quate"|
|:----|:----|:----|
|Taro|Test|Taro,Test|
```
