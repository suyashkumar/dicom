#include <stdlib.h>

unsigned char * J2K_Decode(unsigned char * dataIn, size_t length, int width, int height, size_t * lenOut);
unsigned char * JPEG_Decode(unsigned char * dataIn, size_t length, int bitsAllocated, int width, int height, int samplesPerPixel, int planarConfiguration, int ybr, size_t * lenOut);
unsigned char * JPEGLS_Decode(unsigned char * dataIn, size_t length, int bitsAllocated, int width, int height, int samplesPerPixel, int planarConfiguration, size_t * lenOut);
unsigned char * J2K_Encode(unsigned char * dataIn, int bitsAllocated, int width, int height, int samplesPerPixel, int planarConfiguration, size_t * lenOut);
