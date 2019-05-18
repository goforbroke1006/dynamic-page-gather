package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"time"

	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

var (
	targetUrl    = flag.String("target-url", "", "")
	keepOpenTtl  = flag.Uint64("keep-open", 60, "How many time keep page open (in seconds)")
	gatherPeriod = flag.Uint("gather-period", 1000, "Dumping result period (in milliseconds)")
	outputFile   = flag.String("output", "", "")
)

const pageReadyTimeout = 2500

func init() {
	flag.Parse()
}

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var deadLine time.Time
	//deadLine := time.Now().Add(time.Duration(*keepOpenTtl) * time.Second)
	//deadLine = deadLine.Add(time.Millisecond * time.Duration(pageReadyTimeout))

	sleepDur := time.Duration(*gatherPeriod) * time.Millisecond

	err := chromedp.Run(ctx, chromedp.Tasks{
		network.Enable(),
		network.SetExtraHTTPHeaders(network.Headers(map[string]interface{}{
			"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3",
			"Accept-Language": "ru,en-US;q=0.9,en;q=0.8,uk;q=0.7",
			"User-Agent":      "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36",
		})),

		chromedp.Navigate(*targetUrl),
		chromedp.Sleep(time.Millisecond * time.Duration(pageReadyTimeout)),
		chromedp.ActionFunc(func(ctx context.Context) error {
			deadLine = time.Now().Add(time.Duration(*keepOpenTtl) * time.Second)
			return nil
		}),
		chromedp.ActionFunc(func(ctx context.Context) error {
			go func() {
				for {
					startGather := time.Now()
					node, err := dom.GetDocument().Do(ctx)

					if nil == node {
						break
					}

					htmlData, err := dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
					if err = ioutil.WriteFile(*outputFile, []byte(htmlData), 0644); err != nil {
						log.Fatal(err)
					}
					log.Println("Write tmp result to disk")
					endGather := time.Now()

					time.Sleep(sleepDur - endGather.Sub(startGather))
				}
			}()

			for {
				if time.Now().After(deadLine) {
					log.Println("Bye-bye")
					break
				} else {
					time.Sleep(time.Second)
					continue
				}
			}

			return nil
		}),
	})

	if nil != err {
		log.Fatal(err.Error())
	}

}
