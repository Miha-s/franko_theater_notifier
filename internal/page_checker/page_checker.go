package pagechecker

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

type PageUpdatedCallback interface {
	OnPageUpdated(url string)
}

type PageChecker struct {
	url      string
	callback PageUpdatedCallback
}

func NewPageChecker(url string) *PageChecker {
	return &PageChecker{url: url}
}

func (p *PageChecker) RegisterPageUpdatedCallback(callback PageUpdatedCallback) {
	p.callback = callback
}

func (p *PageChecker) RunPageChecking() {
	prev, err := p.getDatesSection()
	if err != nil {
		log.Print("Failed to load page")
	}

	for {
		time.Sleep(30 * time.Second)

		updated, err := p.getDatesSection()
		if err != nil {
			log.Print("Failed to load page")

			continue
		}

		if !bytes.Equal(updated, prev) {
			log.Print("Page has been updated!")
			p.callback.OnPageUpdated(p.url)
		}

		prev = updated
	}
}

func (p *PageChecker) getDatesSection() ([]byte, error) {
	resp, err := http.Get(p.url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the HTML
	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// Extract the body section
	bodyContent, err := extractULByClass(doc, "performanceHero__dates-list")
	if err != nil {
		return nil, err
	}

	return bodyContent, nil
}

// extractULByClass traverses the HTML tree and finds a <ul> element with the given class
func extractULByClass(n *html.Node, className string) ([]byte, error) {
	if n.Type == html.ElementNode && n.Data == "ul" {
		for _, attr := range n.Attr {
			if attr.Key == "class" && hasClass(attr.Val, className) {
				// Render the <ul> element to a byte slice
				var buf bytes.Buffer
				if err := html.Render(&buf, n); err != nil {
					return nil, err
				}
				return buf.Bytes(), nil
			}
		}
	}

	// Recursively search for the <ul> element
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if ul, err := extractULByClass(c, className); err == nil && ul != nil {
			return ul, nil
		}
	}

	return nil, fmt.Errorf("no <ul> with class %s found", className)
}

// hasClass checks if the class attribute contains the specific className
func hasClass(classAttr string, className string) bool {
	// Classes are space-separated, so we need to split and check
	classes := strings.Fields(classAttr)
	for _, class := range classes {
		if class == className {
			return true
		}
	}
	return false
}
