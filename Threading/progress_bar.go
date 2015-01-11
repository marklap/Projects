package main

import (
    "encoding/json"
    "flag"
    "fmt"
    // "io"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "path"
    "runtime"
    // "strings"
)

const (
    google_search_url   string = "http://api.duckduckgo.com/?q=%s&format=json"
    default_max_links   int    = 10
    default_search_term string = "monkey"
)

var logger *log.Logger

func init() {
    // can use an io.MultiWriter() here to write to both log file and stdout
    logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}

var search_term string
var max_links int

func init() {
    usage := `
Usage: %s [--term TERM] [--links LINKS]

  --term TERM     the search term to use
                  [default: "%s"]

  --links LINKS   the maximum number of links to fetch within the page
                  [default: %d]

`

    flag.StringVar(&search_term, "term", default_search_term, "Term to search for")
    flag.IntVar(&max_links, "links", default_max_links, "How many links to fetch from the search results")
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, usage, path.Base(os.Args[0]), default_search_term, default_max_links)
    }

    flag.Parse()

    runtime.GOMAXPROCS(runtime.NumCPU())
}

type searchResults struct {
    DefinitionSource string
    Heading          string
    ImageWidth       int
    RelatedTopics    []interface{}
    DefinitionURL    string
    AbstractURL      string
}

func google_search(url string) (links []string, err error) {
    resp, err := http.Get(url)
    if err != nil {
        log.Println("Doh! Couldn't complete the HTTP GET to the search engine")
        return nil, err
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Println("Doh! Couldn't read the body of the response from the search results")
        return nil, err
    }

    var json_body searchResults
    err = json.Unmarshal(body, &json_body)
    if err != nil {
        log.Println("Failed to unmarshal response")
        return nil, err
    }

    var urls []string
    for _, related_topics := range json_body.RelatedTopics {
        rts := related_topics.(map[string]interface{})
        if topics, ok := rts["Topics"]; ok {
            ts := topics.([]interface{})
            for _, topic := range ts {
                t := topic.(map[string]interface{})
                if url, ok := t["FirstURL"]; ok {
                    urls = append(urls, url.(string))
                }
            }
        }
    }

    return urls, nil
}

func fetch_serial(links []string) {
    for _, link := range links {
        log.Println("Fetching: %s", link)
        _, err := http.Get(link)
        if err != nil {
            log.Println("Error fetching:", link)
        }
        log.Println("[FINISH]", link)
    }
}

func fetch_parallel(workChan chan string, links []string) {

    for i := 0; i < runtime.NumCPU(); i++ {
        go fetch_link(workChan)
    }

    for _, link := range links {
        log.Println("Adding", link, "to channel")
        workChan <- link
    }

}

func fetch_link(links chan string) {

    chan_closed := false
    for {
        if chan_closed {
            return
        }
        select {
        case link, chan_ok := <-links:
            if !chan_ok {
                log.Println("Channel closed")
                chan_closed = true
            } else {
                log.Println("Fetching:", link)
                _, err := http.Get(link)
                if err != nil {
                    log.Println("Error fetching:", link)
                }
                log.Println("[FINISH]", link)
            }
        }
    }
}

func main() {
    log.Println("Search term is:", search_term)
    log.Println("Maximum number of links to fetch:", max_links)

    links, err := google_search(fmt.Sprintf(google_search_url, search_term))
    if err != nil {
        log.Panicln("Something bad happened with the search")
    }

    if links == nil {
        log.Panicln("We didn't get back a string array... doh!")
    }

    workChan := make(chan string)
    // fetch_serial(links)
    fetch_parallel(workChan, links)
}
