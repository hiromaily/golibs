package semaphore

import (
	"fmt"
	"sync"
	"time"
)

var companies []WellKnownCompany

// WellKnownCompany is company's name and country
type WellKnownCompany struct {
	Name    string
	Country string
}

// Display is to display members of WellKnownCompany
func (w WellKnownCompany) Display() {
	fmt.Printf("Name: %s, Country: %s\n", w.Name, w.Country)
}

func init() {
	companies = []WellKnownCompany{
		{Name: "Google", Country: "USA"},
		{Name: "Amazon", Country: "USA"},
		{Name: "Apple", Country: "USA"},
		{Name: "Facebook", Country: "USA"},
		{Name: "Netflix", Country: "USA"},
		{Name: "Linkedin", Country: "USA"},
		{Name: "Instagram", Country: "USA"},
		{Name: "IBM", Country: "USA"},
		{Name: "Twitter", Country: "USA"},
		{Name: "Airbnb", Country: "USA"},
		{Name: "Alibaba", Country: "China"},
	}
}

// Semaphore is semaphone
func Semaphore(n int) {
	wg := &sync.WaitGroup{}
	chanSemaphore := make(chan bool, n)

	//si.Teachers
	for i := 0; i < 100; i++ {
		wg.Add(1)
		chanSemaphore <- true

		//chanSemaphore <- true
		go func(idx int) {
			defer func() {
				<-chanSemaphore
				wg.Done()
			}()
			//concurrent func
			fmt.Println(idx)
			time.Sleep(1 * time.Second)
		}(i)
	}
	fmt.Println("waiting all go routines ...")
	wg.Wait()
}

// Semaphore2 is semaphone
func Semaphore2(n int) {
	wg := &sync.WaitGroup{}
	chanSemaphore := make(chan bool, n)

	//si.Teachers
	for _, company := range companies {
		wg.Add(1)
		chanSemaphore <- true

		//chanSemaphore <- true
		go func(c WellKnownCompany) {
			defer func() {
				<-chanSemaphore
				wg.Done()
			}()
			//concurrent func
			c.Display()
			time.Sleep(1 * time.Second)
		}(company)
	}
	wg.Wait()
}
