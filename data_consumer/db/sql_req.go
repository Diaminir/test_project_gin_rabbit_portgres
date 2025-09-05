package db

import (
	"context"
	"data_consumer/dto"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Postgres struct {
	db *pgx.Conn
}

func NewDbLogin() (*Postgres, error) {
	connFig := "postgres://qwert:12345@postgres:5432/orders"
	conn, err := pgx.Connect(context.Background(), connFig)
	if err != nil {
		return &Postgres{}, nil
	}
	return &Postgres{
		db: conn,
	}, nil
}

func (pg *Postgres) DbClose() {
	pg.db.Close(context.Background())
}

func (pg *Postgres) DbPut(consRabbit dto.ConsumeRabbitDTO) error {
	Sql := "INSERT INTO orders(title, price, created_at) VALUES ($1, $2, $3)"
	_, err := pg.db.Exec(context.Background(), Sql, consRabbit.CarInfo, consRabbit.Price, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return err
	}
	return nil
}

func (pg *Postgres) DbSearch(c *gin.Context) (pgx.Rows, error) {
	query := "SELECT * FROM orders"
	var conditions []string
	var args []interface{}
	argPos := 1

	minPrice := c.DefaultQuery("minPrice", "")
	maxPrice := c.DefaultQuery("maxPrice", "")

	if minPrice != "" {
		if _, err := strconv.Atoi(minPrice); err != nil {
			return nil, fmt.Errorf("invalid minPrice parameter")
		}
		conditions = append(conditions, fmt.Sprintf("price >= $%d", argPos))
		args = append(args, minPrice)
		argPos++
	}

	if maxPrice != "" {
		if _, err := strconv.Atoi(maxPrice); err != nil {
			return nil, fmt.Errorf("invalid maxPrice parameter")
		}
		conditions = append(conditions, fmt.Sprintf("price <= $%d", argPos))
		args = append(args, maxPrice)
		argPos++
	}

	minDate := c.DefaultQuery("minDate", "")
	maxDate := c.DefaultQuery("maxDate", "")

	if minDate != "" {
		if _, err := time.Parse("2006-01-02", minDate); err != nil {
			return nil, fmt.Errorf("invalid minDate format, use YYYY-MM-DD")
		}
		conditions = append(conditions, fmt.Sprintf("created_at >= $%d", argPos))
		args = append(args, minDate)
		argPos++
	}

	if maxDate != "" {
		if _, err := time.Parse("2006-01-02", maxDate); err != nil {
			return nil, fmt.Errorf("invalid maxDate format, use YYYY-MM-DD")
		}
		conditions = append(conditions, fmt.Sprintf("created_at <= $%d", argPos))
		args = append(args, maxDate)
		argPos++
	}

	searchName := c.DefaultQuery("searchName", "")

	if searchName != "" {
		conditions = append(conditions, fmt.Sprintf("title ILIKE $%d", argPos))
		args = append(args, "%"+searchName+"%")
		argPos++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	orderBy := getOrderByClause(c)
	if orderBy != "" {
		query += " ORDER BY " + orderBy
	}

	page, limit := getPaginationParams(c)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argPos, argPos+1)
	args = append(args, limit, (page-1)*limit)

	return pg.db.Query(context.Background(), query, args...)
}

func getOrderByClause(c *gin.Context) string {
	orderParams := map[string]string{
		"DateMaxMin":  "created_at DESC",
		"DateMinMax":  "created_at ASC",
		"PriceMaxMin": "price DESC",
		"PriceMinMax": "price ASC",
	}

	for param, order := range orderParams {
		if c.DefaultQuery(param, "") != "" {
			return order
		}
	}

	return ""
}

func getPaginationParams(c *gin.Context) (int, int) {
	defaultPage := 1
	defaultLimit := 5
	maxLimit := 100

	pageStr := c.DefaultQuery("page", strconv.Itoa(defaultPage))
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = defaultPage
	}

	limitStr := c.DefaultQuery("limit", strconv.Itoa(defaultLimit))
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = defaultLimit
	}

	if limit > maxLimit {
		limit = maxLimit
	}

	return page, limit
}

func (pg *Postgres) GetTotalCount(c *gin.Context) (int, error) {
	query := "SELECT COUNT(*) FROM orders"
	var conditions []string
	var args []interface{}
	argPos := 1

	minPrice := c.DefaultQuery("minPrice", "")
	maxPrice := c.DefaultQuery("maxPrice", "")

	if minPrice != "" {
		conditions = append(conditions, fmt.Sprintf("price >= $%d", argPos))
		args = append(args, minPrice)
		argPos++
	}

	if maxPrice != "" {
		conditions = append(conditions, fmt.Sprintf("price <= $%d", argPos))
		args = append(args, maxPrice)
		argPos++
	}

	minDate := c.DefaultQuery("minDate", "")
	maxDate := c.DefaultQuery("maxDate", "")

	if minDate != "" {
		conditions = append(conditions, fmt.Sprintf("created_at >= $%d", argPos))
		args = append(args, minDate)
		argPos++
	}

	if maxDate != "" {
		conditions = append(conditions, fmt.Sprintf("created_at <= $%d", argPos))
		args = append(args, maxDate)
		argPos++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	var totalCount int
	err := pg.db.QueryRow(context.Background(), query, args...).Scan(&totalCount)
	return totalCount, err
}
