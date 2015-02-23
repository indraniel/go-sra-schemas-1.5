package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	srasubmission "github.com/indraniel/go-sra-schemas-1.5/SRA.submission.xsd_go"
)

type SraSubmission struct {
	XMLName xml.Name `xml:"SUBMISSION"`
	srasubmission.TSubmissionType
}

func (ss SraSubmission) String() string {
	xml, err := xml.MarshalIndent(ss, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}

func (ss SraSubmission) JSONString() string {
	json, err := json.MarshalIndent(ss, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func main() {
	data := `
            <?xml version="1.0" encoding="UTF-8"?>
            <SUBMISSION alias="72755" center_name="WUGSC" lab_name="" accession="SRA164053">
              <IDENTIFIERS>
                <PRIMARY_ID>SRA164053</PRIMARY_ID>
                <SUBMITTER_ID namespace="WUGSC">72755</SUBMITTER_ID>
              </IDENTIFIERS>
              <CONTACTS>
                <CONTACT name="LIMS" inform_on_status="mailto:lims@genome.wustl.edu" inform_on_error="mailto:lims@genome.wustl.edu"/>
              </CONTACTS>
              <ACTIONS>
                <ACTION>
                  <ADD source="339369.experiment.xml" schema="experiment"/>
                </ACTION>
                <ACTION>
                  <ADD source="A-MAYO-MY000493-BL-MAY-1276_dbgap_1.analysis.xml" schema="analysis"/>
                </ACTION>
                <ACTION>
                  <ADD source="339381.run.xml" schema="run"/>
                </ACTION>
                <ACTION>
                  <ADD source="339390.run.xml" schema="run"/>
                </ACTION>
                <ACTION>
                  <PROTECT/>
                </ACTION>
                <ACTION>
                  <RELEASE/>
                </ACTION>
              </ACTIONS>
            </SUBMISSION>
    `

	var ss SraSubmission
	xml.Unmarshal([]byte(data), &ss)

	// Changing an attribute
	submission := ss.TSubmissionType
	submission.Contacts.Contacts[0].Name = "Dr. Who"

	fmt.Printf("%v\n", ss)
	fmt.Printf("%v\n", ss.JSONString())
}
