
package orders

import (
    "context"
    "github.com/google/uuid"
    "github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
    Create(context.Context, *Order) error
    GetByID(context.Context, uuid.UUID) (*Order, error)
}

type PGStore struct {
    db *pgxpool.Pool
}

func NewStore(url string) (*PGStore, error) {
    db, err := pgxpool.New(context.Background(), url)
    if err != nil {
        return nil, err
    }
    return &PGStore{db: db}, nil
}

func (s *PGStore) Create(ctx context.Context, o *Order) error {
    _, err := s.db.Exec(ctx, `INSERT INTO orders(id, item, created_at) VALUES($1,$2,$3)`, o.ID, o.Item, o.CreatedAt)
    return err
}

func (s *PGStore) GetByID(ctx context.Context, id uuid.UUID) (*Order, error) {
    row := s.db.QueryRow(ctx, `SELECT id, item, created_at FROM orders WHERE id=$1`, id)
    o := &Order{}
    err := row.Scan(&o.ID, &o.Item, &o.CreatedAt)
    if err != nil {
        return nil, ErrNotFound
    }
    return o, nil
}
