Email Extractor
---

This Go program is designed to extract email addresses from a list of Elastic Search Databases. 

It takes in a list of URLs and will search the Database of each URL to find email addresses.



Install
----

```
go install github.com/RandomRobbieBF/elastic-search-email-extractor@latest
```


Usage
----
To run the program, you can provide a list of URLs as command line arguments:

```
go run email_extractor.go http://111.111.111.111:9200
```

Alternatively, you can provide a text file containing a list of URLs, one per line:

```
go run email_extractor.go url_list.txt
```


The program will ignore any bad SSL certificates, and will only extract email addresses and logs the extracted email address along with it's source URL to a file `output.txt`




Code Overview
----
The searchURLs function is where the bulk of the work is done. It takes in a URL and performs an HTTP GET request to retrieve the HTML. It then uses the regexp package to extract email addresses from the HTML. Finally, it writes the found email addresses along with the URL to output.txt.

The main function is responsible for checking the command line arguments and either reading URLs from a text file or using the URLs provided as command line arguments. It then calls the searchURLs function for each URL provided.

The removeDuplicates function is used to remove duplicate email addresses that might be found. This function uses a map to track unique elements, and then converts the map keys back to a slice.
