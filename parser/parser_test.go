package parser_test

import (
	_ "github.com/yarbelk/book_of_ages/parser"
	. "github.com/yarbelk/book_of_ages/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
	"encoding/xml"
)

var _ = Describe("Parser", func() {
	var reader *strings.Reader

	Context("world regions and sites", func() {
		It("should parse a region", func() {
			region_string := "<region><id>1</id><name>the sunny sea</name><type>Ocean</type></region> <region><id>2</id><name>the shadowy sea</name><type>Ocean</type></region>"
			reader = strings.NewReader(region_string)
			region := Region{}
			decoder := xml.NewDecoder(reader)
			region.Decode(decoder)
			Expect(region.Id).To(Equal(1))
			Expect(region.Name).To(Equal("the sunny sea"))
			Expect(region.Type).To(Equal("Ocean"))
		})

		It("should populate underground regions", func() {
			underground_region_string := "<underground_region><id>0</id><type>cavern</type><depth>1</depth></underground_region>"
			reader = strings.NewReader(underground_region_string)
			underground_region := UndergroundRegion{}
			decoder := xml.NewDecoder(reader)
			underground_region.Decode(decoder)
			Expect(underground_region.Id).To(Equal(0))
			Expect(underground_region.Depth).To(Equal(1))
			Expect(underground_region.Type).To(Equal("cavern"))
		})

		It("should populate sites", func() {
			site_string := "<site> <id>1</id> <type>cave</type> <name>dippeddeep</name> <coords>44,174</coords> <structures> </structures> </site>"
			reader = strings.NewReader(site_string)
			site := Site{}
			decoder := xml.NewDecoder(reader)
			site.Decode(decoder)
			Expect(site.Id).To(Equal(1))
			Expect(site.Name).To(Equal("dippeddeep"))
			Expect(site.Coords).To(Equal("44,174"))
			Expect(site.X_coord).To(Equal(44))
			Expect(site.Y_coord).To(Equal(174))
			Expect(site.Type).To(Equal("cave"))
		})
	})
})
