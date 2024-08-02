package main

import "sync"

type Image struct {
	File string
	FileName string
}

type ConcurrentImageMultiMap struct {
	images map[string][]Image
	imagesLock sync.Mutex 
}

func NewConcurrentImageMultiMap() *ConcurrentImageMultiMap {
	return &ConcurrentImageMultiMap{
		images: make(map[string][]Image),
	}
}

func (m *ConcurrentImageMultiMap) Put(team string, image Image) {
	m.imagesLock.Lock()
	defer m.imagesLock.Unlock()

	m.images[team] = append(m.images[team], image)
}

func (m *ConcurrentImageMultiMap) Get (team string)[]Image {
	m.imagesLock.Lock()
	defer m.imagesLock.Unlock()
	
	return m.images[team]
}
