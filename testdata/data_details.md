# Test files Details

This document contains citations and further details for the test DICOMs used
here. Eventually I would like to store a way more expansive set of test DICOMs
in a cloud bucket that we can pull down for testing (this would allow for more 
DICOMs and larger files). This would, however add some complexity to testing 
this repository.  


## Files from The Cancer Imaging Archive

The sub-bullets mention potentially interesting characteristics of the test file.
Some of the interesting characteristics may apply to more than one file, but may only
be mentioned in one of them for brevity.

* [1.dcm](1.dcm) (from [#77](https://github.com/suyashkumar/dicom/issues/77)) 
  * Modality: PET 
  * Native Pixel Data
  * Doubly Nested Sequences
  * Icon pixel data in addition to typical pixel data 
* [2.dcm](2.dcm) (from [#77](https://github.com/suyashkumar/dicom/issues/77))
  * Modality: PET
  * Other items similar to 1.dcm
* [3.dcm](3.dcm)
  * Modality: MRI 
  * Native pixel data
* [4.dcm](4.dcm)
  * Modality: X-Ray
  * Native pixel data
  * COVID-19 positive
* [5.dcm](5.dcm)
  * Modality: CT
  * Multiple frames
  * Native pixel data
### Relevant Citations
#### For files 1.dcm, 2.dcm:
##### Data Citation:
Kinahan, Paul; Muzi, Mark; Bialecki, Brian; Coombs, Laura. (2017). Data from ACRIN-FLT-Breast. The Cancer Imaging Archive. http://doi.org/10.7937/K9/TCIA.2017.ol20zmxg

##### Publication Citation
Kostakoglu L ,  Duan F ,  Idowu MO ,  Jolles PR ,  Bear HD ,  Muzi M ,  Cormack J ,  Muzi JP ,  Pryma DA ,  Specht JM ,  Hovanessian-Larsen L ,  Miliziano J ,  Mallett S ,  Shields AF ,  Mankoff DA ;  ACRIN 668 Investigative Team . A Phase II Study of 3'-Deoxy-3'-18F-Fluorothymidine PET in the Assessment of Early Response of Breast Cancer to Neoadjuvant Chemotherapy: Results from ACRIN 6688. J Nucl Med. 2015 Nov;56(11):1681-9. doi: 10.2967/jnumed.115.160663. Epub 2015 Sep 10. 


#### For file 3.dcm

##### Data Citation:
Choyke P, Turkbey B, Pinto P, Merino M, Wood B. (2016). Data From PROSTATE-MRI. The Cancer Imaging Archive. http://doi.org/10.7937/K9/TCIA.2016.6046GUDv


#### TCIA Citation
Clark K, Vendt B, Smith K, Freymann J, Kirby J, Koppel P, Moore S, Phillips S, Maffitt D, Pringle M, Tarbox L, Prior F. The Cancer Imaging Archive (TCIA): Maintaining and Operating a Public Information Repository, Journal of Digital Imaging, Volume 26, Number 6, December, 2013, pp 1045-1057.

#### For file 4.dcm

##### Data Citation:
Desai, S., Baghal, A., Wongsurawat, T., Al-Shukri, S., Gates, K., Farmer, P., Rutherford, M., Blake, G.D., Nolan, T., Powell, T., Sexton, K., Bennett, W., Prior, F. (2020). Data from Chest Imaging with Clinical and Genomic Correlates Representing a Rural COVID-19 Positive Population [Data set]. The Cancer Imaging Archive. DOI: https://doi.org/10.7937/tcia.2020.py71-5978.

#### TCIA Citation
Clark K, Vendt B, Smith K, Freymann J, Kirby J, Koppel P, Moore S, Phillips S, Maffitt D, Pringle M, Tarbox L, Prior F. The Cancer Imaging Archive (TCIA): Maintaining and Operating a Public Information Repository, Journal of Digital Imaging, Volume 26, Number 6, December, 2013, pp 1045-1057. DOI: 10.1007/s10278-013-9622-7

#### File 5.dcm
This file was sourced from [cornerstone](https://github.com/cornerstonejs/dicomParser/blob/master/testImages/encapsulated/multi-frame/CT0012.explicit_little_endian.dcm) 
(which is MIT licensed, see the license reproduced in included_licenses.md)
