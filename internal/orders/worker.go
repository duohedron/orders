
package orders

import "log"

func StartWorker(ch <-chan Event) {
    for e := range ch {
        log.Println("Processed event:", e)
    }
}
