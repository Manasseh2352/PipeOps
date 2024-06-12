// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createActivityStmt, err = db.PrepareContext(ctx, createActivity); err != nil {
		return nil, fmt.Errorf("error preparing query CreateActivity: %w", err)
	}
	if q.createEmergencyContactStmt, err = db.PrepareContext(ctx, createEmergencyContact); err != nil {
		return nil, fmt.Errorf("error preparing query CreateEmergencyContact: %w", err)
	}
	if q.createTravelChecklistStmt, err = db.PrepareContext(ctx, createTravelChecklist); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTravelChecklist: %w", err)
	}
	if q.createTripStmt, err = db.PrepareContext(ctx, createTrip); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTrip: %w", err)
	}
	if q.createTripBookingStmt, err = db.PrepareContext(ctx, createTripBooking); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTripBooking: %w", err)
	}
	if q.createTripLogsStmt, err = db.PrepareContext(ctx, createTripLogs); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTripLogs: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteActivityStmt, err = db.PrepareContext(ctx, deleteActivity); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteActivity: %w", err)
	}
	if q.deleteItineraryStmt, err = db.PrepareContext(ctx, deleteItinerary); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteItinerary: %w", err)
	}
	if q.deleteTravelChecklistStmt, err = db.PrepareContext(ctx, deleteTravelChecklist); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTravelChecklist: %w", err)
	}
	if q.deleteTripBookingStmt, err = db.PrepareContext(ctx, deleteTripBooking); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTripBooking: %w", err)
	}
	if q.deleteTripLogsStmt, err = db.PrepareContext(ctx, deleteTripLogs); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTripLogs: %w", err)
	}
	if q.deleteUserStmt, err = db.PrepareContext(ctx, deleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.getATripLogUpdateStmt, err = db.PrepareContext(ctx, getATripLogUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query GetATripLogUpdate: %w", err)
	}
	if q.getActivityUpdateStmt, err = db.PrepareContext(ctx, getActivityUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query GetActivityUpdate: %w", err)
	}
	if q.getEmergencyContactUpdateStmt, err = db.PrepareContext(ctx, getEmergencyContactUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query GetEmergencyContactUpdate: %w", err)
	}
	if q.getTravelChecklistStmt, err = db.PrepareContext(ctx, getTravelChecklist); err != nil {
		return nil, fmt.Errorf("error preparing query GetTravelChecklist: %w", err)
	}
	if q.getTripBookingUpdateStmt, err = db.PrepareContext(ctx, getTripBookingUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query GetTripBookingUpdate: %w", err)
	}
	if q.getTripUpdateStmt, err = db.PrepareContext(ctx, getTripUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query GetTripUpdate: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, getUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.getUserForUpdateStmt, err = db.PrepareContext(ctx, getUserForUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserForUpdate: %w", err)
	}
	if q.listActivitiesStmt, err = db.PrepareContext(ctx, listActivities); err != nil {
		return nil, fmt.Errorf("error preparing query ListActivities: %w", err)
	}
	if q.listTravelChecklistStmt, err = db.PrepareContext(ctx, listTravelChecklist); err != nil {
		return nil, fmt.Errorf("error preparing query ListTravelChecklist: %w", err)
	}
	if q.listTripBookingStmt, err = db.PrepareContext(ctx, listTripBooking); err != nil {
		return nil, fmt.Errorf("error preparing query ListTripBooking: %w", err)
	}
	if q.listTripLogsStmt, err = db.PrepareContext(ctx, listTripLogs); err != nil {
		return nil, fmt.Errorf("error preparing query ListTripLogs: %w", err)
	}
	if q.listTripsStmt, err = db.PrepareContext(ctx, listTrips); err != nil {
		return nil, fmt.Errorf("error preparing query ListTrips: %w", err)
	}
	if q.listsAccountsStmt, err = db.PrepareContext(ctx, listsAccounts); err != nil {
		return nil, fmt.Errorf("error preparing query ListsAccounts: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createActivityStmt != nil {
		if cerr := q.createActivityStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createActivityStmt: %w", cerr)
		}
	}
	if q.createEmergencyContactStmt != nil {
		if cerr := q.createEmergencyContactStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createEmergencyContactStmt: %w", cerr)
		}
	}
	if q.createTravelChecklistStmt != nil {
		if cerr := q.createTravelChecklistStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTravelChecklistStmt: %w", cerr)
		}
	}
	if q.createTripStmt != nil {
		if cerr := q.createTripStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTripStmt: %w", cerr)
		}
	}
	if q.createTripBookingStmt != nil {
		if cerr := q.createTripBookingStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTripBookingStmt: %w", cerr)
		}
	}
	if q.createTripLogsStmt != nil {
		if cerr := q.createTripLogsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTripLogsStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteActivityStmt != nil {
		if cerr := q.deleteActivityStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteActivityStmt: %w", cerr)
		}
	}
	if q.deleteItineraryStmt != nil {
		if cerr := q.deleteItineraryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteItineraryStmt: %w", cerr)
		}
	}
	if q.deleteTravelChecklistStmt != nil {
		if cerr := q.deleteTravelChecklistStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTravelChecklistStmt: %w", cerr)
		}
	}
	if q.deleteTripBookingStmt != nil {
		if cerr := q.deleteTripBookingStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTripBookingStmt: %w", cerr)
		}
	}
	if q.deleteTripLogsStmt != nil {
		if cerr := q.deleteTripLogsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTripLogsStmt: %w", cerr)
		}
	}
	if q.deleteUserStmt != nil {
		if cerr := q.deleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserStmt: %w", cerr)
		}
	}
	if q.getATripLogUpdateStmt != nil {
		if cerr := q.getATripLogUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getATripLogUpdateStmt: %w", cerr)
		}
	}
	if q.getActivityUpdateStmt != nil {
		if cerr := q.getActivityUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getActivityUpdateStmt: %w", cerr)
		}
	}
	if q.getEmergencyContactUpdateStmt != nil {
		if cerr := q.getEmergencyContactUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getEmergencyContactUpdateStmt: %w", cerr)
		}
	}
	if q.getTravelChecklistStmt != nil {
		if cerr := q.getTravelChecklistStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTravelChecklistStmt: %w", cerr)
		}
	}
	if q.getTripBookingUpdateStmt != nil {
		if cerr := q.getTripBookingUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTripBookingUpdateStmt: %w", cerr)
		}
	}
	if q.getTripUpdateStmt != nil {
		if cerr := q.getTripUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getTripUpdateStmt: %w", cerr)
		}
	}
	if q.getUserStmt != nil {
		if cerr := q.getUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStmt: %w", cerr)
		}
	}
	if q.getUserForUpdateStmt != nil {
		if cerr := q.getUserForUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserForUpdateStmt: %w", cerr)
		}
	}
	if q.listActivitiesStmt != nil {
		if cerr := q.listActivitiesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listActivitiesStmt: %w", cerr)
		}
	}
	if q.listTravelChecklistStmt != nil {
		if cerr := q.listTravelChecklistStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listTravelChecklistStmt: %w", cerr)
		}
	}
	if q.listTripBookingStmt != nil {
		if cerr := q.listTripBookingStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listTripBookingStmt: %w", cerr)
		}
	}
	if q.listTripLogsStmt != nil {
		if cerr := q.listTripLogsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listTripLogsStmt: %w", cerr)
		}
	}
	if q.listTripsStmt != nil {
		if cerr := q.listTripsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listTripsStmt: %w", cerr)
		}
	}
	if q.listsAccountsStmt != nil {
		if cerr := q.listsAccountsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listsAccountsStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                            DBTX
	tx                            *sql.Tx
	createActivityStmt            *sql.Stmt
	createEmergencyContactStmt    *sql.Stmt
	createTravelChecklistStmt     *sql.Stmt
	createTripStmt                *sql.Stmt
	createTripBookingStmt         *sql.Stmt
	createTripLogsStmt            *sql.Stmt
	createUserStmt                *sql.Stmt
	deleteActivityStmt            *sql.Stmt
	deleteItineraryStmt           *sql.Stmt
	deleteTravelChecklistStmt     *sql.Stmt
	deleteTripBookingStmt         *sql.Stmt
	deleteTripLogsStmt            *sql.Stmt
	deleteUserStmt                *sql.Stmt
	getATripLogUpdateStmt         *sql.Stmt
	getActivityUpdateStmt         *sql.Stmt
	getEmergencyContactUpdateStmt *sql.Stmt
	getTravelChecklistStmt        *sql.Stmt
	getTripBookingUpdateStmt      *sql.Stmt
	getTripUpdateStmt             *sql.Stmt
	getUserStmt                   *sql.Stmt
	getUserForUpdateStmt          *sql.Stmt
	listActivitiesStmt            *sql.Stmt
	listTravelChecklistStmt       *sql.Stmt
	listTripBookingStmt           *sql.Stmt
	listTripLogsStmt              *sql.Stmt
	listTripsStmt                 *sql.Stmt
	listsAccountsStmt             *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                            tx,
		tx:                            tx,
		createActivityStmt:            q.createActivityStmt,
		createEmergencyContactStmt:    q.createEmergencyContactStmt,
		createTravelChecklistStmt:     q.createTravelChecklistStmt,
		createTripStmt:                q.createTripStmt,
		createTripBookingStmt:         q.createTripBookingStmt,
		createTripLogsStmt:            q.createTripLogsStmt,
		createUserStmt:                q.createUserStmt,
		deleteActivityStmt:            q.deleteActivityStmt,
		deleteItineraryStmt:           q.deleteItineraryStmt,
		deleteTravelChecklistStmt:     q.deleteTravelChecklistStmt,
		deleteTripBookingStmt:         q.deleteTripBookingStmt,
		deleteTripLogsStmt:            q.deleteTripLogsStmt,
		deleteUserStmt:                q.deleteUserStmt,
		getATripLogUpdateStmt:         q.getATripLogUpdateStmt,
		getActivityUpdateStmt:         q.getActivityUpdateStmt,
		getEmergencyContactUpdateStmt: q.getEmergencyContactUpdateStmt,
		getTravelChecklistStmt:        q.getTravelChecklistStmt,
		getTripBookingUpdateStmt:      q.getTripBookingUpdateStmt,
		getTripUpdateStmt:             q.getTripUpdateStmt,
		getUserStmt:                   q.getUserStmt,
		getUserForUpdateStmt:          q.getUserForUpdateStmt,
		listActivitiesStmt:            q.listActivitiesStmt,
		listTravelChecklistStmt:       q.listTravelChecklistStmt,
		listTripBookingStmt:           q.listTripBookingStmt,
		listTripLogsStmt:              q.listTripLogsStmt,
		listTripsStmt:                 q.listTripsStmt,
		listsAccountsStmt:             q.listsAccountsStmt,
	}
}
