package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	srarun "github.com/indraniel/go-sra-schemas-1.5/SRA.run.xsd_go"
)

type SraRun struct {
	XMLName xml.Name `xml:"RUN_SET"`
	srarun.TxsdRunSet
}

func (sr SraRun) String() string {
	xml, err := xml.MarshalIndent(sr, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}

func (sr SraRun) JSONString() string {
	json, err := json.MarshalIndent(sr, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func main() {
	data := `
        <?xml version="1.0" encoding="UTF-8"?>
        <RUN_SET xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
          <RUN alias="150949" center_name="WUGSC" run_date="2011-10-18T02:37:00Z" run_center="WUGSC" accession="SRR648183">
            <IDENTIFIERS>
              <PRIMARY_ID>SRR648183</PRIMARY_ID>
              <SECONDARY_ID>SRZ031726</SECONDARY_ID>
              <SUBMITTER_ID namespace="WUGSC">150949</SUBMITTER_ID>
            </IDENTIFIERS>
            <EXPERIMENT_REF accession="SRX216593" refname="150935" refcenter="WUGSC">
              <IDENTIFIERS>
                <PRIMARY_ID>SRX216593</PRIMARY_ID>
                <SUBMITTER_ID namespace="WUGSC">150935</SUBMITTER_ID>
              </IDENTIFIERS>
            </EXPERIMENT_REF>
            <DATA_BLOCK>
              <FILES>
                <FILE filename="116008179.bam" filetype="bam" checksum="92651eb2a910f69d9b7e5aaeab06f278" checksum_method="MD5"/>
              </FILES>
            </DATA_BLOCK>
            <RUN_ATTRIBUTES>
              <RUN_ATTRIBUTE>
                <TAG>assembly</TAG>
                <VALUE>GRCh37-lite</VALUE>
              </RUN_ATTRIBUTE>
            </RUN_ATTRIBUTES>
          </RUN>
        </RUN_SET>
    `

	var sr SraRun
	xml.Unmarshal([]byte(data), &sr)

	// Changing an attribute
	runs := sr.TxsdRunSet.Runs
	runs[0].Title = "The Most Awesome Run Ever"

	fmt.Printf("%v\n", sr)
	fmt.Printf("%v\n", sr.JSONString())
}
