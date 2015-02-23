package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	srastudy "github.com/indraniel/go-sra-schemas-1.5/SRA.study.xsd_go"
)

type SraStudy struct {
	XMLName xml.Name `xml:"STUDY_SET"`
	srastudy.TStudySetType
}

func (ss SraStudy) String() string {
	xml, err := xml.MarshalIndent(ss, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}

func (ss SraStudy) JSONString() string {
	json, err := json.MarshalIndent(ss, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func main() {
	data := `
            <?xml version="1.0" encoding="UTF-8"?>
            <STUDY_SET xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
              <STUDY xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" alias="2302745993" center_name="WUGSC" accession="SRP000121">
                <IDENTIFIERS>
                  <PRIMARY_ID>SRP000121</PRIMARY_ID>
                  <SUBMITTER_ID namespace="WUGSC">2302745993</SUBMITTER_ID>
                </IDENTIFIERS>
                <DESCRIPTOR>
                  <STUDY_TITLE>Tupaia belangeri  Transcriptome Study</STUDY_TITLE>
                  <STUDY_TYPE existing_study_type="Transcriptome Analysis"/>
                  <STUDY_ABSTRACT>The tree shrew species Tupaia belangeri, will be sequenced as part of an initiative to expand sequencing of mammalian genomes for the purpose of expanding the annotation of the human genome. The plan is to create a draft sequence assembly from 6X whole genome coverage of a female tree shrew. The tree shrew DNA source was obtained from Dr. Eberhard Fuchs, German Primate Center, Goettingen, Germany. For more information on the mammalian genomes represented in this initiative please visit here. The National Human Genome Research Institute (NHGRI), National Institutes of Health (NIH) is providing funding for the sequence characterization of the tree shrew genome.</STUDY_ABSTRACT>
                  <CENTER_PROJECT_NAME>2302745993</CENTER_PROJECT_NAME>
                  <RELATED_STUDIES>
                    <RELATED_STUDY>
                      <RELATED_LINK>
                        <DB>bioproject</DB>
                        <ID>20273</ID>
                        <LABEL>PRJNA20273</LABEL>
                      </RELATED_LINK>
                      <IS_PRIMARY>true</IS_PRIMARY>
                    </RELATED_STUDY>
                  </RELATED_STUDIES>
                  <STUDY_DESCRIPTION>none provided</STUDY_DESCRIPTION>
                </DESCRIPTOR>
                <STUDY_LINKS>
                  <STUDY_LINK>
                    <URL_LINK>
                      <LABEL>Genome Project at Washington University</LABEL>
                      <URL>http://genome.wustl.edu/genomes/view/tupaia_belangeri/</URL>
                    </URL_LINK>
                  </STUDY_LINK>
                </STUDY_LINKS>
              </STUDY>
            </STUDY_SET>
    `

	var ss SraStudy
	xml.Unmarshal([]byte(data), &ss)

	// Changing an attribute
	studies := ss.TStudySetType.Studies
	studies[0].Descriptor.StudyTitle = "The Most Awesome Run Ever"

	fmt.Printf("%v\n", ss)
	fmt.Printf("%v\n", ss.JSONString())
}
