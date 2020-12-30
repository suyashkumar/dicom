package tag

// Code generated from './pkg/tag/codegen'. DO NOT EDIT.

var tagDict = make(map[Tag]Info)
var tagDictByKeyword = make(map[string]Info)

func initTagDicts() {
	var thisInfo Info

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x0000},
		Name: "CommandGroupLength",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x0002},
		Name: "AffectedSOPClassUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x0003},
		Name: "RequestedSOPClassUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x0100},
		Name: "CommandField",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x0110},
		Name: "MessageID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x0120},
		Name: "MessageIDBeingRespondedTo",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x0600},
		Name: "MoveDestination",
		VR:   "AE",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x0700},
		Name: "Priority",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x0800},
		Name: "CommandDataSetType",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x0900},
		Name: "Status",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x0901},
		Name: "OffendingElement",
		VR:   "AT",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x0902},
		Name: "ErrorComment",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x0903},
		Name: "ErrorID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x1000},
		Name: "AffectedSOPInstanceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x1001},
		Name: "RequestedSOPInstanceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x1002},
		Name: "EventTypeID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x1005},
		Name: "AttributeIdentifierList",
		VR:   "AT",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x1008},
		Name: "ActionTypeID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x1020},
		Name: "NumberOfRemainingSuboperations",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x1021},
		Name: "NumberOfCompletedSuboperations",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x1022},
		Name: "NumberOfFailedSuboperations",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x1023},
		Name: "NumberOfWarningSuboperations",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x1030},
		Name: "MoveOriginatorApplicationEntityTitle",
		VR:   "AE",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0000, 0x1031},
		Name: "MoveOriginatorMessageID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0005},
		Name: "SpecificCharacterSet",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0006},
		Name: "LanguageCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0008},
		Name: "ImageType",
		VR:   "CS",
		VM:   "2-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0012},
		Name: "InstanceCreationDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0013},
		Name: "InstanceCreationTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0014},
		Name: "InstanceCreatorUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0015},
		Name: "InstanceCoercionDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0016},
		Name: "SOPClassUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0018},
		Name: "SOPInstanceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x001a},
		Name: "RelatedGeneralSOPClassUID",
		VR:   "UI",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x001b},
		Name: "OriginalSpecializedSOPClassUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0020},
		Name: "StudyDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0021},
		Name: "SeriesDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0022},
		Name: "AcquisitionDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0023},
		Name: "ContentDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x002a},
		Name: "AcquisitionDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0030},
		Name: "StudyTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0031},
		Name: "SeriesTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0032},
		Name: "AcquisitionTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0033},
		Name: "ContentTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0050},
		Name: "AccessionNumber",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0051},
		Name: "IssuerOfAccessionNumberSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0052},
		Name: "QueryRetrieveLevel",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0053},
		Name: "QueryRetrieveView",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0054},
		Name: "RetrieveAETitle",
		VR:   "AE",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0055},
		Name: "StationAETitle",
		VR:   "AE",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0056},
		Name: "InstanceAvailability",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0058},
		Name: "FailedSOPInstanceUIDList",
		VR:   "UI",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0060},
		Name: "Modality",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0061},
		Name: "ModalitiesInStudy",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0062},
		Name: "SOPClassesInStudy",
		VR:   "UI",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0063},
		Name: "AnatomicRegionsInStudyCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0064},
		Name: "ConversionType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0068},
		Name: "PresentationIntentType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0070},
		Name: "Manufacturer",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0080},
		Name: "InstitutionName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0081},
		Name: "InstitutionAddress",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0082},
		Name: "InstitutionCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0090},
		Name: "ReferringPhysicianName",
		VR:   "PN",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0092},
		Name: "ReferringPhysicianAddress",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0094},
		Name: "ReferringPhysicianTelephoneNumbers",
		VR:   "SH",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0096},
		Name: "ReferringPhysicianIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x009c},
		Name: "ConsultingPhysicianName",
		VR:   "PN",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x009d},
		Name: "ConsultingPhysicianIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0100},
		Name: "CodeValue",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0101},
		Name: "ExtendedCodeValue",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0102},
		Name: "CodingSchemeDesignator",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0103},
		Name: "CodingSchemeVersion",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0104},
		Name: "CodeMeaning",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0105},
		Name: "MappingResource",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0106},
		Name: "ContextGroupVersion",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0107},
		Name: "ContextGroupLocalVersion",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0108},
		Name: "ExtendedCodeMeaning",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0109},
		Name: "CodingSchemeResourcesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x010a},
		Name: "CodingSchemeURLType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x010b},
		Name: "ContextGroupExtensionFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x010c},
		Name: "CodingSchemeUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x010d},
		Name: "ContextGroupExtensionCreatorUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x010e},
		Name: "CodingSchemeURL",
		VR:   "UR",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x010f},
		Name: "ContextIdentifier",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0110},
		Name: "CodingSchemeIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0112},
		Name: "CodingSchemeRegistry",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0114},
		Name: "CodingSchemeExternalID",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0115},
		Name: "CodingSchemeName",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0116},
		Name: "CodingSchemeResponsibleOrganization",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0117},
		Name: "ContextUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0118},
		Name: "MappingResourceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0119},
		Name: "LongCodeValue",
		VR:   "UC",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0120},
		Name: "URNCodeValue",
		VR:   "UR",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0121},
		Name: "EquivalentCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0122},
		Name: "MappingResourceName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0123},
		Name: "ContextGroupIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0124},
		Name: "MappingResourceIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0201},
		Name: "TimezoneOffsetFromUTC",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0220},
		Name: "ResponsibleGroupCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0221},
		Name: "EquipmentModality",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0222},
		Name: "ManufacturerRelatedModelGroup",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0300},
		Name: "PrivateDataElementCharacteristicsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0301},
		Name: "PrivateGroupReference",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0302},
		Name: "PrivateCreatorReference",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0303},
		Name: "BlockIdentifyingInformationStatus",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0304},
		Name: "NonidentifyingPrivateElements",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0306},
		Name: "IdentifyingPrivateElements",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0305},
		Name: "DeidentificationActionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0307},
		Name: "DeidentificationAction",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0308},
		Name: "PrivateDataElement",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0309},
		Name: "PrivateDataElementValueMultiplicity",
		VR:   "UL",
		VM:   "1-3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x030a},
		Name: "PrivateDataElementValueRepresentation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x030b},
		Name: "PrivateDataElementNumberOfItems",
		VR:   "UL",
		VM:   "1-2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x030c},
		Name: "PrivateDataElementName",
		VR:   "UC",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x030d},
		Name: "PrivateDataElementKeyword",
		VR:   "UC",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x030e},
		Name: "PrivateDataElementDescription",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x030f},
		Name: "PrivateDataElementEncoding",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x0310},
		Name: "PrivateDataElementDefinitionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1010},
		Name: "StationName",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1030},
		Name: "StudyDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1032},
		Name: "ProcedureCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x103e},
		Name: "SeriesDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x103f},
		Name: "SeriesDescriptionCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1040},
		Name: "InstitutionalDepartmentName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1041},
		Name: "InstitutionalDepartmentTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1048},
		Name: "PhysiciansOfRecord",
		VR:   "PN",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1049},
		Name: "PhysiciansOfRecordIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1050},
		Name: "PerformingPhysicianName",
		VR:   "PN",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1052},
		Name: "PerformingPhysicianIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1060},
		Name: "NameOfPhysiciansReadingStudy",
		VR:   "PN",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1062},
		Name: "PhysiciansReadingStudyIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1070},
		Name: "OperatorsName",
		VR:   "PN",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1072},
		Name: "OperatorIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1080},
		Name: "AdmittingDiagnosesDescription",
		VR:   "LO",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1084},
		Name: "AdmittingDiagnosesCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1090},
		Name: "ManufacturerModelName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1110},
		Name: "ReferencedStudySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1111},
		Name: "ReferencedPerformedProcedureStepSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1115},
		Name: "ReferencedSeriesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1120},
		Name: "ReferencedPatientSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1125},
		Name: "ReferencedVisitSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1134},
		Name: "ReferencedStereometricInstanceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x113a},
		Name: "ReferencedWaveformSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1140},
		Name: "ReferencedImageSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x114a},
		Name: "ReferencedInstanceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x114b},
		Name: "ReferencedRealWorldValueMappingInstanceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1150},
		Name: "ReferencedSOPClassUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1155},
		Name: "ReferencedSOPInstanceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1156},
		Name: "DefinitionSourceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x115a},
		Name: "SOPClassesSupported",
		VR:   "UI",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1160},
		Name: "ReferencedFrameNumber",
		VR:   "IS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1161},
		Name: "SimpleFrameList",
		VR:   "UL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1162},
		Name: "CalculatedFrameList",
		VR:   "UL",
		VM:   "3-3n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1163},
		Name: "TimeRange",
		VR:   "FD",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1164},
		Name: "FrameExtractionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1167},
		Name: "MultiFrameSourceSOPInstanceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1190},
		Name: "RetrieveURL",
		VR:   "UR",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1195},
		Name: "TransactionUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1196},
		Name: "WarningReason",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1197},
		Name: "FailureReason",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1198},
		Name: "FailedSOPSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1199},
		Name: "ReferencedSOPSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x119a},
		Name: "OtherFailuresSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1200},
		Name: "StudiesContainingOtherReferencedInstancesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x1250},
		Name: "RelatedSeriesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x2111},
		Name: "DerivationDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x2112},
		Name: "SourceImageSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x2120},
		Name: "StageName",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x2122},
		Name: "StageNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x2124},
		Name: "NumberOfStages",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x2127},
		Name: "ViewName",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x2128},
		Name: "ViewNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x2129},
		Name: "NumberOfEventTimers",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x212a},
		Name: "NumberOfViewsInStage",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x2130},
		Name: "EventElapsedTimes",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x2132},
		Name: "EventTimerNames",
		VR:   "LO",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x2133},
		Name: "EventTimerSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x2134},
		Name: "EventTimeOffset",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x2135},
		Name: "EventCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x2142},
		Name: "StartTrim",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x2143},
		Name: "StopTrim",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x2144},
		Name: "RecommendedDisplayFrameRate",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x2218},
		Name: "AnatomicRegionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x2220},
		Name: "AnatomicRegionModifierSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x2228},
		Name: "PrimaryAnatomicStructureSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x2230},
		Name: "PrimaryAnatomicStructureModifierSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x3001},
		Name: "AlternateRepresentationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x3002},
		Name: "AvailableTransferSyntaxUID",
		VR:   "UI",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x3010},
		Name: "IrradiationEventUID",
		VR:   "UI",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x3011},
		Name: "SourceIrradiationEventSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x3012},
		Name: "RadiopharmaceuticalAdministrationEventUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x9007},
		Name: "FrameType",
		VR:   "CS",
		VM:   "4",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x9092},
		Name: "ReferencedImageEvidenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x9121},
		Name: "ReferencedRawDataSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x9123},
		Name: "CreatorVersionUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x9124},
		Name: "DerivationImageSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x9154},
		Name: "SourceImageEvidenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x9205},
		Name: "PixelPresentation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x9206},
		Name: "VolumetricProperties",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x9207},
		Name: "VolumeBasedCalculationTechnique",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x9208},
		Name: "ComplexImageComponent",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x9209},
		Name: "AcquisitionContrast",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x9215},
		Name: "DerivationCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x9237},
		Name: "ReferencedPresentationStateSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x9410},
		Name: "ReferencedOtherPlaneSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x9458},
		Name: "FrameDisplaySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x9459},
		Name: "RecommendedDisplayFrameRateInFloat",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0008, 0x9460},
		Name: "SkipFrameRangeFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0010},
		Name: "PatientName",
		VR:   "PN",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0020},
		Name: "PatientID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0021},
		Name: "IssuerOfPatientID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0022},
		Name: "TypeOfPatientID",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0024},
		Name: "IssuerOfPatientIDQualifiersSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0026},
		Name: "SourcePatientGroupIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0027},
		Name: "GroupOfPatientsIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0028},
		Name: "SubjectRelativePositionInImage",
		VR:   "US",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0030},
		Name: "PatientBirthDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0032},
		Name: "PatientBirthTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0033},
		Name: "PatientBirthDateInAlternativeCalendar",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0034},
		Name: "PatientDeathDateInAlternativeCalendar",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0035},
		Name: "PatientAlternativeCalendar",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0040},
		Name: "PatientSex",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0050},
		Name: "PatientInsurancePlanCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0101},
		Name: "PatientPrimaryLanguageCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0102},
		Name: "PatientPrimaryLanguageModifierCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0200},
		Name: "QualityControlSubject",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0201},
		Name: "QualityControlSubjectTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0212},
		Name: "StrainDescription",
		VR:   "UC",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0213},
		Name: "StrainNomenclature",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0214},
		Name: "StrainStockNumber",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0215},
		Name: "StrainSourceRegistryCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0216},
		Name: "StrainStockSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0217},
		Name: "StrainSource",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0218},
		Name: "StrainAdditionalInformation",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0219},
		Name: "StrainCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0221},
		Name: "GeneticModificationsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0222},
		Name: "GeneticModificationsDescription",
		VR:   "UC",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0223},
		Name: "GeneticModificationsNomenclature",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x0229},
		Name: "GeneticModificationsCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x1001},
		Name: "OtherPatientNames",
		VR:   "PN",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x1002},
		Name: "OtherPatientIDsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x1005},
		Name: "PatientBirthName",
		VR:   "PN",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x1010},
		Name: "PatientAge",
		VR:   "AS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x1020},
		Name: "PatientSize",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x1021},
		Name: "PatientSizeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x1022},
		Name: "PatientBodyMassIndex",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x1023},
		Name: "MeasuredAPDimension",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x1024},
		Name: "MeasuredLateralDimension",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x1030},
		Name: "PatientWeight",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x1040},
		Name: "PatientAddress",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x1060},
		Name: "PatientMotherBirthName",
		VR:   "PN",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x1080},
		Name: "MilitaryRank",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x1081},
		Name: "BranchOfService",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x1100},
		Name: "ReferencedPatientPhotoSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x2000},
		Name: "MedicalAlerts",
		VR:   "LO",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x2110},
		Name: "Allergies",
		VR:   "LO",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x2150},
		Name: "CountryOfResidence",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x2152},
		Name: "RegionOfResidence",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x2154},
		Name: "PatientTelephoneNumbers",
		VR:   "SH",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x2155},
		Name: "PatientTelecomInformation",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x2160},
		Name: "EthnicGroup",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x2180},
		Name: "Occupation",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x21a0},
		Name: "SmokingStatus",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x21b0},
		Name: "AdditionalPatientHistory",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x21c0},
		Name: "PregnancyStatus",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x21d0},
		Name: "LastMenstrualDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x21f0},
		Name: "PatientReligiousPreference",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x2201},
		Name: "PatientSpeciesDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x2202},
		Name: "PatientSpeciesCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x2203},
		Name: "PatientSexNeutered",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x2210},
		Name: "AnatomicalOrientationType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x2292},
		Name: "PatientBreedDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x2293},
		Name: "PatientBreedCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x2294},
		Name: "BreedRegistrationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x2295},
		Name: "BreedRegistrationNumber",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x2296},
		Name: "BreedRegistryCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x2297},
		Name: "ResponsiblePerson",
		VR:   "PN",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x2298},
		Name: "ResponsiblePersonRole",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x2299},
		Name: "ResponsibleOrganization",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x4000},
		Name: "PatientComments",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0010, 0x9431},
		Name: "ExaminedBodyThickness",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0010},
		Name: "ClinicalTrialSponsorName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0020},
		Name: "ClinicalTrialProtocolID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0021},
		Name: "ClinicalTrialProtocolName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0030},
		Name: "ClinicalTrialSiteID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0031},
		Name: "ClinicalTrialSiteName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0040},
		Name: "ClinicalTrialSubjectID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0042},
		Name: "ClinicalTrialSubjectReadingID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0050},
		Name: "ClinicalTrialTimePointID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0051},
		Name: "ClinicalTrialTimePointDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0052},
		Name: "LongitudinalTemporalOffsetFromEvent",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0053},
		Name: "LongitudinalTemporalEventType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0060},
		Name: "ClinicalTrialCoordinatingCenterName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0062},
		Name: "PatientIdentityRemoved",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0063},
		Name: "DeidentificationMethod",
		VR:   "LO",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0064},
		Name: "DeidentificationMethodCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0071},
		Name: "ClinicalTrialSeriesID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0072},
		Name: "ClinicalTrialSeriesDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0081},
		Name: "ClinicalTrialProtocolEthicsCommitteeName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0082},
		Name: "ClinicalTrialProtocolEthicsCommitteeApprovalNumber",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0083},
		Name: "ConsentForClinicalTrialUseSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0084},
		Name: "DistributionType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0085},
		Name: "ConsentForDistributionFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0086},
		Name: "EthicsCommitteeApprovalEffectivenessStartDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0012, 0x0087},
		Name: "EthicsCommitteeApprovalEffectivenessEndDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x0025},
		Name: "ComponentManufacturingProcedure",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x0028},
		Name: "ComponentManufacturer",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x0030},
		Name: "MaterialThickness",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x0032},
		Name: "MaterialPipeDiameter",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x0034},
		Name: "MaterialIsolationDiameter",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x0042},
		Name: "MaterialGrade",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x0044},
		Name: "MaterialPropertiesDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x0046},
		Name: "MaterialNotes",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x0050},
		Name: "ComponentShape",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x0052},
		Name: "CurvatureType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x0054},
		Name: "OuterDiameter",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x0056},
		Name: "InnerDiameter",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x0100},
		Name: "ComponentWelderIDs",
		VR:   "LO",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x0101},
		Name: "SecondaryApprovalStatus",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x0102},
		Name: "SecondaryReviewDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x0103},
		Name: "SecondaryReviewTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x0104},
		Name: "SecondaryReviewerName",
		VR:   "PN",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x0105},
		Name: "RepairID",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x0106},
		Name: "MultipleComponentApprovalSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x0107},
		Name: "OtherApprovalStatus",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x0108},
		Name: "OtherSecondaryApprovalStatus",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x1010},
		Name: "ActualEnvironmentalConditions",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x1020},
		Name: "ExpiryDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x1040},
		Name: "EnvironmentalConditions",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x2002},
		Name: "EvaluatorSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x2004},
		Name: "EvaluatorNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x2006},
		Name: "EvaluatorName",
		VR:   "PN",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x2008},
		Name: "EvaluationAttempt",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x2012},
		Name: "IndicationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x2014},
		Name: "IndicationNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x2016},
		Name: "IndicationLabel",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x2018},
		Name: "IndicationDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x201a},
		Name: "IndicationType",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x201c},
		Name: "IndicationDisposition",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x201e},
		Name: "IndicationROISequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x2030},
		Name: "IndicationPhysicalPropertySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x2032},
		Name: "PropertyLabel",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x2202},
		Name: "CoordinateSystemNumberOfAxes",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x2204},
		Name: "CoordinateSystemAxesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x2206},
		Name: "CoordinateSystemAxisDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x2208},
		Name: "CoordinateSystemDataSetMapping",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x220a},
		Name: "CoordinateSystemAxisNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x220c},
		Name: "CoordinateSystemAxisType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x220e},
		Name: "CoordinateSystemAxisUnits",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x2210},
		Name: "CoordinateSystemAxisValues",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x2220},
		Name: "CoordinateSystemTransformSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x2222},
		Name: "TransformDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x2224},
		Name: "TransformNumberOfAxes",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x2226},
		Name: "TransformOrderOfAxes",
		VR:   "IS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x2228},
		Name: "TransformedAxisUnits",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x222a},
		Name: "CoordinateSystemTransformRotationAndScaleMatrix",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x222c},
		Name: "CoordinateSystemTransformTranslationMatrix",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x3011},
		Name: "InternalDetectorFrameTime",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x3012},
		Name: "NumberOfFramesIntegrated",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x3020},
		Name: "DetectorTemperatureSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x3022},
		Name: "SensorName",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x3024},
		Name: "HorizontalOffsetOfSensor",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x3026},
		Name: "VerticalOffsetOfSensor",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x3028},
		Name: "SensorTemperature",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x3040},
		Name: "DarkCurrentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x3050},
		Name: "DarkCurrentCounts",
		VR:   "OW",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x3060},
		Name: "GainCorrectionReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x3070},
		Name: "AirCounts",
		VR:   "OW",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x3071},
		Name: "KVUsedInGainCalibration",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x3072},
		Name: "MAUsedInGainCalibration",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x3073},
		Name: "NumberOfFramesUsedForIntegration",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x3074},
		Name: "FilterMaterialUsedInGainCalibration",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x3075},
		Name: "FilterThicknessUsedInGainCalibration",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x3076},
		Name: "DateOfGainCalibration",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x3077},
		Name: "TimeOfGainCalibration",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x3080},
		Name: "BadPixelImage",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x3099},
		Name: "CalibrationNotes",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4002},
		Name: "PulserEquipmentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4004},
		Name: "PulserType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4006},
		Name: "PulserNotes",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4008},
		Name: "ReceiverEquipmentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x400a},
		Name: "AmplifierType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x400c},
		Name: "ReceiverNotes",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x400e},
		Name: "PreAmplifierEquipmentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x400f},
		Name: "PreAmplifierNotes",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4010},
		Name: "TransmitTransducerSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4011},
		Name: "ReceiveTransducerSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4012},
		Name: "NumberOfElements",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4013},
		Name: "ElementShape",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4014},
		Name: "ElementDimensionA",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4015},
		Name: "ElementDimensionB",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4016},
		Name: "ElementPitchA",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4017},
		Name: "MeasuredBeamDimensionA",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4018},
		Name: "MeasuredBeamDimensionB",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4019},
		Name: "LocationOfMeasuredBeamDiameter",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x401a},
		Name: "NominalFrequency",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x401b},
		Name: "MeasuredCenterFrequency",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x401c},
		Name: "MeasuredBandwidth",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x401d},
		Name: "ElementPitchB",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4020},
		Name: "PulserSettingsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4022},
		Name: "PulseWidth",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4024},
		Name: "ExcitationFrequency",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4026},
		Name: "ModulationType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4028},
		Name: "Damping",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4030},
		Name: "ReceiverSettingsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4031},
		Name: "AcquiredSoundpathLength",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4032},
		Name: "AcquisitionCompressionType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4033},
		Name: "AcquisitionSampleSize",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4034},
		Name: "RectifierSmoothing",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4035},
		Name: "DACSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4036},
		Name: "DACType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4038},
		Name: "DACGainPoints",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x403a},
		Name: "DACTimePoints",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x403c},
		Name: "DACAmplitude",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4040},
		Name: "PreAmplifierSettingsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4050},
		Name: "TransmitTransducerSettingsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4051},
		Name: "ReceiveTransducerSettingsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4052},
		Name: "IncidentAngle",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4054},
		Name: "CouplingTechnique",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4056},
		Name: "CouplingMedium",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4057},
		Name: "CouplingVelocity",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4058},
		Name: "ProbeCenterLocationX",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4059},
		Name: "ProbeCenterLocationZ",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x405a},
		Name: "SoundPathLength",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x405c},
		Name: "DelayLawIdentifier",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4060},
		Name: "GateSettingsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4062},
		Name: "GateThreshold",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4064},
		Name: "VelocityOfSound",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4070},
		Name: "CalibrationSettingsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4072},
		Name: "CalibrationProcedure",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4074},
		Name: "ProcedureVersion",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4076},
		Name: "ProcedureCreationDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4078},
		Name: "ProcedureExpirationDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x407a},
		Name: "ProcedureLastModifiedDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x407c},
		Name: "CalibrationTime",
		VR:   "TM",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x407e},
		Name: "CalibrationDate",
		VR:   "DA",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4080},
		Name: "ProbeDriveEquipmentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4081},
		Name: "DriveType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4082},
		Name: "ProbeDriveNotes",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4083},
		Name: "DriveProbeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4084},
		Name: "ProbeInductance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4085},
		Name: "ProbeResistance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4086},
		Name: "ReceiveProbeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4087},
		Name: "ProbeDriveSettingsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4088},
		Name: "BridgeResistors",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4089},
		Name: "ProbeOrientationAngle",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x408b},
		Name: "UserSelectedGainY",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x408c},
		Name: "UserSelectedPhase",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x408d},
		Name: "UserSelectedOffsetX",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x408e},
		Name: "UserSelectedOffsetY",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4091},
		Name: "ChannelSettingsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x4092},
		Name: "ChannelThreshold",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x409a},
		Name: "ScannerSettingsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x409b},
		Name: "ScanProcedure",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x409c},
		Name: "TranslationRateX",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x409d},
		Name: "TranslationRateY",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x409f},
		Name: "ChannelOverlap",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x40a0},
		Name: "ImageQualityIndicatorType",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x40a1},
		Name: "ImageQualityIndicatorMaterial",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x40a2},
		Name: "ImageQualityIndicatorSize",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5002},
		Name: "LINACEnergy",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5004},
		Name: "LINACOutput",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5100},
		Name: "ActiveAperture",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5101},
		Name: "TotalAperture",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5102},
		Name: "ApertureElevation",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5103},
		Name: "MainLobeAngle",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5104},
		Name: "MainRoofAngle",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5105},
		Name: "ConnectorType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5106},
		Name: "WedgeModelNumber",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5107},
		Name: "WedgeAngleFloat",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5108},
		Name: "WedgeRoofAngle",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5109},
		Name: "WedgeElement1Position",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x510a},
		Name: "WedgeMaterialVelocity",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x510b},
		Name: "WedgeMaterial",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x510c},
		Name: "WedgeOffsetZ",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x510d},
		Name: "WedgeOriginOffsetX",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x510e},
		Name: "WedgeTimeDelay",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x510f},
		Name: "WedgeName",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5110},
		Name: "WedgeManufacturerName",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5111},
		Name: "WedgeDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5112},
		Name: "NominalBeamAngle",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5113},
		Name: "WedgeOffsetX",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5114},
		Name: "WedgeOffsetY",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5115},
		Name: "WedgeTotalLength",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5116},
		Name: "WedgeInContactLength",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5117},
		Name: "WedgeFrontGap",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5118},
		Name: "WedgeTotalHeight",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x5119},
		Name: "WedgeFrontHeight",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x511a},
		Name: "WedgeRearHeight",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x511b},
		Name: "WedgeTotalWidth",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x511c},
		Name: "WedgeInContactWidth",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x511d},
		Name: "WedgeChamferHeight",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x511e},
		Name: "WedgeCurve",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0014, 0x511f},
		Name: "RadiusAlongWedge",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0001},
		Name: "WhitePoint",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0002},
		Name: "PrimaryChromaticities",
		VR:   "DS",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0003},
		Name: "BatteryLevel",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0004},
		Name: "ExposureTimeInSeconds",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0005},
		Name: "FNumber",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0006},
		Name: "OECFRows",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0007},
		Name: "OECFColumns",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0008},
		Name: "OECFColumnNames",
		VR:   "UC",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0009},
		Name: "OECFValues",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x000a},
		Name: "SpatialFrequencyResponseRows",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x000b},
		Name: "SpatialFrequencyResponseColumns",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x000c},
		Name: "SpatialFrequencyResponseColumnNames",
		VR:   "UC",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x000d},
		Name: "SpatialFrequencyResponseValues",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x000e},
		Name: "ColorFilterArrayPatternRows",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x000f},
		Name: "ColorFilterArrayPatternColumns",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0010},
		Name: "ColorFilterArrayPatternValues",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0011},
		Name: "FlashFiringStatus",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0012},
		Name: "FlashReturnStatus",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0013},
		Name: "FlashMode",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0014},
		Name: "FlashFunctionPresent",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0015},
		Name: "FlashRedEyeMode",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0016},
		Name: "ExposureProgram",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0017},
		Name: "SpectralSensitivity",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0018},
		Name: "PhotographicSensitivity",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0019},
		Name: "SelfTimerMode",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x001a},
		Name: "SensitivityType",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x001b},
		Name: "StandardOutputSensitivity",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x001c},
		Name: "RecommendedExposureIndex",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x001d},
		Name: "ISOSpeed",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x001e},
		Name: "ISOSpeedLatitudeyyy",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x001f},
		Name: "ISOSpeedLatitudezzz",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0020},
		Name: "EXIFVersion",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0021},
		Name: "ShutterSpeedValue",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0022},
		Name: "ApertureValue",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0023},
		Name: "BrightnessValue",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0024},
		Name: "ExposureBiasValue",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0025},
		Name: "MaxApertureValue",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0026},
		Name: "SubjectDistance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0027},
		Name: "MeteringMode",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0028},
		Name: "LightSource",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0029},
		Name: "FocalLength",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x002a},
		Name: "SubjectArea",
		VR:   "IS",
		VM:   "2-4",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x002b},
		Name: "MakerNote",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0030},
		Name: "Temperature",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0031},
		Name: "Humidity",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0032},
		Name: "Pressure",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0033},
		Name: "WaterDepth",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0034},
		Name: "Acceleration",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0035},
		Name: "CameraElevationAngle",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0036},
		Name: "FlashEnergy",
		VR:   "DS",
		VM:   "1-2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0037},
		Name: "SubjectLocation",
		VR:   "IS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0038},
		Name: "PhotographicExposureIndex",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0039},
		Name: "SensingMethod",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x003a},
		Name: "FileSource",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x003b},
		Name: "SceneType",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0041},
		Name: "CustomRendered",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0042},
		Name: "ExposureMode",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0043},
		Name: "WhiteBalance",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0044},
		Name: "DigitalZoomRatio",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0045},
		Name: "FocalLengthIn35mmFilm",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0046},
		Name: "SceneCaptureType",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0047},
		Name: "GainControl",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0048},
		Name: "Contrast",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0049},
		Name: "Saturation",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x004a},
		Name: "Sharpness",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x004b},
		Name: "DeviceSettingDescription",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x004c},
		Name: "SubjectDistanceRange",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x004d},
		Name: "CameraOwnerName",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x004e},
		Name: "LensSpecification",
		VR:   "DS",
		VM:   "4",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x004f},
		Name: "LensMake",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0050},
		Name: "LensModel",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0051},
		Name: "LensSerialNumber",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0061},
		Name: "InteroperabilityIndex",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0062},
		Name: "InteroperabilityVersion",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0070},
		Name: "GPSVersionID",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0071},
		Name: "GPSLatitudeRef",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0072},
		Name: "GPSLatitude",
		VR:   "DS",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0073},
		Name: "GPSLongitudeRef",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0074},
		Name: "GPSLongitude",
		VR:   "DS",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0075},
		Name: "GPSAltitudeRef",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0076},
		Name: "GPSAltitude",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0077},
		Name: "GPSTimeStamp",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0078},
		Name: "GPSSatellites",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0079},
		Name: "GPSStatus",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x007a},
		Name: "GPSMeasureMode",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x007b},
		Name: "GPSDOP",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x007c},
		Name: "GPSSpeedRef",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x007d},
		Name: "GPSSpeed",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x007e},
		Name: "GPSTrackRef",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x007f},
		Name: "GPSTrack",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0080},
		Name: "GPSImgDirectionRef",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0081},
		Name: "GPSImgDirection",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0082},
		Name: "GPSMapDatum",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0083},
		Name: "GPSDestLatitudeRef",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0084},
		Name: "GPSDestLatitude",
		VR:   "DS",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0085},
		Name: "GPSDestLongitudeRef",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0086},
		Name: "GPSDestLongitude",
		VR:   "DS",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0087},
		Name: "GPSDestBearingRef",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0088},
		Name: "GPSDestBearing",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x0089},
		Name: "GPSDestDistanceRef",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x008a},
		Name: "GPSDestDistance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x008b},
		Name: "GPSProcessingMethod",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x008c},
		Name: "GPSAreaInformation",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x008d},
		Name: "GPSDateStamp",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0016, 0x008e},
		Name: "GPSDifferential",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0010},
		Name: "ContrastBolusAgent",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0012},
		Name: "ContrastBolusAgentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0013},
		Name: "ContrastBolusT1Relaxivity",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0014},
		Name: "ContrastBolusAdministrationRouteSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0015},
		Name: "BodyPartExamined",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0020},
		Name: "ScanningSequence",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0021},
		Name: "SequenceVariant",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0022},
		Name: "ScanOptions",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0023},
		Name: "MRAcquisitionType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0024},
		Name: "SequenceName",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0025},
		Name: "AngioFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0026},
		Name: "InterventionDrugInformationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0027},
		Name: "InterventionDrugStopTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0028},
		Name: "InterventionDrugDose",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0029},
		Name: "InterventionDrugCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x002a},
		Name: "AdditionalDrugSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0031},
		Name: "Radiopharmaceutical",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0034},
		Name: "InterventionDrugName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0035},
		Name: "InterventionDrugStartTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0036},
		Name: "InterventionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0038},
		Name: "InterventionStatus",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x003a},
		Name: "InterventionDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0040},
		Name: "CineRate",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0042},
		Name: "InitialCineRunState",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0050},
		Name: "SliceThickness",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0060},
		Name: "KVP",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0070},
		Name: "CountsAccumulated",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0071},
		Name: "AcquisitionTerminationCondition",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0072},
		Name: "EffectiveDuration",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0073},
		Name: "AcquisitionStartCondition",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0074},
		Name: "AcquisitionStartConditionData",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0075},
		Name: "AcquisitionTerminationConditionData",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0080},
		Name: "RepetitionTime",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0081},
		Name: "EchoTime",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0082},
		Name: "InversionTime",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0083},
		Name: "NumberOfAverages",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0084},
		Name: "ImagingFrequency",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0085},
		Name: "ImagedNucleus",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0086},
		Name: "EchoNumbers",
		VR:   "IS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0087},
		Name: "MagneticFieldStrength",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0088},
		Name: "SpacingBetweenSlices",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0089},
		Name: "NumberOfPhaseEncodingSteps",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0090},
		Name: "DataCollectionDiameter",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0091},
		Name: "EchoTrainLength",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0093},
		Name: "PercentSampling",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0094},
		Name: "PercentPhaseFieldOfView",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x0095},
		Name: "PixelBandwidth",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1000},
		Name: "DeviceSerialNumber",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1002},
		Name: "DeviceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1003},
		Name: "DeviceID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1004},
		Name: "PlateID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1005},
		Name: "GeneratorID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1006},
		Name: "GridID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1007},
		Name: "CassetteID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1008},
		Name: "GantryID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1009},
		Name: "UniqueDeviceIdentifier",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x100a},
		Name: "UDISequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x100b},
		Name: "ManufacturerDeviceClassUID",
		VR:   "UI",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1010},
		Name: "SecondaryCaptureDeviceID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1012},
		Name: "DateOfSecondaryCapture",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1014},
		Name: "TimeOfSecondaryCapture",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1016},
		Name: "SecondaryCaptureDeviceManufacturer",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1018},
		Name: "SecondaryCaptureDeviceManufacturerModelName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1019},
		Name: "SecondaryCaptureDeviceSoftwareVersions",
		VR:   "LO",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1020},
		Name: "SoftwareVersions",
		VR:   "LO",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1022},
		Name: "VideoImageFormatAcquired",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1023},
		Name: "DigitalImageFormatAcquired",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1030},
		Name: "ProtocolName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1040},
		Name: "ContrastBolusRoute",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1041},
		Name: "ContrastBolusVolume",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1042},
		Name: "ContrastBolusStartTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1043},
		Name: "ContrastBolusStopTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1044},
		Name: "ContrastBolusTotalDose",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1045},
		Name: "SyringeCounts",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1046},
		Name: "ContrastFlowRate",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1047},
		Name: "ContrastFlowDuration",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1048},
		Name: "ContrastBolusIngredient",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1049},
		Name: "ContrastBolusIngredientConcentration",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1050},
		Name: "SpatialResolution",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1060},
		Name: "TriggerTime",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1061},
		Name: "TriggerSourceOrType",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1062},
		Name: "NominalInterval",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1063},
		Name: "FrameTime",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1064},
		Name: "CardiacFramingType",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1065},
		Name: "FrameTimeVector",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1066},
		Name: "FrameDelay",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1067},
		Name: "ImageTriggerDelay",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1068},
		Name: "MultiplexGroupTimeOffset",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1069},
		Name: "TriggerTimeOffset",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x106a},
		Name: "SynchronizationTrigger",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x106c},
		Name: "SynchronizationChannel",
		VR:   "US",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x106e},
		Name: "TriggerSamplePosition",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1070},
		Name: "RadiopharmaceuticalRoute",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1071},
		Name: "RadiopharmaceuticalVolume",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1072},
		Name: "RadiopharmaceuticalStartTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1073},
		Name: "RadiopharmaceuticalStopTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1074},
		Name: "RadionuclideTotalDose",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1075},
		Name: "RadionuclideHalfLife",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1076},
		Name: "RadionuclidePositronFraction",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1077},
		Name: "RadiopharmaceuticalSpecificActivity",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1078},
		Name: "RadiopharmaceuticalStartDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1079},
		Name: "RadiopharmaceuticalStopDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1080},
		Name: "BeatRejectionFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1081},
		Name: "LowRRValue",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1082},
		Name: "HighRRValue",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1083},
		Name: "IntervalsAcquired",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1084},
		Name: "IntervalsRejected",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1085},
		Name: "PVCRejection",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1086},
		Name: "SkipBeats",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1088},
		Name: "HeartRate",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1090},
		Name: "CardiacNumberOfImages",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1094},
		Name: "TriggerWindow",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1100},
		Name: "ReconstructionDiameter",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1110},
		Name: "DistanceSourceToDetector",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1111},
		Name: "DistanceSourceToPatient",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1114},
		Name: "EstimatedRadiographicMagnificationFactor",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1120},
		Name: "GantryDetectorTilt",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1121},
		Name: "GantryDetectorSlew",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1130},
		Name: "TableHeight",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1131},
		Name: "TableTraverse",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1134},
		Name: "TableMotion",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1135},
		Name: "TableVerticalIncrement",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1136},
		Name: "TableLateralIncrement",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1137},
		Name: "TableLongitudinalIncrement",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1138},
		Name: "TableAngle",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x113a},
		Name: "TableType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1140},
		Name: "RotationDirection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1142},
		Name: "RadialPosition",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1143},
		Name: "ScanArc",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1144},
		Name: "AngularStep",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1145},
		Name: "CenterOfRotationOffset",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1147},
		Name: "FieldOfViewShape",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1149},
		Name: "FieldOfViewDimensions",
		VR:   "IS",
		VM:   "1-2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1150},
		Name: "ExposureTime",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1151},
		Name: "XRayTubeCurrent",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1152},
		Name: "Exposure",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1153},
		Name: "ExposureInuAs",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1154},
		Name: "AveragePulseWidth",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1155},
		Name: "RadiationSetting",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1156},
		Name: "RectificationType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x115a},
		Name: "RadiationMode",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x115e},
		Name: "ImageAndFluoroscopyAreaDoseProduct",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1160},
		Name: "FilterType",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1161},
		Name: "TypeOfFilters",
		VR:   "LO",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1162},
		Name: "IntensifierSize",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1164},
		Name: "ImagerPixelSpacing",
		VR:   "DS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1166},
		Name: "Grid",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1170},
		Name: "GeneratorPower",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1180},
		Name: "CollimatorGridName",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1181},
		Name: "CollimatorType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1182},
		Name: "FocalDistance",
		VR:   "IS",
		VM:   "1-2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1183},
		Name: "XFocusCenter",
		VR:   "DS",
		VM:   "1-2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1184},
		Name: "YFocusCenter",
		VR:   "DS",
		VM:   "1-2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1190},
		Name: "FocalSpots",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1191},
		Name: "AnodeTargetMaterial",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x11a0},
		Name: "BodyPartThickness",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x11a2},
		Name: "CompressionForce",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x11a3},
		Name: "CompressionPressure",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x11a4},
		Name: "PaddleDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x11a5},
		Name: "CompressionContactArea",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1200},
		Name: "DateOfLastCalibration",
		VR:   "DA",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1201},
		Name: "TimeOfLastCalibration",
		VR:   "TM",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1202},
		Name: "DateTimeOfLastCalibration",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1210},
		Name: "ConvolutionKernel",
		VR:   "SH",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1242},
		Name: "ActualFrameDuration",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1243},
		Name: "CountRate",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1244},
		Name: "PreferredPlaybackSequencing",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1250},
		Name: "ReceiveCoilName",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1251},
		Name: "TransmitCoilName",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1260},
		Name: "PlateType",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1261},
		Name: "PhosphorType",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1271},
		Name: "WaterEquivalentDiameter",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1272},
		Name: "WaterEquivalentDiameterCalculationMethodCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1300},
		Name: "ScanVelocity",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1301},
		Name: "WholeBodyTechnique",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1302},
		Name: "ScanLength",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1310},
		Name: "AcquisitionMatrix",
		VR:   "US",
		VM:   "4",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1312},
		Name: "InPlanePhaseEncodingDirection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1314},
		Name: "FlipAngle",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1315},
		Name: "VariableFlipAngleFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1316},
		Name: "SAR",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1318},
		Name: "dBdt",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1320},
		Name: "B1rms",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1400},
		Name: "AcquisitionDeviceProcessingDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1401},
		Name: "AcquisitionDeviceProcessingCode",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1402},
		Name: "CassetteOrientation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1403},
		Name: "CassetteSize",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1404},
		Name: "ExposuresOnPlate",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1405},
		Name: "RelativeXRayExposure",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1411},
		Name: "ExposureIndex",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1412},
		Name: "TargetExposureIndex",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1413},
		Name: "DeviationIndex",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1450},
		Name: "ColumnAngulation",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1460},
		Name: "TomoLayerHeight",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1470},
		Name: "TomoAngle",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1480},
		Name: "TomoTime",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1490},
		Name: "TomoType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1491},
		Name: "TomoClass",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1495},
		Name: "NumberOfTomosynthesisSourceImages",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1500},
		Name: "PositionerMotion",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1508},
		Name: "PositionerType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1510},
		Name: "PositionerPrimaryAngle",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1511},
		Name: "PositionerSecondaryAngle",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1520},
		Name: "PositionerPrimaryAngleIncrement",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1521},
		Name: "PositionerSecondaryAngleIncrement",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1530},
		Name: "DetectorPrimaryAngle",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1531},
		Name: "DetectorSecondaryAngle",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1600},
		Name: "ShutterShape",
		VR:   "CS",
		VM:   "1-3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1602},
		Name: "ShutterLeftVerticalEdge",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1604},
		Name: "ShutterRightVerticalEdge",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1606},
		Name: "ShutterUpperHorizontalEdge",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1608},
		Name: "ShutterLowerHorizontalEdge",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1610},
		Name: "CenterOfCircularShutter",
		VR:   "IS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1612},
		Name: "RadiusOfCircularShutter",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1620},
		Name: "VerticesOfThePolygonalShutter",
		VR:   "IS",
		VM:   "2-2n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1622},
		Name: "ShutterPresentationValue",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1623},
		Name: "ShutterOverlayGroup",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1624},
		Name: "ShutterPresentationColorCIELabValue",
		VR:   "US",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1630},
		Name: "OutlineShapeType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1631},
		Name: "OutlineLeftVerticalEdge",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1632},
		Name: "OutlineRightVerticalEdge",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1633},
		Name: "OutlineUpperHorizontalEdge",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1634},
		Name: "OutlineLowerHorizontalEdge",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1635},
		Name: "CenterOfCircularOutline",
		VR:   "FD",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1636},
		Name: "DiameterOfCircularOutline",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1637},
		Name: "NumberOfPolygonalVertices",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1638},
		Name: "VerticesOfThePolygonalOutline",
		VR:   "OF",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1700},
		Name: "CollimatorShape",
		VR:   "CS",
		VM:   "1-3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1702},
		Name: "CollimatorLeftVerticalEdge",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1704},
		Name: "CollimatorRightVerticalEdge",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1706},
		Name: "CollimatorUpperHorizontalEdge",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1708},
		Name: "CollimatorLowerHorizontalEdge",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1710},
		Name: "CenterOfCircularCollimator",
		VR:   "IS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1712},
		Name: "RadiusOfCircularCollimator",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1720},
		Name: "VerticesOfThePolygonalCollimator",
		VR:   "IS",
		VM:   "2-2n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1800},
		Name: "AcquisitionTimeSynchronized",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1801},
		Name: "TimeSource",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1802},
		Name: "TimeDistributionProtocol",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x1803},
		Name: "NTPSourceAddress",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x2001},
		Name: "PageNumberVector",
		VR:   "IS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x2002},
		Name: "FrameLabelVector",
		VR:   "SH",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x2003},
		Name: "FramePrimaryAngleVector",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x2004},
		Name: "FrameSecondaryAngleVector",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x2005},
		Name: "SliceLocationVector",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x2006},
		Name: "DisplayWindowLabelVector",
		VR:   "SH",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x2010},
		Name: "NominalScannedPixelSpacing",
		VR:   "DS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x2020},
		Name: "DigitizingDeviceTransportDirection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x2030},
		Name: "RotationOfScannedFilm",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x2041},
		Name: "BiopsyTargetSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x2042},
		Name: "TargetUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x2043},
		Name: "LocalizingCursorPosition",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x2044},
		Name: "CalculatedTargetPosition",
		VR:   "FL",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x2045},
		Name: "TargetLabel",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x2046},
		Name: "DisplayedZValue",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x3100},
		Name: "IVUSAcquisition",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x3101},
		Name: "IVUSPullbackRate",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x3102},
		Name: "IVUSGatedRate",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x3103},
		Name: "IVUSPullbackStartFrameNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x3104},
		Name: "IVUSPullbackStopFrameNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x3105},
		Name: "LesionNumber",
		VR:   "IS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x5000},
		Name: "OutputPower",
		VR:   "SH",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x5010},
		Name: "TransducerData",
		VR:   "LO",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x5012},
		Name: "FocusDepth",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x5020},
		Name: "ProcessingFunction",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x5022},
		Name: "MechanicalIndex",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x5024},
		Name: "BoneThermalIndex",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x5026},
		Name: "CranialThermalIndex",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x5027},
		Name: "SoftTissueThermalIndex",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x5028},
		Name: "SoftTissueFocusThermalIndex",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x5029},
		Name: "SoftTissueSurfaceThermalIndex",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x5050},
		Name: "DepthOfScanField",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x5100},
		Name: "PatientPosition",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x5101},
		Name: "ViewPosition",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x5104},
		Name: "ProjectionEponymousNameCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6000},
		Name: "Sensitivity",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6011},
		Name: "SequenceOfUltrasoundRegions",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6012},
		Name: "RegionSpatialFormat",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6014},
		Name: "RegionDataType",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6016},
		Name: "RegionFlags",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6018},
		Name: "RegionLocationMinX0",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x601a},
		Name: "RegionLocationMinY0",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x601c},
		Name: "RegionLocationMaxX1",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x601e},
		Name: "RegionLocationMaxY1",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6020},
		Name: "ReferencePixelX0",
		VR:   "SL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6022},
		Name: "ReferencePixelY0",
		VR:   "SL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6024},
		Name: "PhysicalUnitsXDirection",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6026},
		Name: "PhysicalUnitsYDirection",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6028},
		Name: "ReferencePixelPhysicalValueX",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x602a},
		Name: "ReferencePixelPhysicalValueY",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x602c},
		Name: "PhysicalDeltaX",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x602e},
		Name: "PhysicalDeltaY",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6030},
		Name: "TransducerFrequency",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6031},
		Name: "TransducerType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6032},
		Name: "PulseRepetitionFrequency",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6034},
		Name: "DopplerCorrectionAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6036},
		Name: "SteeringAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6039},
		Name: "DopplerSampleVolumeXPosition",
		VR:   "SL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x603b},
		Name: "DopplerSampleVolumeYPosition",
		VR:   "SL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x603d},
		Name: "TMLinePositionX0",
		VR:   "SL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x603f},
		Name: "TMLinePositionY0",
		VR:   "SL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6041},
		Name: "TMLinePositionX1",
		VR:   "SL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6043},
		Name: "TMLinePositionY1",
		VR:   "SL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6044},
		Name: "PixelComponentOrganization",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6046},
		Name: "PixelComponentMask",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6048},
		Name: "PixelComponentRangeStart",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x604a},
		Name: "PixelComponentRangeStop",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x604c},
		Name: "PixelComponentPhysicalUnits",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x604e},
		Name: "PixelComponentDataType",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6050},
		Name: "NumberOfTableBreakPoints",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6052},
		Name: "TableOfXBreakPoints",
		VR:   "UL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6054},
		Name: "TableOfYBreakPoints",
		VR:   "FD",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6056},
		Name: "NumberOfTableEntries",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6058},
		Name: "TableOfPixelValues",
		VR:   "UL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x605a},
		Name: "TableOfParameterValues",
		VR:   "FL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x6060},
		Name: "RWaveTimeVector",
		VR:   "FL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7000},
		Name: "DetectorConditionsNominalFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7001},
		Name: "DetectorTemperature",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7004},
		Name: "DetectorType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7005},
		Name: "DetectorConfiguration",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7006},
		Name: "DetectorDescription",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7008},
		Name: "DetectorMode",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x700a},
		Name: "DetectorID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x700c},
		Name: "DateOfLastDetectorCalibration",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x700e},
		Name: "TimeOfLastDetectorCalibration",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7010},
		Name: "ExposuresOnDetectorSinceLastCalibration",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7011},
		Name: "ExposuresOnDetectorSinceManufactured",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7012},
		Name: "DetectorTimeSinceLastExposure",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7014},
		Name: "DetectorActiveTime",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7016},
		Name: "DetectorActivationOffsetFromExposure",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x701a},
		Name: "DetectorBinning",
		VR:   "DS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7020},
		Name: "DetectorElementPhysicalSize",
		VR:   "DS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7022},
		Name: "DetectorElementSpacing",
		VR:   "DS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7024},
		Name: "DetectorActiveShape",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7026},
		Name: "DetectorActiveDimensions",
		VR:   "DS",
		VM:   "1-2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7028},
		Name: "DetectorActiveOrigin",
		VR:   "DS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x702a},
		Name: "DetectorManufacturerName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x702b},
		Name: "DetectorManufacturerModelName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7030},
		Name: "FieldOfViewOrigin",
		VR:   "DS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7032},
		Name: "FieldOfViewRotation",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7034},
		Name: "FieldOfViewHorizontalFlip",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7036},
		Name: "PixelDataAreaOriginRelativeToFOV",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7038},
		Name: "PixelDataAreaRotationAngleRelativeToFOV",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7040},
		Name: "GridAbsorbingMaterial",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7041},
		Name: "GridSpacingMaterial",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7042},
		Name: "GridThickness",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7044},
		Name: "GridPitch",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7046},
		Name: "GridAspectRatio",
		VR:   "IS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7048},
		Name: "GridPeriod",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x704c},
		Name: "GridFocalDistance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7050},
		Name: "FilterMaterial",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7052},
		Name: "FilterThicknessMinimum",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7054},
		Name: "FilterThicknessMaximum",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7056},
		Name: "FilterBeamPathLengthMinimum",
		VR:   "FL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7058},
		Name: "FilterBeamPathLengthMaximum",
		VR:   "FL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7060},
		Name: "ExposureControlMode",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7062},
		Name: "ExposureControlModeDescription",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7064},
		Name: "ExposureStatus",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x7065},
		Name: "PhototimerSetting",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x8150},
		Name: "ExposureTimeInuS",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x8151},
		Name: "XRayTubeCurrentInuA",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9004},
		Name: "ContentQualification",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9005},
		Name: "PulseSequenceName",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9006},
		Name: "MRImagingModifierSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9008},
		Name: "EchoPulseSequence",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9009},
		Name: "InversionRecovery",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9010},
		Name: "FlowCompensation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9011},
		Name: "MultipleSpinEcho",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9012},
		Name: "MultiPlanarExcitation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9014},
		Name: "PhaseContrast",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9015},
		Name: "TimeOfFlightContrast",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9016},
		Name: "Spoiling",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9017},
		Name: "SteadyStatePulseSequence",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9018},
		Name: "EchoPlanarPulseSequence",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9019},
		Name: "TagAngleFirstAxis",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9020},
		Name: "MagnetizationTransfer",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9021},
		Name: "T2Preparation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9022},
		Name: "BloodSignalNulling",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9024},
		Name: "SaturationRecovery",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9025},
		Name: "SpectrallySelectedSuppression",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9026},
		Name: "SpectrallySelectedExcitation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9027},
		Name: "SpatialPresaturation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9028},
		Name: "Tagging",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9029},
		Name: "OversamplingPhase",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9030},
		Name: "TagSpacingFirstDimension",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9032},
		Name: "GeometryOfKSpaceTraversal",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9033},
		Name: "SegmentedKSpaceTraversal",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9034},
		Name: "RectilinearPhaseEncodeReordering",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9035},
		Name: "TagThickness",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9036},
		Name: "PartialFourierDirection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9037},
		Name: "CardiacSynchronizationTechnique",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9041},
		Name: "ReceiveCoilManufacturerName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9042},
		Name: "MRReceiveCoilSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9043},
		Name: "ReceiveCoilType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9044},
		Name: "QuadratureReceiveCoil",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9045},
		Name: "MultiCoilDefinitionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9046},
		Name: "MultiCoilConfiguration",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9047},
		Name: "MultiCoilElementName",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9048},
		Name: "MultiCoilElementUsed",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9049},
		Name: "MRTransmitCoilSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9050},
		Name: "TransmitCoilManufacturerName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9051},
		Name: "TransmitCoilType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9052},
		Name: "SpectralWidth",
		VR:   "FD",
		VM:   "1-2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9053},
		Name: "ChemicalShiftReference",
		VR:   "FD",
		VM:   "1-2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9054},
		Name: "VolumeLocalizationTechnique",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9058},
		Name: "MRAcquisitionFrequencyEncodingSteps",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9059},
		Name: "Decoupling",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9060},
		Name: "DecoupledNucleus",
		VR:   "CS",
		VM:   "1-2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9061},
		Name: "DecouplingFrequency",
		VR:   "FD",
		VM:   "1-2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9062},
		Name: "DecouplingMethod",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9063},
		Name: "DecouplingChemicalShiftReference",
		VR:   "FD",
		VM:   "1-2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9064},
		Name: "KSpaceFiltering",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9065},
		Name: "TimeDomainFiltering",
		VR:   "CS",
		VM:   "1-2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9066},
		Name: "NumberOfZeroFills",
		VR:   "US",
		VM:   "1-2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9067},
		Name: "BaselineCorrection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9069},
		Name: "ParallelReductionFactorInPlane",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9070},
		Name: "CardiacRRIntervalSpecified",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9073},
		Name: "AcquisitionDuration",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9074},
		Name: "FrameAcquisitionDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9075},
		Name: "DiffusionDirectionality",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9076},
		Name: "DiffusionGradientDirectionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9077},
		Name: "ParallelAcquisition",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9078},
		Name: "ParallelAcquisitionTechnique",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9079},
		Name: "InversionTimes",
		VR:   "FD",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9080},
		Name: "MetaboliteMapDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9081},
		Name: "PartialFourier",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9082},
		Name: "EffectiveEchoTime",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9083},
		Name: "MetaboliteMapCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9084},
		Name: "ChemicalShiftSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9085},
		Name: "CardiacSignalSource",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9087},
		Name: "DiffusionBValue",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9089},
		Name: "DiffusionGradientOrientation",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9090},
		Name: "VelocityEncodingDirection",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9091},
		Name: "VelocityEncodingMinimumValue",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9092},
		Name: "VelocityEncodingAcquisitionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9093},
		Name: "NumberOfKSpaceTrajectories",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9094},
		Name: "CoverageOfKSpace",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9095},
		Name: "SpectroscopyAcquisitionPhaseRows",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9098},
		Name: "TransmitterFrequency",
		VR:   "FD",
		VM:   "1-2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9100},
		Name: "ResonantNucleus",
		VR:   "CS",
		VM:   "1-2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9101},
		Name: "FrequencyCorrection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9103},
		Name: "MRSpectroscopyFOVGeometrySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9104},
		Name: "SlabThickness",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9105},
		Name: "SlabOrientation",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9106},
		Name: "MidSlabPosition",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9107},
		Name: "MRSpatialSaturationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9112},
		Name: "MRTimingAndRelatedParametersSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9114},
		Name: "MREchoSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9115},
		Name: "MRModifierSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9117},
		Name: "MRDiffusionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9118},
		Name: "CardiacSynchronizationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9119},
		Name: "MRAveragesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9125},
		Name: "MRFOVGeometrySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9126},
		Name: "VolumeLocalizationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9127},
		Name: "SpectroscopyAcquisitionDataColumns",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9147},
		Name: "DiffusionAnisotropyType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9151},
		Name: "FrameReferenceDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9152},
		Name: "MRMetaboliteMapSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9155},
		Name: "ParallelReductionFactorOutOfPlane",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9159},
		Name: "SpectroscopyAcquisitionOutOfPlanePhaseSteps",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9168},
		Name: "ParallelReductionFactorSecondInPlane",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9169},
		Name: "CardiacBeatRejectionTechnique",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9170},
		Name: "RespiratoryMotionCompensationTechnique",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9171},
		Name: "RespiratorySignalSource",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9172},
		Name: "BulkMotionCompensationTechnique",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9173},
		Name: "BulkMotionSignalSource",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9174},
		Name: "ApplicableSafetyStandardAgency",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9175},
		Name: "ApplicableSafetyStandardDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9176},
		Name: "OperatingModeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9177},
		Name: "OperatingModeType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9178},
		Name: "OperatingMode",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9179},
		Name: "SpecificAbsorptionRateDefinition",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9180},
		Name: "GradientOutputType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9181},
		Name: "SpecificAbsorptionRateValue",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9182},
		Name: "GradientOutput",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9183},
		Name: "FlowCompensationDirection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9184},
		Name: "TaggingDelay",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9185},
		Name: "RespiratoryMotionCompensationTechniqueDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9186},
		Name: "RespiratorySignalSourceID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9197},
		Name: "MRVelocityEncodingSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9198},
		Name: "FirstOrderPhaseCorrection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9199},
		Name: "WaterReferencedPhaseCorrection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9200},
		Name: "MRSpectroscopyAcquisitionType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9214},
		Name: "RespiratoryCyclePosition",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9217},
		Name: "VelocityEncodingMaximumValue",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9218},
		Name: "TagSpacingSecondDimension",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9219},
		Name: "TagAngleSecondAxis",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9220},
		Name: "FrameAcquisitionDuration",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9226},
		Name: "MRImageFrameTypeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9227},
		Name: "MRSpectroscopyFrameTypeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9231},
		Name: "MRAcquisitionPhaseEncodingStepsInPlane",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9232},
		Name: "MRAcquisitionPhaseEncodingStepsOutOfPlane",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9234},
		Name: "SpectroscopyAcquisitionPhaseColumns",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9236},
		Name: "CardiacCyclePosition",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9239},
		Name: "SpecificAbsorptionRateSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9240},
		Name: "RFEchoTrainLength",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9241},
		Name: "GradientEchoTrainLength",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9250},
		Name: "ArterialSpinLabelingContrast",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9251},
		Name: "MRArterialSpinLabelingSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9252},
		Name: "ASLTechniqueDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9253},
		Name: "ASLSlabNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9254},
		Name: "ASLSlabThickness",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9255},
		Name: "ASLSlabOrientation",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9256},
		Name: "ASLMidSlabPosition",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9257},
		Name: "ASLContext",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9258},
		Name: "ASLPulseTrainDuration",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9259},
		Name: "ASLCrusherFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x925a},
		Name: "ASLCrusherFlowLimit",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x925b},
		Name: "ASLCrusherDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x925c},
		Name: "ASLBolusCutoffFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x925d},
		Name: "ASLBolusCutoffTimingSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x925e},
		Name: "ASLBolusCutoffTechnique",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x925f},
		Name: "ASLBolusCutoffDelayTime",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9260},
		Name: "ASLSlabSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9295},
		Name: "ChemicalShiftMinimumIntegrationLimitInppm",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9296},
		Name: "ChemicalShiftMaximumIntegrationLimitInppm",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9297},
		Name: "WaterReferenceAcquisition",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9298},
		Name: "EchoPeakPosition",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9301},
		Name: "CTAcquisitionTypeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9302},
		Name: "AcquisitionType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9303},
		Name: "TubeAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9304},
		Name: "CTAcquisitionDetailsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9305},
		Name: "RevolutionTime",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9306},
		Name: "SingleCollimationWidth",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9307},
		Name: "TotalCollimationWidth",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9308},
		Name: "CTTableDynamicsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9309},
		Name: "TableSpeed",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9310},
		Name: "TableFeedPerRotation",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9311},
		Name: "SpiralPitchFactor",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9312},
		Name: "CTGeometrySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9313},
		Name: "DataCollectionCenterPatient",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9314},
		Name: "CTReconstructionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9315},
		Name: "ReconstructionAlgorithm",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9316},
		Name: "ConvolutionKernelGroup",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9317},
		Name: "ReconstructionFieldOfView",
		VR:   "FD",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9318},
		Name: "ReconstructionTargetCenterPatient",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9319},
		Name: "ReconstructionAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9320},
		Name: "ImageFilter",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9321},
		Name: "CTExposureSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9322},
		Name: "ReconstructionPixelSpacing",
		VR:   "FD",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9323},
		Name: "ExposureModulationType",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9324},
		Name: "EstimatedDoseSaving",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9325},
		Name: "CTXRayDetailsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9326},
		Name: "CTPositionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9327},
		Name: "TablePosition",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9328},
		Name: "ExposureTimeInms",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9329},
		Name: "CTImageFrameTypeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9330},
		Name: "XRayTubeCurrentInmA",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9332},
		Name: "ExposureInmAs",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9333},
		Name: "ConstantVolumeFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9334},
		Name: "FluoroscopyFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9335},
		Name: "DistanceSourceToDataCollectionCenter",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9337},
		Name: "ContrastBolusAgentNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9338},
		Name: "ContrastBolusIngredientCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9340},
		Name: "ContrastAdministrationProfileSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9341},
		Name: "ContrastBolusUsageSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9342},
		Name: "ContrastBolusAgentAdministered",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9343},
		Name: "ContrastBolusAgentDetected",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9344},
		Name: "ContrastBolusAgentPhase",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9345},
		Name: "CTDIvol",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9346},
		Name: "CTDIPhantomTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9351},
		Name: "CalciumScoringMassFactorPatient",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9352},
		Name: "CalciumScoringMassFactorDevice",
		VR:   "FL",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9353},
		Name: "EnergyWeightingFactor",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9360},
		Name: "CTAdditionalXRaySourceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9361},
		Name: "MultienergyCTAcquisition",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9362},
		Name: "MultienergyCTAcquisitionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9363},
		Name: "MultienergyCTProcessingSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9364},
		Name: "MultienergyCTCharacteristicsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9365},
		Name: "MultienergyCTXRaySourceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9366},
		Name: "XRaySourceIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9367},
		Name: "XRaySourceID",
		VR:   "UC",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9368},
		Name: "MultienergySourceTechnique",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9369},
		Name: "SourceStartDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x936a},
		Name: "SourceEndDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x936b},
		Name: "SwitchingPhaseNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x936c},
		Name: "SwitchingPhaseNominalDuration",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x936d},
		Name: "SwitchingPhaseTransitionDuration",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x936e},
		Name: "EffectiveBinEnergy",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x936f},
		Name: "MultienergyCTXRayDetectorSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9370},
		Name: "XRayDetectorIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9371},
		Name: "XRayDetectorID",
		VR:   "UC",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9372},
		Name: "MultienergyDetectorType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9373},
		Name: "XRayDetectorLabel",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9374},
		Name: "NominalMaxEnergy",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9375},
		Name: "NominalMinEnergy",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9376},
		Name: "ReferencedXRayDetectorIndex",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9377},
		Name: "ReferencedXRaySourceIndex",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9378},
		Name: "ReferencedPathIndex",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9379},
		Name: "MultienergyCTPathSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x937a},
		Name: "MultienergyCTPathIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x937b},
		Name: "MultienergyAcquisitionDescription",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x937c},
		Name: "MonoenergeticEnergyEquivalent",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x937d},
		Name: "MaterialCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x937e},
		Name: "DecompositionMethod",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x937f},
		Name: "DecompositionDescription",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9380},
		Name: "DecompositionAlgorithmIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9381},
		Name: "DecompositionMaterialSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9382},
		Name: "MaterialAttenuationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9383},
		Name: "PhotonEnergy",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9384},
		Name: "XRayMassAttenuationCoefficient",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9401},
		Name: "ProjectionPixelCalibrationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9402},
		Name: "DistanceSourceToIsocenter",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9403},
		Name: "DistanceObjectToTableTop",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9404},
		Name: "ObjectPixelSpacingInCenterOfBeam",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9405},
		Name: "PositionerPositionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9406},
		Name: "TablePositionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9407},
		Name: "CollimatorShapeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9410},
		Name: "PlanesInAcquisition",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9412},
		Name: "XAXRFFrameCharacteristicsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9417},
		Name: "FrameAcquisitionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9420},
		Name: "XRayReceptorType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9423},
		Name: "AcquisitionProtocolName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9424},
		Name: "AcquisitionProtocolDescription",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9425},
		Name: "ContrastBolusIngredientOpaque",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9426},
		Name: "DistanceReceptorPlaneToDetectorHousing",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9427},
		Name: "IntensifierActiveShape",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9428},
		Name: "IntensifierActiveDimensions",
		VR:   "FL",
		VM:   "1-2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9429},
		Name: "PhysicalDetectorSize",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9430},
		Name: "PositionOfIsocenterProjection",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9432},
		Name: "FieldOfViewSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9433},
		Name: "FieldOfViewDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9434},
		Name: "ExposureControlSensingRegionsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9435},
		Name: "ExposureControlSensingRegionShape",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9436},
		Name: "ExposureControlSensingRegionLeftVerticalEdge",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9437},
		Name: "ExposureControlSensingRegionRightVerticalEdge",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9438},
		Name: "ExposureControlSensingRegionUpperHorizontalEdge",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9439},
		Name: "ExposureControlSensingRegionLowerHorizontalEdge",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9440},
		Name: "CenterOfCircularExposureControlSensingRegion",
		VR:   "SS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9441},
		Name: "RadiusOfCircularExposureControlSensingRegion",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9442},
		Name: "VerticesOfThePolygonalExposureControlSensingRegion",
		VR:   "SS",
		VM:   "2-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9447},
		Name: "ColumnAngulationPatient",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9449},
		Name: "BeamAngle",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9451},
		Name: "FrameDetectorParametersSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9452},
		Name: "CalculatedAnatomyThickness",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9455},
		Name: "CalibrationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9456},
		Name: "ObjectThicknessSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9457},
		Name: "PlaneIdentification",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9461},
		Name: "FieldOfViewDimensionsInFloat",
		VR:   "FL",
		VM:   "1-2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9462},
		Name: "IsocenterReferenceSystemSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9463},
		Name: "PositionerIsocenterPrimaryAngle",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9464},
		Name: "PositionerIsocenterSecondaryAngle",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9465},
		Name: "PositionerIsocenterDetectorRotationAngle",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9466},
		Name: "TableXPositionToIsocenter",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9467},
		Name: "TableYPositionToIsocenter",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9468},
		Name: "TableZPositionToIsocenter",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9469},
		Name: "TableHorizontalRotationAngle",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9470},
		Name: "TableHeadTiltAngle",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9471},
		Name: "TableCradleTiltAngle",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9472},
		Name: "FrameDisplayShutterSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9473},
		Name: "AcquiredImageAreaDoseProduct",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9474},
		Name: "CArmPositionerTabletopRelationship",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9476},
		Name: "XRayGeometrySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9477},
		Name: "IrradiationEventIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9504},
		Name: "XRay3DFrameTypeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9506},
		Name: "ContributingSourcesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9507},
		Name: "XRay3DAcquisitionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9508},
		Name: "PrimaryPositionerScanArc",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9509},
		Name: "SecondaryPositionerScanArc",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9510},
		Name: "PrimaryPositionerScanStartAngle",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9511},
		Name: "SecondaryPositionerScanStartAngle",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9514},
		Name: "PrimaryPositionerIncrement",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9515},
		Name: "SecondaryPositionerIncrement",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9516},
		Name: "StartAcquisitionDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9517},
		Name: "EndAcquisitionDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9518},
		Name: "PrimaryPositionerIncrementSign",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9519},
		Name: "SecondaryPositionerIncrementSign",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9524},
		Name: "ApplicationName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9525},
		Name: "ApplicationVersion",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9526},
		Name: "ApplicationManufacturer",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9527},
		Name: "AlgorithmType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9528},
		Name: "AlgorithmDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9530},
		Name: "XRay3DReconstructionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9531},
		Name: "ReconstructionDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9538},
		Name: "PerProjectionAcquisitionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9541},
		Name: "DetectorPositionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9542},
		Name: "XRayAcquisitionDoseSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9543},
		Name: "XRaySourceIsocenterPrimaryAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9544},
		Name: "XRaySourceIsocenterSecondaryAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9545},
		Name: "BreastSupportIsocenterPrimaryAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9546},
		Name: "BreastSupportIsocenterSecondaryAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9547},
		Name: "BreastSupportXPositionToIsocenter",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9548},
		Name: "BreastSupportYPositionToIsocenter",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9549},
		Name: "BreastSupportZPositionToIsocenter",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9550},
		Name: "DetectorIsocenterPrimaryAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9551},
		Name: "DetectorIsocenterSecondaryAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9552},
		Name: "DetectorXPositionToIsocenter",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9553},
		Name: "DetectorYPositionToIsocenter",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9554},
		Name: "DetectorZPositionToIsocenter",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9555},
		Name: "XRayGridSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9556},
		Name: "XRayFilterSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9557},
		Name: "DetectorActiveAreaTLHCPosition",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9558},
		Name: "DetectorActiveAreaOrientation",
		VR:   "FD",
		VM:   "6",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9559},
		Name: "PositionerPrimaryAngleDirection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9601},
		Name: "DiffusionBMatrixSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9602},
		Name: "DiffusionBValueXX",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9603},
		Name: "DiffusionBValueXY",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9604},
		Name: "DiffusionBValueXZ",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9605},
		Name: "DiffusionBValueYY",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9606},
		Name: "DiffusionBValueYZ",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9607},
		Name: "DiffusionBValueZZ",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9621},
		Name: "FunctionalMRSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9622},
		Name: "FunctionalSettlingPhaseFramesPresent",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9623},
		Name: "FunctionalSyncPulse",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9624},
		Name: "SettlingPhaseFrame",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9701},
		Name: "DecayCorrectionDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9715},
		Name: "StartDensityThreshold",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9716},
		Name: "StartRelativeDensityDifferenceThreshold",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9717},
		Name: "StartCardiacTriggerCountThreshold",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9718},
		Name: "StartRespiratoryTriggerCountThreshold",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9719},
		Name: "TerminationCountsThreshold",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9720},
		Name: "TerminationDensityThreshold",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9721},
		Name: "TerminationRelativeDensityThreshold",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9722},
		Name: "TerminationTimeThreshold",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9723},
		Name: "TerminationCardiacTriggerCountThreshold",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9724},
		Name: "TerminationRespiratoryTriggerCountThreshold",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9725},
		Name: "DetectorGeometry",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9726},
		Name: "TransverseDetectorSeparation",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9727},
		Name: "AxialDetectorDimension",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9729},
		Name: "RadiopharmaceuticalAgentNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9732},
		Name: "PETFrameAcquisitionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9733},
		Name: "PETDetectorMotionDetailsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9734},
		Name: "PETTableDynamicsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9735},
		Name: "PETPositionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9736},
		Name: "PETFrameCorrectionFactorsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9737},
		Name: "RadiopharmaceuticalUsageSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9738},
		Name: "AttenuationCorrectionSource",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9739},
		Name: "NumberOfIterations",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9740},
		Name: "NumberOfSubsets",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9749},
		Name: "PETReconstructionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9751},
		Name: "PETFrameTypeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9755},
		Name: "TimeOfFlightInformationUsed",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9756},
		Name: "ReconstructionType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9758},
		Name: "DecayCorrected",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9759},
		Name: "AttenuationCorrected",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9760},
		Name: "ScatterCorrected",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9761},
		Name: "DeadTimeCorrected",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9762},
		Name: "GantryMotionCorrected",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9763},
		Name: "PatientMotionCorrected",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9764},
		Name: "CountLossNormalizationCorrected",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9765},
		Name: "RandomsCorrected",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9766},
		Name: "NonUniformRadialSamplingCorrected",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9767},
		Name: "SensitivityCalibrated",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9768},
		Name: "DetectorNormalizationCorrection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9769},
		Name: "IterativeReconstructionMethod",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9770},
		Name: "AttenuationCorrectionTemporalRelationship",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9771},
		Name: "PatientPhysiologicalStateSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9772},
		Name: "PatientPhysiologicalStateCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9801},
		Name: "DepthsOfFocus",
		VR:   "FD",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9803},
		Name: "ExcludedIntervalsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9804},
		Name: "ExclusionStartDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9805},
		Name: "ExclusionDuration",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9806},
		Name: "USImageDescriptionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9807},
		Name: "ImageDataTypeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9808},
		Name: "DataType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9809},
		Name: "TransducerScanPatternCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x980b},
		Name: "AliasedDataType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x980c},
		Name: "PositionMeasuringDeviceUsed",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x980d},
		Name: "TransducerGeometryCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x980e},
		Name: "TransducerBeamSteeringCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x980f},
		Name: "TransducerApplicationCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9810},
		Name: "ZeroVelocityPixelValue",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9900},
		Name: "ReferenceLocationLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9901},
		Name: "ReferenceLocationDescription",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9902},
		Name: "ReferenceBasisCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9903},
		Name: "ReferenceGeometryCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9904},
		Name: "OffsetDistance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9905},
		Name: "OffsetDirection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9906},
		Name: "PotentialScheduledProtocolCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9907},
		Name: "PotentialRequestedProcedureCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9908},
		Name: "PotentialReasonsForProcedure",
		VR:   "UC",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9909},
		Name: "PotentialReasonsForProcedureCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x990a},
		Name: "PotentialDiagnosticTasks",
		VR:   "UC",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x990b},
		Name: "ContraindicationsCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x990c},
		Name: "ReferencedDefinedProtocolSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x990d},
		Name: "ReferencedPerformedProtocolSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x990e},
		Name: "PredecessorProtocolSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x990f},
		Name: "ProtocolPlanningInformation",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9910},
		Name: "ProtocolDesignRationale",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9911},
		Name: "PatientSpecificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9912},
		Name: "ModelSpecificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9913},
		Name: "ParametersSpecificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9914},
		Name: "InstructionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9915},
		Name: "InstructionIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9916},
		Name: "InstructionText",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9917},
		Name: "InstructionDescription",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9918},
		Name: "InstructionPerformedFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9919},
		Name: "InstructionPerformedDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x991a},
		Name: "InstructionPerformanceComment",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x991b},
		Name: "PatientPositioningInstructionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x991c},
		Name: "PositioningMethodCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x991d},
		Name: "PositioningLandmarkSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x991e},
		Name: "TargetFrameOfReferenceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x991f},
		Name: "AcquisitionProtocolElementSpecificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9920},
		Name: "AcquisitionProtocolElementSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9921},
		Name: "ProtocolElementNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9922},
		Name: "ProtocolElementName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9923},
		Name: "ProtocolElementCharacteristicsSummary",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9924},
		Name: "ProtocolElementPurpose",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9930},
		Name: "AcquisitionMotion",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9931},
		Name: "AcquisitionStartLocationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9932},
		Name: "AcquisitionEndLocationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9933},
		Name: "ReconstructionProtocolElementSpecificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9934},
		Name: "ReconstructionProtocolElementSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9935},
		Name: "StorageProtocolElementSpecificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9936},
		Name: "StorageProtocolElementSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9937},
		Name: "RequestedSeriesDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9938},
		Name: "SourceAcquisitionProtocolElementNumber",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9939},
		Name: "SourceAcquisitionBeamNumber",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x993a},
		Name: "SourceReconstructionProtocolElementNumber",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x993b},
		Name: "ReconstructionStartLocationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x993c},
		Name: "ReconstructionEndLocationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x993d},
		Name: "ReconstructionAlgorithmSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x993e},
		Name: "ReconstructionTargetCenterLocationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9941},
		Name: "ImageFilterDescription",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9942},
		Name: "CTDIvolNotificationTrigger",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9943},
		Name: "DLPNotificationTrigger",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9944},
		Name: "AutoKVPSelectionType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9945},
		Name: "AutoKVPUpperBound",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9946},
		Name: "AutoKVPLowerBound",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0x9947},
		Name: "ProtocolDefinedPatientPosition",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0xa001},
		Name: "ContributingEquipmentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0xa002},
		Name: "ContributionDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0018, 0xa003},
		Name: "ContributionDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x000d},
		Name: "StudyInstanceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x000e},
		Name: "SeriesInstanceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x0010},
		Name: "StudyID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x0011},
		Name: "SeriesNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x0012},
		Name: "AcquisitionNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x0013},
		Name: "InstanceNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x0019},
		Name: "ItemNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x0020},
		Name: "PatientOrientation",
		VR:   "CS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x0032},
		Name: "ImagePositionPatient",
		VR:   "DS",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x0037},
		Name: "ImageOrientationPatient",
		VR:   "DS",
		VM:   "6",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x0052},
		Name: "FrameOfReferenceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x0060},
		Name: "Laterality",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x0062},
		Name: "ImageLaterality",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x0100},
		Name: "TemporalPositionIdentifier",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x0105},
		Name: "NumberOfTemporalPositions",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x0110},
		Name: "TemporalResolution",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x0200},
		Name: "SynchronizationFrameOfReferenceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x0242},
		Name: "SOPInstanceUIDOfConcatenationSource",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x1002},
		Name: "ImagesInAcquisition",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x103f},
		Name: "TargetPositionReferenceIndicator",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x1040},
		Name: "PositionReferenceIndicator",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x1041},
		Name: "SliceLocation",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x1200},
		Name: "NumberOfPatientRelatedStudies",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x1202},
		Name: "NumberOfPatientRelatedSeries",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x1204},
		Name: "NumberOfPatientRelatedInstances",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x1206},
		Name: "NumberOfStudyRelatedSeries",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x1208},
		Name: "NumberOfStudyRelatedInstances",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x1209},
		Name: "NumberOfSeriesRelatedInstances",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x4000},
		Name: "ImageComments",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9056},
		Name: "StackID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9057},
		Name: "InStackPositionNumber",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9071},
		Name: "FrameAnatomySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9072},
		Name: "FrameLaterality",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9111},
		Name: "FrameContentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9113},
		Name: "PlanePositionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9116},
		Name: "PlaneOrientationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9128},
		Name: "TemporalPositionIndex",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9153},
		Name: "NominalCardiacTriggerDelayTime",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9154},
		Name: "NominalCardiacTriggerTimePriorToRPeak",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9155},
		Name: "ActualCardiacTriggerTimePriorToRPeak",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9156},
		Name: "FrameAcquisitionNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9157},
		Name: "DimensionIndexValues",
		VR:   "UL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9158},
		Name: "FrameComments",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9161},
		Name: "ConcatenationUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9162},
		Name: "InConcatenationNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9163},
		Name: "InConcatenationTotalNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9164},
		Name: "DimensionOrganizationUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9165},
		Name: "DimensionIndexPointer",
		VR:   "AT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9167},
		Name: "FunctionalGroupPointer",
		VR:   "AT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9170},
		Name: "UnassignedSharedConvertedAttributesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9171},
		Name: "UnassignedPerFrameConvertedAttributesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9172},
		Name: "ConversionSourceAttributesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9213},
		Name: "DimensionIndexPrivateCreator",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9221},
		Name: "DimensionOrganizationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9222},
		Name: "DimensionIndexSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9228},
		Name: "ConcatenationFrameOffsetNumber",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9238},
		Name: "FunctionalGroupPrivateCreator",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9241},
		Name: "NominalPercentageOfCardiacPhase",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9245},
		Name: "NominalPercentageOfRespiratoryPhase",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9246},
		Name: "StartingRespiratoryAmplitude",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9247},
		Name: "StartingRespiratoryPhase",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9248},
		Name: "EndingRespiratoryAmplitude",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9249},
		Name: "EndingRespiratoryPhase",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9250},
		Name: "RespiratoryTriggerType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9251},
		Name: "RRIntervalTimeNominal",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9252},
		Name: "ActualCardiacTriggerDelayTime",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9253},
		Name: "RespiratorySynchronizationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9254},
		Name: "RespiratoryIntervalTime",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9255},
		Name: "NominalRespiratoryTriggerDelayTime",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9256},
		Name: "RespiratoryTriggerDelayThreshold",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9257},
		Name: "ActualRespiratoryTriggerDelayTime",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9301},
		Name: "ImagePositionVolume",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9302},
		Name: "ImageOrientationVolume",
		VR:   "FD",
		VM:   "6",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9307},
		Name: "UltrasoundAcquisitionGeometry",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9308},
		Name: "ApexPosition",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9309},
		Name: "VolumeToTransducerMappingMatrix",
		VR:   "FD",
		VM:   "16",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x930a},
		Name: "VolumeToTableMappingMatrix",
		VR:   "FD",
		VM:   "16",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x930b},
		Name: "VolumeToTransducerRelationship",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x930c},
		Name: "PatientFrameOfReferenceSource",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x930d},
		Name: "TemporalPositionTimeOffset",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x930e},
		Name: "PlanePositionVolumeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x930f},
		Name: "PlaneOrientationVolumeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9310},
		Name: "TemporalPositionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9311},
		Name: "DimensionOrganizationType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9312},
		Name: "VolumeFrameOfReferenceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9313},
		Name: "TableFrameOfReferenceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9421},
		Name: "DimensionDescriptionLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9450},
		Name: "PatientOrientationInFrameSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9453},
		Name: "FrameLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9518},
		Name: "AcquisitionIndex",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9529},
		Name: "ContributingSOPInstancesReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0020, 0x9536},
		Name: "ReconstructionIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0001},
		Name: "LightPathFilterPassThroughWavelength",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0002},
		Name: "LightPathFilterPassBand",
		VR:   "US",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0003},
		Name: "ImagePathFilterPassThroughWavelength",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0004},
		Name: "ImagePathFilterPassBand",
		VR:   "US",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0005},
		Name: "PatientEyeMovementCommanded",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0006},
		Name: "PatientEyeMovementCommandCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0007},
		Name: "SphericalLensPower",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0008},
		Name: "CylinderLensPower",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0009},
		Name: "CylinderAxis",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x000a},
		Name: "EmmetropicMagnification",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x000b},
		Name: "IntraOcularPressure",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x000c},
		Name: "HorizontalFieldOfView",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x000d},
		Name: "PupilDilated",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x000e},
		Name: "DegreeOfDilation",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0010},
		Name: "StereoBaselineAngle",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0011},
		Name: "StereoBaselineDisplacement",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0012},
		Name: "StereoHorizontalPixelOffset",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0013},
		Name: "StereoVerticalPixelOffset",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0014},
		Name: "StereoRotation",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0015},
		Name: "AcquisitionDeviceTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0016},
		Name: "IlluminationTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0017},
		Name: "LightPathFilterTypeStackCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0018},
		Name: "ImagePathFilterTypeStackCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0019},
		Name: "LensesCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x001a},
		Name: "ChannelDescriptionCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x001b},
		Name: "RefractiveStateSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x001c},
		Name: "MydriaticAgentCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x001d},
		Name: "RelativeImagePositionCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x001e},
		Name: "CameraAngleOfView",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0020},
		Name: "StereoPairsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0021},
		Name: "LeftImageSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0022},
		Name: "RightImageSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0028},
		Name: "StereoPairsPresent",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0030},
		Name: "AxialLengthOfTheEye",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0031},
		Name: "OphthalmicFrameLocationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0032},
		Name: "ReferenceCoordinates",
		VR:   "FL",
		VM:   "2-2n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0035},
		Name: "DepthSpatialResolution",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0036},
		Name: "MaximumDepthDistortion",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0037},
		Name: "AlongScanSpatialResolution",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0038},
		Name: "MaximumAlongScanDistortion",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0039},
		Name: "OphthalmicImageOrientation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0041},
		Name: "DepthOfTransverseImage",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0042},
		Name: "MydriaticAgentConcentrationUnitsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0048},
		Name: "AcrossScanSpatialResolution",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0049},
		Name: "MaximumAcrossScanDistortion",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x004e},
		Name: "MydriaticAgentConcentration",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0055},
		Name: "IlluminationWaveLength",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0056},
		Name: "IlluminationPower",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0057},
		Name: "IlluminationBandwidth",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x0058},
		Name: "MydriaticAgentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1007},
		Name: "OphthalmicAxialMeasurementsRightEyeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1008},
		Name: "OphthalmicAxialMeasurementsLeftEyeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1009},
		Name: "OphthalmicAxialMeasurementsDeviceType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1010},
		Name: "OphthalmicAxialLengthMeasurementsType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1012},
		Name: "OphthalmicAxialLengthSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1019},
		Name: "OphthalmicAxialLength",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1024},
		Name: "LensStatusCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1025},
		Name: "VitreousStatusCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1028},
		Name: "IOLFormulaCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1029},
		Name: "IOLFormulaDetail",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1033},
		Name: "KeratometerIndex",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1035},
		Name: "SourceOfOphthalmicAxialLengthCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1036},
		Name: "SourceOfCornealSizeDataCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1037},
		Name: "TargetRefraction",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1039},
		Name: "RefractiveProcedureOccurred",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1040},
		Name: "RefractiveSurgeryTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1044},
		Name: "OphthalmicUltrasoundMethodCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1045},
		Name: "SurgicallyInducedAstigmatismSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1046},
		Name: "TypeOfOpticalCorrection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1047},
		Name: "ToricIOLPowerSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1048},
		Name: "PredictedToricErrorSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1049},
		Name: "PreSelectedForImplantation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x104a},
		Name: "ToricIOLPowerForExactEmmetropiaSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x104b},
		Name: "ToricIOLPowerForExactTargetRefractionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1050},
		Name: "OphthalmicAxialLengthMeasurementsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1053},
		Name: "IOLPower",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1054},
		Name: "PredictedRefractiveError",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1059},
		Name: "OphthalmicAxialLengthVelocity",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1065},
		Name: "LensStatusDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1066},
		Name: "VitreousStatusDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1090},
		Name: "IOLPowerSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1092},
		Name: "LensConstantSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1093},
		Name: "IOLManufacturer",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1095},
		Name: "ImplantName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1096},
		Name: "KeratometryMeasurementTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1097},
		Name: "ImplantPartNumber",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1100},
		Name: "ReferencedOphthalmicAxialMeasurementsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1101},
		Name: "OphthalmicAxialLengthMeasurementsSegmentNameCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1103},
		Name: "RefractiveErrorBeforeRefractiveSurgeryCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1121},
		Name: "IOLPowerForExactEmmetropia",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1122},
		Name: "IOLPowerForExactTargetRefraction",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1125},
		Name: "AnteriorChamberDepthDefinitionCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1127},
		Name: "LensThicknessSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1128},
		Name: "AnteriorChamberDepthSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x112a},
		Name: "CalculationCommentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x112b},
		Name: "CalculationCommentType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x112c},
		Name: "CalculationComment",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1130},
		Name: "LensThickness",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1131},
		Name: "AnteriorChamberDepth",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1132},
		Name: "SourceOfLensThicknessDataCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1133},
		Name: "SourceOfAnteriorChamberDepthDataCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1134},
		Name: "SourceOfRefractiveMeasurementsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1135},
		Name: "SourceOfRefractiveMeasurementsCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1140},
		Name: "OphthalmicAxialLengthMeasurementModified",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1150},
		Name: "OphthalmicAxialLengthDataSourceCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1155},
		Name: "SignalToNoiseRatio",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1159},
		Name: "OphthalmicAxialLengthDataSourceDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1210},
		Name: "OphthalmicAxialLengthMeasurementsTotalLengthSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1211},
		Name: "OphthalmicAxialLengthMeasurementsSegmentalLengthSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1212},
		Name: "OphthalmicAxialLengthMeasurementsLengthSummationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1220},
		Name: "UltrasoundOphthalmicAxialLengthMeasurementsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1225},
		Name: "OpticalOphthalmicAxialLengthMeasurementsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1230},
		Name: "UltrasoundSelectedOphthalmicAxialLengthSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1250},
		Name: "OphthalmicAxialLengthSelectionMethodCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1255},
		Name: "OpticalSelectedOphthalmicAxialLengthSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1257},
		Name: "SelectedSegmentalOphthalmicAxialLengthSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1260},
		Name: "SelectedTotalOphthalmicAxialLengthSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1262},
		Name: "OphthalmicAxialLengthQualityMetricSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1300},
		Name: "IntraocularLensCalculationsRightEyeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1310},
		Name: "IntraocularLensCalculationsLeftEyeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1330},
		Name: "ReferencedOphthalmicAxialLengthMeasurementQCImageSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1415},
		Name: "OphthalmicMappingDeviceType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1420},
		Name: "AcquisitionMethodCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1423},
		Name: "AcquisitionMethodAlgorithmSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1436},
		Name: "OphthalmicThicknessMapTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1443},
		Name: "OphthalmicThicknessMappingNormalsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1445},
		Name: "RetinalThicknessDefinitionCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1450},
		Name: "PixelValueMappingToCodedConceptSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1452},
		Name: "MappedPixelValue",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1454},
		Name: "PixelValueMappingExplanation",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1458},
		Name: "OphthalmicThicknessMapQualityThresholdSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1460},
		Name: "OphthalmicThicknessMapThresholdQualityRating",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1463},
		Name: "AnatomicStructureReferencePoint",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1465},
		Name: "RegistrationToLocalizerSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1466},
		Name: "RegisteredLocalizerUnits",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1467},
		Name: "RegisteredLocalizerTopLeftHandCorner",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1468},
		Name: "RegisteredLocalizerBottomRightHandCorner",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1470},
		Name: "OphthalmicThicknessMapQualityRatingSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1472},
		Name: "RelevantOPTAttributesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1512},
		Name: "TransformationMethodCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1513},
		Name: "TransformationAlgorithmSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1515},
		Name: "OphthalmicAxialLengthMethod",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1517},
		Name: "OphthalmicFOV",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1518},
		Name: "TwoDimensionalToThreeDimensionalMapSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1525},
		Name: "WideFieldOphthalmicPhotographyQualityRatingSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1526},
		Name: "WideFieldOphthalmicPhotographyQualityThresholdSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1527},
		Name: "WideFieldOphthalmicPhotographyThresholdQualityRating",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1528},
		Name: "XCoordinatesCenterPixelViewAngle",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1529},
		Name: "YCoordinatesCenterPixelViewAngle",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1530},
		Name: "NumberOfMapPoints",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1531},
		Name: "TwoDimensionalToThreeDimensionalMapData",
		VR:   "OF",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1612},
		Name: "DerivationAlgorithmSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1615},
		Name: "OphthalmicImageTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1616},
		Name: "OphthalmicImageTypeDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1618},
		Name: "ScanPatternTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1620},
		Name: "ReferencedSurfaceMeshIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1622},
		Name: "OphthalmicVolumetricPropertiesFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1624},
		Name: "OphthalmicAnatomicReferencePointXCoordinate",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1626},
		Name: "OphthalmicAnatomicReferencePointYCoordinate",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1628},
		Name: "OphthalmicEnFaceImageQualityRatingSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1630},
		Name: "QualityThreshold",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1640},
		Name: "OCTBscanAnalysisAcquisitionParametersSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1642},
		Name: "NumberofBscansPerFrame",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1643},
		Name: "BscanSlabThickness",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1644},
		Name: "DistanceBetweenBscanSlabs",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1645},
		Name: "BscanCycleTime",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1646},
		Name: "BscanCycleTimeVector",
		VR:   "FL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1649},
		Name: "AscanRate",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1650},
		Name: "BscanRate",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0022, 0x1658},
		Name: "SurfaceMeshZPixelOffset",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0010},
		Name: "VisualFieldHorizontalExtent",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0011},
		Name: "VisualFieldVerticalExtent",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0012},
		Name: "VisualFieldShape",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0016},
		Name: "ScreeningTestModeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0018},
		Name: "MaximumStimulusLuminance",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0020},
		Name: "BackgroundLuminance",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0021},
		Name: "StimulusColorCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0024},
		Name: "BackgroundIlluminationColorCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0025},
		Name: "StimulusArea",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0028},
		Name: "StimulusPresentationTime",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0032},
		Name: "FixationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0033},
		Name: "FixationMonitoringCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0034},
		Name: "VisualFieldCatchTrialSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0035},
		Name: "FixationCheckedQuantity",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0036},
		Name: "PatientNotProperlyFixatedQuantity",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0037},
		Name: "PresentedVisualStimuliDataFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0038},
		Name: "NumberOfVisualStimuli",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0039},
		Name: "ExcessiveFixationLossesDataFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0040},
		Name: "ExcessiveFixationLosses",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0042},
		Name: "StimuliRetestingQuantity",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0044},
		Name: "CommentsOnPatientPerformanceOfVisualField",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0045},
		Name: "FalseNegativesEstimateFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0046},
		Name: "FalseNegativesEstimate",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0048},
		Name: "NegativeCatchTrialsQuantity",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0050},
		Name: "FalseNegativesQuantity",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0051},
		Name: "ExcessiveFalseNegativesDataFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0052},
		Name: "ExcessiveFalseNegatives",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0053},
		Name: "FalsePositivesEstimateFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0054},
		Name: "FalsePositivesEstimate",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0055},
		Name: "CatchTrialsDataFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0056},
		Name: "PositiveCatchTrialsQuantity",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0057},
		Name: "TestPointNormalsDataFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0058},
		Name: "TestPointNormalsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0059},
		Name: "GlobalDeviationProbabilityNormalsFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0060},
		Name: "FalsePositivesQuantity",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0061},
		Name: "ExcessiveFalsePositivesDataFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0062},
		Name: "ExcessiveFalsePositives",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0063},
		Name: "VisualFieldTestNormalsFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0064},
		Name: "ResultsNormalsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0065},
		Name: "AgeCorrectedSensitivityDeviationAlgorithmSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0066},
		Name: "GlobalDeviationFromNormal",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0067},
		Name: "GeneralizedDefectSensitivityDeviationAlgorithmSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0068},
		Name: "LocalizedDeviationFromNormal",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0069},
		Name: "PatientReliabilityIndicator",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0070},
		Name: "VisualFieldMeanSensitivity",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0071},
		Name: "GlobalDeviationProbability",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0072},
		Name: "LocalDeviationProbabilityNormalsFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0073},
		Name: "LocalizedDeviationProbability",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0074},
		Name: "ShortTermFluctuationCalculated",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0075},
		Name: "ShortTermFluctuation",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0076},
		Name: "ShortTermFluctuationProbabilityCalculated",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0077},
		Name: "ShortTermFluctuationProbability",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0078},
		Name: "CorrectedLocalizedDeviationFromNormalCalculated",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0079},
		Name: "CorrectedLocalizedDeviationFromNormal",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0080},
		Name: "CorrectedLocalizedDeviationFromNormalProbabilityCalculated",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0081},
		Name: "CorrectedLocalizedDeviationFromNormalProbability",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0083},
		Name: "GlobalDeviationProbabilitySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0085},
		Name: "LocalizedDeviationProbabilitySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0086},
		Name: "FovealSensitivityMeasured",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0087},
		Name: "FovealSensitivity",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0088},
		Name: "VisualFieldTestDuration",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0089},
		Name: "VisualFieldTestPointSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0090},
		Name: "VisualFieldTestPointXCoordinate",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0091},
		Name: "VisualFieldTestPointYCoordinate",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0092},
		Name: "AgeCorrectedSensitivityDeviationValue",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0093},
		Name: "StimulusResults",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0094},
		Name: "SensitivityValue",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0095},
		Name: "RetestStimulusSeen",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0096},
		Name: "RetestSensitivityValue",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0097},
		Name: "VisualFieldTestPointNormalsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0098},
		Name: "QuantifiedDefect",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0100},
		Name: "AgeCorrectedSensitivityDeviationProbabilityValue",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0102},
		Name: "GeneralizedDefectCorrectedSensitivityDeviationFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0103},
		Name: "GeneralizedDefectCorrectedSensitivityDeviationValue",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0104},
		Name: "GeneralizedDefectCorrectedSensitivityDeviationProbabilityValue",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0105},
		Name: "MinimumSensitivityValue",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0106},
		Name: "BlindSpotLocalized",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0107},
		Name: "BlindSpotXCoordinate",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0108},
		Name: "BlindSpotYCoordinate",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0110},
		Name: "VisualAcuityMeasurementSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0112},
		Name: "RefractiveParametersUsedOnPatientSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0113},
		Name: "MeasurementLaterality",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0114},
		Name: "OphthalmicPatientClinicalInformationLeftEyeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0115},
		Name: "OphthalmicPatientClinicalInformationRightEyeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0117},
		Name: "FovealPointNormativeDataFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0118},
		Name: "FovealPointProbabilityValue",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0120},
		Name: "ScreeningBaselineMeasured",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0122},
		Name: "ScreeningBaselineMeasuredSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0124},
		Name: "ScreeningBaselineType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0126},
		Name: "ScreeningBaselineValue",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0202},
		Name: "AlgorithmSource",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0306},
		Name: "DataSetName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0307},
		Name: "DataSetVersion",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0308},
		Name: "DataSetSource",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0309},
		Name: "DataSetDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0317},
		Name: "VisualFieldTestReliabilityGlobalIndexSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0320},
		Name: "VisualFieldGlobalResultsIndexSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0325},
		Name: "DataObservationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0338},
		Name: "IndexNormalsFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0341},
		Name: "IndexProbability",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0024, 0x0344},
		Name: "IndexProbabilitySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0002},
		Name: "SamplesPerPixel",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0003},
		Name: "SamplesPerPixelUsed",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0004},
		Name: "PhotometricInterpretation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0006},
		Name: "PlanarConfiguration",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0008},
		Name: "NumberOfFrames",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0009},
		Name: "FrameIncrementPointer",
		VR:   "AT",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x000a},
		Name: "FrameDimensionPointer",
		VR:   "AT",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0010},
		Name: "Rows",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0011},
		Name: "Columns",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0014},
		Name: "UltrasoundColorDataPresent",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0030},
		Name: "PixelSpacing",
		VR:   "DS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0031},
		Name: "ZoomFactor",
		VR:   "DS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0032},
		Name: "ZoomCenter",
		VR:   "DS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0034},
		Name: "PixelAspectRatio",
		VR:   "IS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0051},
		Name: "CorrectedImage",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0100},
		Name: "BitsAllocated",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0101},
		Name: "BitsStored",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0102},
		Name: "HighBit",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0103},
		Name: "PixelRepresentation",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0106},
		Name: "SmallestImagePixelValue",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0107},
		Name: "LargestImagePixelValue",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0108},
		Name: "SmallestPixelValueInSeries",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0109},
		Name: "LargestPixelValueInSeries",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0120},
		Name: "PixelPaddingValue",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0121},
		Name: "PixelPaddingRangeLimit",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0122},
		Name: "FloatPixelPaddingValue",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0123},
		Name: "DoubleFloatPixelPaddingValue",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0124},
		Name: "FloatPixelPaddingRangeLimit",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0125},
		Name: "DoubleFloatPixelPaddingRangeLimit",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0300},
		Name: "QualityControlImage",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0301},
		Name: "BurnedInAnnotation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0302},
		Name: "RecognizableVisualFeatures",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0303},
		Name: "LongitudinalTemporalInformationModified",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0304},
		Name: "ReferencedColorPaletteInstanceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0a02},
		Name: "PixelSpacingCalibrationType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x0a04},
		Name: "PixelSpacingCalibrationDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1040},
		Name: "PixelIntensityRelationship",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1041},
		Name: "PixelIntensityRelationshipSign",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1050},
		Name: "WindowCenter",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1051},
		Name: "WindowWidth",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1052},
		Name: "RescaleIntercept",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1053},
		Name: "RescaleSlope",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1054},
		Name: "RescaleType",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1055},
		Name: "WindowCenterWidthExplanation",
		VR:   "LO",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1056},
		Name: "VOILUTFunction",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1090},
		Name: "RecommendedViewingMode",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1101},
		Name: "RedPaletteColorLookupTableDescriptor",
		VR:   "SS",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1102},
		Name: "GreenPaletteColorLookupTableDescriptor",
		VR:   "SS",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1103},
		Name: "BluePaletteColorLookupTableDescriptor",
		VR:   "SS",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1104},
		Name: "AlphaPaletteColorLookupTableDescriptor",
		VR:   "US",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1199},
		Name: "PaletteColorLookupTableUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1201},
		Name: "RedPaletteColorLookupTableData",
		VR:   "OW",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1202},
		Name: "GreenPaletteColorLookupTableData",
		VR:   "OW",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1203},
		Name: "BluePaletteColorLookupTableData",
		VR:   "OW",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1204},
		Name: "AlphaPaletteColorLookupTableData",
		VR:   "OW",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1221},
		Name: "SegmentedRedPaletteColorLookupTableData",
		VR:   "OW",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1222},
		Name: "SegmentedGreenPaletteColorLookupTableData",
		VR:   "OW",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1223},
		Name: "SegmentedBluePaletteColorLookupTableData",
		VR:   "OW",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1224},
		Name: "SegmentedAlphaPaletteColorLookupTableData",
		VR:   "OW",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1230},
		Name: "StoredValueColorRangeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1231},
		Name: "MinimumStoredValueMapped",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1232},
		Name: "MaximumStoredValueMapped",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1300},
		Name: "BreastImplantPresent",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1350},
		Name: "PartialView",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1351},
		Name: "PartialViewDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1352},
		Name: "PartialViewCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x135a},
		Name: "SpatialLocationsPreserved",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1401},
		Name: "DataFrameAssignmentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1402},
		Name: "DataPathAssignment",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1403},
		Name: "BitsMappedToColorLookupTable",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1404},
		Name: "BlendingLUT1Sequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1405},
		Name: "BlendingLUT1TransferFunction",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1406},
		Name: "BlendingWeightConstant",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1407},
		Name: "BlendingLookupTableDescriptor",
		VR:   "US",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1408},
		Name: "BlendingLookupTableData",
		VR:   "OW",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x140b},
		Name: "EnhancedPaletteColorLookupTableSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x140c},
		Name: "BlendingLUT2Sequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x140d},
		Name: "BlendingLUT2TransferFunction",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x140e},
		Name: "DataPathID",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x140f},
		Name: "RGBLUTTransferFunction",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x1410},
		Name: "AlphaLUTTransferFunction",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x2000},
		Name: "ICCProfile",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x2002},
		Name: "ColorSpace",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x2110},
		Name: "LossyImageCompression",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x2112},
		Name: "LossyImageCompressionRatio",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x2114},
		Name: "LossyImageCompressionMethod",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x3000},
		Name: "ModalityLUTSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x3002},
		Name: "LUTDescriptor",
		VR:   "SS",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x3003},
		Name: "LUTExplanation",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x3004},
		Name: "ModalityLUTType",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x3006},
		Name: "LUTData",
		VR:   "OW",
		VM:   "1-n or 1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x3010},
		Name: "VOILUTSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x3110},
		Name: "SoftcopyVOILUTSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x6010},
		Name: "RepresentativeFrameNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x6020},
		Name: "FrameNumbersOfInterest",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x6022},
		Name: "FrameOfInterestDescription",
		VR:   "LO",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x6023},
		Name: "FrameOfInterestType",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x6040},
		Name: "RWavePointer",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x6100},
		Name: "MaskSubtractionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x6101},
		Name: "MaskOperation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x6102},
		Name: "ApplicableFrameRange",
		VR:   "US",
		VM:   "2-2n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x6110},
		Name: "MaskFrameNumbers",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x6112},
		Name: "ContrastFrameAveraging",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x6114},
		Name: "MaskSubPixelShift",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x6120},
		Name: "TIDOffset",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x6190},
		Name: "MaskOperationExplanation",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7000},
		Name: "EquipmentAdministratorSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7001},
		Name: "NumberOfDisplaySubsystems",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7002},
		Name: "CurrentConfigurationID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7003},
		Name: "DisplaySubsystemID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7004},
		Name: "DisplaySubsystemName",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7005},
		Name: "DisplaySubsystemDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7006},
		Name: "SystemStatus",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7007},
		Name: "SystemStatusComment",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7008},
		Name: "TargetLuminanceCharacteristicsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7009},
		Name: "LuminanceCharacteristicsID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x700a},
		Name: "DisplaySubsystemConfigurationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x700b},
		Name: "ConfigurationID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x700c},
		Name: "ConfigurationName",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x700d},
		Name: "ConfigurationDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x700e},
		Name: "ReferencedTargetLuminanceCharacteristicsID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x700f},
		Name: "QAResultsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7010},
		Name: "DisplaySubsystemQAResultsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7011},
		Name: "ConfigurationQAResultsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7012},
		Name: "MeasurementEquipmentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7013},
		Name: "MeasurementFunctions",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7014},
		Name: "MeasurementEquipmentType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7015},
		Name: "VisualEvaluationResultSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7016},
		Name: "DisplayCalibrationResultSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7017},
		Name: "DDLValue",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7018},
		Name: "CIExyWhitePoint",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7019},
		Name: "DisplayFunctionType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x701a},
		Name: "GammaValue",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x701b},
		Name: "NumberOfLuminancePoints",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x701c},
		Name: "LuminanceResponseSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x701d},
		Name: "TargetMinimumLuminance",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x701e},
		Name: "TargetMaximumLuminance",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x701f},
		Name: "LuminanceValue",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7020},
		Name: "LuminanceResponseDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7021},
		Name: "WhitePointFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7022},
		Name: "DisplayDeviceTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7023},
		Name: "DisplaySubsystemSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7024},
		Name: "LuminanceResultSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7025},
		Name: "AmbientLightValueSource",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7026},
		Name: "MeasuredCharacteristics",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7027},
		Name: "LuminanceUniformityResultSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7028},
		Name: "VisualEvaluationTestSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7029},
		Name: "TestResult",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x702a},
		Name: "TestResultComment",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x702b},
		Name: "TestImageValidation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x702c},
		Name: "TestPatternCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x702d},
		Name: "MeasurementPatternCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x702e},
		Name: "VisualEvaluationMethodCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x7fe0},
		Name: "PixelDataProviderURL",
		VR:   "UR",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9001},
		Name: "DataPointRows",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9002},
		Name: "DataPointColumns",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9003},
		Name: "SignalDomainColumns",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9108},
		Name: "DataRepresentation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9110},
		Name: "PixelMeasuresSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9132},
		Name: "FrameVOILUTSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9145},
		Name: "PixelValueTransformationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9235},
		Name: "SignalDomainRows",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9411},
		Name: "DisplayFilterPercentage",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9415},
		Name: "FramePixelShiftSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9416},
		Name: "SubtractionItemID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9422},
		Name: "PixelIntensityRelationshipLUTSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9443},
		Name: "FramePixelDataPropertiesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9444},
		Name: "GeometricalProperties",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9445},
		Name: "GeometricMaximumDistortion",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9446},
		Name: "ImageProcessingApplied",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9454},
		Name: "MaskSelectionMode",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9474},
		Name: "LUTFunction",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9478},
		Name: "MaskVisibilityPercentage",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9501},
		Name: "PixelShiftSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9502},
		Name: "RegionPixelShiftSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9503},
		Name: "VerticesOfTheRegion",
		VR:   "SS",
		VM:   "2-2n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9505},
		Name: "MultiFramePresentationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9506},
		Name: "PixelShiftFrameRange",
		VR:   "US",
		VM:   "2-2n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9507},
		Name: "LUTFrameRange",
		VR:   "US",
		VM:   "2-2n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9520},
		Name: "ImageToEquipmentMappingMatrix",
		VR:   "DS",
		VM:   "16",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0028, 0x9537},
		Name: "EquipmentCoordinateSystemIdentification",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0032, 0x1031},
		Name: "RequestingPhysicianIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0032, 0x1032},
		Name: "RequestingPhysician",
		VR:   "PN",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0032, 0x1033},
		Name: "RequestingService",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0032, 0x1034},
		Name: "RequestingServiceCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0032, 0x1060},
		Name: "RequestedProcedureDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0032, 0x1064},
		Name: "RequestedProcedureCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0032, 0x1066},
		Name: "ReasonForVisit",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0032, 0x1067},
		Name: "ReasonForVisitCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0032, 0x1070},
		Name: "RequestedContrastAgent",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0034, 0x0001},
		Name: "FlowIdentifierSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0034, 0x0002},
		Name: "FlowIdentifier",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0034, 0x0003},
		Name: "FlowTransferSyntaxUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0034, 0x0004},
		Name: "FlowRTPSamplingRate",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0034, 0x0005},
		Name: "SourceIdentifier",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0034, 0x0007},
		Name: "FrameOriginTimestamp",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0034, 0x0008},
		Name: "IncludesImagingSubject",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0034, 0x0009},
		Name: "FrameUsefulnessGroupSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0034, 0x000a},
		Name: "RealTimeBulkDataFlowSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0034, 0x000b},
		Name: "CameraPositionGroupSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0034, 0x000c},
		Name: "IncludesInformation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0034, 0x000d},
		Name: "TimeOfFrameGroupSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0038, 0x0004},
		Name: "ReferencedPatientAliasSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0038, 0x0008},
		Name: "VisitStatusID",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0038, 0x0010},
		Name: "AdmissionID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0038, 0x0014},
		Name: "IssuerOfAdmissionIDSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0038, 0x0016},
		Name: "RouteOfAdmissions",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0038, 0x0020},
		Name: "AdmittingDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0038, 0x0021},
		Name: "AdmittingTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0038, 0x0050},
		Name: "SpecialNeeds",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0038, 0x0060},
		Name: "ServiceEpisodeID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0038, 0x0062},
		Name: "ServiceEpisodeDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0038, 0x0064},
		Name: "IssuerOfServiceEpisodeIDSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0038, 0x0100},
		Name: "PertinentDocumentsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0038, 0x0101},
		Name: "PertinentResourcesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0038, 0x0102},
		Name: "ResourceDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0038, 0x0300},
		Name: "CurrentPatientLocation",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0038, 0x0400},
		Name: "PatientInstitutionResidence",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0038, 0x0500},
		Name: "PatientState",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0038, 0x0502},
		Name: "PatientClinicalTrialParticipationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0038, 0x4000},
		Name: "VisitComments",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0004},
		Name: "WaveformOriginality",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0005},
		Name: "NumberOfWaveformChannels",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0010},
		Name: "NumberOfWaveformSamples",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x001a},
		Name: "SamplingFrequency",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0020},
		Name: "MultiplexGroupLabel",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0200},
		Name: "ChannelDefinitionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0202},
		Name: "WaveformChannelNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0203},
		Name: "ChannelLabel",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0205},
		Name: "ChannelStatus",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0208},
		Name: "ChannelSourceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0209},
		Name: "ChannelSourceModifiersSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x020a},
		Name: "SourceWaveformSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x020c},
		Name: "ChannelDerivationDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0210},
		Name: "ChannelSensitivity",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0211},
		Name: "ChannelSensitivityUnitsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0212},
		Name: "ChannelSensitivityCorrectionFactor",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0213},
		Name: "ChannelBaseline",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0214},
		Name: "ChannelTimeSkew",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0215},
		Name: "ChannelSampleSkew",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0218},
		Name: "ChannelOffset",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x021a},
		Name: "WaveformBitsStored",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0220},
		Name: "FilterLowFrequency",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0221},
		Name: "FilterHighFrequency",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0222},
		Name: "NotchFilterFrequency",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0223},
		Name: "NotchFilterBandwidth",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0230},
		Name: "WaveformDataDisplayScale",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0231},
		Name: "WaveformDisplayBackgroundCIELabValue",
		VR:   "US",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0240},
		Name: "WaveformPresentationGroupSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0241},
		Name: "PresentationGroupNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0242},
		Name: "ChannelDisplaySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0244},
		Name: "ChannelRecommendedDisplayCIELabValue",
		VR:   "US",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0245},
		Name: "ChannelPosition",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0246},
		Name: "DisplayShadingFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0247},
		Name: "FractionalChannelDisplayScale",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0248},
		Name: "AbsoluteChannelDisplayScale",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0300},
		Name: "MultiplexedAudioChannelsDescriptionCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0301},
		Name: "ChannelIdentificationCode",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x003a, 0x0302},
		Name: "ChannelMode",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0001},
		Name: "ScheduledStationAETitle",
		VR:   "AE",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0002},
		Name: "ScheduledProcedureStepStartDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0003},
		Name: "ScheduledProcedureStepStartTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0004},
		Name: "ScheduledProcedureStepEndDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0005},
		Name: "ScheduledProcedureStepEndTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0006},
		Name: "ScheduledPerformingPhysicianName",
		VR:   "PN",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0007},
		Name: "ScheduledProcedureStepDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0008},
		Name: "ScheduledProtocolCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0009},
		Name: "ScheduledProcedureStepID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x000a},
		Name: "StageCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x000b},
		Name: "ScheduledPerformingPhysicianIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0010},
		Name: "ScheduledStationName",
		VR:   "SH",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0011},
		Name: "ScheduledProcedureStepLocation",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0012},
		Name: "PreMedication",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0020},
		Name: "ScheduledProcedureStepStatus",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0026},
		Name: "OrderPlacerIdentifierSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0027},
		Name: "OrderFillerIdentifierSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0031},
		Name: "LocalNamespaceEntityID",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0032},
		Name: "UniversalEntityID",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0033},
		Name: "UniversalEntityIDType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0035},
		Name: "IdentifierTypeCode",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0036},
		Name: "AssigningFacilitySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0039},
		Name: "AssigningJurisdictionCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x003a},
		Name: "AssigningAgencyOrDepartmentCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0100},
		Name: "ScheduledProcedureStepSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0220},
		Name: "ReferencedNonImageCompositeSOPInstanceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0241},
		Name: "PerformedStationAETitle",
		VR:   "AE",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0242},
		Name: "PerformedStationName",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0243},
		Name: "PerformedLocation",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0244},
		Name: "PerformedProcedureStepStartDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0245},
		Name: "PerformedProcedureStepStartTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0250},
		Name: "PerformedProcedureStepEndDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0251},
		Name: "PerformedProcedureStepEndTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0252},
		Name: "PerformedProcedureStepStatus",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0253},
		Name: "PerformedProcedureStepID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0254},
		Name: "PerformedProcedureStepDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0255},
		Name: "PerformedProcedureTypeDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0260},
		Name: "PerformedProtocolCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0261},
		Name: "PerformedProtocolType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0270},
		Name: "ScheduledStepAttributesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0275},
		Name: "RequestAttributesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0280},
		Name: "CommentsOnThePerformedProcedureStep",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0281},
		Name: "PerformedProcedureStepDiscontinuationReasonCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0293},
		Name: "QuantitySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0294},
		Name: "Quantity",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0295},
		Name: "MeasuringUnitsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0296},
		Name: "BillingItemSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0302},
		Name: "EntranceDose",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0303},
		Name: "ExposedArea",
		VR:   "US",
		VM:   "1-2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0306},
		Name: "DistanceSourceToEntrance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0310},
		Name: "CommentsOnRadiationDose",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0312},
		Name: "XRayOutput",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0314},
		Name: "HalfValueLayer",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0316},
		Name: "OrganDose",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0318},
		Name: "OrganExposed",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0320},
		Name: "BillingProcedureStepSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0321},
		Name: "FilmConsumptionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0324},
		Name: "BillingSuppliesAndDevicesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0340},
		Name: "PerformedSeriesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0400},
		Name: "CommentsOnTheScheduledProcedureStep",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0440},
		Name: "ProtocolContextSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0441},
		Name: "ContentItemModifierSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0500},
		Name: "ScheduledSpecimenSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0512},
		Name: "ContainerIdentifier",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0513},
		Name: "IssuerOfTheContainerIdentifierSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0515},
		Name: "AlternateContainerIdentifierSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0518},
		Name: "ContainerTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x051a},
		Name: "ContainerDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0520},
		Name: "ContainerComponentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0551},
		Name: "SpecimenIdentifier",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0554},
		Name: "SpecimenUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0555},
		Name: "AcquisitionContextSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0556},
		Name: "AcquisitionContextDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x059a},
		Name: "SpecimenTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0560},
		Name: "SpecimenDescriptionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0562},
		Name: "IssuerOfTheSpecimenIdentifierSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0600},
		Name: "SpecimenShortDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0602},
		Name: "SpecimenDetailedDescription",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0610},
		Name: "SpecimenPreparationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0612},
		Name: "SpecimenPreparationStepContentItemSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0620},
		Name: "SpecimenLocalizationContentItemSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x0710},
		Name: "WholeSlideMicroscopyImageFrameTypeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x071a},
		Name: "ImageCenterPointCoordinatesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x072a},
		Name: "XOffsetInSlideCoordinateSystem",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x073a},
		Name: "YOffsetInSlideCoordinateSystem",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x074a},
		Name: "ZOffsetInSlideCoordinateSystem",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x08ea},
		Name: "MeasurementUnitsCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x1001},
		Name: "RequestedProcedureID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x1002},
		Name: "ReasonForTheRequestedProcedure",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x1003},
		Name: "RequestedProcedurePriority",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x1004},
		Name: "PatientTransportArrangements",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x1005},
		Name: "RequestedProcedureLocation",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x1008},
		Name: "ConfidentialityCode",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x1009},
		Name: "ReportingPriority",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x100a},
		Name: "ReasonForRequestedProcedureCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x1010},
		Name: "NamesOfIntendedRecipientsOfResults",
		VR:   "PN",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x1011},
		Name: "IntendedRecipientsOfResultsIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x1012},
		Name: "ReasonForPerformedProcedureCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x1101},
		Name: "PersonIdentificationCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x1102},
		Name: "PersonAddress",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x1103},
		Name: "PersonTelephoneNumbers",
		VR:   "LO",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x1104},
		Name: "PersonTelecomInformation",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x1400},
		Name: "RequestedProcedureComments",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x2004},
		Name: "IssueDateOfImagingServiceRequest",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x2005},
		Name: "IssueTimeOfImagingServiceRequest",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x2008},
		Name: "OrderEnteredBy",
		VR:   "PN",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x2009},
		Name: "OrderEntererLocation",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x2010},
		Name: "OrderCallbackPhoneNumber",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x2011},
		Name: "OrderCallbackTelecomInformation",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x2016},
		Name: "PlacerOrderNumberImagingServiceRequest",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x2017},
		Name: "FillerOrderNumberImagingServiceRequest",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x2400},
		Name: "ImagingServiceRequestComments",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x3001},
		Name: "ConfidentialityConstraintOnPatientDataDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4005},
		Name: "ScheduledProcedureStepStartDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4008},
		Name: "ScheduledProcedureStepExpirationDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4009},
		Name: "HumanPerformerCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4010},
		Name: "ScheduledProcedureStepModificationDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4011},
		Name: "ExpectedCompletionDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4018},
		Name: "ScheduledWorkitemCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4019},
		Name: "PerformedWorkitemCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4021},
		Name: "InputInformationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4025},
		Name: "ScheduledStationNameCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4026},
		Name: "ScheduledStationClassCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4027},
		Name: "ScheduledStationGeographicLocationCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4028},
		Name: "PerformedStationNameCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4029},
		Name: "PerformedStationClassCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4030},
		Name: "PerformedStationGeographicLocationCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4033},
		Name: "OutputInformationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4034},
		Name: "ScheduledHumanPerformersSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4035},
		Name: "ActualHumanPerformersSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4036},
		Name: "HumanPerformerOrganization",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4037},
		Name: "HumanPerformerName",
		VR:   "PN",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4040},
		Name: "RawDataHandling",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4041},
		Name: "InputReadinessState",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4050},
		Name: "PerformedProcedureStepStartDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4051},
		Name: "PerformedProcedureStepEndDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4052},
		Name: "ProcedureStepCancellationDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4070},
		Name: "OutputDestinationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4071},
		Name: "DICOMStorageSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4072},
		Name: "STOWRSStorageSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4073},
		Name: "StorageURL",
		VR:   "UR",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x4074},
		Name: "XDSStorageSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x8302},
		Name: "EntranceDoseInmGy",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x8303},
		Name: "EntranceDoseDerivation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x9092},
		Name: "ParametricMapFrameTypeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x9094},
		Name: "ReferencedImageRealWorldValueMappingSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x9096},
		Name: "RealWorldValueMappingSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x9098},
		Name: "PixelValueMappingCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x9210},
		Name: "LUTLabel",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x9211},
		Name: "RealWorldValueLastValueMapped",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x9212},
		Name: "RealWorldValueLUTData",
		VR:   "FD",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x9213},
		Name: "DoubleFloatRealWorldValueLastValueMapped",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x9214},
		Name: "DoubleFloatRealWorldValueFirstValueMapped",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x9216},
		Name: "RealWorldValueFirstValueMapped",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x9220},
		Name: "QuantityDefinitionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x9224},
		Name: "RealWorldValueIntercept",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0x9225},
		Name: "RealWorldValueSlope",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa010},
		Name: "RelationshipType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa027},
		Name: "VerifyingOrganization",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa030},
		Name: "VerificationDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa032},
		Name: "ObservationDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa040},
		Name: "ValueType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa043},
		Name: "ConceptNameCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa050},
		Name: "ContinuityOfContent",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa073},
		Name: "VerifyingObserverSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa075},
		Name: "VerifyingObserverName",
		VR:   "PN",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa078},
		Name: "AuthorObserverSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa07a},
		Name: "ParticipantSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa07c},
		Name: "CustodialOrganizationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa080},
		Name: "ParticipationType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa082},
		Name: "ParticipationDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa084},
		Name: "ObserverType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa088},
		Name: "VerifyingObserverIdentificationCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa0b0},
		Name: "ReferencedWaveformChannels",
		VR:   "US",
		VM:   "2-2n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa120},
		Name: "DateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa121},
		Name: "Date",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa122},
		Name: "Time",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa123},
		Name: "PersonName",
		VR:   "PN",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa124},
		Name: "UID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa130},
		Name: "TemporalRangeType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa132},
		Name: "ReferencedSamplePositions",
		VR:   "UL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa138},
		Name: "ReferencedTimeOffsets",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa13a},
		Name: "ReferencedDateTime",
		VR:   "DT",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa160},
		Name: "TextValue",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa161},
		Name: "FloatingPointValue",
		VR:   "FD",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa162},
		Name: "RationalNumeratorValue",
		VR:   "SL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa163},
		Name: "RationalDenominatorValue",
		VR:   "UL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa168},
		Name: "ConceptCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa170},
		Name: "PurposeOfReferenceCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa171},
		Name: "ObservationUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa180},
		Name: "AnnotationGroupNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa195},
		Name: "ModifierCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa300},
		Name: "MeasuredValueSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa301},
		Name: "NumericValueQualifierCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa30a},
		Name: "NumericValue",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa360},
		Name: "PredecessorDocumentsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa370},
		Name: "ReferencedRequestSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa372},
		Name: "PerformedProcedureCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa375},
		Name: "CurrentRequestedProcedureEvidenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa385},
		Name: "PertinentOtherEvidenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa390},
		Name: "HL7StructuredDocumentReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa491},
		Name: "CompletionFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa492},
		Name: "CompletionFlagDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa493},
		Name: "VerificationFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa494},
		Name: "ArchiveRequested",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa496},
		Name: "PreliminaryFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa504},
		Name: "ContentTemplateSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa525},
		Name: "IdenticalDocumentsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xa730},
		Name: "ContentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xb020},
		Name: "WaveformAnnotationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xdb00},
		Name: "TemplateIdentifier",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xdb73},
		Name: "ReferencedContentItemIdentifier",
		VR:   "UL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xe001},
		Name: "HL7InstanceIdentifier",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xe004},
		Name: "HL7DocumentEffectiveTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xe006},
		Name: "HL7DocumentTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xe008},
		Name: "DocumentClassCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xe010},
		Name: "RetrieveURI",
		VR:   "UR",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xe011},
		Name: "RetrieveLocationUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xe020},
		Name: "TypeOfInstances",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xe021},
		Name: "DICOMRetrievalSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xe022},
		Name: "DICOMMediaRetrievalSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xe023},
		Name: "WADORetrievalSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xe024},
		Name: "XDSRetrievalSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xe025},
		Name: "WADORSRetrievalSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xe030},
		Name: "RepositoryUniqueID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0040, 0xe031},
		Name: "HomeCommunityID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0042, 0x0010},
		Name: "DocumentTitle",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0042, 0x0011},
		Name: "EncapsulatedDocument",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0042, 0x0012},
		Name: "MIMETypeOfEncapsulatedDocument",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0042, 0x0013},
		Name: "SourceInstanceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0042, 0x0014},
		Name: "ListOfMIMETypes",
		VR:   "LO",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0042, 0x0015},
		Name: "EncapsulatedDocumentLength",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0001},
		Name: "ProductPackageIdentifier",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0002},
		Name: "SubstanceAdministrationApproval",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0003},
		Name: "ApprovalStatusFurtherDescription",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0004},
		Name: "ApprovalStatusDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0007},
		Name: "ProductTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0008},
		Name: "ProductName",
		VR:   "LO",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0009},
		Name: "ProductDescription",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x000a},
		Name: "ProductLotIdentifier",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x000b},
		Name: "ProductExpirationDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0010},
		Name: "SubstanceAdministrationDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0011},
		Name: "SubstanceAdministrationNotes",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0012},
		Name: "SubstanceAdministrationDeviceID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0013},
		Name: "ProductParameterSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0019},
		Name: "SubstanceAdministrationParameterSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0100},
		Name: "ApprovalSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0101},
		Name: "AssertionCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0102},
		Name: "AssertionUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0103},
		Name: "AsserterIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0104},
		Name: "AssertionDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0105},
		Name: "AssertionExpirationDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0106},
		Name: "AssertionComments",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0107},
		Name: "RelatedAssertionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0108},
		Name: "ReferencedAssertionUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x0109},
		Name: "ApprovalSubjectSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0044, 0x010a},
		Name: "OrganizationalRoleCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0012},
		Name: "LensDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0014},
		Name: "RightLensSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0015},
		Name: "LeftLensSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0016},
		Name: "UnspecifiedLateralityLensSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0018},
		Name: "CylinderSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0028},
		Name: "PrismSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0030},
		Name: "HorizontalPrismPower",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0032},
		Name: "HorizontalPrismBase",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0034},
		Name: "VerticalPrismPower",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0036},
		Name: "VerticalPrismBase",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0038},
		Name: "LensSegmentType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0040},
		Name: "OpticalTransmittance",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0042},
		Name: "ChannelWidth",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0044},
		Name: "PupilSize",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0046},
		Name: "CornealSize",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0047},
		Name: "CornealSizeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0050},
		Name: "AutorefractionRightEyeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0052},
		Name: "AutorefractionLeftEyeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0060},
		Name: "DistancePupillaryDistance",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0062},
		Name: "NearPupillaryDistance",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0063},
		Name: "IntermediatePupillaryDistance",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0064},
		Name: "OtherPupillaryDistance",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0070},
		Name: "KeratometryRightEyeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0071},
		Name: "KeratometryLeftEyeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0074},
		Name: "SteepKeratometricAxisSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0075},
		Name: "RadiusOfCurvature",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0076},
		Name: "KeratometricPower",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0077},
		Name: "KeratometricAxis",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0080},
		Name: "FlatKeratometricAxisSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0092},
		Name: "BackgroundColor",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0094},
		Name: "Optotype",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0095},
		Name: "OptotypePresentation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0097},
		Name: "SubjectiveRefractionRightEyeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0098},
		Name: "SubjectiveRefractionLeftEyeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0100},
		Name: "AddNearSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0101},
		Name: "AddIntermediateSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0102},
		Name: "AddOtherSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0104},
		Name: "AddPower",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0106},
		Name: "ViewingDistance",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0110},
		Name: "CorneaMeasurementsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0111},
		Name: "SourceOfCorneaMeasurementDataCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0112},
		Name: "SteepCornealAxisSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0113},
		Name: "FlatCornealAxisSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0114},
		Name: "CornealPower",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0115},
		Name: "CornealAxis",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0116},
		Name: "CorneaMeasurementMethodCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0117},
		Name: "RefractiveIndexOfCornea",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0118},
		Name: "RefractiveIndexOfAqueousHumor",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0121},
		Name: "VisualAcuityTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0122},
		Name: "VisualAcuityRightEyeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0123},
		Name: "VisualAcuityLeftEyeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0124},
		Name: "VisualAcuityBothEyesOpenSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0125},
		Name: "ViewingDistanceType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0135},
		Name: "VisualAcuityModifiers",
		VR:   "SS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0137},
		Name: "DecimalVisualAcuity",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0139},
		Name: "OptotypeDetailedDefinition",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0145},
		Name: "ReferencedRefractiveMeasurementsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0146},
		Name: "SpherePower",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0147},
		Name: "CylinderPower",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0201},
		Name: "CornealTopographySurface",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0202},
		Name: "CornealVertexLocation",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0203},
		Name: "PupilCentroidXCoordinate",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0204},
		Name: "PupilCentroidYCoordinate",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0205},
		Name: "EquivalentPupilRadius",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0207},
		Name: "CornealTopographyMapTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0208},
		Name: "VerticesOfTheOutlineOfPupil",
		VR:   "IS",
		VM:   "2-2n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0210},
		Name: "CornealTopographyMappingNormalsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0211},
		Name: "MaximumCornealCurvatureSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0212},
		Name: "MaximumCornealCurvature",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0213},
		Name: "MaximumCornealCurvatureLocation",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0215},
		Name: "MinimumKeratometricSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0218},
		Name: "SimulatedKeratometricCylinderSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0220},
		Name: "AverageCornealPower",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0224},
		Name: "CornealISValue",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0227},
		Name: "AnalyzedArea",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0230},
		Name: "SurfaceRegularityIndex",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0232},
		Name: "SurfaceAsymmetryIndex",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0234},
		Name: "CornealEccentricityIndex",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0236},
		Name: "KeratoconusPredictionIndex",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0238},
		Name: "DecimalPotentialVisualAcuity",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0242},
		Name: "CornealTopographyMapQualityEvaluation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0244},
		Name: "SourceImageCornealProcessedDataSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0247},
		Name: "CornealPointLocation",
		VR:   "FL",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0248},
		Name: "CornealPointEstimated",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0249},
		Name: "AxialPower",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0250},
		Name: "TangentialPower",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0251},
		Name: "RefractivePower",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0252},
		Name: "RelativeElevation",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0046, 0x0253},
		Name: "CornealWavefront",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0001},
		Name: "ImagedVolumeWidth",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0002},
		Name: "ImagedVolumeHeight",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0003},
		Name: "ImagedVolumeDepth",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0006},
		Name: "TotalPixelMatrixColumns",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0007},
		Name: "TotalPixelMatrixRows",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0008},
		Name: "TotalPixelMatrixOriginSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0010},
		Name: "SpecimenLabelInImage",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0011},
		Name: "FocusMethod",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0012},
		Name: "ExtendedDepthOfField",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0013},
		Name: "NumberOfFocalPlanes",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0014},
		Name: "DistanceBetweenFocalPlanes",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0015},
		Name: "RecommendedAbsentPixelCIELabValue",
		VR:   "US",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0100},
		Name: "IlluminatorTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0102},
		Name: "ImageOrientationSlide",
		VR:   "DS",
		VM:   "6",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0105},
		Name: "OpticalPathSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0106},
		Name: "OpticalPathIdentifier",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0107},
		Name: "OpticalPathDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0108},
		Name: "IlluminationColorCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0110},
		Name: "SpecimenReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0111},
		Name: "CondenserLensPower",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0112},
		Name: "ObjectiveLensPower",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0113},
		Name: "ObjectiveLensNumericalAperture",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0120},
		Name: "PaletteColorLookupTableSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0200},
		Name: "ReferencedImageNavigationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0201},
		Name: "TopLeftHandCornerOfLocalizerArea",
		VR:   "US",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0202},
		Name: "BottomRightHandCornerOfLocalizerArea",
		VR:   "US",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0207},
		Name: "OpticalPathIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x021a},
		Name: "PlanePositionSlideSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x021e},
		Name: "ColumnPositionInTotalImagePixelMatrix",
		VR:   "SL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x021f},
		Name: "RowPositionInTotalImagePixelMatrix",
		VR:   "SL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0301},
		Name: "PixelOriginInterpretation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0302},
		Name: "NumberOfOpticalPaths",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0048, 0x0303},
		Name: "TotalPixelMatrixFocalPlanes",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0050, 0x0004},
		Name: "CalibrationImage",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0050, 0x0010},
		Name: "DeviceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0050, 0x0012},
		Name: "ContainerComponentTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0050, 0x0013},
		Name: "ContainerComponentThickness",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0050, 0x0014},
		Name: "DeviceLength",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0050, 0x0015},
		Name: "ContainerComponentWidth",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0050, 0x0016},
		Name: "DeviceDiameter",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0050, 0x0017},
		Name: "DeviceDiameterUnits",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0050, 0x0018},
		Name: "DeviceVolume",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0050, 0x0019},
		Name: "InterMarkerDistance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0050, 0x001a},
		Name: "ContainerComponentMaterial",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0050, 0x001b},
		Name: "ContainerComponentID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0050, 0x001c},
		Name: "ContainerComponentLength",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0050, 0x001d},
		Name: "ContainerComponentDiameter",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0050, 0x001e},
		Name: "ContainerComponentDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0050, 0x0020},
		Name: "DeviceDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0050, 0x0021},
		Name: "LongDeviceDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0001},
		Name: "ContrastBolusIngredientPercentByVolume",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0002},
		Name: "OCTFocalDistance",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0003},
		Name: "BeamSpotSize",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0004},
		Name: "EffectiveRefractiveIndex",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0006},
		Name: "OCTAcquisitionDomain",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0007},
		Name: "OCTOpticalCenterWavelength",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0008},
		Name: "AxialResolution",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0009},
		Name: "RangingDepth",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0011},
		Name: "ALineRate",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0012},
		Name: "ALinesPerFrame",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0013},
		Name: "CatheterRotationalRate",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0014},
		Name: "ALinePixelSpacing",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0016},
		Name: "ModeOfPercutaneousAccessSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0025},
		Name: "IntravascularOCTFrameTypeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0026},
		Name: "OCTZOffsetApplied",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0027},
		Name: "IntravascularFrameContentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0028},
		Name: "IntravascularLongitudinalDistance",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0029},
		Name: "IntravascularOCTFrameContentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0030},
		Name: "OCTZOffsetCorrection",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0031},
		Name: "CatheterDirectionOfRotation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0033},
		Name: "SeamLineLocation",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0034},
		Name: "FirstALineLocation",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0036},
		Name: "SeamLineIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0038},
		Name: "NumberOfPaddedALines",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x0039},
		Name: "InterpolationType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0052, 0x003a},
		Name: "RefractiveIndexApplied",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0010},
		Name: "EnergyWindowVector",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0011},
		Name: "NumberOfEnergyWindows",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0012},
		Name: "EnergyWindowInformationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0013},
		Name: "EnergyWindowRangeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0014},
		Name: "EnergyWindowLowerLimit",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0015},
		Name: "EnergyWindowUpperLimit",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0016},
		Name: "RadiopharmaceuticalInformationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0017},
		Name: "ResidualSyringeCounts",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0018},
		Name: "EnergyWindowName",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0020},
		Name: "DetectorVector",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0021},
		Name: "NumberOfDetectors",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0022},
		Name: "DetectorInformationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0030},
		Name: "PhaseVector",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0031},
		Name: "NumberOfPhases",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0032},
		Name: "PhaseInformationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0033},
		Name: "NumberOfFramesInPhase",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0036},
		Name: "PhaseDelay",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0038},
		Name: "PauseBetweenFrames",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0039},
		Name: "PhaseDescription",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0050},
		Name: "RotationVector",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0051},
		Name: "NumberOfRotations",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0052},
		Name: "RotationInformationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0053},
		Name: "NumberOfFramesInRotation",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0060},
		Name: "RRIntervalVector",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0061},
		Name: "NumberOfRRIntervals",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0062},
		Name: "GatedInformationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0063},
		Name: "DataInformationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0070},
		Name: "TimeSlotVector",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0071},
		Name: "NumberOfTimeSlots",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0072},
		Name: "TimeSlotInformationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0073},
		Name: "TimeSlotTime",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0080},
		Name: "SliceVector",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0081},
		Name: "NumberOfSlices",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0090},
		Name: "AngularViewVector",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0100},
		Name: "TimeSliceVector",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0101},
		Name: "NumberOfTimeSlices",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0200},
		Name: "StartAngle",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0202},
		Name: "TypeOfDetectorMotion",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0210},
		Name: "TriggerVector",
		VR:   "IS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0211},
		Name: "NumberOfTriggersInPhase",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0220},
		Name: "ViewCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0222},
		Name: "ViewModifierCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0300},
		Name: "RadionuclideCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0302},
		Name: "AdministrationRouteCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0304},
		Name: "RadiopharmaceuticalCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0306},
		Name: "CalibrationDataSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0308},
		Name: "EnergyWindowNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0400},
		Name: "ImageID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0410},
		Name: "PatientOrientationCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0412},
		Name: "PatientOrientationModifierCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0414},
		Name: "PatientGantryRelationshipCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0500},
		Name: "SliceProgressionDirection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x0501},
		Name: "ScanProgressionDirection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1000},
		Name: "SeriesType",
		VR:   "CS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1001},
		Name: "Units",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1002},
		Name: "CountsSource",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1004},
		Name: "ReprojectionMethod",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1006},
		Name: "SUVType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1100},
		Name: "RandomsCorrectionMethod",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1101},
		Name: "AttenuationCorrectionMethod",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1102},
		Name: "DecayCorrection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1103},
		Name: "ReconstructionMethod",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1104},
		Name: "DetectorLinesOfResponseUsed",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1105},
		Name: "ScatterCorrectionMethod",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1200},
		Name: "AxialAcceptance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1201},
		Name: "AxialMash",
		VR:   "IS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1202},
		Name: "TransverseMash",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1203},
		Name: "DetectorElementSize",
		VR:   "DS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1210},
		Name: "CoincidenceWindowWidth",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1220},
		Name: "SecondaryCountsType",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1300},
		Name: "FrameReferenceTime",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1310},
		Name: "PrimaryPromptsCountsAccumulated",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1311},
		Name: "SecondaryCountsAccumulated",
		VR:   "IS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1320},
		Name: "SliceSensitivityFactor",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1321},
		Name: "DecayFactor",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1322},
		Name: "DoseCalibrationFactor",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1323},
		Name: "ScatterFractionFactor",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1324},
		Name: "DeadTimeFactor",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0054, 0x1330},
		Name: "ImageIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0060, 0x3000},
		Name: "HistogramSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0060, 0x3002},
		Name: "HistogramNumberOfBins",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0060, 0x3004},
		Name: "HistogramFirstBinValue",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0060, 0x3006},
		Name: "HistogramLastBinValue",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0060, 0x3008},
		Name: "HistogramBinWidth",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0060, 0x3010},
		Name: "HistogramExplanation",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0060, 0x3020},
		Name: "HistogramData",
		VR:   "UL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0062, 0x0001},
		Name: "SegmentationType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0062, 0x0002},
		Name: "SegmentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0062, 0x0003},
		Name: "SegmentedPropertyCategoryCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0062, 0x0004},
		Name: "SegmentNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0062, 0x0005},
		Name: "SegmentLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0062, 0x0006},
		Name: "SegmentDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0062, 0x0007},
		Name: "SegmentationAlgorithmIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0062, 0x0008},
		Name: "SegmentAlgorithmType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0062, 0x0009},
		Name: "SegmentAlgorithmName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0062, 0x000a},
		Name: "SegmentIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0062, 0x000b},
		Name: "ReferencedSegmentNumber",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0062, 0x000c},
		Name: "RecommendedDisplayGrayscaleValue",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0062, 0x000d},
		Name: "RecommendedDisplayCIELabValue",
		VR:   "US",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0062, 0x000e},
		Name: "MaximumFractionalValue",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0062, 0x000f},
		Name: "SegmentedPropertyTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0062, 0x0010},
		Name: "SegmentationFractionalType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0062, 0x0011},
		Name: "SegmentedPropertyTypeModifierCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0062, 0x0012},
		Name: "UsedSegmentsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0062, 0x0013},
		Name: "SegmentsOverlap",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0062, 0x0020},
		Name: "TrackingID",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0062, 0x0021},
		Name: "TrackingUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0064, 0x0002},
		Name: "DeformableRegistrationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0064, 0x0003},
		Name: "SourceFrameOfReferenceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0064, 0x0005},
		Name: "DeformableRegistrationGridSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0064, 0x0007},
		Name: "GridDimensions",
		VR:   "UL",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0064, 0x0008},
		Name: "GridResolution",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0064, 0x0009},
		Name: "VectorGridData",
		VR:   "OF",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0064, 0x000f},
		Name: "PreDeformationMatrixRegistrationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0064, 0x0010},
		Name: "PostDeformationMatrixRegistrationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0001},
		Name: "NumberOfSurfaces",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0002},
		Name: "SurfaceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0003},
		Name: "SurfaceNumber",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0004},
		Name: "SurfaceComments",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0009},
		Name: "SurfaceProcessing",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x000a},
		Name: "SurfaceProcessingRatio",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x000b},
		Name: "SurfaceProcessingDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x000c},
		Name: "RecommendedPresentationOpacity",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x000d},
		Name: "RecommendedPresentationType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x000e},
		Name: "FiniteVolume",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0010},
		Name: "Manifold",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0011},
		Name: "SurfacePointsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0012},
		Name: "SurfacePointsNormalsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0013},
		Name: "SurfaceMeshPrimitivesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0015},
		Name: "NumberOfSurfacePoints",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0016},
		Name: "PointCoordinatesData",
		VR:   "OF",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0017},
		Name: "PointPositionAccuracy",
		VR:   "FL",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0018},
		Name: "MeanPointDistance",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0019},
		Name: "MaximumPointDistance",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x001a},
		Name: "PointsBoundingBoxCoordinates",
		VR:   "FL",
		VM:   "6",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x001b},
		Name: "AxisOfRotation",
		VR:   "FL",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x001c},
		Name: "CenterOfRotation",
		VR:   "FL",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x001e},
		Name: "NumberOfVectors",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x001f},
		Name: "VectorDimensionality",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0020},
		Name: "VectorAccuracy",
		VR:   "FL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0021},
		Name: "VectorCoordinateData",
		VR:   "OF",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0026},
		Name: "TriangleStripSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0027},
		Name: "TriangleFanSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0028},
		Name: "LineSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x002a},
		Name: "SurfaceCount",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x002b},
		Name: "ReferencedSurfaceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x002c},
		Name: "ReferencedSurfaceNumber",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x002d},
		Name: "SegmentSurfaceGenerationAlgorithmIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x002e},
		Name: "SegmentSurfaceSourceInstanceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x002f},
		Name: "AlgorithmFamilyCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0030},
		Name: "AlgorithmNameCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0031},
		Name: "AlgorithmVersion",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0032},
		Name: "AlgorithmParameters",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0034},
		Name: "FacetSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0035},
		Name: "SurfaceProcessingAlgorithmIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0036},
		Name: "AlgorithmName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0037},
		Name: "RecommendedPointRadius",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0038},
		Name: "RecommendedLineThickness",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0040},
		Name: "LongPrimitivePointIndexList",
		VR:   "OL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0041},
		Name: "LongTrianglePointIndexList",
		VR:   "OL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0042},
		Name: "LongEdgePointIndexList",
		VR:   "OL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0043},
		Name: "LongVertexPointIndexList",
		VR:   "OL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0101},
		Name: "TrackSetSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0102},
		Name: "TrackSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0103},
		Name: "RecommendedDisplayCIELabValueList",
		VR:   "OW",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0104},
		Name: "TrackingAlgorithmIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0105},
		Name: "TrackSetNumber",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0106},
		Name: "TrackSetLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0107},
		Name: "TrackSetDescription",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0108},
		Name: "TrackSetAnatomicalTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0121},
		Name: "MeasurementsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0124},
		Name: "TrackSetStatisticsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0125},
		Name: "FloatingPointValues",
		VR:   "OF",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0129},
		Name: "TrackPointIndexList",
		VR:   "OL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0130},
		Name: "TrackStatisticsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0132},
		Name: "MeasurementValuesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0133},
		Name: "DiffusionAcquisitionCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0066, 0x0134},
		Name: "DiffusionModelCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6210},
		Name: "ImplantSize",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6221},
		Name: "ImplantTemplateVersion",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6222},
		Name: "ReplacedImplantTemplateSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6223},
		Name: "ImplantType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6224},
		Name: "DerivationImplantTemplateSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6225},
		Name: "OriginalImplantTemplateSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6226},
		Name: "EffectiveDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6230},
		Name: "ImplantTargetAnatomySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6260},
		Name: "InformationFromManufacturerSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6265},
		Name: "NotificationFromManufacturerSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6270},
		Name: "InformationIssueDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6280},
		Name: "InformationSummary",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x62a0},
		Name: "ImplantRegulatoryDisapprovalCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x62a5},
		Name: "OverallTemplateSpatialTolerance",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x62c0},
		Name: "HPGLDocumentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x62d0},
		Name: "HPGLDocumentID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x62d5},
		Name: "HPGLDocumentLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x62e0},
		Name: "ViewOrientationCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x62f0},
		Name: "ViewOrientationModifierCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x62f2},
		Name: "HPGLDocumentScaling",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6300},
		Name: "HPGLDocument",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6310},
		Name: "HPGLContourPenNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6320},
		Name: "HPGLPenSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6330},
		Name: "HPGLPenNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6340},
		Name: "HPGLPenLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6345},
		Name: "HPGLPenDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6346},
		Name: "RecommendedRotationPoint",
		VR:   "FD",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6347},
		Name: "BoundingRectangle",
		VR:   "FD",
		VM:   "4",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6350},
		Name: "ImplantTemplate3DModelSurfaceNumber",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6360},
		Name: "SurfaceModelDescriptionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6380},
		Name: "SurfaceModelLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6390},
		Name: "SurfaceModelScalingFactor",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x63a0},
		Name: "MaterialsCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x63a4},
		Name: "CoatingMaterialsCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x63a8},
		Name: "ImplantTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x63ac},
		Name: "FixationMethodCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x63b0},
		Name: "MatingFeatureSetsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x63c0},
		Name: "MatingFeatureSetID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x63d0},
		Name: "MatingFeatureSetLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x63e0},
		Name: "MatingFeatureSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x63f0},
		Name: "MatingFeatureID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6400},
		Name: "MatingFeatureDegreeOfFreedomSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6410},
		Name: "DegreeOfFreedomID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6420},
		Name: "DegreeOfFreedomType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6430},
		Name: "TwoDMatingFeatureCoordinatesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6440},
		Name: "ReferencedHPGLDocumentID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6450},
		Name: "TwoDMatingPoint",
		VR:   "FD",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6460},
		Name: "TwoDMatingAxes",
		VR:   "FD",
		VM:   "4",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6470},
		Name: "TwoDDegreeOfFreedomSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6490},
		Name: "ThreeDDegreeOfFreedomAxis",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x64a0},
		Name: "RangeOfFreedom",
		VR:   "FD",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x64c0},
		Name: "ThreeDMatingPoint",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x64d0},
		Name: "ThreeDMatingAxes",
		VR:   "FD",
		VM:   "9",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x64f0},
		Name: "TwoDDegreeOfFreedomAxis",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6500},
		Name: "PlanningLandmarkPointSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6510},
		Name: "PlanningLandmarkLineSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6520},
		Name: "PlanningLandmarkPlaneSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6530},
		Name: "PlanningLandmarkID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6540},
		Name: "PlanningLandmarkDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6545},
		Name: "PlanningLandmarkIdentificationCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6550},
		Name: "TwoDPointCoordinatesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6560},
		Name: "TwoDPointCoordinates",
		VR:   "FD",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6590},
		Name: "ThreeDPointCoordinates",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x65a0},
		Name: "TwoDLineCoordinatesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x65b0},
		Name: "TwoDLineCoordinates",
		VR:   "FD",
		VM:   "4",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x65d0},
		Name: "ThreeDLineCoordinates",
		VR:   "FD",
		VM:   "6",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x65e0},
		Name: "TwoDPlaneCoordinatesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x65f0},
		Name: "TwoDPlaneIntersection",
		VR:   "FD",
		VM:   "4",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6610},
		Name: "ThreeDPlaneOrigin",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x6620},
		Name: "ThreeDPlaneNormal",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x7001},
		Name: "ModelModification",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x7002},
		Name: "ModelMirroring",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x7003},
		Name: "ModelUsageCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x7004},
		Name: "ModelGroupUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0068, 0x7005},
		Name: "RelativeURIReferenceWithinEncapsulatedDocument",
		VR:   "UR",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0001},
		Name: "GraphicAnnotationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0002},
		Name: "GraphicLayer",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0003},
		Name: "BoundingBoxAnnotationUnits",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0004},
		Name: "AnchorPointAnnotationUnits",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0005},
		Name: "GraphicAnnotationUnits",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0006},
		Name: "UnformattedTextValue",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0008},
		Name: "TextObjectSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0009},
		Name: "GraphicObjectSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0010},
		Name: "BoundingBoxTopLeftHandCorner",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0011},
		Name: "BoundingBoxBottomRightHandCorner",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0012},
		Name: "BoundingBoxTextHorizontalJustification",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0014},
		Name: "AnchorPoint",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0015},
		Name: "AnchorPointVisibility",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0020},
		Name: "GraphicDimensions",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0021},
		Name: "NumberOfGraphicPoints",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0022},
		Name: "GraphicData",
		VR:   "FL",
		VM:   "2-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0023},
		Name: "GraphicType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0024},
		Name: "GraphicFilled",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0041},
		Name: "ImageHorizontalFlip",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0042},
		Name: "ImageRotation",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0052},
		Name: "DisplayedAreaTopLeftHandCorner",
		VR:   "SL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0053},
		Name: "DisplayedAreaBottomRightHandCorner",
		VR:   "SL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x005a},
		Name: "DisplayedAreaSelectionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0060},
		Name: "GraphicLayerSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0062},
		Name: "GraphicLayerOrder",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0066},
		Name: "GraphicLayerRecommendedDisplayGrayscaleValue",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0068},
		Name: "GraphicLayerDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0080},
		Name: "ContentLabel",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0081},
		Name: "ContentDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0082},
		Name: "PresentationCreationDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0083},
		Name: "PresentationCreationTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0084},
		Name: "ContentCreatorName",
		VR:   "PN",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0086},
		Name: "ContentCreatorIdentificationCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0087},
		Name: "AlternateContentDescriptionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0100},
		Name: "PresentationSizeMode",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0101},
		Name: "PresentationPixelSpacing",
		VR:   "DS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0102},
		Name: "PresentationPixelAspectRatio",
		VR:   "IS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0103},
		Name: "PresentationPixelMagnificationRatio",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0207},
		Name: "GraphicGroupLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0208},
		Name: "GraphicGroupDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0209},
		Name: "CompoundGraphicSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0226},
		Name: "CompoundGraphicInstanceID",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0227},
		Name: "FontName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0228},
		Name: "FontNameType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0229},
		Name: "CSSFontName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0230},
		Name: "RotationAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0231},
		Name: "TextStyleSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0232},
		Name: "LineStyleSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0233},
		Name: "FillStyleSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0234},
		Name: "GraphicGroupSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0241},
		Name: "TextColorCIELabValue",
		VR:   "US",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0242},
		Name: "HorizontalAlignment",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0243},
		Name: "VerticalAlignment",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0244},
		Name: "ShadowStyle",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0245},
		Name: "ShadowOffsetX",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0246},
		Name: "ShadowOffsetY",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0247},
		Name: "ShadowColorCIELabValue",
		VR:   "US",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0248},
		Name: "Underlined",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0249},
		Name: "Bold",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0250},
		Name: "Italic",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0251},
		Name: "PatternOnColorCIELabValue",
		VR:   "US",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0252},
		Name: "PatternOffColorCIELabValue",
		VR:   "US",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0253},
		Name: "LineThickness",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0254},
		Name: "LineDashingStyle",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0255},
		Name: "LinePattern",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0256},
		Name: "FillPattern",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0257},
		Name: "FillMode",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0258},
		Name: "ShadowOpacity",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0261},
		Name: "GapLength",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0262},
		Name: "DiameterOfVisibility",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0273},
		Name: "RotationPoint",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0274},
		Name: "TickAlignment",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0278},
		Name: "ShowTickLabel",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0279},
		Name: "TickLabelAlignment",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0282},
		Name: "CompoundGraphicUnits",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0284},
		Name: "PatternOnOpacity",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0285},
		Name: "PatternOffOpacity",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0287},
		Name: "MajorTicksSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0288},
		Name: "TickPosition",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0289},
		Name: "TickLabel",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0294},
		Name: "CompoundGraphicType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0295},
		Name: "GraphicGroupID",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0306},
		Name: "ShapeType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0308},
		Name: "RegistrationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0309},
		Name: "MatrixRegistrationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x030a},
		Name: "MatrixSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x030b},
		Name: "FrameOfReferenceToDisplayedCoordinateSystemTransformationMatrix",
		VR:   "FD",
		VM:   "16",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x030c},
		Name: "FrameOfReferenceTransformationMatrixType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x030d},
		Name: "RegistrationTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x030f},
		Name: "FiducialDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0310},
		Name: "FiducialIdentifier",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0311},
		Name: "FiducialIdentifierCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0312},
		Name: "ContourUncertaintyRadius",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0314},
		Name: "UsedFiducialsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0318},
		Name: "GraphicCoordinatesDataSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x031a},
		Name: "FiducialUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x031b},
		Name: "ReferencedFiducialUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x031c},
		Name: "FiducialSetSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x031e},
		Name: "FiducialSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x031f},
		Name: "FiducialsPropertyCategoryCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0401},
		Name: "GraphicLayerRecommendedDisplayCIELabValue",
		VR:   "US",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0402},
		Name: "BlendingSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0403},
		Name: "RelativeOpacity",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0404},
		Name: "ReferencedSpatialRegistrationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x0405},
		Name: "BlendingPosition",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1101},
		Name: "PresentationDisplayCollectionUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1102},
		Name: "PresentationSequenceCollectionUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1103},
		Name: "PresentationSequencePositionIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1104},
		Name: "RenderedImageReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1201},
		Name: "VolumetricPresentationStateInputSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1202},
		Name: "PresentationInputType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1203},
		Name: "InputSequencePositionIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1204},
		Name: "Crop",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1205},
		Name: "CroppingSpecificationIndex",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1207},
		Name: "VolumetricPresentationInputNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1208},
		Name: "ImageVolumeGeometry",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1209},
		Name: "VolumetricPresentationInputSetUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x120a},
		Name: "VolumetricPresentationInputSetSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x120b},
		Name: "GlobalCrop",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x120c},
		Name: "GlobalCroppingSpecificationIndex",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x120d},
		Name: "RenderingMethod",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1301},
		Name: "VolumeCroppingSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1302},
		Name: "VolumeCroppingMethod",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1303},
		Name: "BoundingBoxCrop",
		VR:   "FD",
		VM:   "6",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1304},
		Name: "ObliqueCroppingPlaneSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1305},
		Name: "Plane",
		VR:   "FD",
		VM:   "4",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1306},
		Name: "PlaneNormal",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1309},
		Name: "CroppingSpecificationNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1501},
		Name: "MultiPlanarReconstructionStyle",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1502},
		Name: "MPRThicknessType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1503},
		Name: "MPRSlabThickness",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1505},
		Name: "MPRTopLeftHandCorner",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1507},
		Name: "MPRViewWidthDirection",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1508},
		Name: "MPRViewWidth",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x150c},
		Name: "NumberOfVolumetricCurvePoints",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x150d},
		Name: "VolumetricCurvePoints",
		VR:   "OD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1511},
		Name: "MPRViewHeightDirection",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1512},
		Name: "MPRViewHeight",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1602},
		Name: "RenderProjection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1603},
		Name: "ViewpointPosition",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1604},
		Name: "ViewpointLookAtPoint",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1605},
		Name: "ViewpointUpDirection",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1606},
		Name: "RenderFieldOfView",
		VR:   "FD",
		VM:   "6",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1607},
		Name: "SamplingStepSize",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1701},
		Name: "ShadingStyle",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1702},
		Name: "AmbientReflectionIntensity",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1703},
		Name: "LightDirection",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1704},
		Name: "DiffuseReflectionIntensity",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1705},
		Name: "SpecularReflectionIntensity",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1706},
		Name: "Shininess",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1801},
		Name: "PresentationStateClassificationComponentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1802},
		Name: "ComponentType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1803},
		Name: "ComponentInputSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1804},
		Name: "VolumetricPresentationInputIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1805},
		Name: "PresentationStateCompositorComponentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1806},
		Name: "WeightingTransferFunctionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1807},
		Name: "WeightingLookupTableDescriptor",
		VR:   "US",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1808},
		Name: "WeightingLookupTableData",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1901},
		Name: "VolumetricAnnotationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1903},
		Name: "ReferencedStructuredContextSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1904},
		Name: "ReferencedContentItem",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1905},
		Name: "VolumetricPresentationInputAnnotationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1907},
		Name: "AnnotationClipping",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1a01},
		Name: "PresentationAnimationStyle",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1a03},
		Name: "RecommendedAnimationRate",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1a04},
		Name: "AnimationCurveSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1a05},
		Name: "AnimationStepSize",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1a06},
		Name: "SwivelRange",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1a07},
		Name: "VolumetricCurveUpDirections",
		VR:   "OD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1a08},
		Name: "VolumeStreamSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1a09},
		Name: "RGBATransferFunctionDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1b01},
		Name: "AdvancedBlendingSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1b02},
		Name: "BlendingInputNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1b03},
		Name: "BlendingDisplayInputSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1b04},
		Name: "BlendingDisplaySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1b06},
		Name: "BlendingMode",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1b07},
		Name: "TimeSeriesBlending",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1b08},
		Name: "GeometryForDisplay",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1b11},
		Name: "ThresholdSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1b12},
		Name: "ThresholdValueSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1b13},
		Name: "ThresholdType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0070, 0x1b14},
		Name: "ThresholdValue",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0002},
		Name: "HangingProtocolName",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0004},
		Name: "HangingProtocolDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0006},
		Name: "HangingProtocolLevel",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0008},
		Name: "HangingProtocolCreator",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x000a},
		Name: "HangingProtocolCreationDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x000c},
		Name: "HangingProtocolDefinitionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x000e},
		Name: "HangingProtocolUserIdentificationCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0010},
		Name: "HangingProtocolUserGroupName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0012},
		Name: "SourceHangingProtocolSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0014},
		Name: "NumberOfPriorsReferenced",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0020},
		Name: "ImageSetsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0022},
		Name: "ImageSetSelectorSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0024},
		Name: "ImageSetSelectorUsageFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0026},
		Name: "SelectorAttribute",
		VR:   "AT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0028},
		Name: "SelectorValueNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0030},
		Name: "TimeBasedImageSetsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0032},
		Name: "ImageSetNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0034},
		Name: "ImageSetSelectorCategory",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0038},
		Name: "RelativeTime",
		VR:   "US",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x003a},
		Name: "RelativeTimeUnits",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x003c},
		Name: "AbstractPriorValue",
		VR:   "SS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x003e},
		Name: "AbstractPriorCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0040},
		Name: "ImageSetLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0050},
		Name: "SelectorAttributeVR",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0052},
		Name: "SelectorSequencePointer",
		VR:   "AT",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0054},
		Name: "SelectorSequencePointerPrivateCreator",
		VR:   "LO",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0056},
		Name: "SelectorAttributePrivateCreator",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x005e},
		Name: "SelectorAEValue",
		VR:   "AE",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x005f},
		Name: "SelectorASValue",
		VR:   "AS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0060},
		Name: "SelectorATValue",
		VR:   "AT",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0061},
		Name: "SelectorDAValue",
		VR:   "DA",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0062},
		Name: "SelectorCSValue",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0063},
		Name: "SelectorDTValue",
		VR:   "DT",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0064},
		Name: "SelectorISValue",
		VR:   "IS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0065},
		Name: "SelectorOBValue",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0066},
		Name: "SelectorLOValue",
		VR:   "LO",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0067},
		Name: "SelectorOFValue",
		VR:   "OF",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0068},
		Name: "SelectorLTValue",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0069},
		Name: "SelectorOWValue",
		VR:   "OW",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x006a},
		Name: "SelectorPNValue",
		VR:   "PN",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x006b},
		Name: "SelectorTMValue",
		VR:   "TM",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x006c},
		Name: "SelectorSHValue",
		VR:   "SH",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x006d},
		Name: "SelectorUNValue",
		VR:   "UN",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x006e},
		Name: "SelectorSTValue",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x006f},
		Name: "SelectorUCValue",
		VR:   "UC",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0070},
		Name: "SelectorUTValue",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0071},
		Name: "SelectorURValue",
		VR:   "UR",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0072},
		Name: "SelectorDSValue",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0073},
		Name: "SelectorODValue",
		VR:   "OD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0074},
		Name: "SelectorFDValue",
		VR:   "FD",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0075},
		Name: "SelectorOLValue",
		VR:   "OL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0076},
		Name: "SelectorFLValue",
		VR:   "FL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0078},
		Name: "SelectorULValue",
		VR:   "UL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x007a},
		Name: "SelectorUSValue",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x007c},
		Name: "SelectorSLValue",
		VR:   "SL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x007e},
		Name: "SelectorSSValue",
		VR:   "SS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x007f},
		Name: "SelectorUIValue",
		VR:   "UI",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0080},
		Name: "SelectorCodeSequenceValue",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0100},
		Name: "NumberOfScreens",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0102},
		Name: "NominalScreenDefinitionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0104},
		Name: "NumberOfVerticalPixels",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0106},
		Name: "NumberOfHorizontalPixels",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0108},
		Name: "DisplayEnvironmentSpatialPosition",
		VR:   "FD",
		VM:   "4",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x010a},
		Name: "ScreenMinimumGrayscaleBitDepth",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x010c},
		Name: "ScreenMinimumColorBitDepth",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x010e},
		Name: "ApplicationMaximumRepaintTime",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0200},
		Name: "DisplaySetsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0202},
		Name: "DisplaySetNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0203},
		Name: "DisplaySetLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0204},
		Name: "DisplaySetPresentationGroup",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0206},
		Name: "DisplaySetPresentationGroupDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0208},
		Name: "PartialDataDisplayHandling",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0210},
		Name: "SynchronizedScrollingSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0212},
		Name: "DisplaySetScrollingGroup",
		VR:   "US",
		VM:   "2-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0214},
		Name: "NavigationIndicatorSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0216},
		Name: "NavigationDisplaySet",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0218},
		Name: "ReferenceDisplaySets",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0300},
		Name: "ImageBoxesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0302},
		Name: "ImageBoxNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0304},
		Name: "ImageBoxLayoutType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0306},
		Name: "ImageBoxTileHorizontalDimension",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0308},
		Name: "ImageBoxTileVerticalDimension",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0310},
		Name: "ImageBoxScrollDirection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0312},
		Name: "ImageBoxSmallScrollType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0314},
		Name: "ImageBoxSmallScrollAmount",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0316},
		Name: "ImageBoxLargeScrollType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0318},
		Name: "ImageBoxLargeScrollAmount",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0320},
		Name: "ImageBoxOverlapPriority",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0330},
		Name: "CineRelativeToRealTime",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0400},
		Name: "FilterOperationsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0402},
		Name: "FilterByCategory",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0404},
		Name: "FilterByAttributePresence",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0406},
		Name: "FilterByOperator",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0420},
		Name: "StructuredDisplayBackgroundCIELabValue",
		VR:   "US",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0421},
		Name: "EmptyImageBoxCIELabValue",
		VR:   "US",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0422},
		Name: "StructuredDisplayImageBoxSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0424},
		Name: "StructuredDisplayTextBoxSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0427},
		Name: "ReferencedFirstFrameSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0430},
		Name: "ImageBoxSynchronizationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0432},
		Name: "SynchronizedImageBoxList",
		VR:   "US",
		VM:   "2-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0434},
		Name: "TypeOfSynchronization",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0500},
		Name: "BlendingOperationType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0510},
		Name: "ReformattingOperationType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0512},
		Name: "ReformattingThickness",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0514},
		Name: "ReformattingInterval",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0516},
		Name: "ReformattingOperationInitialViewDirection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0520},
		Name: "ThreeDRenderingType",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0600},
		Name: "SortingOperationsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0602},
		Name: "SortByCategory",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0604},
		Name: "SortingDirection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0700},
		Name: "DisplaySetPatientOrientation",
		VR:   "CS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0702},
		Name: "VOIType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0704},
		Name: "PseudoColorType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0705},
		Name: "PseudoColorPaletteInstanceReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0706},
		Name: "ShowGrayscaleInverted",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0710},
		Name: "ShowImageTrueSizeFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0712},
		Name: "ShowGraphicAnnotationFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0714},
		Name: "ShowPatientDemographicsFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0716},
		Name: "ShowAcquisitionTechniquesFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0717},
		Name: "DisplaySetHorizontalJustification",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0072, 0x0718},
		Name: "DisplaySetVerticalJustification",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x0120},
		Name: "ContinuationStartMeterset",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x0121},
		Name: "ContinuationEndMeterset",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1000},
		Name: "ProcedureStepState",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1002},
		Name: "ProcedureStepProgressInformationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1004},
		Name: "ProcedureStepProgress",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1006},
		Name: "ProcedureStepProgressDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1007},
		Name: "ProcedureStepProgressParametersSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1008},
		Name: "ProcedureStepCommunicationsURISequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x100a},
		Name: "ContactURI",
		VR:   "UR",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x100c},
		Name: "ContactDisplayName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x100e},
		Name: "ProcedureStepDiscontinuationReasonCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1020},
		Name: "BeamTaskSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1022},
		Name: "BeamTaskType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1025},
		Name: "AutosequenceFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1026},
		Name: "TableTopVerticalAdjustedPosition",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1027},
		Name: "TableTopLongitudinalAdjustedPosition",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1028},
		Name: "TableTopLateralAdjustedPosition",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x102a},
		Name: "PatientSupportAdjustedAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x102b},
		Name: "TableTopEccentricAdjustedAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x102c},
		Name: "TableTopPitchAdjustedAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x102d},
		Name: "TableTopRollAdjustedAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1030},
		Name: "DeliveryVerificationImageSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1032},
		Name: "VerificationImageTiming",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1034},
		Name: "DoubleExposureFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1036},
		Name: "DoubleExposureOrdering",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1040},
		Name: "RelatedReferenceRTImageSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1042},
		Name: "GeneralMachineVerificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1044},
		Name: "ConventionalMachineVerificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1046},
		Name: "IonMachineVerificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1048},
		Name: "FailedAttributesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x104a},
		Name: "OverriddenAttributesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x104c},
		Name: "ConventionalControlPointVerificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x104e},
		Name: "IonControlPointVerificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1050},
		Name: "AttributeOccurrenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1052},
		Name: "AttributeOccurrencePointer",
		VR:   "AT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1054},
		Name: "AttributeItemSelector",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1056},
		Name: "AttributeOccurrencePrivateCreator",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1057},
		Name: "SelectorSequencePointerItems",
		VR:   "IS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1200},
		Name: "ScheduledProcedureStepPriority",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1202},
		Name: "WorklistLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1204},
		Name: "ProcedureStepLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1210},
		Name: "ScheduledProcessingParametersSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1212},
		Name: "PerformedProcessingParametersSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1216},
		Name: "UnifiedProcedureStepPerformedProcedureSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1224},
		Name: "ReplacedProcedureStepSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1230},
		Name: "DeletionLock",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1234},
		Name: "ReceivingAE",
		VR:   "AE",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1236},
		Name: "RequestingAE",
		VR:   "AE",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1238},
		Name: "ReasonForCancellation",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1242},
		Name: "SCPStatus",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1244},
		Name: "SubscriptionListStatus",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1246},
		Name: "UnifiedProcedureStepListStatus",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1324},
		Name: "BeamOrderIndex",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1338},
		Name: "DoubleExposureMeterset",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x133a},
		Name: "DoubleExposureFieldDelta",
		VR:   "FD",
		VM:   "4",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1401},
		Name: "BrachyTaskSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1402},
		Name: "ContinuationStartTotalReferenceAirKerma",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1403},
		Name: "ContinuationEndTotalReferenceAirKerma",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1404},
		Name: "ContinuationPulseNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1405},
		Name: "ChannelDeliveryOrderSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1406},
		Name: "ReferencedChannelNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1407},
		Name: "StartCumulativeTimeWeight",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1408},
		Name: "EndCumulativeTimeWeight",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x1409},
		Name: "OmittedChannelSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x140a},
		Name: "ReasonForChannelOmission",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x140b},
		Name: "ReasonForChannelOmissionDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x140c},
		Name: "ChannelDeliveryOrderIndex",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x140d},
		Name: "ChannelDeliveryContinuationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0074, 0x140e},
		Name: "OmittedApplicationSetupSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x0001},
		Name: "ImplantAssemblyTemplateName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x0003},
		Name: "ImplantAssemblyTemplateIssuer",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x0006},
		Name: "ImplantAssemblyTemplateVersion",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x0008},
		Name: "ReplacedImplantAssemblyTemplateSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x000a},
		Name: "ImplantAssemblyTemplateType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x000c},
		Name: "OriginalImplantAssemblyTemplateSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x000e},
		Name: "DerivationImplantAssemblyTemplateSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x0010},
		Name: "ImplantAssemblyTemplateTargetAnatomySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x0020},
		Name: "ProcedureTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x0030},
		Name: "SurgicalTechnique",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x0032},
		Name: "ComponentTypesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x0034},
		Name: "ComponentTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x0036},
		Name: "ExclusiveComponentType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x0038},
		Name: "MandatoryComponentType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x0040},
		Name: "ComponentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x0055},
		Name: "ComponentID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x0060},
		Name: "ComponentAssemblySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x0070},
		Name: "Component1ReferencedID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x0080},
		Name: "Component1ReferencedMatingFeatureSetID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x0090},
		Name: "Component1ReferencedMatingFeatureID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x00a0},
		Name: "Component2ReferencedID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x00b0},
		Name: "Component2ReferencedMatingFeatureSetID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0076, 0x00c0},
		Name: "Component2ReferencedMatingFeatureID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0078, 0x0001},
		Name: "ImplantTemplateGroupName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0078, 0x0010},
		Name: "ImplantTemplateGroupDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0078, 0x0020},
		Name: "ImplantTemplateGroupIssuer",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0078, 0x0024},
		Name: "ImplantTemplateGroupVersion",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0078, 0x0026},
		Name: "ReplacedImplantTemplateGroupSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0078, 0x0028},
		Name: "ImplantTemplateGroupTargetAnatomySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0078, 0x002a},
		Name: "ImplantTemplateGroupMembersSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0078, 0x002e},
		Name: "ImplantTemplateGroupMemberID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0078, 0x0050},
		Name: "ThreeDImplantTemplateGroupMemberMatchingPoint",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0078, 0x0060},
		Name: "ThreeDImplantTemplateGroupMemberMatchingAxes",
		VR:   "FD",
		VM:   "9",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0078, 0x0070},
		Name: "ImplantTemplateGroupMemberMatching2DCoordinatesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0078, 0x0090},
		Name: "TwoDImplantTemplateGroupMemberMatchingPoint",
		VR:   "FD",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0078, 0x00a0},
		Name: "TwoDImplantTemplateGroupMemberMatchingAxes",
		VR:   "FD",
		VM:   "4",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0078, 0x00b0},
		Name: "ImplantTemplateGroupVariationDimensionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0078, 0x00b2},
		Name: "ImplantTemplateGroupVariationDimensionName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0078, 0x00b4},
		Name: "ImplantTemplateGroupVariationDimensionRankSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0078, 0x00b6},
		Name: "ReferencedImplantTemplateGroupMemberID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0078, 0x00b8},
		Name: "ImplantTemplateGroupVariationDimensionRank",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0080, 0x0001},
		Name: "SurfaceScanAcquisitionTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0080, 0x0002},
		Name: "SurfaceScanModeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0080, 0x0003},
		Name: "RegistrationMethodCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0080, 0x0004},
		Name: "ShotDurationTime",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0080, 0x0005},
		Name: "ShotOffsetTime",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0080, 0x0006},
		Name: "SurfacePointPresentationValueData",
		VR:   "US",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0080, 0x0007},
		Name: "SurfacePointColorCIELabValueData",
		VR:   "US",
		VM:   "3-3n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0080, 0x0008},
		Name: "UVMappingSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0080, 0x0009},
		Name: "TextureLabel",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0080, 0x0010},
		Name: "UValueData",
		VR:   "OF",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0080, 0x0011},
		Name: "VValueData",
		VR:   "OF",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0080, 0x0012},
		Name: "ReferencedTextureSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0080, 0x0013},
		Name: "ReferencedSurfaceDataSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0001},
		Name: "AssessmentSummary",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0003},
		Name: "AssessmentSummaryDescription",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0004},
		Name: "AssessedSOPInstanceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0005},
		Name: "ReferencedComparisonSOPInstanceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0006},
		Name: "NumberOfAssessmentObservations",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0007},
		Name: "AssessmentObservationsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0008},
		Name: "ObservationSignificance",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x000a},
		Name: "ObservationDescription",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x000c},
		Name: "StructuredConstraintObservationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0010},
		Name: "AssessedAttributeValueSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0016},
		Name: "AssessmentSetID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0017},
		Name: "AssessmentRequesterSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0018},
		Name: "SelectorAttributeName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0019},
		Name: "SelectorAttributeKeyword",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0021},
		Name: "AssessmentTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0022},
		Name: "ObservationBasisCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0023},
		Name: "AssessmentLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0032},
		Name: "ConstraintType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0033},
		Name: "SpecificationSelectionGuidance",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0034},
		Name: "ConstraintValueSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0035},
		Name: "RecommendedDefaultValueSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0036},
		Name: "ConstraintViolationSignificance",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0037},
		Name: "ConstraintViolationCondition",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0082, 0x0038},
		Name: "ModifiableConstraintFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0088, 0x0130},
		Name: "StorageMediaFileSetID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0088, 0x0140},
		Name: "StorageMediaFileSetUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0088, 0x0200},
		Name: "IconImageSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0100, 0x0410},
		Name: "SOPInstanceStatus",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0100, 0x0420},
		Name: "SOPAuthorizationDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0100, 0x0424},
		Name: "SOPAuthorizationComment",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0100, 0x0426},
		Name: "AuthorizationEquipmentCertificationNumber",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0005},
		Name: "MACIDNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0010},
		Name: "MACCalculationTransferSyntaxUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0015},
		Name: "MACAlgorithm",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0020},
		Name: "DataElementsSigned",
		VR:   "AT",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0100},
		Name: "DigitalSignatureUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0105},
		Name: "DigitalSignatureDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0110},
		Name: "CertificateType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0115},
		Name: "CertificateOfSigner",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0120},
		Name: "Signature",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0305},
		Name: "CertifiedTimestampType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0310},
		Name: "CertifiedTimestamp",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0401},
		Name: "DigitalSignaturePurposeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0402},
		Name: "ReferencedDigitalSignatureSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0403},
		Name: "ReferencedSOPInstanceMACSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0404},
		Name: "MAC",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0500},
		Name: "EncryptedAttributesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0510},
		Name: "EncryptedContentTransferSyntaxUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0520},
		Name: "EncryptedContent",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0550},
		Name: "ModifiedAttributesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0551},
		Name: "NonconformingModifiedAttributesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0552},
		Name: "NonconformingDataElementValue",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0561},
		Name: "OriginalAttributesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0562},
		Name: "AttributeModificationDateTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0563},
		Name: "ModifyingSystem",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0564},
		Name: "SourceOfPreviousValues",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0565},
		Name: "ReasonForTheAttributeModification",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0400, 0x0600},
		Name: "InstanceOriginStatus",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2000, 0x0010},
		Name: "NumberOfCopies",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2000, 0x001e},
		Name: "PrinterConfigurationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2000, 0x0020},
		Name: "PrintPriority",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2000, 0x0030},
		Name: "MediumType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2000, 0x0040},
		Name: "FilmDestination",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2000, 0x0050},
		Name: "FilmSessionLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2000, 0x0060},
		Name: "MemoryAllocation",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2000, 0x0061},
		Name: "MaximumMemoryAllocation",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2000, 0x00a0},
		Name: "MemoryBitDepth",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2000, 0x00a1},
		Name: "PrintingBitDepth",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2000, 0x00a2},
		Name: "MediaInstalledSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2000, 0x00a4},
		Name: "OtherMediaAvailableSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2000, 0x00a8},
		Name: "SupportedImageDisplayFormatsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2000, 0x0500},
		Name: "ReferencedFilmBoxSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x0010},
		Name: "ImageDisplayFormat",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x0030},
		Name: "AnnotationDisplayFormatID",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x0040},
		Name: "FilmOrientation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x0050},
		Name: "FilmSizeID",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x0052},
		Name: "PrinterResolutionID",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x0054},
		Name: "DefaultPrinterResolutionID",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x0060},
		Name: "MagnificationType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x0080},
		Name: "SmoothingType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x00a6},
		Name: "DefaultMagnificationType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x00a7},
		Name: "OtherMagnificationTypesAvailable",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x00a8},
		Name: "DefaultSmoothingType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x00a9},
		Name: "OtherSmoothingTypesAvailable",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x0100},
		Name: "BorderDensity",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x0110},
		Name: "EmptyImageDensity",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x0120},
		Name: "MinDensity",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x0130},
		Name: "MaxDensity",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x0140},
		Name: "Trim",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x0150},
		Name: "ConfigurationInformation",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x0152},
		Name: "ConfigurationInformationDescription",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x0154},
		Name: "MaximumCollatedFilms",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x015e},
		Name: "Illumination",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x0160},
		Name: "ReflectedAmbientLight",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x0376},
		Name: "PrinterPixelSpacing",
		VR:   "DS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x0500},
		Name: "ReferencedFilmSessionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x0510},
		Name: "ReferencedImageBoxSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2010, 0x0520},
		Name: "ReferencedBasicAnnotationBoxSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2020, 0x0010},
		Name: "ImageBoxPosition",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2020, 0x0020},
		Name: "Polarity",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2020, 0x0030},
		Name: "RequestedImageSize",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2020, 0x0040},
		Name: "RequestedDecimateCropBehavior",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2020, 0x0050},
		Name: "RequestedResolutionID",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2020, 0x00a0},
		Name: "RequestedImageSizeFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2020, 0x00a2},
		Name: "DecimateCropResult",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2020, 0x0110},
		Name: "BasicGrayscaleImageSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2020, 0x0111},
		Name: "BasicColorImageSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2030, 0x0010},
		Name: "AnnotationPosition",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2030, 0x0020},
		Name: "TextString",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2050, 0x0010},
		Name: "PresentationLUTSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2050, 0x0020},
		Name: "PresentationLUTShape",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2050, 0x0500},
		Name: "ReferencedPresentationLUTSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2100, 0x0020},
		Name: "ExecutionStatus",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2100, 0x0030},
		Name: "ExecutionStatusInfo",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2100, 0x0040},
		Name: "CreationDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2100, 0x0050},
		Name: "CreationTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2100, 0x0070},
		Name: "Originator",
		VR:   "AE",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2100, 0x0140},
		Name: "DestinationAE",
		VR:   "AE",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2100, 0x0160},
		Name: "OwnerID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2100, 0x0170},
		Name: "NumberOfFilms",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2110, 0x0010},
		Name: "PrinterStatus",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2110, 0x0020},
		Name: "PrinterStatusInfo",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2110, 0x0030},
		Name: "PrinterName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2200, 0x0001},
		Name: "LabelUsingInformationExtractedFromInstances",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2200, 0x0002},
		Name: "LabelText",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2200, 0x0003},
		Name: "LabelStyleSelection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2200, 0x0004},
		Name: "MediaDisposition",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2200, 0x0005},
		Name: "BarcodeValue",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2200, 0x0006},
		Name: "BarcodeSymbology",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2200, 0x0007},
		Name: "AllowMediaSplitting",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2200, 0x0008},
		Name: "IncludeNonDICOMObjects",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2200, 0x0009},
		Name: "IncludeDisplayApplication",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2200, 0x000a},
		Name: "PreserveCompositeInstancesAfterMediaCreation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2200, 0x000b},
		Name: "TotalNumberOfPiecesOfMediaCreated",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2200, 0x000c},
		Name: "RequestedMediaApplicationProfile",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2200, 0x000d},
		Name: "ReferencedStorageMediaSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2200, 0x000e},
		Name: "FailureAttributes",
		VR:   "AT",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2200, 0x000f},
		Name: "AllowLossyCompression",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x2200, 0x0020},
		Name: "RequestPriority",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x0002},
		Name: "RTImageLabel",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x0003},
		Name: "RTImageName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x0004},
		Name: "RTImageDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x000a},
		Name: "ReportedValuesOrigin",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x000c},
		Name: "RTImagePlane",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x000d},
		Name: "XRayImageReceptorTranslation",
		VR:   "DS",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x000e},
		Name: "XRayImageReceptorAngle",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x0010},
		Name: "RTImageOrientation",
		VR:   "DS",
		VM:   "6",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x0011},
		Name: "ImagePlanePixelSpacing",
		VR:   "DS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x0012},
		Name: "RTImagePosition",
		VR:   "DS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x0020},
		Name: "RadiationMachineName",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x0022},
		Name: "RadiationMachineSAD",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x0024},
		Name: "RadiationMachineSSD",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x0026},
		Name: "RTImageSID",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x0028},
		Name: "SourceToReferenceObjectDistance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x0029},
		Name: "FractionNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x0030},
		Name: "ExposureSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x0032},
		Name: "MetersetExposure",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x0034},
		Name: "DiaphragmPosition",
		VR:   "DS",
		VM:   "4",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x0040},
		Name: "FluenceMapSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x0041},
		Name: "FluenceDataSource",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x0042},
		Name: "FluenceDataScale",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x0050},
		Name: "PrimaryFluenceModeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x0051},
		Name: "FluenceMode",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3002, 0x0052},
		Name: "FluenceModeID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x0001},
		Name: "DVHType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x0002},
		Name: "DoseUnits",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x0004},
		Name: "DoseType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x0005},
		Name: "SpatialTransformOfDose",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x0006},
		Name: "DoseComment",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x0008},
		Name: "NormalizationPoint",
		VR:   "DS",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x000a},
		Name: "DoseSummationType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x000c},
		Name: "GridFrameOffsetVector",
		VR:   "DS",
		VM:   "2-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x000e},
		Name: "DoseGridScaling",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x0010},
		Name: "RTDoseROISequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x0012},
		Name: "DoseValue",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x0014},
		Name: "TissueHeterogeneityCorrection",
		VR:   "CS",
		VM:   "1-3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x0040},
		Name: "DVHNormalizationPoint",
		VR:   "DS",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x0042},
		Name: "DVHNormalizationDoseValue",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x0050},
		Name: "DVHSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x0052},
		Name: "DVHDoseScaling",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x0054},
		Name: "DVHVolumeUnits",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x0056},
		Name: "DVHNumberOfBins",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x0058},
		Name: "DVHData",
		VR:   "DS",
		VM:   "2-2n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x0060},
		Name: "DVHReferencedROISequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x0062},
		Name: "DVHROIContributionType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x0070},
		Name: "DVHMinimumDose",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x0072},
		Name: "DVHMaximumDose",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3004, 0x0074},
		Name: "DVHMeanDose",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0002},
		Name: "StructureSetLabel",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0004},
		Name: "StructureSetName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0006},
		Name: "StructureSetDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0008},
		Name: "StructureSetDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0009},
		Name: "StructureSetTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0010},
		Name: "ReferencedFrameOfReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0012},
		Name: "RTReferencedStudySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0014},
		Name: "RTReferencedSeriesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0016},
		Name: "ContourImageSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0018},
		Name: "PredecessorStructureSetSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0020},
		Name: "StructureSetROISequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0022},
		Name: "ROINumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0024},
		Name: "ReferencedFrameOfReferenceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0026},
		Name: "ROIName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0028},
		Name: "ROIDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x002a},
		Name: "ROIDisplayColor",
		VR:   "IS",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x002c},
		Name: "ROIVolume",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0030},
		Name: "RTRelatedROISequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0033},
		Name: "RTROIRelationship",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0036},
		Name: "ROIGenerationAlgorithm",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0037},
		Name: "ROIDerivationAlgorithmIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0038},
		Name: "ROIGenerationDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0039},
		Name: "ROIContourSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0040},
		Name: "ContourSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0042},
		Name: "ContourGeometricType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0044},
		Name: "ContourSlabThickness",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0045},
		Name: "ContourOffsetVector",
		VR:   "DS",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0046},
		Name: "NumberOfContourPoints",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0048},
		Name: "ContourNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0049},
		Name: "AttachedContours",
		VR:   "IS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0050},
		Name: "ContourData",
		VR:   "DS",
		VM:   "3-3n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0080},
		Name: "RTROIObservationsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0082},
		Name: "ObservationNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0084},
		Name: "ReferencedROINumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0085},
		Name: "ROIObservationLabel",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0086},
		Name: "RTROIIdentificationCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x0088},
		Name: "ROIObservationDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x00a0},
		Name: "RelatedRTROIObservationsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x00a4},
		Name: "RTROIInterpretedType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x00a6},
		Name: "ROIInterpreter",
		VR:   "PN",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x00b0},
		Name: "ROIPhysicalPropertiesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x00b2},
		Name: "ROIPhysicalProperty",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x00b4},
		Name: "ROIPhysicalPropertyValue",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x00b6},
		Name: "ROIElementalCompositionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x00b7},
		Name: "ROIElementalCompositionAtomicNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x00b8},
		Name: "ROIElementalCompositionAtomicMassFraction",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x00c6},
		Name: "FrameOfReferenceTransformationMatrix",
		VR:   "DS",
		VM:   "16",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x00c8},
		Name: "FrameOfReferenceTransformationComment",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x00c9},
		Name: "PatientLocationCoordinatesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x00ca},
		Name: "PatientLocationCoordinatesCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3006, 0x00cb},
		Name: "PatientSupportPositionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0010},
		Name: "MeasuredDoseReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0012},
		Name: "MeasuredDoseDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0014},
		Name: "MeasuredDoseType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0016},
		Name: "MeasuredDoseValue",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0020},
		Name: "TreatmentSessionBeamSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0021},
		Name: "TreatmentSessionIonBeamSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0022},
		Name: "CurrentFractionNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0024},
		Name: "TreatmentControlPointDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0025},
		Name: "TreatmentControlPointTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x002a},
		Name: "TreatmentTerminationStatus",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x002b},
		Name: "TreatmentTerminationCode",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x002c},
		Name: "TreatmentVerificationStatus",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0030},
		Name: "ReferencedTreatmentRecordSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0032},
		Name: "SpecifiedPrimaryMeterset",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0033},
		Name: "SpecifiedSecondaryMeterset",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0036},
		Name: "DeliveredPrimaryMeterset",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0037},
		Name: "DeliveredSecondaryMeterset",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x003a},
		Name: "SpecifiedTreatmentTime",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x003b},
		Name: "DeliveredTreatmentTime",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0040},
		Name: "ControlPointDeliverySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0041},
		Name: "IonControlPointDeliverySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0042},
		Name: "SpecifiedMeterset",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0044},
		Name: "DeliveredMeterset",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0045},
		Name: "MetersetRateSet",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0046},
		Name: "MetersetRateDelivered",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0047},
		Name: "ScanSpotMetersetsDelivered",
		VR:   "FL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0048},
		Name: "DoseRateDelivered",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0050},
		Name: "TreatmentSummaryCalculatedDoseReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0052},
		Name: "CumulativeDoseToDoseReference",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0054},
		Name: "FirstTreatmentDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0056},
		Name: "MostRecentTreatmentDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x005a},
		Name: "NumberOfFractionsDelivered",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0060},
		Name: "OverrideSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0061},
		Name: "ParameterSequencePointer",
		VR:   "AT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0062},
		Name: "OverrideParameterPointer",
		VR:   "AT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0063},
		Name: "ParameterItemIndex",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0064},
		Name: "MeasuredDoseReferenceNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0065},
		Name: "ParameterPointer",
		VR:   "AT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0066},
		Name: "OverrideReason",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0067},
		Name: "ParameterValueNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0068},
		Name: "CorrectedParameterSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x006a},
		Name: "CorrectionValue",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0070},
		Name: "CalculatedDoseReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0072},
		Name: "CalculatedDoseReferenceNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0074},
		Name: "CalculatedDoseReferenceDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0076},
		Name: "CalculatedDoseReferenceDoseValue",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0078},
		Name: "StartMeterset",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x007a},
		Name: "EndMeterset",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0080},
		Name: "ReferencedMeasuredDoseReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0082},
		Name: "ReferencedMeasuredDoseReferenceNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0090},
		Name: "ReferencedCalculatedDoseReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0092},
		Name: "ReferencedCalculatedDoseReferenceNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x00a0},
		Name: "BeamLimitingDeviceLeafPairsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x00b0},
		Name: "RecordedWedgeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x00c0},
		Name: "RecordedCompensatorSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x00d0},
		Name: "RecordedBlockSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x00e0},
		Name: "TreatmentSummaryMeasuredDoseReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x00f0},
		Name: "RecordedSnoutSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x00f2},
		Name: "RecordedRangeShifterSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x00f4},
		Name: "RecordedLateralSpreadingDeviceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x00f6},
		Name: "RecordedRangeModulatorSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0100},
		Name: "RecordedSourceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0105},
		Name: "SourceSerialNumber",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0110},
		Name: "TreatmentSessionApplicationSetupSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0116},
		Name: "ApplicationSetupCheck",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0120},
		Name: "RecordedBrachyAccessoryDeviceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0122},
		Name: "ReferencedBrachyAccessoryDeviceNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0130},
		Name: "RecordedChannelSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0132},
		Name: "SpecifiedChannelTotalTime",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0134},
		Name: "DeliveredChannelTotalTime",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0136},
		Name: "SpecifiedNumberOfPulses",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0138},
		Name: "DeliveredNumberOfPulses",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x013a},
		Name: "SpecifiedPulseRepetitionInterval",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x013c},
		Name: "DeliveredPulseRepetitionInterval",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0140},
		Name: "RecordedSourceApplicatorSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0142},
		Name: "ReferencedSourceApplicatorNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0150},
		Name: "RecordedChannelShieldSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0152},
		Name: "ReferencedChannelShieldNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0160},
		Name: "BrachyControlPointDeliveredSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0162},
		Name: "SafePositionExitDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0164},
		Name: "SafePositionExitTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0166},
		Name: "SafePositionReturnDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0168},
		Name: "SafePositionReturnTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0171},
		Name: "PulseSpecificBrachyControlPointDeliveredSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0172},
		Name: "PulseNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0173},
		Name: "BrachyPulseControlPointDeliveredSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0200},
		Name: "CurrentTreatmentStatus",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0202},
		Name: "TreatmentStatusComment",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0220},
		Name: "FractionGroupSummarySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0223},
		Name: "ReferencedFractionNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0224},
		Name: "FractionGroupType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0230},
		Name: "BeamStopperPosition",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0240},
		Name: "FractionStatusSummarySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0250},
		Name: "TreatmentDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3008, 0x0251},
		Name: "TreatmentTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0002},
		Name: "RTPlanLabel",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0003},
		Name: "RTPlanName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0004},
		Name: "RTPlanDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0006},
		Name: "RTPlanDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0007},
		Name: "RTPlanTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0009},
		Name: "TreatmentProtocols",
		VR:   "LO",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x000a},
		Name: "PlanIntent",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x000b},
		Name: "TreatmentSites",
		VR:   "LO",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x000c},
		Name: "RTPlanGeometry",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x000e},
		Name: "PrescriptionDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0010},
		Name: "DoseReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0012},
		Name: "DoseReferenceNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0013},
		Name: "DoseReferenceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0014},
		Name: "DoseReferenceStructureType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0015},
		Name: "NominalBeamEnergyUnit",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0016},
		Name: "DoseReferenceDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0018},
		Name: "DoseReferencePointCoordinates",
		VR:   "DS",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x001a},
		Name: "NominalPriorDose",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0020},
		Name: "DoseReferenceType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0021},
		Name: "ConstraintWeight",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0022},
		Name: "DeliveryWarningDose",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0023},
		Name: "DeliveryMaximumDose",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0025},
		Name: "TargetMinimumDose",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0026},
		Name: "TargetPrescriptionDose",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0027},
		Name: "TargetMaximumDose",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0028},
		Name: "TargetUnderdoseVolumeFraction",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x002a},
		Name: "OrganAtRiskFullVolumeDose",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x002b},
		Name: "OrganAtRiskLimitDose",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x002c},
		Name: "OrganAtRiskMaximumDose",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x002d},
		Name: "OrganAtRiskOverdoseVolumeFraction",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0040},
		Name: "ToleranceTableSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0042},
		Name: "ToleranceTableNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0043},
		Name: "ToleranceTableLabel",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0044},
		Name: "GantryAngleTolerance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0046},
		Name: "BeamLimitingDeviceAngleTolerance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0048},
		Name: "BeamLimitingDeviceToleranceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x004a},
		Name: "BeamLimitingDevicePositionTolerance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x004b},
		Name: "SnoutPositionTolerance",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x004c},
		Name: "PatientSupportAngleTolerance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x004e},
		Name: "TableTopEccentricAngleTolerance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x004f},
		Name: "TableTopPitchAngleTolerance",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0050},
		Name: "TableTopRollAngleTolerance",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0051},
		Name: "TableTopVerticalPositionTolerance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0052},
		Name: "TableTopLongitudinalPositionTolerance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0053},
		Name: "TableTopLateralPositionTolerance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0055},
		Name: "RTPlanRelationship",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0070},
		Name: "FractionGroupSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0071},
		Name: "FractionGroupNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0072},
		Name: "FractionGroupDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0078},
		Name: "NumberOfFractionsPlanned",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0079},
		Name: "NumberOfFractionPatternDigitsPerDay",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x007a},
		Name: "RepeatFractionCycleLength",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x007b},
		Name: "FractionPattern",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0080},
		Name: "NumberOfBeams",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0083},
		Name: "ReferencedDoseReferenceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0084},
		Name: "BeamDose",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0086},
		Name: "BeamMeterset",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0088},
		Name: "BeamDosePointDepth",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0089},
		Name: "BeamDosePointEquivalentDepth",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x008a},
		Name: "BeamDosePointSSD",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x008b},
		Name: "BeamDoseMeaning",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x008c},
		Name: "BeamDoseVerificationControlPointSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0090},
		Name: "BeamDoseType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0091},
		Name: "AlternateBeamDose",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0092},
		Name: "AlternateBeamDoseType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0093},
		Name: "DepthValueAveragingFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0094},
		Name: "BeamDosePointSourceToExternalContourDistance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00a0},
		Name: "NumberOfBrachyApplicationSetups",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00a2},
		Name: "BrachyApplicationSetupDoseSpecificationPoint",
		VR:   "DS",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00a4},
		Name: "BrachyApplicationSetupDose",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00b0},
		Name: "BeamSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00b2},
		Name: "TreatmentMachineName",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00b3},
		Name: "PrimaryDosimeterUnit",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00b4},
		Name: "SourceAxisDistance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00b6},
		Name: "BeamLimitingDeviceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00b8},
		Name: "RTBeamLimitingDeviceType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00ba},
		Name: "SourceToBeamLimitingDeviceDistance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00bb},
		Name: "IsocenterToBeamLimitingDeviceDistance",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00bc},
		Name: "NumberOfLeafJawPairs",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00be},
		Name: "LeafPositionBoundaries",
		VR:   "DS",
		VM:   "3-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00c0},
		Name: "BeamNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00c2},
		Name: "BeamName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00c3},
		Name: "BeamDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00c4},
		Name: "BeamType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00c5},
		Name: "BeamDeliveryDurationLimit",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00c6},
		Name: "RadiationType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00c7},
		Name: "HighDoseTechniqueType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00c8},
		Name: "ReferenceImageNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00ca},
		Name: "PlannedVerificationImageSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00cc},
		Name: "ImagingDeviceSpecificAcquisitionParameters",
		VR:   "LO",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00ce},
		Name: "TreatmentDeliveryType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00d0},
		Name: "NumberOfWedges",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00d1},
		Name: "WedgeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00d2},
		Name: "WedgeNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00d3},
		Name: "WedgeType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00d4},
		Name: "WedgeID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00d5},
		Name: "WedgeAngle",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00d6},
		Name: "WedgeFactor",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00d7},
		Name: "TotalWedgeTrayWaterEquivalentThickness",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00d8},
		Name: "WedgeOrientation",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00d9},
		Name: "IsocenterToWedgeTrayDistance",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00da},
		Name: "SourceToWedgeTrayDistance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00db},
		Name: "WedgeThinEdgePosition",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00dc},
		Name: "BolusID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00dd},
		Name: "BolusDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00de},
		Name: "EffectiveWedgeAngle",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00e0},
		Name: "NumberOfCompensators",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00e1},
		Name: "MaterialID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00e2},
		Name: "TotalCompensatorTrayFactor",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00e3},
		Name: "CompensatorSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00e4},
		Name: "CompensatorNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00e5},
		Name: "CompensatorID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00e6},
		Name: "SourceToCompensatorTrayDistance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00e7},
		Name: "CompensatorRows",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00e8},
		Name: "CompensatorColumns",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00e9},
		Name: "CompensatorPixelSpacing",
		VR:   "DS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00ea},
		Name: "CompensatorPosition",
		VR:   "DS",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00eb},
		Name: "CompensatorTransmissionData",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00ec},
		Name: "CompensatorThicknessData",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00ed},
		Name: "NumberOfBoli",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00ee},
		Name: "CompensatorType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00ef},
		Name: "CompensatorTrayID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00f0},
		Name: "NumberOfBlocks",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00f2},
		Name: "TotalBlockTrayFactor",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00f3},
		Name: "TotalBlockTrayWaterEquivalentThickness",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00f4},
		Name: "BlockSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00f5},
		Name: "BlockTrayID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00f6},
		Name: "SourceToBlockTrayDistance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00f7},
		Name: "IsocenterToBlockTrayDistance",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00f8},
		Name: "BlockType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00f9},
		Name: "AccessoryCode",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00fa},
		Name: "BlockDivergence",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00fb},
		Name: "BlockMountingPosition",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00fc},
		Name: "BlockNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x00fe},
		Name: "BlockName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0100},
		Name: "BlockThickness",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0102},
		Name: "BlockTransmission",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0104},
		Name: "BlockNumberOfPoints",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0106},
		Name: "BlockData",
		VR:   "DS",
		VM:   "2-2n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0107},
		Name: "ApplicatorSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0108},
		Name: "ApplicatorID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0109},
		Name: "ApplicatorType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x010a},
		Name: "ApplicatorDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x010c},
		Name: "CumulativeDoseReferenceCoefficient",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x010e},
		Name: "FinalCumulativeMetersetWeight",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0110},
		Name: "NumberOfControlPoints",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0111},
		Name: "ControlPointSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0112},
		Name: "ControlPointIndex",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0114},
		Name: "NominalBeamEnergy",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0115},
		Name: "DoseRateSet",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0116},
		Name: "WedgePositionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0118},
		Name: "WedgePosition",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x011a},
		Name: "BeamLimitingDevicePositionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x011c},
		Name: "LeafJawPositions",
		VR:   "DS",
		VM:   "2-2n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x011e},
		Name: "GantryAngle",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x011f},
		Name: "GantryRotationDirection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0120},
		Name: "BeamLimitingDeviceAngle",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0121},
		Name: "BeamLimitingDeviceRotationDirection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0122},
		Name: "PatientSupportAngle",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0123},
		Name: "PatientSupportRotationDirection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0124},
		Name: "TableTopEccentricAxisDistance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0125},
		Name: "TableTopEccentricAngle",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0126},
		Name: "TableTopEccentricRotationDirection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0128},
		Name: "TableTopVerticalPosition",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0129},
		Name: "TableTopLongitudinalPosition",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x012a},
		Name: "TableTopLateralPosition",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x012c},
		Name: "IsocenterPosition",
		VR:   "DS",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x012e},
		Name: "SurfaceEntryPoint",
		VR:   "DS",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0130},
		Name: "SourceToSurfaceDistance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0131},
		Name: "AverageBeamDosePointSourceToExternalContourDistance",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0132},
		Name: "SourceToExternalContourDistance",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0133},
		Name: "ExternalContourEntryPoint",
		VR:   "FL",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0134},
		Name: "CumulativeMetersetWeight",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0140},
		Name: "TableTopPitchAngle",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0142},
		Name: "TableTopPitchRotationDirection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0144},
		Name: "TableTopRollAngle",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0146},
		Name: "TableTopRollRotationDirection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0148},
		Name: "HeadFixationAngle",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x014a},
		Name: "GantryPitchAngle",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x014c},
		Name: "GantryPitchRotationDirection",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x014e},
		Name: "GantryPitchAngleTolerance",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0150},
		Name: "FixationEye",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0151},
		Name: "ChairHeadFramePosition",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0152},
		Name: "HeadFixationAngleTolerance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0153},
		Name: "ChairHeadFramePositionTolerance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0154},
		Name: "FixationLightAzimuthalAngleTolerance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0155},
		Name: "FixationLightPolarAngleTolerance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0180},
		Name: "PatientSetupSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0182},
		Name: "PatientSetupNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0183},
		Name: "PatientSetupLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0184},
		Name: "PatientAdditionalPosition",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0190},
		Name: "FixationDeviceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0192},
		Name: "FixationDeviceType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0194},
		Name: "FixationDeviceLabel",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0196},
		Name: "FixationDeviceDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0198},
		Name: "FixationDevicePosition",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0199},
		Name: "FixationDevicePitchAngle",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x019a},
		Name: "FixationDeviceRollAngle",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x01a0},
		Name: "ShieldingDeviceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x01a2},
		Name: "ShieldingDeviceType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x01a4},
		Name: "ShieldingDeviceLabel",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x01a6},
		Name: "ShieldingDeviceDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x01a8},
		Name: "ShieldingDevicePosition",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x01b0},
		Name: "SetupTechnique",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x01b2},
		Name: "SetupTechniqueDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x01b4},
		Name: "SetupDeviceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x01b6},
		Name: "SetupDeviceType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x01b8},
		Name: "SetupDeviceLabel",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x01ba},
		Name: "SetupDeviceDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x01bc},
		Name: "SetupDeviceParameter",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x01d0},
		Name: "SetupReferenceDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x01d2},
		Name: "TableTopVerticalSetupDisplacement",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x01d4},
		Name: "TableTopLongitudinalSetupDisplacement",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x01d6},
		Name: "TableTopLateralSetupDisplacement",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0200},
		Name: "BrachyTreatmentTechnique",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0202},
		Name: "BrachyTreatmentType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0206},
		Name: "TreatmentMachineSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0210},
		Name: "SourceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0212},
		Name: "SourceNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0214},
		Name: "SourceType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0216},
		Name: "SourceManufacturer",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0218},
		Name: "ActiveSourceDiameter",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x021a},
		Name: "ActiveSourceLength",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x021b},
		Name: "SourceModelID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x021c},
		Name: "SourceDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0222},
		Name: "SourceEncapsulationNominalThickness",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0224},
		Name: "SourceEncapsulationNominalTransmission",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0226},
		Name: "SourceIsotopeName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0228},
		Name: "SourceIsotopeHalfLife",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0229},
		Name: "SourceStrengthUnits",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x022a},
		Name: "ReferenceAirKermaRate",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x022b},
		Name: "SourceStrength",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x022c},
		Name: "SourceStrengthReferenceDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x022e},
		Name: "SourceStrengthReferenceTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0230},
		Name: "ApplicationSetupSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0232},
		Name: "ApplicationSetupType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0234},
		Name: "ApplicationSetupNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0236},
		Name: "ApplicationSetupName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0238},
		Name: "ApplicationSetupManufacturer",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0240},
		Name: "TemplateNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0242},
		Name: "TemplateType",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0244},
		Name: "TemplateName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0250},
		Name: "TotalReferenceAirKerma",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0260},
		Name: "BrachyAccessoryDeviceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0262},
		Name: "BrachyAccessoryDeviceNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0263},
		Name: "BrachyAccessoryDeviceID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0264},
		Name: "BrachyAccessoryDeviceType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0266},
		Name: "BrachyAccessoryDeviceName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x026a},
		Name: "BrachyAccessoryDeviceNominalThickness",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x026c},
		Name: "BrachyAccessoryDeviceNominalTransmission",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0271},
		Name: "ChannelEffectiveLength",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0272},
		Name: "ChannelInnerLength",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0273},
		Name: "AfterloaderChannelID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0274},
		Name: "SourceApplicatorTipLength",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0280},
		Name: "ChannelSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0282},
		Name: "ChannelNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0284},
		Name: "ChannelLength",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0286},
		Name: "ChannelTotalTime",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0288},
		Name: "SourceMovementType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x028a},
		Name: "NumberOfPulses",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x028c},
		Name: "PulseRepetitionInterval",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0290},
		Name: "SourceApplicatorNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0291},
		Name: "SourceApplicatorID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0292},
		Name: "SourceApplicatorType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0294},
		Name: "SourceApplicatorName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0296},
		Name: "SourceApplicatorLength",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0298},
		Name: "SourceApplicatorManufacturer",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x029c},
		Name: "SourceApplicatorWallNominalThickness",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x029e},
		Name: "SourceApplicatorWallNominalTransmission",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02a0},
		Name: "SourceApplicatorStepSize",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02a2},
		Name: "TransferTubeNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02a4},
		Name: "TransferTubeLength",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02b0},
		Name: "ChannelShieldSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02b2},
		Name: "ChannelShieldNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02b3},
		Name: "ChannelShieldID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02b4},
		Name: "ChannelShieldName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02b8},
		Name: "ChannelShieldNominalThickness",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02ba},
		Name: "ChannelShieldNominalTransmission",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02c8},
		Name: "FinalCumulativeTimeWeight",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02d0},
		Name: "BrachyControlPointSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02d2},
		Name: "ControlPointRelativePosition",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02d4},
		Name: "ControlPoint3DPosition",
		VR:   "DS",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02d6},
		Name: "CumulativeTimeWeight",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02e0},
		Name: "CompensatorDivergence",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02e1},
		Name: "CompensatorMountingPosition",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02e2},
		Name: "SourceToCompensatorDistance",
		VR:   "DS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02e3},
		Name: "TotalCompensatorTrayWaterEquivalentThickness",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02e4},
		Name: "IsocenterToCompensatorTrayDistance",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02e5},
		Name: "CompensatorColumnOffset",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02e6},
		Name: "IsocenterToCompensatorDistances",
		VR:   "FL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02e7},
		Name: "CompensatorRelativeStoppingPowerRatio",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02e8},
		Name: "CompensatorMillingToolDiameter",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02ea},
		Name: "IonRangeCompensatorSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x02eb},
		Name: "CompensatorDescription",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0302},
		Name: "RadiationMassNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0304},
		Name: "RadiationAtomicNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0306},
		Name: "RadiationChargeState",
		VR:   "SS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0308},
		Name: "ScanMode",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0309},
		Name: "ModulatedScanModeType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x030a},
		Name: "VirtualSourceAxisDistances",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x030c},
		Name: "SnoutSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x030d},
		Name: "SnoutPosition",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x030f},
		Name: "SnoutID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0312},
		Name: "NumberOfRangeShifters",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0314},
		Name: "RangeShifterSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0316},
		Name: "RangeShifterNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0318},
		Name: "RangeShifterID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0320},
		Name: "RangeShifterType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0322},
		Name: "RangeShifterDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0330},
		Name: "NumberOfLateralSpreadingDevices",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0332},
		Name: "LateralSpreadingDeviceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0334},
		Name: "LateralSpreadingDeviceNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0336},
		Name: "LateralSpreadingDeviceID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0338},
		Name: "LateralSpreadingDeviceType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x033a},
		Name: "LateralSpreadingDeviceDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x033c},
		Name: "LateralSpreadingDeviceWaterEquivalentThickness",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0340},
		Name: "NumberOfRangeModulators",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0342},
		Name: "RangeModulatorSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0344},
		Name: "RangeModulatorNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0346},
		Name: "RangeModulatorID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0348},
		Name: "RangeModulatorType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x034a},
		Name: "RangeModulatorDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x034c},
		Name: "BeamCurrentModulationID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0350},
		Name: "PatientSupportType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0352},
		Name: "PatientSupportID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0354},
		Name: "PatientSupportAccessoryCode",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0355},
		Name: "TrayAccessoryCode",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0356},
		Name: "FixationLightAzimuthalAngle",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0358},
		Name: "FixationLightPolarAngle",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x035a},
		Name: "MetersetRate",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0360},
		Name: "RangeShifterSettingsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0362},
		Name: "RangeShifterSetting",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0364},
		Name: "IsocenterToRangeShifterDistance",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0366},
		Name: "RangeShifterWaterEquivalentThickness",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0370},
		Name: "LateralSpreadingDeviceSettingsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0372},
		Name: "LateralSpreadingDeviceSetting",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0374},
		Name: "IsocenterToLateralSpreadingDeviceDistance",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0380},
		Name: "RangeModulatorSettingsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0382},
		Name: "RangeModulatorGatingStartValue",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0384},
		Name: "RangeModulatorGatingStopValue",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0386},
		Name: "RangeModulatorGatingStartWaterEquivalentThickness",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0388},
		Name: "RangeModulatorGatingStopWaterEquivalentThickness",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x038a},
		Name: "IsocenterToRangeModulatorDistance",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x038f},
		Name: "ScanSpotTimeOffset",
		VR:   "FL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0390},
		Name: "ScanSpotTuneID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0391},
		Name: "ScanSpotPrescribedIndices",
		VR:   "IS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0392},
		Name: "NumberOfScanSpotPositions",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0393},
		Name: "ScanSpotReordered",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0394},
		Name: "ScanSpotPositionMap",
		VR:   "FL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0395},
		Name: "ScanSpotReorderingAllowed",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0396},
		Name: "ScanSpotMetersetWeights",
		VR:   "FL",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0398},
		Name: "ScanningSpotSize",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x039a},
		Name: "NumberOfPaintings",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x03a0},
		Name: "IonToleranceTableSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x03a2},
		Name: "IonBeamSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x03a4},
		Name: "IonBeamLimitingDeviceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x03a6},
		Name: "IonBlockSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x03a8},
		Name: "IonControlPointSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x03aa},
		Name: "IonWedgeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x03ac},
		Name: "IonWedgePositionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0401},
		Name: "ReferencedSetupImageSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0402},
		Name: "SetupImageComment",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0410},
		Name: "MotionSynchronizationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0412},
		Name: "ControlPointOrientation",
		VR:   "FL",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0420},
		Name: "GeneralAccessorySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0421},
		Name: "GeneralAccessoryID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0422},
		Name: "GeneralAccessoryDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0423},
		Name: "GeneralAccessoryType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0424},
		Name: "GeneralAccessoryNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0425},
		Name: "SourceToGeneralAccessoryDistance",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0426},
		Name: "IsocenterToGeneralAccessoryDistance",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0431},
		Name: "ApplicatorGeometrySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0432},
		Name: "ApplicatorApertureShape",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0433},
		Name: "ApplicatorOpening",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0434},
		Name: "ApplicatorOpeningX",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0435},
		Name: "ApplicatorOpeningY",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0436},
		Name: "SourceToApplicatorMountingPositionDistance",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0440},
		Name: "NumberOfBlockSlabItems",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0441},
		Name: "BlockSlabSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0442},
		Name: "BlockSlabThickness",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0443},
		Name: "BlockSlabNumber",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0450},
		Name: "DeviceMotionControlSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0451},
		Name: "DeviceMotionExecutionMode",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0452},
		Name: "DeviceMotionObservationMode",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0453},
		Name: "DeviceMotionParameterCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0501},
		Name: "DistalDepthFraction",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0502},
		Name: "DistalDepth",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0503},
		Name: "NominalRangeModulationFractions",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0504},
		Name: "NominalRangeModulatedRegionDepths",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0505},
		Name: "DepthDoseParametersSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0506},
		Name: "DeliveredDepthDoseParametersSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0507},
		Name: "DeliveredDistalDepthFraction",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0508},
		Name: "DeliveredDistalDepth",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0509},
		Name: "DeliveredNominalRangeModulationFractions",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0510},
		Name: "DeliveredNominalRangeModulatedRegionDepths",
		VR:   "FL",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0511},
		Name: "DeliveredReferenceDoseDefinition",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0512},
		Name: "ReferenceDoseDefinition",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0600},
		Name: "RTControlPointIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0601},
		Name: "RadiationGenerationModeIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0602},
		Name: "ReferencedDefinedDeviceIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0603},
		Name: "RadiationDoseIdentificationIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0604},
		Name: "NumberOfRTControlPoints",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0605},
		Name: "ReferencedRadiationGenerationModeIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0606},
		Name: "TreatmentPositionIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0607},
		Name: "ReferencedDeviceIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0608},
		Name: "TreatmentPositionGroupLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0609},
		Name: "TreatmentPositionGroupUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x060a},
		Name: "TreatmentPositionGroupSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x060b},
		Name: "ReferencedTreatmentPositionIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x060c},
		Name: "ReferencedRadiationDoseIdentificationIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x060d},
		Name: "RTAccessoryHolderWaterEquivalentThickness",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x060e},
		Name: "ReferencedRTAccessoryHolderDeviceIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x060f},
		Name: "RTAccessoryHolderSlotExistenceFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0610},
		Name: "RTAccessoryHolderSlotSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0611},
		Name: "RTAccessoryHolderSlotID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0612},
		Name: "RTAccessoryHolderSlotDistance",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0613},
		Name: "RTAccessorySlotDistance",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0614},
		Name: "RTAccessoryHolderDefinitionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0615},
		Name: "RTAccessoryDeviceSlotID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0616},
		Name: "RTRadiationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0617},
		Name: "RadiationDoseSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0618},
		Name: "RadiationDoseIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0619},
		Name: "RadiationDoseIdentificationLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x061a},
		Name: "ReferenceDoseType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x061b},
		Name: "PrimaryDoseValueIndicator",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x061c},
		Name: "DoseValuesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x061d},
		Name: "DoseValuePurpose",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x061e},
		Name: "ReferenceDosePointCoordinates",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x061f},
		Name: "RadiationDoseValuesParametersSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0620},
		Name: "MetersetToDoseMappingSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0621},
		Name: "ExpectedInVivoMeasurementValuesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0622},
		Name: "ExpectedInVivoMeasurementValueIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0623},
		Name: "RadiationDoseInVivoMeasurementLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0624},
		Name: "RadiationDoseCentralAxisDisplacement",
		VR:   "FD",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0625},
		Name: "RadiationDoseValue",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0626},
		Name: "RadiationDoseSourceToSkinDistance",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0627},
		Name: "RadiationDoseMeasurementPointCoordinates",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0628},
		Name: "RadiationDoseSourceToExternalContourDistance",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0629},
		Name: "RTToleranceSetSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x062a},
		Name: "RTToleranceSetLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x062b},
		Name: "AttributeToleranceValuesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x062c},
		Name: "ToleranceValue",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x062d},
		Name: "PatientSupportPositionToleranceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x062e},
		Name: "TreatmentTimeLimit",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x062f},
		Name: "CArmPhotonElectronControlPointSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0630},
		Name: "ReferencedRTRadiationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0631},
		Name: "ReferencedRTInstanceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0632},
		Name: "ReferencedRTPatientSetupSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0634},
		Name: "SourceToPatientSurfaceDistance",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0635},
		Name: "TreatmentMachineSpecialModeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0636},
		Name: "IntendedNumberOfFractions",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0637},
		Name: "RTRadiationSetIntent",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0638},
		Name: "RTRadiationPhysicalAndGeometricContentDetailFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0639},
		Name: "RTRecordFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x063a},
		Name: "TreatmentDeviceIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x063b},
		Name: "ReferencedRTPhysicianIntentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x063c},
		Name: "CumulativeMeterset",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x063d},
		Name: "DeliveryRate",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x063e},
		Name: "DeliveryRateUnitSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x063f},
		Name: "TreatmentPositionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0640},
		Name: "RadiationSourceAxisDistance",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0641},
		Name: "NumberOfRTBeamLimitingDevices",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0642},
		Name: "RTBeamLimitingDeviceProximalDistance",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0643},
		Name: "RTBeamLimitingDeviceDistalDistance",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0644},
		Name: "ParallelRTBeamDelimiterDeviceOrientationLabelCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0645},
		Name: "BeamModifierOrientationAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0646},
		Name: "FixedRTBeamDelimiterDeviceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0647},
		Name: "ParallelRTBeamDelimiterDeviceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0648},
		Name: "NumberOfParallelRTBeamDelimiters",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0649},
		Name: "ParallelRTBeamDelimiterBoundaries",
		VR:   "FD",
		VM:   "2-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x064a},
		Name: "ParallelRTBeamDelimiterPositions",
		VR:   "FD",
		VM:   "2-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x064b},
		Name: "RTBeamLimitingDeviceOffset",
		VR:   "FD",
		VM:   "2",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x064c},
		Name: "RTBeamDelimiterGeometrySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x064d},
		Name: "RTBeamLimitingDeviceDefinitionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x064e},
		Name: "ParallelRTBeamDelimiterOpeningMode",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x064f},
		Name: "ParallelRTBeamDelimiterLeafMountingSide",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0650},
		Name: "PatientSetupUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0651},
		Name: "WedgeDefinitionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0652},
		Name: "RadiationBeamWedgeAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0653},
		Name: "RadiationBeamWedgeThinEdgeDistance",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0654},
		Name: "RadiationBeamEffectiveWedgeAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0655},
		Name: "NumberOfWedgePositions",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0656},
		Name: "RTBeamLimitingDeviceOpeningSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0657},
		Name: "NumberOfRTBeamLimitingDeviceOpenings",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0658},
		Name: "RadiationDosimeterUnitSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0659},
		Name: "RTDeviceDistanceReferenceLocationCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x065a},
		Name: "RadiationDeviceConfigurationAndCommissioningKeySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x065b},
		Name: "PatientSupportPositionParameterSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x065c},
		Name: "PatientSupportPositionSpecificationMethod",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x065d},
		Name: "PatientSupportPositionDeviceParameterSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x065e},
		Name: "DeviceOrderIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x065f},
		Name: "PatientSupportPositionParameterOrderIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0660},
		Name: "PatientSupportPositionDeviceToleranceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0661},
		Name: "PatientSupportPositionToleranceOrderIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0662},
		Name: "CompensatorDefinitionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0663},
		Name: "CompensatorMapOrientation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0664},
		Name: "CompensatorProximalThicknessMap",
		VR:   "OF",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0665},
		Name: "CompensatorDistalThicknessMap",
		VR:   "OF",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0666},
		Name: "CompensatorBasePlaneOffset",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0667},
		Name: "CompensatorShapeFabricationCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0668},
		Name: "CompensatorShapeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0669},
		Name: "RadiationBeamCompensatorMillingToolDiameter",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x066a},
		Name: "BlockDefinitionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x066b},
		Name: "BlockEdgeData",
		VR:   "OF",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x066c},
		Name: "BlockOrientation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x066d},
		Name: "RadiationBeamBlockThickness",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x066e},
		Name: "RadiationBeamBlockSlabThickness",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x066f},
		Name: "BlockEdgeDataSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0670},
		Name: "NumberOfRTAccessoryHolders",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0671},
		Name: "GeneralAccessoryDefinitionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0672},
		Name: "NumberOfGeneralAccessories",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0673},
		Name: "BolusDefinitionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0674},
		Name: "NumberOfBoluses",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0675},
		Name: "EquipmentFrameOfReferenceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0676},
		Name: "EquipmentFrameOfReferenceDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0677},
		Name: "EquipmentReferencePointCoordinatesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0678},
		Name: "EquipmentReferencePointCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0679},
		Name: "RTBeamLimitingDeviceAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x067a},
		Name: "SourceRollAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x067b},
		Name: "RadiationGenerationModeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x067c},
		Name: "RadiationGenerationModeLabel",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x067d},
		Name: "RadiationGenerationModeDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x067e},
		Name: "RadiationGenerationModeMachineCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x067f},
		Name: "RadiationTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0680},
		Name: "NominalEnergy",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0681},
		Name: "MinimumNominalEnergy",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0682},
		Name: "MaximumNominalEnergy",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0683},
		Name: "RadiationFluenceModifierCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0684},
		Name: "EnergyUnitCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0685},
		Name: "NumberOfRadiationGenerationModes",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0686},
		Name: "PatientSupportDevicesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0687},
		Name: "NumberOfPatientSupportDevices",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0688},
		Name: "RTBeamModifierDefinitionDistance",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x0689},
		Name: "BeamAreaLimitSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300a, 0x068a},
		Name: "ReferencedRTPrescriptionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x0002},
		Name: "ReferencedRTPlanSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x0004},
		Name: "ReferencedBeamSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x0006},
		Name: "ReferencedBeamNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x0007},
		Name: "ReferencedReferenceImageNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x0008},
		Name: "StartCumulativeMetersetWeight",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x0009},
		Name: "EndCumulativeMetersetWeight",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x000a},
		Name: "ReferencedBrachyApplicationSetupSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x000c},
		Name: "ReferencedBrachyApplicationSetupNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x000e},
		Name: "ReferencedSourceNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x0020},
		Name: "ReferencedFractionGroupSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x0022},
		Name: "ReferencedFractionGroupNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x0040},
		Name: "ReferencedVerificationImageSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x0042},
		Name: "ReferencedReferenceImageSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x0050},
		Name: "ReferencedDoseReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x0051},
		Name: "ReferencedDoseReferenceNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x0055},
		Name: "BrachyReferencedDoseReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x0060},
		Name: "ReferencedStructureSetSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x006a},
		Name: "ReferencedPatientSetupNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x0080},
		Name: "ReferencedDoseSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x00a0},
		Name: "ReferencedToleranceTableNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x00b0},
		Name: "ReferencedBolusSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x00c0},
		Name: "ReferencedWedgeNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x00d0},
		Name: "ReferencedCompensatorNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x00e0},
		Name: "ReferencedBlockNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x00f0},
		Name: "ReferencedControlPointIndex",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x00f2},
		Name: "ReferencedControlPointSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x00f4},
		Name: "ReferencedStartControlPointIndex",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x00f6},
		Name: "ReferencedStopControlPointIndex",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x0100},
		Name: "ReferencedRangeShifterNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x0102},
		Name: "ReferencedLateralSpreadingDeviceNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x0104},
		Name: "ReferencedRangeModulatorNumber",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x0111},
		Name: "OmittedBeamTaskSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x0112},
		Name: "ReasonForOmission",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300c, 0x0113},
		Name: "ReasonForOmissionDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300e, 0x0002},
		Name: "ApprovalStatus",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300e, 0x0004},
		Name: "ReviewDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300e, 0x0005},
		Name: "ReviewTime",
		VR:   "TM",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x300e, 0x0008},
		Name: "ReviewerName",
		VR:   "PN",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0001},
		Name: "RadiobiologicalDoseEffectSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0002},
		Name: "RadiobiologicalDoseEffectFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0003},
		Name: "EffectiveDoseCalculationMethodCategoryCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0004},
		Name: "EffectiveDoseCalculationMethodCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0005},
		Name: "EffectiveDoseCalculationMethodDescription",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0006},
		Name: "ConceptualVolumeUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0007},
		Name: "OriginatingSOPInstanceReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0008},
		Name: "ConceptualVolumeConstituentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0009},
		Name: "EquivalentConceptualVolumeInstanceReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x000a},
		Name: "EquivalentConceptualVolumesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x000b},
		Name: "ReferencedConceptualVolumeUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x000c},
		Name: "ConceptualVolumeCombinationExpression",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x000d},
		Name: "ConceptualVolumeConstituentIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x000e},
		Name: "ConceptualVolumeCombinationFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x000f},
		Name: "ConceptualVolumeCombinationDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0010},
		Name: "ConceptualVolumeSegmentationDefinedFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0011},
		Name: "ConceptualVolumeSegmentationReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0012},
		Name: "ConceptualVolumeConstituentSegmentationReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0013},
		Name: "ConstituentConceptualVolumeUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0014},
		Name: "DerivationConceptualVolumeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0015},
		Name: "SourceConceptualVolumeUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0016},
		Name: "ConceptualVolumeDerivationAlgorithmSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0017},
		Name: "ConceptualVolumeDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0018},
		Name: "SourceConceptualVolumeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0019},
		Name: "AuthorIdentificationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x001a},
		Name: "ManufacturerModelVersion",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x001b},
		Name: "DeviceAlternateIdentifier",
		VR:   "UC",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x001c},
		Name: "DeviceAlternateIdentifierType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x001d},
		Name: "DeviceAlternateIdentifierFormat",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x001e},
		Name: "SegmentationCreationTemplateLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x001f},
		Name: "SegmentationTemplateUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0020},
		Name: "ReferencedSegmentReferenceIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0021},
		Name: "SegmentReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0022},
		Name: "SegmentReferenceIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0023},
		Name: "DirectSegmentReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0024},
		Name: "CombinationSegmentReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0025},
		Name: "ConceptualVolumeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0026},
		Name: "SegmentedRTAccessoryDeviceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0027},
		Name: "SegmentCharacteristicsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0028},
		Name: "RelatedSegmentCharacteristicsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0029},
		Name: "SegmentCharacteristicsPrecedence",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x002a},
		Name: "RTSegmentAnnotationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x002b},
		Name: "SegmentAnnotationCategoryCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x002c},
		Name: "SegmentAnnotationTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x002d},
		Name: "DeviceLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x002e},
		Name: "DeviceTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x002f},
		Name: "SegmentAnnotationTypeModifierCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0030},
		Name: "PatientEquipmentRelationshipCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0031},
		Name: "ReferencedFiducialsUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0032},
		Name: "PatientTreatmentOrientationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0033},
		Name: "UserContentLabel",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0034},
		Name: "UserContentLongLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0035},
		Name: "EntityLabel",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0036},
		Name: "EntityName",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0037},
		Name: "EntityDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0038},
		Name: "EntityLongLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0039},
		Name: "DeviceIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x003a},
		Name: "RTTreatmentPhaseIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x003b},
		Name: "RTTreatmentPhaseUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x003c},
		Name: "RTPrescriptionIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x003d},
		Name: "RTSegmentAnnotationIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x003e},
		Name: "BasisRTTreatmentPhaseIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x003f},
		Name: "RelatedRTTreatmentPhaseIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0040},
		Name: "ReferencedRTTreatmentPhaseIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0041},
		Name: "ReferencedRTPrescriptionIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0042},
		Name: "ReferencedParentRTPrescriptionIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0043},
		Name: "ManufacturerDeviceIdentifier",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0044},
		Name: "InstanceLevelReferencedPerformedProcedureStepSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0045},
		Name: "RTTreatmentPhaseIntentPresenceFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0046},
		Name: "RadiotherapyTreatmentType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0047},
		Name: "TeletherapyRadiationType",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0048},
		Name: "BrachytherapySourceType",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0049},
		Name: "ReferencedRTTreatmentPhaseSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x004a},
		Name: "ReferencedDirectSegmentInstanceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x004b},
		Name: "IntendedRTTreatmentPhaseSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x004c},
		Name: "IntendedPhaseStartDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x004d},
		Name: "IntendedPhaseEndDate",
		VR:   "DA",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x004e},
		Name: "RTTreatmentPhaseIntervalSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x004f},
		Name: "TemporalRelationshipIntervalAnchor",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0050},
		Name: "MinimumNumberOfIntervalDays",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0051},
		Name: "MaximumNumberOfIntervalDays",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0052},
		Name: "PertinentSOPClassesInStudy",
		VR:   "UI",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0053},
		Name: "PertinentSOPClassesInSeries",
		VR:   "UI",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0054},
		Name: "RTPrescriptionLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0055},
		Name: "RTPhysicianIntentPredecessorSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0056},
		Name: "RTTreatmentApproachLabel",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0057},
		Name: "RTPhysicianIntentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0058},
		Name: "RTPhysicianIntentIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0059},
		Name: "RTTreatmentIntentType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x005a},
		Name: "RTPhysicianIntentNarrative",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x005b},
		Name: "RTProtocolCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x005c},
		Name: "ReasonForSuperseding",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x005d},
		Name: "RTDiagnosisCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x005e},
		Name: "ReferencedRTPhysicianIntentIndex",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x005f},
		Name: "RTPhysicianIntentInputInstanceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0060},
		Name: "RTAnatomicPrescriptionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0061},
		Name: "PriorTreatmentDoseDescription",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0062},
		Name: "PriorTreatmentReferenceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0063},
		Name: "DosimetricObjectiveEvaluationScope",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0064},
		Name: "TherapeuticRoleCategoryCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0065},
		Name: "TherapeuticRoleTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0066},
		Name: "ConceptualVolumeOptimizationPrecedence",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0067},
		Name: "ConceptualVolumeCategoryCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0068},
		Name: "ConceptualVolumeBlockingConstraint",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0069},
		Name: "ConceptualVolumeTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x006a},
		Name: "ConceptualVolumeTypeModifierCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x006b},
		Name: "RTPrescriptionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x006c},
		Name: "DosimetricObjectiveSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x006d},
		Name: "DosimetricObjectiveTypeCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x006e},
		Name: "DosimetricObjectiveUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x006f},
		Name: "ReferencedDosimetricObjectiveUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0070},
		Name: "DosimetricObjectiveParameterSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0071},
		Name: "ReferencedDosimetricObjectivesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0073},
		Name: "AbsoluteDosimetricObjectiveFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0074},
		Name: "DosimetricObjectiveWeight",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0075},
		Name: "DosimetricObjectivePurpose",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0076},
		Name: "PlanningInputInformationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0077},
		Name: "TreatmentSite",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0078},
		Name: "TreatmentSiteCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0079},
		Name: "FractionPatternSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x007a},
		Name: "TreatmentTechniqueNotes",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x007b},
		Name: "PrescriptionNotes",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x007c},
		Name: "NumberOfIntervalFractions",
		VR:   "IS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x007d},
		Name: "NumberOfFractions",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x007e},
		Name: "IntendedDeliveryDuration",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x007f},
		Name: "FractionationNotes",
		VR:   "UT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0080},
		Name: "RTTreatmentTechniqueCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0081},
		Name: "PrescriptionNotesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0082},
		Name: "FractionBasedRelationshipSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0083},
		Name: "FractionBasedRelationshipIntervalAnchor",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0084},
		Name: "MinimumHoursBetweenFractions",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0085},
		Name: "IntendedFractionStartTime",
		VR:   "TM",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0086},
		Name: "IntendedStartDayOfWeek",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0087},
		Name: "WeekdayFractionPatternSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0088},
		Name: "DeliveryTimeStructureCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0089},
		Name: "TreatmentSiteModifierCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0090},
		Name: "RoboticBaseLocationIndicator",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0091},
		Name: "RoboticPathNodeSetCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0092},
		Name: "RoboticNodeIdentifier",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0093},
		Name: "RTTreatmentSourceCoordinates",
		VR:   "FD",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0094},
		Name: "RadiationSourceCoordinateSystemYawAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0095},
		Name: "RadiationSourceCoordinateSystemRollAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0096},
		Name: "RadiationSourceCoordinateSystemPitchAngle",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0097},
		Name: "RoboticPathControlPointSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0098},
		Name: "TomotherapeuticControlPointSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x0099},
		Name: "TomotherapeuticLeafOpenDurations",
		VR:   "FD",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x3010, 0x009a},
		Name: "TomotherapeuticLeafInitialClosedDurations",
		VR:   "FD",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x0001},
		Name: "LowEnergyDetectors",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x0002},
		Name: "HighEnergyDetectors",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x0004},
		Name: "DetectorGeometrySequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1001},
		Name: "ThreatROIVoxelSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1004},
		Name: "ThreatROIBase",
		VR:   "FL",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1005},
		Name: "ThreatROIExtents",
		VR:   "FL",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1006},
		Name: "ThreatROIBitmap",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1007},
		Name: "RouteSegmentID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1008},
		Name: "GantryType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1009},
		Name: "OOIOwnerType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x100a},
		Name: "RouteSegmentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1010},
		Name: "PotentialThreatObjectID",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1011},
		Name: "ThreatSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1012},
		Name: "ThreatCategory",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1013},
		Name: "ThreatCategoryDescription",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1014},
		Name: "ATDAbilityAssessment",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1015},
		Name: "ATDAssessmentFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1016},
		Name: "ATDAssessmentProbability",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1017},
		Name: "Mass",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1018},
		Name: "Density",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1019},
		Name: "ZEffective",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x101a},
		Name: "BoardingPassID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x101b},
		Name: "CenterOfMass",
		VR:   "FL",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x101c},
		Name: "CenterOfPTO",
		VR:   "FL",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x101d},
		Name: "BoundingPolygon",
		VR:   "FL",
		VM:   "6-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x101e},
		Name: "RouteSegmentStartLocationID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x101f},
		Name: "RouteSegmentEndLocationID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1020},
		Name: "RouteSegmentLocationIDType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1021},
		Name: "AbortReason",
		VR:   "CS",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1023},
		Name: "VolumeOfPTO",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1024},
		Name: "AbortFlag",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1025},
		Name: "RouteSegmentStartTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1026},
		Name: "RouteSegmentEndTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1027},
		Name: "TDRType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1028},
		Name: "InternationalRouteSegment",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1029},
		Name: "ThreatDetectionAlgorithmandVersion",
		VR:   "LO",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x102a},
		Name: "AssignedLocation",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x102b},
		Name: "AlarmDecisionTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1031},
		Name: "AlarmDecision",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1033},
		Name: "NumberOfTotalObjects",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1034},
		Name: "NumberOfAlarmObjects",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1037},
		Name: "PTORepresentationSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1038},
		Name: "ATDAssessmentSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1039},
		Name: "TIPType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x103a},
		Name: "DICOSVersion",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1041},
		Name: "OOIOwnerCreationTime",
		VR:   "DT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1042},
		Name: "OOIType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1043},
		Name: "OOISize",
		VR:   "FL",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1044},
		Name: "AcquisitionStatus",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1045},
		Name: "BasisMaterialsCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1046},
		Name: "PhantomType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1047},
		Name: "OOIOwnerSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1048},
		Name: "ScanType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1051},
		Name: "ItineraryID",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1052},
		Name: "ItineraryIDType",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1053},
		Name: "ItineraryIDAssigningAuthority",
		VR:   "LO",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1054},
		Name: "RouteID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1055},
		Name: "RouteIDAssigningAuthority",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1056},
		Name: "InboundArrivalType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1058},
		Name: "CarrierID",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1059},
		Name: "CarrierIDAssigningAuthority",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1060},
		Name: "SourceOrientation",
		VR:   "FL",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1061},
		Name: "SourcePosition",
		VR:   "FL",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1062},
		Name: "BeltHeight",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1064},
		Name: "AlgorithmRoutingCodeSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1067},
		Name: "TransportClassification",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1068},
		Name: "OOITypeDescriptor",
		VR:   "LT",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1069},
		Name: "TotalProcessingTime",
		VR:   "FL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x106c},
		Name: "DetectorCalibrationData",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x106d},
		Name: "AdditionalScreeningPerformed",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x106e},
		Name: "AdditionalInspectionSelectionCriteria",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x106f},
		Name: "AdditionalInspectionMethodSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1070},
		Name: "AITDeviceType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1071},
		Name: "QRMeasurementsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1072},
		Name: "TargetMaterialSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1073},
		Name: "SNRThreshold",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1075},
		Name: "ImageScaleRepresentation",
		VR:   "DS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1076},
		Name: "ReferencedPTOSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1077},
		Name: "ReferencedTDRInstanceSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1078},
		Name: "PTOLocationDescription",
		VR:   "ST",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x1079},
		Name: "AnomalyLocatorIndicatorSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x107a},
		Name: "AnomalyLocatorIndicator",
		VR:   "FL",
		VM:   "3",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x107b},
		Name: "PTORegionSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x107c},
		Name: "InspectionSelectionCriteria",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x107d},
		Name: "SecondaryInspectionMethodSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4010, 0x107e},
		Name: "PRCSToRCSOrientation",
		VR:   "DS",
		VM:   "6",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x4ffe, 0x0001},
		Name: "MACParametersSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x5200, 0x9229},
		Name: "SharedFunctionalGroupsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x5200, 0x9230},
		Name: "PerFrameFunctionalGroupsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x5400, 0x0100},
		Name: "WaveformSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x5400, 0x0110},
		Name: "ChannelMinimumValue",
		VR:   "OW",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x5400, 0x0112},
		Name: "ChannelMaximumValue",
		VR:   "OW",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x5400, 0x1004},
		Name: "WaveformBitsAllocated",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x5400, 0x1006},
		Name: "WaveformSampleInterpretation",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x5400, 0x100a},
		Name: "WaveformPaddingValue",
		VR:   "OW",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x5400, 0x1010},
		Name: "WaveformData",
		VR:   "OW",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x5600, 0x0010},
		Name: "FirstOrderPhaseCorrectionAngle",
		VR:   "OF",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x5600, 0x0020},
		Name: "SpectroscopyData",
		VR:   "OF",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x7fe0, 0x0001},
		Name: "ExtendedOffsetTable",
		VR:   "OV",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x7fe0, 0x0002},
		Name: "ExtendedOffsetTableLengths",
		VR:   "OV",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x7fe0, 0x0008},
		Name: "FloatPixelData",
		VR:   "OF",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x7fe0, 0x0009},
		Name: "DoubleFloatPixelData",
		VR:   "OD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x7fe0, 0x0010},
		Name: "PixelData",
		VR:   "OW",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0xfffa, 0xfffa},
		Name: "DigitalSignaturesSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0xfffc, 0xfffc},
		Name: "DataSetTrailingPadding",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0xfffe, 0xe000},
		Name: "Item",
		VR:   "See Note 2",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0xfffe, 0xe00d},
		Name: "ItemDelimitationItem",
		VR:   "See Note 2",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0xfffe, 0xe0dd},
		Name: "SequenceDelimitationItem",
		VR:   "See Note 2",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0000},
		Name: "FileMetaInformationGroupLength",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0001},
		Name: "FileMetaInformationVersion",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0002},
		Name: "MediaStorageSOPClassUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0003},
		Name: "MediaStorageSOPInstanceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0010},
		Name: "TransferSyntaxUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0012},
		Name: "ImplementationClassUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0013},
		Name: "ImplementationVersionName",
		VR:   "SH",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0016},
		Name: "SourceApplicationEntityTitle",
		VR:   "AE",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0017},
		Name: "SendingApplicationEntityTitle",
		VR:   "AE",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0018},
		Name: "ReceivingApplicationEntityTitle",
		VR:   "AE",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0026},
		Name: "SourcePresentationAddress",
		VR:   "UR",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0027},
		Name: "SendingPresentationAddress",
		VR:   "UR",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0028},
		Name: "ReceivingPresentationAddress",
		VR:   "UR",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0031},
		Name: "RTVMetaInformationVersion",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0032},
		Name: "RTVCommunicationSOPClassUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0033},
		Name: "RTVCommunicationSOPInstanceUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0035},
		Name: "RTVSourceIdentifier",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0036},
		Name: "RTVFlowIdentifier",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0037},
		Name: "RTVFlowRTPSamplingRate",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0038},
		Name: "RTVFlowActualFrameDuration",
		VR:   "FD",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0100},
		Name: "PrivateInformationCreatorUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0002, 0x0102},
		Name: "PrivateInformation",
		VR:   "OB",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0004, 0x1130},
		Name: "FileSetID",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0004, 0x1141},
		Name: "FileSetDescriptorFileID",
		VR:   "CS",
		VM:   "1-8",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0004, 0x1142},
		Name: "SpecificCharacterSetOfFileSetDescriptorFile",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0004, 0x1200},
		Name: "OffsetOfTheFirstDirectoryRecordOfTheRootDirectoryEntity",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0004, 0x1202},
		Name: "OffsetOfTheLastDirectoryRecordOfTheRootDirectoryEntity",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0004, 0x1212},
		Name: "FileSetConsistencyFlag",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0004, 0x1220},
		Name: "DirectoryRecordSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0004, 0x1400},
		Name: "OffsetOfTheNextDirectoryRecord",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0004, 0x1410},
		Name: "RecordInUseFlag",
		VR:   "US",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0004, 0x1420},
		Name: "OffsetOfReferencedLowerLevelDirectoryEntity",
		VR:   "UL",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0004, 0x1430},
		Name: "DirectoryRecordType",
		VR:   "CS",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0004, 0x1432},
		Name: "PrivateRecordUID",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0004, 0x1500},
		Name: "ReferencedFileID",
		VR:   "CS",
		VM:   "1-8",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0004, 0x1510},
		Name: "ReferencedSOPClassUIDInFile",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0004, 0x1511},
		Name: "ReferencedSOPInstanceUIDInFile",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0004, 0x1512},
		Name: "ReferencedTransferSyntaxUIDInFile",
		VR:   "UI",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0004, 0x151a},
		Name: "ReferencedRelatedGeneralSOPClassUIDInFile",
		VR:   "UI",
		VM:   "1-n",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

	thisInfo = Info{
		Tag:  Tag{0x0006, 0x0001},
		Name: "CurrentFrameFunctionalGroupsSequence",
		VR:   "SQ",
		VM:   "1",
	}
	tagDict[thisInfo.Tag] = thisInfo
	tagDictByKeyword[thisInfo.Name] = thisInfo

}
