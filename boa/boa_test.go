package boa_test

import (
	. "github.com/yarbelk/book_of_ages/boa"
	"github.com/yarbelk/book_of_ages/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/jmoiron/sqlx"
	"gopkg.in/gorp.v1"
	"strings"
	"fmt"
)

var _ = Describe("Boa", func() {
	var (
		world_reader *strings.Reader
		xml_reader *strings.Reader
	)
	Context("Ingest World Data", func () {
		var (
			world_history string = "Otstuxsmata\nThe Cyclopean Realms\n\n"
			xml_history string = ""
			dbmap *gorp.DbMap
		)

		BeforeEach(func () {
			db.Database = ":memory:"
			fmt.Println("Database: ", db.Database)
			dbmap = db.InitDb()
			db.RunMigrations()
			dbmap.TruncateTables()
			world_reader = strings.NewReader(world_history)
			xml_reader = strings.NewReader(xml_history)
		})

		AfterEach(func () {
			db.CloseDb()
		})

		It("Should persist the World Data", func () {
			var world *db.World = &db.World{}
			IngestData(xml_reader, world_reader)
			dbmap.SelectOne(world, "select * from worlds")

			Expect(world.Id).To(Equal(1))  // 1-indexed
			Expect(world.Name).To(Equal("Otstuxsmata"))
			Expect(world.TranslatedName).To(Equal("The Cyclopean Realms"))
		})
	})
})
