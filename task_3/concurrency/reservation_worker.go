
package concurrency

import (
	"fmt"
	"time"

	"github.com/Ghaby-X/library_manager/services"
)

func ReservationWorker(lm *services.LibraryManager) {
	for reservation := range lm.ReservationChannel {
		go func(res services.Reservation) {
			timer := time.NewTimer(5 * time.Second)
			<-timer.C

			book, prs := lm.Books[res.BookID]
			if !prs {
				return
			}

			book.Mutex.Lock()
			defer book.Mutex.Unlock()

			if book.Status == "Reserved" && book.ReservedBy == res.MemberID {
				book.Status = "Available"
				book.ReservedBy = 0
				fmt.Printf("Reservation for book %d by member %d has been cancelled due to inactivity\n", res.BookID, res.MemberID)
			}
		}(reservation)
	}
}
