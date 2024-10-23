package repositories

import (
	"context"

	"github.com/Ashmn07/ticket-booking/models"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func (r *TicketRepository) GetMany(ctx context.Context) ([]*models.Ticket, error) {
	return nil, nil
}
func (r *TicketRepository) GetOne(ctx context.Context, ticketId uint) (*models.Ticket, error) {
	return nil, nil
}
func (r *TicketRepository) CreateOne(ctx context.Context, ticket *models.Ticket) (*models.Ticket, error) {
	return nil, nil
}
func (r *TicketRepository) UpdateOne(ctx context.Context, ticketId uint, updateData map[string]interface{}) (*models.Ticket, error) {
	return nil, nil
}
func NewTicketRepository(db *gorm.DB) models.TicketRepository {
	return &TicketRepository{
		db: db,
	}
}
