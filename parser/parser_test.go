package parser_test

import (
	. "github.com/yarbelk/book_of_ages/parser"
	. "github.com/yarbelk/book_of_ages/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = Describe("Parser", func() {
	var (
		inputXML string
		world World
		reader *strings.Reader
	)

	Context("world regions and sites", func() {
		BeforeEach(func() {
inputXML=`<?xmlversion="1.0"encoding='UTF-8'?>
<df_world>
<regions>
<region><id>1</id><name>the sunny sea</name><type>Ocean</type></region>
<region><id>2</id><name>the shadowy sea</name><type>Ocean</type></region>
</regions>
<underground_regions>
<underground_region><id>0</id><type>cavern</type><depth>1</depth></underground_region>
<underground_region><id>1</id><type>cavern</type><depth>1</depth></underground_region>
</underground_regions>
</df_world>`
		world = World{Name:"Test World", Id: 0}
		reader = strings.NewReader(inputXML)
		})

		It("should populate regions", func() {
			Parse(&world, reader)
			Expect(len(world.RegionList.Regions)).To(Equal(1))
			region := world.RegionList.Regions[0]
			Expect(region.Id).To(Equal(1))
			Expect(region.Name).To(Equal("the sunny sea"))
			Expect(region.Type).To(Equal("Ocean"))
		})

		It("should populate underground regions", func() {
			Parse(&world, reader)
			Expect(len(world.UndergroundRegionList.UndergroundRegions)).To(Equal(1))
			region := world.UndergroundRegionList.UndergroundRegions[0]
			Expect(region.Id).To(Equal(0))
			Expect(region.Depth).To(Equal(1))
			Expect(region.Type).To(Equal("cavern"))
		})
	})
})
