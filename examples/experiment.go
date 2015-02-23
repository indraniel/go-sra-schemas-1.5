package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	sraexp "github.com/indraniel/go-sra-schemas-1.5/SRA.experiment.xsd_go"
)

type SraExp struct {
	XMLName xml.Name `xml:"EXPERIMENT_SET"`
	sraexp.TExperimentSetType
}

func (se SraExp) String() string {
	xml, err := xml.MarshalIndent(se, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}

func (se SraExp) JSONString() string {
	json, err := json.MarshalIndent(se, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func main() {
	data := `
            <?xml version="1.0" encoding="UTF-8"?>
            <EXPERIMENT_SET xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
              <EXPERIMENT alias="150935" center_name="WUGSC" accession="SRX216593">
                <IDENTIFIERS>
                  <PRIMARY_ID>SRX216593</PRIMARY_ID>
                  <SUBMITTER_ID namespace="WUGSC">150935</SUBMITTER_ID>
                </IDENTIFIERS>
                <TITLE/>
                <STUDY_REF accession="SRP017924">
                  <IDENTIFIERS>
                    <PRIMARY_ID>SRP017924</PRIMARY_ID>
                    <EXTERNAL_ID namespace="dbgap">phs000276</EXTERNAL_ID>
                  </IDENTIFIERS>
                </STUDY_REF>
                <DESIGN>
                  <DESIGN_DESCRIPTION>Illumina sequencing of phs000276 2725296873 paired end WXS library</DESIGN_DESCRIPTION>
                  <SAMPLE_DESCRIPTOR accession="SRS385489">
                    <IDENTIFIERS>
                      <PRIMARY_ID>SRS385489</PRIMARY_ID>
                      <EXTERNAL_ID namespace="phs000276">H_HY-01703.wxs</EXTERNAL_ID>
                    </IDENTIFIERS>
                  </SAMPLE_DESCRIPTOR>
                  <LIBRARY_DESCRIPTOR>
                    <LIBRARY_NAME>2874304214</LIBRARY_NAME>
                    <LIBRARY_STRATEGY>WXS</LIBRARY_STRATEGY>
                    <LIBRARY_SOURCE>GENOMIC</LIBRARY_SOURCE>
                    <LIBRARY_SELECTION>unspecified</LIBRARY_SELECTION>
                    <LIBRARY_LAYOUT>
                      <PAIRED NOMINAL_LENGTH="338" NOMINAL_SDEV="0.0E0"/>
                    </LIBRARY_LAYOUT>
                  </LIBRARY_DESCRIPTOR>
                  <SPOT_DESCRIPTOR>
                    <SPOT_DECODE_SPEC>
                      <READ_SPEC>
                        <READ_INDEX>0</READ_INDEX>
                        <READ_CLASS>Application Read</READ_CLASS>
                        <READ_TYPE>Forward</READ_TYPE>
                        <BASE_COORD>1</BASE_COORD>
                      </READ_SPEC>
                      <READ_SPEC>
                        <READ_INDEX>1</READ_INDEX>
                        <READ_CLASS>Application Read</READ_CLASS>
                        <READ_TYPE>Reverse</READ_TYPE>
                        <BASE_COORD>101</BASE_COORD>
                      </READ_SPEC>
                    </SPOT_DECODE_SPEC>
                  </SPOT_DESCRIPTOR>
                </DESIGN>
                <PLATFORM>
                  <ILLUMINA>
                    <INSTRUMENT_MODEL>Illumina HiSeq 2000</INSTRUMENT_MODEL>
                  </ILLUMINA>
                </PLATFORM>
                <PROCESSING/>
              </EXPERIMENT>
            </EXPERIMENT_SET>
    `

	var se SraExp
	xml.Unmarshal([]byte(data), &se)

	// Changing an attribute
	exps := se.TExperimentSetType.Experiments
	exps[0].Title = "The Most Awesome Experiment Ever"

	fmt.Printf("%v\n", se)
	fmt.Printf("%v\n", se.JSONString())
}
