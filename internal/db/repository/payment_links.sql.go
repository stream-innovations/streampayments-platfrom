// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: payment_links.sql

package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

const createPaymentLink = `-- name: CreatePaymentLink :one
INSERT INTO payment_links (
  uuid,
  slug,
  created_at,
  updated_at,
  merchant_id,
  name,
  description,
  price,
  decimals,
  currency,
  success_action,
  redirect_url,
  success_message,
  is_test
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
RETURNING id, uuid, slug, created_at, updated_at, merchant_id, name, description, price, decimals, currency, success_action, redirect_url, success_message, is_test
`

type CreatePaymentLinkParams struct {
	Uuid           uuid.UUID
	Slug           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	MerchantID     int64
	Name           string
	Description    string
	Price          pgtype.Numeric
	Decimals       int32
	Currency       string
	SuccessAction  string
	RedirectUrl    sql.NullString
	SuccessMessage sql.NullString
	IsTest         bool
}

func (q *Queries) CreatePaymentLink(ctx context.Context, arg CreatePaymentLinkParams) (PaymentLink, error) {
	row := q.db.QueryRow(ctx, createPaymentLink,
		arg.Uuid,
		arg.Slug,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.MerchantID,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.Decimals,
		arg.Currency,
		arg.SuccessAction,
		arg.RedirectUrl,
		arg.SuccessMessage,
		arg.IsTest,
	)
	var i PaymentLink
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Slug,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.MerchantID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Decimals,
		&i.Currency,
		&i.SuccessAction,
		&i.RedirectUrl,
		&i.SuccessMessage,
		&i.IsTest,
	)
	return i, err
}

const deletePaymentLinkByPublicID = `-- name: DeletePaymentLinkByPublicID :exec
delete from payment_links where merchant_id = $1 and uuid = $2
`

type DeletePaymentLinkByPublicIDParams struct {
	MerchantID int64
	Uuid       uuid.UUID
}

func (q *Queries) DeletePaymentLinkByPublicID(ctx context.Context, arg DeletePaymentLinkByPublicIDParams) error {
	_, err := q.db.Exec(ctx, deletePaymentLinkByPublicID, arg.MerchantID, arg.Uuid)
	return err
}

const getPaymentLinkByID = `-- name: GetPaymentLinkByID :one
select id, uuid, slug, created_at, updated_at, merchant_id, name, description, price, decimals, currency, success_action, redirect_url, success_message, is_test from payment_links where merchant_id = $1 and id = $2 limit 1
`

type GetPaymentLinkByIDParams struct {
	MerchantID int64
	ID         int64
}

func (q *Queries) GetPaymentLinkByID(ctx context.Context, arg GetPaymentLinkByIDParams) (PaymentLink, error) {
	row := q.db.QueryRow(ctx, getPaymentLinkByID, arg.MerchantID, arg.ID)
	var i PaymentLink
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Slug,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.MerchantID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Decimals,
		&i.Currency,
		&i.SuccessAction,
		&i.RedirectUrl,
		&i.SuccessMessage,
		&i.IsTest,
	)
	return i, err
}

const getPaymentLinkByPublicID = `-- name: GetPaymentLinkByPublicID :one
select id, uuid, slug, created_at, updated_at, merchant_id, name, description, price, decimals, currency, success_action, redirect_url, success_message, is_test from payment_links where merchant_id = $1 and uuid = $2 limit 1
`

type GetPaymentLinkByPublicIDParams struct {
	MerchantID int64
	Uuid       uuid.UUID
}

func (q *Queries) GetPaymentLinkByPublicID(ctx context.Context, arg GetPaymentLinkByPublicIDParams) (PaymentLink, error) {
	row := q.db.QueryRow(ctx, getPaymentLinkByPublicID, arg.MerchantID, arg.Uuid)
	var i PaymentLink
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Slug,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.MerchantID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Decimals,
		&i.Currency,
		&i.SuccessAction,
		&i.RedirectUrl,
		&i.SuccessMessage,
		&i.IsTest,
	)
	return i, err
}

const getPaymentLinkBySlug = `-- name: GetPaymentLinkBySlug :one
select id, uuid, slug, created_at, updated_at, merchant_id, name, description, price, decimals, currency, success_action, redirect_url, success_message, is_test from payment_links where slug = $1 limit 1
`

func (q *Queries) GetPaymentLinkBySlug(ctx context.Context, slug string) (PaymentLink, error) {
	row := q.db.QueryRow(ctx, getPaymentLinkBySlug, slug)
	var i PaymentLink
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Slug,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.MerchantID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.Decimals,
		&i.Currency,
		&i.SuccessAction,
		&i.RedirectUrl,
		&i.SuccessMessage,
		&i.IsTest,
	)
	return i, err
}

const listPaymentLinks = `-- name: ListPaymentLinks :many
select id, uuid, slug, created_at, updated_at, merchant_id, name, description, price, decimals, currency, success_action, redirect_url, success_message, is_test from payment_links where merchant_id = $1 order by id desc limit $2
`

type ListPaymentLinksParams struct {
	MerchantID int64
	Limit      int32
}

func (q *Queries) ListPaymentLinks(ctx context.Context, arg ListPaymentLinksParams) ([]PaymentLink, error) {
	rows, err := q.db.Query(ctx, listPaymentLinks, arg.MerchantID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PaymentLink
	for rows.Next() {
		var i PaymentLink
		if err := rows.Scan(
			&i.ID,
			&i.Uuid,
			&i.Slug,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.MerchantID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.Decimals,
			&i.Currency,
			&i.SuccessAction,
			&i.RedirectUrl,
			&i.SuccessMessage,
			&i.IsTest,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
