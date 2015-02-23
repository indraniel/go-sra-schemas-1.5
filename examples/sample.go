package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	srasample "github.com/indraniel/go-sra-schemas-1.5/SRA.sample.xsd_go"
)

type SraSample struct {
	XMLName xml.Name `xml:"SAMPLE_SET"`
	srasample.TSampleSetType
}

func (ss SraSample) String() string {
	xml, err := xml.MarshalIndent(ss, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}

func (ss SraSample) JSONString() string {
	json, err := json.MarshalIndent(ss, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func main() {
	data := `
            <?xml version="1.0" encoding="UTF-8"?>
            <SAMPLE_SET xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
              <SAMPLE center_name="phs000276" alias="H_HY-01125.wxs" accession="SRS385983">
                <IDENTIFIERS>
                  <PRIMARY_ID>SRS385983</PRIMARY_ID>
                  <EXTERNAL_ID namespace="BioSample">SAMN01877218</EXTERNAL_ID>
                  <EXTERNAL_ID namespace="dbGaP" label="Sample name">276-H_HY-01125.wxs</EXTERNAL_ID>
                  <EXTERNAL_ID namespace="phs000276" label="submitted sample id">H_HY-01125.wxs</EXTERNAL_ID>
                </IDENTIFIERS>
                <TITLE>DNA sample from a human female participant in the dbGaP study "STAMPEED: Northern Finland Birth Cohort 1966 (NFBC1966)"</TITLE>
                <SAMPLE_NAME>
                  <TAXON_ID>9606</TAXON_ID>
                  <SCIENTIFIC_NAME>Homo sapiens</SCIENTIFIC_NAME>
                </SAMPLE_NAME>
                <SAMPLE_LINKS>
                  <SAMPLE_LINK>
                    <XREF_LINK>
                      <DB>biosample</DB>
                      <ID>1877218</ID>
                      <LABEL>SAMN01877218</LABEL>
                    </XREF_LINK>
                  </SAMPLE_LINK>
                </SAMPLE_LINKS>
                <SAMPLE_ATTRIBUTES>
                  <SAMPLE_ATTRIBUTE>
                    <TAG>gap_accession</TAG>
                    <VALUE>phs000276</VALUE>
                  </SAMPLE_ATTRIBUTE>
                  <SAMPLE_ATTRIBUTE>
                    <TAG>submitter handle</TAG>
                    <VALUE>STAMPEED_NFBC</VALUE>
                  </SAMPLE_ATTRIBUTE>
                  <SAMPLE_ATTRIBUTE>
                    <TAG>biospecimen repository</TAG>
                    <VALUE>STAMPEED_NFBC</VALUE>
                  </SAMPLE_ATTRIBUTE>
                  <SAMPLE_ATTRIBUTE>
                    <TAG>study name</TAG>
                    <VALUE>STAMPEED: Northern Finland Birth Cohort 1966 (NFBC1966)</VALUE>
                  </SAMPLE_ATTRIBUTE>
                  <SAMPLE_ATTRIBUTE>
                    <TAG>study design</TAG>
                    <VALUE>Prospective Longitudinal Cohort</VALUE>
                  </SAMPLE_ATTRIBUTE>
                  <SAMPLE_ATTRIBUTE>
                    <TAG>biospecimen repository sample id</TAG>
                    <VALUE>H_HY-01125.wxs</VALUE>
                  </SAMPLE_ATTRIBUTE>
                  <SAMPLE_ATTRIBUTE>
                    <TAG>submitted sample id</TAG>
                    <VALUE>H_HY-01125.wxs</VALUE>
                  </SAMPLE_ATTRIBUTE>
                  <SAMPLE_ATTRIBUTE>
                    <TAG>submitted subject id</TAG>
                    <VALUE>2740</VALUE>
                  </SAMPLE_ATTRIBUTE>
                  <SAMPLE_ATTRIBUTE>
                    <TAG>gap_sample_id</TAG>
                    <VALUE>896255</VALUE>
                  </SAMPLE_ATTRIBUTE>
                  <SAMPLE_ATTRIBUTE>
                    <TAG>gap_subject_id</TAG>
                    <VALUE>358832</VALUE>
                  </SAMPLE_ATTRIBUTE>
                  <SAMPLE_ATTRIBUTE>
                    <TAG>sex</TAG>
                    <VALUE>female</VALUE>
                  </SAMPLE_ATTRIBUTE>
                  <SAMPLE_ATTRIBUTE>
                    <TAG>analyte type</TAG>
                    <VALUE>DNA</VALUE>
                  </SAMPLE_ATTRIBUTE>
                  <SAMPLE_ATTRIBUTE>
                    <TAG>subject is affected</TAG>
                    <VALUE>No</VALUE>
                  </SAMPLE_ATTRIBUTE>
                  <SAMPLE_ATTRIBUTE>
                    <TAG>molecular data type</TAG>
                    <VALUE>Whole Exome (NGS)</VALUE>
                  </SAMPLE_ATTRIBUTE>
                  <SAMPLE_ATTRIBUTE>
                    <TAG>gap_consent_code</TAG>
                    <VALUE>1</VALUE>
                  </SAMPLE_ATTRIBUTE>
                  <SAMPLE_ATTRIBUTE>
                    <TAG>gap_consent_short_name</TAG>
                    <VALUE>GRU</VALUE>
                  </SAMPLE_ATTRIBUTE>
                  <SAMPLE_ATTRIBUTE>
                    <TAG>BioSampleModel</TAG>
                    <VALUE>Generic</VALUE>
                  </SAMPLE_ATTRIBUTE>
                </SAMPLE_ATTRIBUTES>
              </SAMPLE>
            </SAMPLE_SET>
        `

	var ss SraSample
	xml.Unmarshal([]byte(data), &ss)

	// Changing an attribute
	samples := ss.TSampleSetType.Samples
	samples[0].Title = "The Most Awesome Sample Ever"

	fmt.Printf("%v\n", ss)
	fmt.Printf("%v\n", ss.JSONString())
}
