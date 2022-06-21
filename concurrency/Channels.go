package concurrency

import (
	"log"
	"strconv"
	"time"
)

type WorkerThing struct {
	WorkUnitCh chan string
}

func (w *WorkerThing) StartSystem() {
	log.Println("StartSystem - starting")
	w.WorkUnitCh = make(chan string)
	w.GenerateWorkUnits(100, 1)
	for i := 0; i < 100; i++ {
		w.StartWorker(i)
	}
	<-w.WorkUnitCh
	log.Println("StartSystem - closing")
}

func (w *WorkerThing) StartWorker(workerId int) {
	go func() {
		log.Printf("StartWorker - starting - %d", workerId)
		for workUnit := range w.WorkUnitCh {
			log.Printf("StartWorker - begin processing - %s on workerId %d", workUnit, workerId)
			time.Sleep(5 * time.Second)
			log.Printf("StartWorker - finished processing - %s on workerId %d", workUnit, workerId)
		}
		log.Printf("StartWorker - closing - %d", workerId)
	}()
}

func (w *WorkerThing) GenerateWorkUnits(amount int, delayBetween int) {
	go func() {
		for i := 0; i < amount; i++ {
			w.WorkUnitCh <- "wu_" + strconv.Itoa(i)
			time.Sleep(time.Duration(delayBetween) * time.Second)
		}
		//close(w.WorkUnitCh)
	}()

}
