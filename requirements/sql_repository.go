package requirements

import (
	"database/sql"
	"fmt"

	"github.com/kn100/studio9/logging"
)

//Repository is an interface capable of storing Requirement objects
type Repository interface {
	Add(Requirement) error
	AddCondition(Condition) error
	Update(...Requirement) error
	UpdateCondition(...Condition) error
	Requirement(string) (Requirement, error)
	Requirements() ([]Requirement, error)
	Condition(string) (Condition, error)
	Remove(...string) error
	RemoveCondition(...string) error
}

//SQLRepository is an object that is used to make SQL calls
type SQLRepository struct {
	Db     *sql.DB
	Logger logging.Logger
}

//AddCondition attempts to add the given Condition to the SQL database
func (repo *SQLRepository) AddCondition(con Condition) error {
	stmt, err := repo.Db.Prepare("INSERT INTO epos_item_requirement_conditions (uuid, type, data) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	conData, err := con.Save()
	if err != nil {
		return err
	}

	_, err = stmt.Exec(con.ID(), con.Type(), conData)
	if err != nil {
		return err
	}

	return nil
}

//Add attempts to add the given Requirement to the database
func (repo *SQLRepository) Add(req Requirement) error {
	//check if add condition
	con := req.GetCondition()
	if con != nil {
		err := repo.AddCondition(con)
		if err != nil {
			return err
		}
	}

	stmt, err := repo.Db.Prepare("INSERT INTO epos_item_requirements (uuid, `condition-uuid`, type) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	var conID string

	if con != nil {
		conID = con.ID()
	}
	fmt.Println("SQL VARS", req.ID(), conID, req.Type())
	_, err = stmt.Exec(req.ID(), conID, req.Type())
	if err != nil {
		repo.RemoveCondition(con.ID())
		return err
	}
	return nil
}

//RemoveCondition attempts to remove the Conditions with the specified IDs from the SQL database
func (repo *SQLRepository) RemoveCondition(ids ...string) error {
	fmt.Printf("REMOVECOND %#v", ids)
	for _, id := range ids {
		fmt.Println("Deleting Condition with id", id)
		stmt, err := repo.Db.Prepare("DELETE FROM epos_item_requirement_conditions WHERE uuid=?")
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(id)
		if err != nil {
			return err
		}
	}
	return nil
}

//Remove attempts to remove the Requirements of the specified IDs from the SQL database
func (repo *SQLRepository) Remove(ids ...string) error {
	for _, id := range ids {
		//Remove condition
		req, err := repo.Requirement(id)
		if err != nil {
			return err
		}
		fmt.Printf("checking to see if we have a condition %#v", req.GetCondition())
		if req != nil {
			fmt.Println("Have condition")
			con := req.GetCondition()
			if con != nil {
				fmt.Println("Removing condition")
				repo.RemoveCondition(con.ID())
			}
		}
		stmt, err := repo.Db.Prepare("DELETE FROM epos_item_requirements WHERE uuid=?")
		if err != nil {
			return err
		}

		defer stmt.Close()

		_, err = stmt.Exec(id)
		if err != nil {
			return err
		}
	}

	return nil
}

//Condition attempts to load a Condition from the SQL database with the specified ID
func (repo *SQLRepository) Condition(id string) (Condition, error) {
	var con Condition
	fmt.Println("GOT HERE")
	res, err := repo.Db.Query("SELECT * FROM epos_item_requirement_conditions WHERE uuid=?", id)
	if err != nil {
		return con, err
	}
	defer res.Close()

	if res.Next() {
		var t int
		var data string

		err = res.Scan(&id, &t, &data)
		if err != nil {
			return con, err
		}
		con, err = LoadCondition(t)
		if err != nil {
			return con, err
		}
		err = con.Load([]byte(data))
		if err != nil {
			return con, err
		}
		con.SetID(id)
	}
	return con, nil
}

func (repo *SQLRepository) Requirements() ([]Requirement, error) {
	var reqs []Requirement
	res, err := repo.Db.Query("SELECT uuid FROM epos_item_requirements")
	if err != nil {
		return reqs, err
	}

	defer res.Close()

	for res.Next() {
		var id string
		err = res.Scan(&id)
		if err != nil {
			return reqs, err
		}

		req, err := repo.Requirement(id)
		if err != nil {
			return reqs, err
		}

		reqs = append(reqs, req)
	}

	return reqs, nil
}

//Requirement attempts to load a Requirement from the SQL database with the specified ID
func (repo *SQLRepository) Requirement(id string) (Requirement, error) {
	var req Requirement
	res, err := repo.Db.Query("SELECT * FROM epos_item_requirements WHERE uuid=?", id)
	if err != nil {
		return req, err
	}
	defer res.Close()

	if res.Next() {
		var t int
		var conID string

		err = res.Scan(&id, &conID, &t)
		if err != nil {
			return req, err
		}
		con, err := repo.Condition(conID)
		if err != nil {
			return req, err
		}
		req, err = LoadRequirement(t)
		if err != nil {
			return req, err
		}
		req.SetID(id)

		err = req.Condition(con)
		if err != nil {
			return req, err
		}
	}

	return req, nil
}

//UpdateCondition attempts to update the given Condition and save the changes
//within the SQL database
func (repo *SQLRepository) UpdateCondition(cons ...Condition) error {
	for _, con := range cons {
		stmt, err := repo.Db.Prepare("UPDATE epos_item_requirement_conditions SET type=?, data=? WHERE uuid=?")

		if err != nil {
			return err
		}

		defer stmt.Close()

		conData, err := con.Save()
		if err != nil {
			return err
		}

		_, err = stmt.Exec(con.Type(), conData, con.ID())
		if err != nil {
			return err
		}

	}

	return nil
}

//Update attempts to update the given Requirement and save the changes
//within the SQL database
func (repo *SQLRepository) Update(reqs ...Requirement) error {
	for _, req := range reqs {
		stmt, err := repo.Db.Prepare("UPDATE epos_item_requirements SET type=?, `condition-uuid`=? WHERE uuid=?")

		if err != nil {
			return err
		}

		defer stmt.Close()

		var conID string
		if tmp := req.GetCondition(); tmp != nil {
			err := repo.UpdateCondition(tmp)
			if err != nil {
				return err
			}
			conID = tmp.ID()
		}

		_, err = stmt.Exec(req.Type(), conID, req.ID())
		if err != nil {
			return err
		}
	}

	return nil
}
