package repository

import (
	"context"
	"errors"
	"fmt"
	"reload/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepo struct {
	pool *pgxpool.Pool
}

func NewPostgresRepo(dsn string) (*PostgresRepo, error) {

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("Cannot connect to DB: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("Cannot ping DB: %w", err)
	}

	return &PostgresRepo{
		pool: pool,
	}, nil
}

func scanEvents(rows pgx.Rows) ([]models.Event, error) {

	var userEvents []models.Event
	for rows.Next() {
		var userEvent models.Event
		if err := rows.Scan(&userEvent.EventID, &userEvent.UserID, &userEvent.Activity, &userEvent.ProductID, &userEvent.HappenedAt); err != nil {
			return []models.Event{}, errors.New("Failed to scan row from database " + err.Error())
		}
		userEvents = append(userEvents, userEvent)
	}
	if err := rows.Err(); err != nil {
		return nil, errors.New("Failed to scan row from database " + err.Error())
	}
	return userEvents, nil
}

func (p *PostgresRepo) SaveEvent(event models.EventDTO) error {
	query := `INSERT INTO events (user_id, activity, product_id, happened_at) 
	Values ($1,$2,$3,$4)`

	if _, err := p.pool.Exec(context.Background(), query, event.UserID, event.Activity, event.ProductID, event.HappenedAt); err != nil {
		return errors.New("Failed to save data to database " + err.Error())
	}

	return nil
}

func (p *PostgresRepo) GetEventsByUserId(id int) ([]models.Event, error) {

	query := `SELECT id, user_id, activity, product_id, happened_at from events
	Where user_id = $1 ORDER BY happened_at DESC`

	rows, err := p.pool.Query(context.Background(), query, id)
	if err != nil {
		return []models.Event{}, errors.New("Failed to get data from database " + err.Error())
	}
	defer rows.Close()

	userEvents, err := scanEvents(rows)
	if err != nil {
		return []models.Event{}, err
	}

	return userEvents, nil
}
func (p *PostgresRepo) GetAllEvents() ([]models.Event, error) {
	query := `SELECT id, user_id, activity, product_id, happened_at from events 
	ORDER BY happened_at DESC`

	rows, err := p.pool.Query(context.Background(), query)
	if err != nil {
		return []models.Event{}, errors.New("Failed to get data from database " + err.Error())
	}
	defer rows.Close()

	userEvents, err := scanEvents(rows)
	if err != nil {
		return []models.Event{}, err
	}

	return userEvents, nil
}
