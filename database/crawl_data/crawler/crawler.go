package crawler

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"example.com/database/todb"
	"github.com/gocolly/colly"
)

type Product struct {
	productCategory    string
	productName        string
	productPrice       float64
	productImageSource string
}

func (p *Product) String() string {
	return fmt.Sprintf("%s,%s,%v", strings.Trim(p.productCategory, " \n"),
		strings.Trim(p.productName, " \n"),
		p.productPrice)
}

var validPriceString = regexp.MustCompile(`(.*)\$(\d+),?(\d*\.\d+)(.*)`)

func Crawl(db *sql.DB) {
	var category string
	var name string
	var price float64
	var imageSource string

	c1 := colly.NewCollector(colly.Async(true))
	c2 := colly.NewCollector(colly.Async(true))
	c1.OnHTML("ul.mega-stack", func(h *colly.HTMLElement) {
		h.ForEach("li>a", func(i int, t *colly.HTMLElement) {
			link := t.Attr("href")
			c2.Visit(t.Request.AbsoluteURL(link))
		})

	})
	c2.OnHTML("h1", func(h *colly.HTMLElement) {
		category = h.Text
	})

	c2.OnHTML("div.productindexinner", func(e *colly.HTMLElement) {
		name = e.ChildText("h3")
		priceText := e.ChildText("div.price div.prod-price")

		var err error
		m := validPriceString.FindStringSubmatch(priceText)
		// log.Println(strings.Join(m, ", "))
		if m != nil {
			price, err = strconv.ParseFloat(m[2]+m[3], 64)
			if err != nil {
				log.Printf("Error during extracting price value, %v\n", err)
			}
		}

		imageSource = e.ChildAttr("div.product-info-inner > a", "href")
		absImageSource := e.Request.AbsoluteURL(imageSource)
		aProduct := Product{
			productCategory:    category,
			productName:        name,
			productPrice:       price,
			productImageSource: absImageSource,
		}

		rows, err := todb.InsertToDB(db, aProduct.productCategory, aProduct.productName, aProduct.productPrice, aProduct.productImageSource)
		if rows != 0 {
			log.Printf("Insert for product %s suceed, %v affected\n", aProduct.String(), rows)
		} else {
			log.Printf("Insert for product %s failed, %v\n", aProduct.String(), err)
		}
	})

	c2.OnRequest(func(r *colly.Request) {
		fmt.Println("C2 visiting...", r.URL.String())
		r.Headers.Add("Accept-Language", "en-US")
	})
	c1.OnRequest(func(r *colly.Request) {
		fmt.Println("C1 visiting...", r.URL.String())
		r.Headers.Add("Accept-Language", "en-US")
	})
	c1.OnError(func(r *colly.Response, e error) {
		log.Println("Error ", e)
	})
	c2.OnError(func(r *colly.Response, e error) {
		log.Println("Error ", e)
	})

	c1.Visit("https://interiorsonline.com.au/")
	c1.Wait()
	c2.Wait()

}
