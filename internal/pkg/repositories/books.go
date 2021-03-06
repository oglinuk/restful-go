package repositories

import (
	"encoding/csv"
	"database/sql"
	"log"
	"os"

	"github.com/oglinuk/restful-go/internal/pkg/database"
	"github.com/oglinuk/restful-go/internal/pkg/models"
)

var (
	bookSchema = `
		CREATE TABLE IF NOT EXISTS tblBooks(
			id TEXT PRIMARY KEY NOT NULL,
			title TEXT NOT NULL,
			author TEXT NOT NULL,
			published TEXT NOT NULL,
			genre TEXT NOT NULL,
			read_status TEXT NOT NULL
		);`

	seedFile = "seeds.csv"
)

type BooksRepo struct{
	DB *sql.DB
}

// NewBooksRepo creates a new BooksRepo object using the given db. If the
// db is nil, create a new one with the bookSchema. If there is a
// seedFile present, call br.Seed. Finally return the BooksRepo object.
func NewBooksRepo(db *sql.DB) *BooksRepo {
	var err error

	br := &BooksRepo{
		DB: db,
	}

	if br.DB == nil {
		br.DB = database.Open(bookSchema)
	}

	if _, err = os.Stat(seedFile); err == nil {
		br.seed()
	}

	return br
}

// seed opens the seedFile (if exists), and for each record in the file,
// inserts a NewBook into the database. 
func (br *BooksRepo) seed() {
	log.Println("Seeding database ...")

	f, err := os.Open(seedFile)
	if err != nil {
		log.Fatalf("seed::os.Open: %s\n", err.Error())
	}
	defer f.Close()

	csvr := csv.NewReader(f)
	csvr.TrimLeadingSpace = true
	records, err := csvr.ReadAll()
	if err != nil {
		log.Fatalf("seed::csv.ReadAll: %s\n", err.Error())
	}

	// the first row is always assumed to be column headers 
	for _, r := range records[1:] {
		b := models.NewBook(r[0], r[1], r[2], r[3], r[4])
		log.Printf("[BooksRepo] Inserting %v\n", b)
		_, _ = br.DB.Exec(`INSERT INTO tblBooks(
			id, title, author, published, genre, read_status)
			VALUES(?,?,?,?,?,?)`,
			b.ID, b.Title, b.Author, b.Published, b.Genre, b.ReadStatus)
	}
}

// ===== CREATE ===== //

// Insert (b)ook into the database
func (br *BooksRepo) Insert(b *models.Book) error {
	log.Printf("[BooksRepo] Inserting %v\n", b)
	stmt, err := br.DB.Prepare(`INSERT INTO tblBooks(
		id, title, author, published, genre, read_status)
		VALUES(?,?,?,?,?,?)`)
	if err != nil {
		return err
	}

	 _, err = stmt.Exec(
		b.ID,
		b.Title,
		b.Author,
		b.Published,
		b.Genre,
		b.ReadStatus,
	)
	if err != nil {
		return err
	}

	return nil
}

// ===== RETRIEVE ===== //

// SelectAll books from the database
func (br *BooksRepo) SelectAll() ([]*models.Book, error) {
	var books []*models.Book

	rows, err := br.DB.Query(`SELECT * FROM tblBooks`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		 b := &models.Book{}

		err = rows.Scan(
			&b.ID,
			&b.Title,
			&b.Author,
			&b.Published,
			&b.Genre,
			&b.ReadStatus,
		)
		if err != nil {
			return nil, err
		}

		books = append(books, b)
	}

	return books, nil
}

// Select book by id
func (br *BooksRepo) Select(id string) (*models.Book, error) {
	log.Printf("[BooksRepo] Selecting %s\n", id)

	book := &models.Book{}

	row := br.DB.QueryRow(`SELECT * FROM tblBooks WHERE ID=?`, id)

	err := row.Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.Published,
		&book.Genre,
		&book.ReadStatus,
	)
	if err != nil {
		return nil, err
	}

	return book, nil
}

// ===== UPDATE ===== //

// Update deletes the old book by id and inserts the new given book
func (br *BooksRepo) Update(id string, book *models.Book) (string, error) {
	log.Printf("[BooksRepo] Updating [%s]: %v\n", id, book)

	err := br.Delete(id)
	if err != nil {
		return "", err
	}

	err = br.Insert(book)
	if err != nil {
		return "", err
	}

	return book.ID, nil
}

// ===== DELETE ===== //

// Delete book by id
func (br *BooksRepo) Delete(id string) error {
	log.Printf("[BooksRepo] Deleting %s\n", id)

	stmt, err := br.DB.Prepare(`DELETE FROM tblBooks WHERE ID=?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(id); err != nil {
		return err
	}

	return nil
}
