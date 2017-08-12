package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/ei-grad/hlcup/models"
)

type Loader struct {
	baseURL, fileName string
	wg                sync.WaitGroup
	count             int
}

func (l *Loader) LoadData() {

	// Open a zip archive for reading.
	r, err := zip.OpenReader(l.fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Wait for a server to start
	time.Sleep(1 * time.Second)

	log.Printf("loader: starting")

	t0 := time.Now()

	for _, f := range r.File {
		l.wg.Add(1)
		go l.loadFile(f, 1)
	}

	l.wg.Wait()

	t1 := time.Now()
	log.Printf("loader: stage 1 finished in %s", t1.Sub(t0))

	for _, f := range r.File {
		l.wg.Add(1)
		go l.loadFile(f, 2)
	}

	l.wg.Wait()

	t2 := time.Now()
	log.Printf("loader: stage 2 finished in %s", t2.Sub(t1))
	log.Printf("loader: load finished in %s", t2.Sub(t0))

}

func (l *Loader) loadFile(f *zip.File, stage int) {

	defer l.wg.Done()

	rc, err := f.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer rc.Close()

	decoder := json.NewDecoder(rc)

	// read left_bracket token
	token, err := decoder.Token()
	if err != nil {
		log.Fatalf("Bad start token in %s!", f.Name)
	}
	if _, ok := token.(json.Delim); !ok {
		log.Fatalf("Bad start token in %s!", f.Name)
	}

	for decoder.More() {

		// read key
		token, err = decoder.Token()
		if err != nil {
			log.Fatalf("Bad second token in %s!", f.Name)
		}
		key, ok := token.(string)
		if !ok {
			log.Fatalf("Second token in %s is not string!", f.Name)
		}

		// read left_brace token
		token, err = decoder.Token()
		if err != nil {
			log.Fatalf("Bad start token in %s!", f.Name)
		}
		if _, ok := token.(json.Delim); !ok {
			log.Fatalf("Bad start token in %s!", f.Name)
		}

		type Entity interface {
			MarshalJSON() ([]byte, error)
			UnmarshalJSON([]byte) error
		}

		var v Entity
		var constructor func() Entity

		switch key {
		case strUsers:
			if stage == 1 {
				constructor = func() Entity { return &models.User{} }
			}
		case strLocations:
			if stage == 1 {
				constructor = func() Entity { return &models.Location{} }
			}
		case strVisits:
			if stage == 2 {
				constructor = func() Entity { return &models.Location{} }
			}
		}

		if constructor == nil {
			return
		}

		log.Printf("Loading %s from %s...", key, f.Name)

		for decoder.More() {
			v = constructor()
			err := decoder.Decode(&v)
			if err != nil {
				log.Fatalf("Bad JSON: %s", err)
			}
			body, err := v.MarshalJSON()
			if err != nil {
				log.Fatalf("Can't encode %+v back: %s", v, err)
			}
			l.sendPost(fmt.Sprintf("%s/%s/new", l.baseURL, key), body)
		}
	}

	log.Printf("Loaded %s.", f.Name)
}

func (l *Loader) sendPost(url string, body []byte) {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetMethod(strPost)
	req.SetRequestURI(url)
	req.Header.SetContentType(strApplicationJSON)
	req.SetBody(body)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := fasthttp.Do(req, resp)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode() != 200 {
		log.Fatalf("loader: LOAD FAILED! Got non-200 response:\n%s", resp)
	}

}
