# Third-Party Validation Testing

In order to increase confidence in our reads and writes, this library 
includes tests that call third-party tools to validate DICOM files it 
produces

These tests are not enabled by default when running `make test` or 
`go test ./...`, they must be enabled by certain build tags.

Each section below details the installation requirements and build tags for
each type of validation testing.

To run all tests, make sure you meet the all requirements for each validation
tool listed below and run:

```shell
make test-full
```

## Pydicom

[Pydicom](https://pydicom.github.io/pydicom/stable/) is a widely used and 
popular library for reading and writing dicom files in Python.

### Requirements

[Python 3.6+](https://www.python.org/downloads/)
    
- pydicom

### Dependency Installation

To install dependencies, after installing python and setting up any desired
[virtual environment](https://realpython.com/python-virtual-environments-a-primer/)
run the following command at the library root:

```shell
pip install -r validate_pydicom_requirements.txt
```

### Running the Tests

Running the pydicom validation tests requires the `pydicom` build tag, like
so:

```shell
go test -v -tags=pydicom ./...
```

### Optional Configuration

You can set the python interpreter for the pydicom validation tests to use
by setting the `PYDICOM_TEST_INTERPRETER` environmental variable. This can
be useful for pointing the tests towards your virtual environment in an IDE,
or using a non-default python interpreter on you system.

If no value is found for this environmental variable, `python3` will be 
used.
