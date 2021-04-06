import sys
import pydicom.config
import pydicom.dataset

pydicom.config.replace_un_with_known_vr = False

if __name__ == '__main__':
    """
    This script is called by roundtrip_test.go to check that the go library's writes
    are both valid and match their original values.
    
    The script takes two positional arguments:
    
        arg1: the path to the original DICOM file.
        arg2: the path to the file that has been parsed and re-written.
    """
    original_file: str = sys.argv[1]
    written_file: str = sys.argv[2]

    print("ORIGINAL PATH:", original_file)
    print("WRITTEN PATH:", written_file)

    original_dataset = pydicom.dcmread(original_file)
    print("ORIGINAL DATASET PARSED")

    # We are implicitly testing here that we can load the written file
    written_dataset = pydicom.dcmread(original_file)
    print("GO-DICOM DATASET PARSED")

    # Assert that the number of elements matches.
    assert sum(1 for _ in original_dataset.iterall()) == sum(1 for _ in written_dataset.iterall())

    original_elements = original_dataset.iterall()
    written_elements = written_dataset.iterall()

    # Iterate over the flattened elements.
    for original_elm, written_elm in zip(original_elements, written_elements):
        prompt = ""
        try:
            # Pydicom's DataElement type defines an __eq__ method, so we don't have to
            # do anything beyond assert equality here.
            assert original_elm == written_elm
        except BaseException as err:
            prompt = "FAILURE ON:"
            raise err
        else:
            prompt = "ELEMENT PASSED:"
        finally:
            name = original_elm.keyword
            if not name:
                name = str(original_elm.tag)
            print(f"{prompt} {name} ({original_elm.VR})")
