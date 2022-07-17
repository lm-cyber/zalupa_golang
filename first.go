package main

import (
	"fmt"
	"github.com/tebeka/selenium"
)

func ExampleFindElement() {
	var webDriver selenium.WebDriver
	var err error
	selenium.SetDebug(true)
	caps := selenium.Capabilities{"browserName": "firefox"}
	if webDriver, err = selenium.NewRemote(caps, "http://localhost:8999/wd/hub"); err != nil {
		fmt.Printf("Failed to open session: %s\n", err)
		return
	}
	defer webDriver.Quit()

	err = webDriver.Get("https://www.youtube.com")
	if err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
		return
	}

	if title, err := webDriver.Title(); err == nil {
		fmt.Printf("Page title: %s\n", title)
	} else {
		fmt.Printf("Failed to get page title: %s", err)
		return
	}

	var elem selenium.WebElement
	elem, err = webDriver.FindElement(selenium.ByCSSSelector, ".repo .name")
	if err != nil {
		fmt.Printf("Failed to find element: %s\n", err)
		return
	}

	if text, err := elem.Text(); err == nil {
		fmt.Printf("Repository: %s\n", text)
	} else {
		fmt.Printf("Failed to get text of element: %s\n", err)
		return
	}

	// output:
	// Page title: go-selenium - Sourcegraph
	// Repository: go-selenium
}
