package dicomuid_test

import (
	"testing"

	"github.com/grailbio/go-dicom/dicomuid"
)

func TestStandardUIDs(t *testing.T) {
	if dicomuid.PatientRootQRFind != "1.2.840.10008.5.1.4.1.2.1.1" {
		t.Error(dicomuid.PatientRootQRFind)
	}
}

func TestLookupUID(t *testing.T) {
	u := dicomuid.MustLookup("1.2.840.10008.15.0.4.8")
	if u.Name != "dicomTransferCapability" {
		t.Error(u)
	}
	if u.Type != "LDAP OID" {
		t.Error(u)
	}
}
