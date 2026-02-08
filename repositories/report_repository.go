package repositories

import (
	"database/sql"
	"kasir-api/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GetDailyReport() (*models.DailyReport, error) {
	report := &models.DailyReport{}

	// Get total revenue and total transactions for today
	query := `
		SELECT 
			COALESCE(SUM(total_amount), 0) as total_revenue,
			COUNT(*) as total_transaksi
		FROM transactions
		WHERE DATE(created_at) = CURRENT_DATE
	`
	err := repo.db.QueryRow(query).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	// Get best selling product for today
	bestSellerQuery := `
		SELECT 
			p.name,
			SUM(td.quantity) as qty_terjual
		FROM transaction_details td
		JOIN transactions t ON td.transaction_id = t.id
		JOIN products p ON td.product_id = p.id
		WHERE DATE(t.created_at) = CURRENT_DATE
		GROUP BY p.id, p.name
		ORDER BY qty_terjual DESC
		LIMIT 1
	`

	var bestSeller models.BestSellingProduct
	err = repo.db.QueryRow(bestSellerQuery).Scan(&bestSeller.Nama, &bestSeller.QtyTerjual)
	if err == sql.ErrNoRows {
		// No sales today, return report with nil best seller
		report.ProdukTerlaris = nil
		return report, nil
	}
	if err != nil {
		return nil, err
	}

	report.ProdukTerlaris = &bestSeller
	return report, nil
}

func (repo *ReportRepository) GetReportByDateRange(startDate, endDate string) (*models.DailyReport, error) {
	report := &models.DailyReport{}

	// Get total revenue and total transactions for date range
	query := `
		SELECT 
			COALESCE(SUM(total_amount), 0) as total_revenue,
			COUNT(*) as total_transaksi
		FROM transactions
		WHERE DATE(created_at) >= $1 AND DATE(created_at) <= $2
	`
	err := repo.db.QueryRow(query, startDate, endDate).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	if err != nil {
		return nil, err
	}

	// Get best selling product for date range
	bestSellerQuery := `
		SELECT 
			p.name,
			SUM(td.quantity) as qty_terjual
		FROM transaction_details td
		JOIN transactions t ON td.transaction_id = t.id
		JOIN products p ON td.product_id = p.id
		WHERE DATE(t.created_at) >= $1 AND DATE(t.created_at) <= $2
		GROUP BY p.id, p.name
		ORDER BY qty_terjual DESC
		LIMIT 1
	`

	var bestSeller models.BestSellingProduct
	err = repo.db.QueryRow(bestSellerQuery, startDate, endDate).Scan(&bestSeller.Nama, &bestSeller.QtyTerjual)
	if err == sql.ErrNoRows {
		// No sales in this date range
		report.ProdukTerlaris = nil
		return report, nil
	}
	if err != nil {
		return nil, err
	}

	report.ProdukTerlaris = &bestSeller
	return report, nil
}
