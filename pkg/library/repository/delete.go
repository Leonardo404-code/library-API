package repository

import (
	"context"
	"fmt"

	"library-api/pkg/library"
)

func (r *repository) DeleteBook(filter *library.Filter, deleteFromBucket func() error) error {
	session, err := r.conn.StartSession()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrStartSession, err)
	}

	defer session.EndSession(context.TODO())

	if err = session.StartTransaction(); err != nil {
		return fmt.Errorf("%w: %v", ErrStartTransaction, err)
	}

	query := filter.GenerateQuery()
	collection := r.conn.Database(libraryDB).Collection(booksColl)

	if _, err := collection.DeleteOne(context.TODO(), query); err != nil {
		return err
	}

	if err = deleteFromBucket(); err != nil {
		return err
	}

	if err = session.CommitTransaction(context.TODO()); err != nil {
		return fmt.Errorf("%w: %v", ErrCommitTransaction, err)
	}

	return nil
}
