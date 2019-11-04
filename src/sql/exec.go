package sql

import (
	"database/sql"
	"dsc/fancy_errors"
	"time"
)

func ApplyChangeSet(changeSet map[int][]Query, dryrun bool) error {

	var txList []*sql.Tx

	for i := 0; i < len(changeSet); i++ {
		queryList := changeSet[i]

		// TODO refactor this
		incompleteTx, err := getTx()
		if err != nil {
			err = RevertChangeSet(txList)
			if err != nil {
				return fancy_errors.Wrap(err)
			}

			return fancy_errors.Wrap(err)
		}

		txList = append(txList, incompleteTx)

		for _, query := range queryList {
			err = Exec(incompleteTx, query)
			if err != nil {

				err = RevertChangeSet(txList)
				if err != nil {
					return fancy_errors.Wrap(err)
				}

				return fancy_errors.Wrap(err)
			}
		}

	}

	if dryrun {
		err = RevertChangeSet(txList)
		if err != nil {
			return fancy_errors.Wrap(err)
		}
	} else {
		// TODO pop and rollback other transactions?
		for _, tx := range txList {
			tx.Commit()
		}
	}

	return nil
}

func RevertChangeSet(txList []*sql.Tx) error {

	for _, tx := range txList {
		err := tx.Rollback()
		if err != nil {
			return fancy_errors.Wrap(err)
		}
	}

	return nil
}

func Exec(tx *sql.Tx, query Query) error {

	query.StartTime = time.Now()

	_, err := tx.Exec(query.Content)
	if err != nil {
		return fancy_errors.Wrap(err)
	}

	query.EndTime = time.Now()

	return nil
}
