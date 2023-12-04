package loan

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/miriam-samuels/loan-management-backend/internal/database"
	"github.com/miriam-samuels/loan-management-backend/internal/types"
)

func (loanApp *Loan) CreateLoan(id uuid.UUID, loanId string, borrowerId string, w http.ResponseWriter) (*sql.Stmt, error) {
	//  prepare query statement to create loan application in db
	stmt, err := database.LoanDb.Prepare("INSERT INTO applications (id,loanId, borrowerId, type, term, amount, purpose, has_collateral, collateral, collateral_docs, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")
	return stmt, err
}

func GetLoans(currentUser types.AuthCtxKey, statusCondition string) (*sql.Rows, error) {
	var rows *sql.Rows
	var err error
	if currentUser.Role == "borrower" {
		rows, err = database.LoanDb.Query("SELECT * FROM applications WHERE borrowerId = $1"+statusCondition, currentUser.Id)
	} else if currentUser.Role == "lender" {
		rows, err = database.LoanDb.Query("SELECT * FROM applications" + statusCondition)
	} else {
		return rows, errors.New("unauthorized")
	}

	return rows, err
}

func (loan *Loan) GetLoanById() *sql.Row {
	row := database.LoanDb.QueryRow("SELECT * FROM applications WHERE id = $1", loan.ID)
	return row
}
