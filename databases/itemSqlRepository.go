package databases

import (
	"database/sql"
	"errors"
	"os/exec"
	"sort"
	"strconv"
	"strings"

	"github.com/auenc/gTeller-core/items"
	"github.com/auenc/gTeller-core/logging"
	"github.com/auenc/gTeller-core/requirements"
)

func UUID() (string, error) {
	tmp, err := exec.Command("uuidgen").Output()
	if err != nil {
		return "", err
	}

	return strings.Replace(string(tmp), "\n", "", -1), nil
}

type SQLItemRepository struct {
	reqRepo requirements.Repository
	db      *sql.DB
	logger  logging.Logger
}

func (sqlR *SQLItemRepository) Test() {
	sqlR.logger.LogLine("TESTING SQL ITEM REPOSITORY")

	var str string
	err := sqlR.db.QueryRow("select message from hello WHERE id = 1").Scan(&str)
	if err != nil && err != sql.ErrNoRows {
		sqlR.logger.LogLine(err)
	}
	sqlR.logger.LogLine(str)

	item, err := sqlR.Item("8b9f12e4-8fdc-4cd3-b323-ec61d15757b3")
	if err != nil {
		sqlR.logger.LogLine(err)
	}
	sqlR.logger.LogLine("Got item", item)
}

func (sqlR *SQLItemRepository) AddItemOptions(item items.Item) error {
	for _, req := range item.Requirements {
		err := sqlR.reqRepo.Add(req)
		if err != nil {
			return err
		}
	}
	return nil
}

func (sqlR *SQLItemRepository) Add(item items.Item) error {

	stmt, err := sqlR.db.Prepare("INSERT INTO epos_items (uuid, name, price) VALUES(?, ?, ?);")
	if err != nil {
		sqlR.logger.LogLine(err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(item.Id(), item.Name(), item.PriceRaw)
	if err != nil {
		sqlR.logger.LogLine(err)
		return err
	}

	sqlR.AddItemOptions(item)

	sqlR.logger.LogLine("Added item to database")
	return nil
}

func (sqlR *SQLItemRepository) AddOption(optionCat items.OptionCategory) error {

	stmt, err := sqlR.db.Prepare("INSERT INTO epos_options (uuid, name) VALUES(?, ?)")
	if err != nil {
		sqlR.logger.LogLine(err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(optionCat.ID, optionCat.Name)
	if err != nil {
		sqlR.logger.LogLine(err)
		return err
	}

	for _, option := range optionCat.Options {
		stmt, err = sqlR.db.Prepare("INSERT INTO epos_option_choices (`uuid`, `option-uuid`, `option-name`, `option-price`) VALUES (?, ?, ?, ?);")
		if err != nil {
			sqlR.logger.LogLine(err)
			return err
		}

		_, err = stmt.Exec(option.ID, optionCat.ID, option.Name, strconv.FormatFloat(option.Price, 'f', -1, 64))
		if err != nil {
			sqlR.logger.LogLine(err)
			return err
		}
	}
	return nil
}

func (sqlR *SQLItemRepository) OptionChoices(id string) ([]items.Option, error) {
	choices := make([]items.Option, 0)

	rows, err := sqlR.db.Query("SELECT `uuid`, `option-name`, `option-price` FROM epos_option_choices WHERE `option-uuid` = '" + id + "'")
	if err != nil {
		return choices, err
	}

	defer rows.Close()

	for rows.Next() {
		var uuid string
		var name string
		var priceRaw string
		rows.Scan(&uuid, &name, &priceRaw)
		if price, err := strconv.ParseFloat(priceRaw, 64); err == nil {
			choices = append(choices, items.Option{uuid, name, price})
		} else {
			return choices, err
		}

	}

	//Sorting option
	sort.Sort(items.ByPrice(choices))

	return choices, nil
}

func (sqlR *SQLItemRepository) ItemOptions(itemUuid string) ([]requirements.Requirement, error) {
	reqs := make([]requirements.Requirement, 0)

	rows, err := sqlR.db.Query("SELECT `requirement-uuid` FROM epos_item_requirements_connection WHERE `item-uuid`=?", itemUuid)
	if err != nil {
		return reqs, err
	}

	defer rows.Close()

	for rows.Next() {
		var uuid string
		err = rows.Scan(&uuid)
		if err != nil {
			return reqs, err
		}
		req, err := sqlR.reqRepo.Requirement(uuid)
		if err != nil {
			return reqs, err
		}
		reqs = append(reqs, req)

	}

	return reqs, nil
}

func (sqlR *SQLItemRepository) Option(id string) (items.OptionCategory, error) {
	optionCat := items.OptionCategory{}

	//Getting data to build Option Category
	data, err := sqlR.db.Query("SELECT uuid, name FROM epos_options WHERE uuid='" + id + "'")
	if err != nil {
		return optionCat, err
	}

	defer data.Close()

	if data.Next() {
		var uuid string
		var name string

		//Getting data from result
		err = data.Scan(&uuid, &name)
		if err != nil {
			return optionCat, err
		}
		choices, err := sqlR.OptionChoices(uuid)
		if err != nil {
			return optionCat, err
		}
		return items.OptionCategory{uuid, name, choices}, nil
	}
	return optionCat, nil
}

func (sqlR *SQLItemRepository) Item(id string) (*items.Item, error) {

	rows, err := sqlR.db.Query("SELECT uuid, name, price FROM epos_items WHERE uuid='" + id + "'")
	if err != nil {
		sqlR.logger.LogLine(err)
	}

	defer rows.Close()

	if rows.Next() {
		var uuid string
		var name string
		var priceRAW string
		err = rows.Scan(&uuid, &name, &priceRAW)
		if err != nil {
			return nil, err
		}
		options, err := sqlR.ItemOptions(uuid)
		if err != nil {
			return nil, err
		}
		return &items.Item{uuid, name, priceRAW, nil, options, ""}, nil
	}
	return nil, errors.New("Could not find item")
}

func (sqlR *SQLItemRepository) Items() ([]items.Item, error) {
	var itemList []items.Item

	//Running select query that will return a list of items
	rows, err := sqlR.db.Query("SELECT * FROM epos_items")
	if err != nil {
		return itemList, err
	}

	defer rows.Close()

	//Going through each row to create an item out of the data received
	for rows.Next() {
		var uuid string
		var name string
		var priceRaw string
		err = rows.Scan(&uuid, &name, &priceRaw)
		if err != nil {
			return itemList, err
		}

		options, err := sqlR.ItemOptions(uuid)
		if err != nil {
			return nil, err
		}

		item := items.Item{uuid, name, priceRaw, nil, options, ""}
		itemList = append(itemList, item)
	}
	return itemList, nil
}

func (sqlR *SQLItemRepository) Options() ([]items.OptionCategory, error) {
	options := make([]items.OptionCategory, 0)

	rows, err := sqlR.db.Query("SELECT uuid FROM epos_options")
	if err != nil {
		sqlR.logger.LogLine(err)
		return options, err
	}

	defer rows.Close()

	for rows.Next() {
		var uuid string
		err = rows.Scan(&uuid)
		if err != nil {
			sqlR.logger.LogLine(err)
			return options, err
		}

		option, err := sqlR.Option(uuid)
		if err != nil {
			return options, err
		}
		options = append(options, option)
	}

	return options, nil
}

func (sqlR *SQLItemRepository) Remove(id string) error {

	stmt, err := sqlR.db.Prepare("DELETE FROM epos_items WHERE uuid=?")
	if err != nil {
		sqlR.logger.LogLine(err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		sqlR.logger.LogLine(err)
		return err
	}
	return nil
}

func (sqlR *SQLItemRepository) Update(item items.Item) error {
	stmt, err := sqlR.db.Prepare("UPDATE epos_items SET name=?, price=? WHERE uuid=?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(item.NameRaw, item.PriceRaw, item.ID)
	if err != nil {
		return err
	}
	//Finish this
	return nil
}

func (sqlR *SQLItemRepository) RemoveOption(id string) error {

	stmt, err := sqlR.db.Prepare("DELETE FROM epos_item_options WHERE uuid=?")
	if err != nil {
		sqlR.logger.LogLine(err)
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		sqlR.logger.LogLine(err)
		return err
	}
	return nil
}
