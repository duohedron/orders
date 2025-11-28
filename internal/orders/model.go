
package orders

import (
    "time"
    "github.com/google/uuid"
)

type Order struct {
    ID        uuid.UUID `json:"id"`
    Item      string    `json:"item"`
    CreatedAt time.Time `json:"created_at"`
}
