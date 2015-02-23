package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	sraanalysis "github.com/indraniel/go-sra-schemas-1.5/SRA.analysis.xsd_go"
)

type SraAnalysis struct {
	XMLName xml.Name `xml:"ANALYSIS_SET"`
	sraanalysis.TAnalysisSetType
}

func (sa SraAnalysis) String() string {
	xml, err := xml.MarshalIndent(sa, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(xml)
}

func (sa SraAnalysis) JSONString() string {
	json, err := json.MarshalIndent(sa, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(json)
}

func main() {
	data := `
            <?xml version="1.0" encoding="UTF-8"?>
            <ANALYSIS_SET xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
              <ANALYSIS alias="A-MAYO-MY000493-BL-MAY-1276_dbgap_1" accession="SRZ068470" center_name="WUGSC">
                <IDENTIFIERS>
                  <PRIMARY_ID>SRZ068470</PRIMARY_ID>
                  <SUBMITTER_ID namespace="WUGSC">A-MAYO-MY000493-BL-MAY-1276_dbgap_1</SUBMITTER_ID>
                </IDENTIFIERS>
                <TITLE>Reference alignment of phs000572 sample A-MAYO-MY000493-BL-MAY-1276</TITLE>
                <STUDY_REF accession="SRP028407">
                  <IDENTIFIERS>
                    <PRIMARY_ID>SRP028407</PRIMARY_ID>
                    <EXTERNAL_ID namespace="dbgap">phs000572</EXTERNAL_ID>
                  </IDENTIFIERS>
                </STUDY_REF>
                <DESCRIPTION>GRCh37-lite reference alignment of sample A-MAYO-MY000493-BL-MAY-1276</DESCRIPTION>
                <ANALYSIS_TYPE>
                  <REFERENCE_ALIGNMENT>
                    <ASSEMBLY>
                      <STANDARD short_name="GRCh37-lite"/>
                    </ASSEMBLY>
                    <RUN_LABELS>
                      <RUN accession="SRR1285508" refname="339381" refcenter="WUGSC" read_group_label="2893973296"/>
                      <RUN accession="SRR1285509" refname="339390" refcenter="WUGSC" read_group_label="2893973552"/>
                    </RUN_LABELS>
                    <SEQ_LABELS>
                      <SEQUENCE accession="CM000663.1" seq_label="1"/>
                      <SEQUENCE accession="CM000664.1" seq_label="2"/>
                      <SEQUENCE accession="CM000665.1" seq_label="3"/>
                      <SEQUENCE accession="CM000666.1" seq_label="4"/>
                      <SEQUENCE accession="CM000667.1" seq_label="5"/>
                      <SEQUENCE accession="CM000668.1" seq_label="6"/>
                      <SEQUENCE accession="CM000669.1" seq_label="7"/>
                      <SEQUENCE accession="CM000670.1" seq_label="8"/>
                      <SEQUENCE accession="CM000671.1" seq_label="9"/>
                      <SEQUENCE accession="CM000672.1" seq_label="10"/>
                      <SEQUENCE accession="CM000673.1" seq_label="11"/>
                      <SEQUENCE accession="CM000674.1" seq_label="12"/>
                      <SEQUENCE accession="CM000675.1" seq_label="13"/>
                      <SEQUENCE accession="CM000676.1" seq_label="14"/>
                      <SEQUENCE accession="CM000677.1" seq_label="15"/>
                      <SEQUENCE accession="CM000678.1" seq_label="16"/>
                      <SEQUENCE accession="CM000679.1" seq_label="17"/>
                      <SEQUENCE accession="CM000680.1" seq_label="18"/>
                      <SEQUENCE accession="CM000681.1" seq_label="19"/>
                      <SEQUENCE accession="CM000682.1" seq_label="20"/>
                      <SEQUENCE accession="CM000683.1" seq_label="21"/>
                      <SEQUENCE accession="CM000684.1" seq_label="22"/>
                      <SEQUENCE accession="CM000685.1" seq_label="X"/>
                      <SEQUENCE accession="CM000686.1" seq_label="Y"/>
                      <SEQUENCE accession="J01415.2" seq_label="MT"/>
                      <SEQUENCE accession="GL000207.1" seq_label="GL000207.1"/>
                      <SEQUENCE accession="GL000226.1" seq_label="GL000226.1"/>
                      <SEQUENCE accession="GL000229.1" seq_label="GL000229.1"/>
                      <SEQUENCE accession="GL000231.1" seq_label="GL000231.1"/>
                      <SEQUENCE accession="GL000210.1" seq_label="GL000210.1"/>
                      <SEQUENCE accession="GL000239.1" seq_label="GL000239.1"/>
                      <SEQUENCE accession="GL000235.1" seq_label="GL000235.1"/>
                      <SEQUENCE accession="GL000201.1" seq_label="GL000201.1"/>
                      <SEQUENCE accession="GL000247.1" seq_label="GL000247.1"/>
                      <SEQUENCE accession="GL000245.1" seq_label="GL000245.1"/>
                      <SEQUENCE accession="GL000197.1" seq_label="GL000197.1"/>
                      <SEQUENCE accession="GL000203.1" seq_label="GL000203.1"/>
                      <SEQUENCE accession="GL000246.1" seq_label="GL000246.1"/>
                      <SEQUENCE accession="GL000249.1" seq_label="GL000249.1"/>
                      <SEQUENCE accession="GL000196.1" seq_label="GL000196.1"/>
                      <SEQUENCE accession="GL000248.1" seq_label="GL000248.1"/>
                      <SEQUENCE accession="GL000244.1" seq_label="GL000244.1"/>
                      <SEQUENCE accession="GL000238.1" seq_label="GL000238.1"/>
                      <SEQUENCE accession="GL000202.1" seq_label="GL000202.1"/>
                      <SEQUENCE accession="GL000234.1" seq_label="GL000234.1"/>
                      <SEQUENCE accession="GL000232.1" seq_label="GL000232.1"/>
                      <SEQUENCE accession="GL000206.1" seq_label="GL000206.1"/>
                      <SEQUENCE accession="GL000240.1" seq_label="GL000240.1"/>
                      <SEQUENCE accession="GL000236.1" seq_label="GL000236.1"/>
                      <SEQUENCE accession="GL000241.1" seq_label="GL000241.1"/>
                      <SEQUENCE accession="GL000243.1" seq_label="GL000243.1"/>
                      <SEQUENCE accession="GL000242.1" seq_label="GL000242.1"/>
                      <SEQUENCE accession="GL000230.1" seq_label="GL000230.1"/>
                      <SEQUENCE accession="GL000237.1" seq_label="GL000237.1"/>
                      <SEQUENCE accession="GL000233.1" seq_label="GL000233.1"/>
                      <SEQUENCE accession="GL000204.1" seq_label="GL000204.1"/>
                      <SEQUENCE accession="GL000198.1" seq_label="GL000198.1"/>
                      <SEQUENCE accession="GL000208.1" seq_label="GL000208.1"/>
                      <SEQUENCE accession="GL000191.1" seq_label="GL000191.1"/>
                      <SEQUENCE accession="GL000227.1" seq_label="GL000227.1"/>
                      <SEQUENCE accession="GL000228.1" seq_label="GL000228.1"/>
                      <SEQUENCE accession="GL000214.1" seq_label="GL000214.1"/>
                      <SEQUENCE accession="GL000221.1" seq_label="GL000221.1"/>
                      <SEQUENCE accession="GL000209.1" seq_label="GL000209.1"/>
                      <SEQUENCE accession="GL000218.1" seq_label="GL000218.1"/>
                      <SEQUENCE accession="GL000220.1" seq_label="GL000220.1"/>
                      <SEQUENCE accession="GL000213.1" seq_label="GL000213.1"/>
                      <SEQUENCE accession="GL000211.1" seq_label="GL000211.1"/>
                      <SEQUENCE accession="GL000199.1" seq_label="GL000199.1"/>
                      <SEQUENCE accession="GL000217.1" seq_label="GL000217.1"/>
                      <SEQUENCE accession="GL000216.1" seq_label="GL000216.1"/>
                      <SEQUENCE accession="GL000215.1" seq_label="GL000215.1"/>
                      <SEQUENCE accession="GL000205.1" seq_label="GL000205.1"/>
                      <SEQUENCE accession="GL000219.1" seq_label="GL000219.1"/>
                      <SEQUENCE accession="GL000224.1" seq_label="GL000224.1"/>
                      <SEQUENCE accession="GL000223.1" seq_label="GL000223.1"/>
                      <SEQUENCE accession="GL000195.1" seq_label="GL000195.1"/>
                      <SEQUENCE accession="GL000212.1" seq_label="GL000212.1"/>
                      <SEQUENCE accession="GL000222.1" seq_label="GL000222.1"/>
                      <SEQUENCE accession="GL000200.1" seq_label="GL000200.1"/>
                      <SEQUENCE accession="GL000193.1" seq_label="GL000193.1"/>
                      <SEQUENCE accession="GL000194.1" seq_label="GL000194.1"/>
                      <SEQUENCE accession="GL000225.1" seq_label="GL000225.1"/>
                      <SEQUENCE accession="GL000192.1" seq_label="GL000192.1"/>
                    </SEQ_LABELS>
                    <PROCESSING>
                      <PIPELINE>
                        <PIPE_SECTION>
                          <STEP_INDEX>2893973296</STEP_INDEX>
                          <PREV_STEP_INDEX>NULL</PREV_STEP_INDEX>
                          <PROGRAM>bwa</PROGRAM>
                          <VERSION>0.5.9</VERSION>
                          <NOTES>bwa aln -t4 -q 5; bwa sampe  -a 600 </NOTES>
                        </PIPE_SECTION>
                        <PIPE_SECTION>
                          <STEP_INDEX>2893973552</STEP_INDEX>
                          <PREV_STEP_INDEX>NULL</PREV_STEP_INDEX>
                          <PROGRAM>bwa</PROGRAM>
                          <VERSION>0.5.9</VERSION>
                          <NOTES>bwa aln -t4 -q 5; bwa sampe  -a 600 </NOTES>
                        </PIPE_SECTION>
                        <PIPE_SECTION>
                          <STEP_INDEX>GATK IndelRealigner</STEP_INDEX>
                          <PREV_STEP_INDEX>NULL</PREV_STEP_INDEX>
                          <PROGRAM>knownAlleles=[(RodBinding</PROGRAM>
                          <VERSION>f6dac2d</VERSION>
                          <NOTES>knownAlleles=[(RodBinding name=knownAlleles source=/tmp/ZnQsVFpkYs/139321835.vcf)] targetIntervals=/gscmnt/gc4000/info/model_data/gatk/indel_realigner-blade13-3-4.gsc.wustl.edu-wnash-11686-50e1a4cdc14547b3a104579920c9a87a/b5afb0632fcf40c2ae11798dca6ce8e8.bam.intervals LODThresholdForCleaning=5.0 consensusDeterminationModel=USE_READS entropyThreshold=0.15 maxReadsInMemory=150000 maxIsizeForMovement=3000 maxPositionalMoveAllowed=200 maxConsensuses=30 maxReadsForConsensuses=120 maxReadsForRealignment=20000 noOriginalAlignmentTags=false nWayOut=null generate_nWayOut_md5s=false check_early=false noPGTag=false keepPGTags=false indelsFileForDebugging=null statisticsFileForDebugging=null SNPsFileForDebugging=null</NOTES>
                        </PIPE_SECTION>
                        <PIPE_SECTION>
                          <STEP_INDEX>GATK PrintReads</STEP_INDEX>
                          <PREV_STEP_INDEX>NULL</PREV_STEP_INDEX>
                          <PROGRAM>readGroup=null</PROGRAM>
                          <VERSION>f6dac2d</VERSION>
                          <NOTES>readGroup=null platform=null number=-1 downsample_coverage=1.0 sample_file=[] sample_name=[] simplify=false no_pg_tag=false</NOTES>
                        </PIPE_SECTION>
                        <PIPE_SECTION>
                          <STEP_INDEX>GATK PrintReads</STEP_INDEX>
                          <PREV_STEP_INDEX>NULL</PREV_STEP_INDEX>
                          <PROGRAM>readGroup=null</PROGRAM>
                          <VERSION>f6dac2d</VERSION>
                          <NOTES>readGroup=null platform=null number=-1 downsample_coverage=1.0 sample_file=[] sample_name=[] simplify=false no_pg_tag=false</NOTES>
                        </PIPE_SECTION>
                      </PIPELINE>
                      <DIRECTIVES>
                        <alignment_includes_unaligned_reads>true</alignment_includes_unaligned_reads>
                        <alignment_marks_duplicate_reads>true</alignment_marks_duplicate_reads>
                        <alignment_includes_failed_reads>true</alignment_includes_failed_reads>
                      </DIRECTIVES>
                    </PROCESSING>
                  </REFERENCE_ALIGNMENT>
                </ANALYSIS_TYPE>
                <TARGETS>
                  <IDENTIFIERS>
                    <EXTERNAL_ID namespace="phs000572">A-MAYO-MY000493-BL-MAY-1276</EXTERNAL_ID>
                  </IDENTIFIERS>
                </TARGETS>
                <DATA_BLOCK>
                  <FILES>
                    <FILE filename="a7e5240bad40431e93b3995b6780771f.bam" filetype="bam" checksum_method="MD5" checksum="db52c4e9c93a39552dc52377fa27bf94"/>
                  </FILES>
                </DATA_BLOCK>
              </ANALYSIS>
            </ANALYSIS_SET>
    `

	var sa SraAnalysis
	xml.Unmarshal([]byte(data), &sa)

	// Changing an attribute
	analyses := sa.TAnalysisSetType.Analysises
	analyses[0].Title = "The Most Awesome Analysis Ever"

	fmt.Printf("%v\n", sa)
	fmt.Printf("%v\n", sa.JSONString())
}
