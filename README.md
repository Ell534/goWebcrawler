# goWebcrawler

A web crawler cli application to crawl a specified webpage and return the number of links to each webpage found on the site. A report will be printed in the following format:

```
================================
  REPORT for https://example.com
================================
Found 3 internal links to https://example.com/page1
Found 2 internal links to https://example.com/page2
...
```

assuming you already have go installed, if not go to 
```
https://go.dev/dl/
```
to install go for your OS.


Clone the repo
```
git clone https://github.com/Ell534/goWebcrawler.git
```


Then run
```
go build -o crawler
```
to build the program.


To run the crawler you must provide a url, a maximum concurrency and a maximum number of pages to crawl in the following format:
```
./crawler <URL> <max concurrency> <max pages>
```
Example:
```
./crawler https://example.com 3 10
```
