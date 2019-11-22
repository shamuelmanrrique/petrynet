package centralsim

import (
	"log"
	"testing"
)

func TestEvent(t *testing.T) {
	//t.Skip("skipping test evento.")
	log.Printf("Test Insertcion lista eventos")
	le := EventList{}
	le.Insert(Event{1, 1, 1})
	le.PrintEvent()
	le.Insert(Event{1, 1, 1})
	le.PrintEvent()
	le.Insert(Event{0, 3, 3})
	le.PrintEvent()
}
