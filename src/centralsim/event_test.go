package centralsim

import (
	"log"
	"testing"
)

func TestEvent(t *testing.T) {
	//t.Skip("skipping test evento.")
	log.Printf("Test insercion lista eventos")
	le := EventList{}
	le.Inser(Event{1, 1, 1})
	le.PrintEvent()
	le.Inser(Event{1, 1, 1})
	le.PrintEvent()
	le.Inser(Event{0, 3, 3})
	le.PrintEvent()
}
