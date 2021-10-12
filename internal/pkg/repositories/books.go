package repositories

import (
	"bufio"
	"database/sql"
	"log"
	"os"
	"strings"

	"github.com/oglinuk/restful-go/internal/pkg/database"
	"github.com/oglinuk/restful-go/internal/pkg/models"
)

var (
	defaultBookSchema = `
		CREATE TABLE IF NOT EXISTS tblBooks(
			id TEXT PRIMARY KEY NOT NULL,
			title TEXT NOT NULL,
			author TEXT NOT NULL,
			published TEXT NOT NULL
		);`

	seedFile = "seeds.txt"
)

type BooksRepo struct{
	DB *sql.DB
}

// NewBooksRepo creates a new BooksRepo object using the given db. If the
// db is nil, create a new one with the defaultBookSchema. If there is a
// seedFile present, call br.Seed. Finally return the BooksRepo object.
func NewBooksRepo(db *sql.DB) *BooksRepo {
	var err error

	br := &BooksRepo{
		DB: db,
	}

	if br.DB == nil {
		br.DB = database.Open(defaultBookSchema)
	}

	if _, err = os.Stat(seedFile); err == nil {
		log.Println("Found a seed file.")
		br.seed()
	} else {
		log.Println("No seed file found ...")
	}

	return br
}

// seed opens the seedFile, and for each line of the file, inserts a
// NewBook into the database. 
func (br *BooksRepo) seed() {
	log.Println("Seeding database ...")

	f, err := os.Open(seedFile)
	if err != nil {
		log.Fatalf("seed::os.Open: %s\n", err.Error())
	}
	defer f.Close()

	bs := bufio.NewScanner(f)
	for bs.Scan() {
		parts := strings.Split(bs.Text(), ",")
		if err != nil {
			log.Fatalf("seed::time.Parse: %s\n", err.Error())
		}

		b := models.NewBook(parts[0], parts[1], parts[2])
		log.Printf("Inserting %v\n", b)
		_, _ = br.DB.Exec(`INSERT INTO tblBooks(id, title, author, published)
			VALUES(?,?,?,?)`, b.ID, b.Title, b.Author, b.Published)
	}
}

// ===== CREATE ===== //

// Insert (b)ook into the database
func (br *BooksRepo) Insert(b *models.Book) error {
	log.Printf("Inserting %v\n", b)
	stmt, err := br.DB.Prepare(`INSERT INTO tblBooks(
		id, title, author, published) VALUES(?, ?, ?, ?)`)
	if err != nil {
		return err
	}

	 _, err = stmt.Exec(b.ID, b.Title, b.Author, b.Published)
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

		err = rows.Scan(&b.ID, &b.Title, &b.Author, &b.Published)
		if err != nil {
			return nil, err
		}

		books = append(books, b)
	}

	return books, nil
}