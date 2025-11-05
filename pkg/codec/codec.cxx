// Copyright (c) 2023 Segmed Inc.
#include "openjpeg.h"
#include "dcmtk/dcmjpeg/djcparam.h"
#include "dcmtk/dcmjpeg/djeijg8.h"
#include "dcmtk/dcmjpeg/djeijg12.h"
#include "dcmtk/dcmjpeg/djeijg16.h"
#include "dcmtk/dcmjpeg/djdijg8.h"
#include "dcmtk/dcmjpeg/djdijg12.h"
#include "dcmtk/dcmjpeg/djdijg16.h"
#include "charls/charls.h"
#include "codec.h"

/**
 * width: cols
 * height: rows
 * numcomps/samplesPerPixel: number of components/samples/planes (RGB-3/Grey-1) in the image
 * prec/bitsAllocated: number of bits per component per pixel
 * sgnd/pixelRepresentation: signed (1) / unsigned (0)
 * planarConfiguration: color-by-plane (1) or color-by-pixel (0)
* */
unsigned char * JPEGLS_Decode(unsigned char * dataSrc, size_t length, int bitsAllocated, int width, int height, int samplesPerPixel, int planarConfiguration, size_t * lenOut)
{
    charls::jpegls_decoder decoder;
    decoder.source(dataSrc, length);
    decoder.read_header();

    const size_t size{static_cast<size_t>(decoder.destination_size())};
    BYTE *newdata = new BYTE[decoder.destination_size()];

    decoder.decode(newdata, size);
    *lenOut = size;
    return newdata;
}

unsigned char * JPEG_Decode(unsigned char * dataSrc, size_t length, int bitsAllocated, int width, int height, int samplesPerPixel, int planarConfiguration, int ybr, size_t * lenOut)
{
    if (bitsAllocated == 8)
    {
        DJCodecParameter param(
            ECC_lossyRGB,
            ybr ? EDC_always : EDC_never,
            EUC_never,
            planarConfiguration ? EPC_colorByPlane : EPC_colorByPixel
            );

        int outSize = width*height*samplesPerPixel;
        BYTE *outBuf = new BYTE[outSize];

        DJDecompressIJG8Bit *decop = new DJDecompressIJG8Bit(param, ybr);
        OFCondition check1 = decop->init();
        OFCondition check2 = decop->decode(dataSrc, length, outBuf, outSize, false);
        delete decop;

        *lenOut = outSize;
        return outBuf;
    }
    else if (bitsAllocated == 12)
    {
        DJCodecParameter param(
            ECC_lossyRGB,
            ybr ? EDC_always : EDC_never,
            EUC_never,
            planarConfiguration ? EPC_colorByPlane : EPC_colorByPixel
            );

        int outSize = width*height*samplesPerPixel*2;
        BYTE *outBuf = new BYTE[outSize];

        DJDecompressIJG12Bit *decop = new DJDecompressIJG12Bit(param, ybr);
        OFCondition check1 = decop->init();
        OFCondition check2 = decop->decode(dataSrc, length, outBuf, outSize, false);
        delete decop;

        *lenOut = outSize;
        return outBuf;
    }
    else if (bitsAllocated == 16)
    {
        DJCodecParameter param(
            ECC_lossyRGB,
            ybr ? EDC_always : EDC_never,
            EUC_never,
            planarConfiguration ? EPC_colorByPlane : EPC_colorByPixel
            );

        int outSize = width*height*samplesPerPixel*2;
        BYTE *outBuf = new BYTE[outSize];

        DJDecompressIJG16Bit *decop = new DJDecompressIJG16Bit(param, ybr);
        OFCondition check1 = decop->init();
        OFCondition check2 = decop->decode(dataSrc, length, outBuf, outSize, false);
        delete decop;

        *lenOut = outSize;
        return outBuf;
    }

    return NULL;
}

unsigned char * J2K_Decode(unsigned char * dataSrc, size_t length, int width, int height, size_t * lenOut)
{
    // WARNING: OpenJPEG is very picky when there is a trailing 00 at the end of the JPC
    // so we need to make sure to remove it.
    // Marker 0xffd9 EOI End of Image (JPEG 2000 EOC End of codestream)
    // gdcmData/D_CLUNIE_CT1_J2KR.dcm contains a trailing 0xFF which apparently is ok...
    // Ref: https://github.com/malaterre/GDCM/blob/master/Source/MediaStorageAndFileFormat/gdcmJPEG2000Codec.cxx#L637
    while( length > 0 && dataSrc[length-1] != 0xd9 )
    {
        length--;
    }
    // what if 0xd9 is never found ?
    if( !( length > 0 && dataSrc[length-1] == 0xd9 ) )
    {
        return NULL;
    }
    // https://github.com/malaterre/GDCM/blob/master/Source/MediaStorageAndFileFormat/gdcmJPEG2000Codec.cxx#L656
    OPJ_CODEC_FORMAT decod_format;
    const char jp2magic[] = "\x00\x00\x00\x0C\x6A\x50\x20\x20\x0D\x0A\x87\x0A";
    if( memcmp( dataSrc, jp2magic, sizeof(jp2magic) ) == 0 )
    {
        /* JPEG-2000 compressed image data ... sigh */
        fprintf(stderr, "J2K start like JPEG-2000 compressed image data instead of codestream" );
        decod_format = OPJ_CODEC_JP2;
    }
    else
    {
        /* JPEG-2000 codestream */
        decod_format = OPJ_CODEC_J2K;
    }
    opj_codec_t * l_codec = opj_create_decompress(decod_format);

    opj_input_memory_stream fsrc;
    fsrc.dataSize = length;
    fsrc.offset = 0;
    fsrc.pData = dataSrc;

    opj_stream_t * l_stream = opj_stream_default_create(1);
    if (!l_stream){
        return NULL;
    }
    opj_stream_set_read_function(l_stream, opj_input_memory_stream_read);
    opj_stream_set_seek_function(l_stream, opj_input_memory_stream_seek);
    opj_stream_set_skip_function(l_stream, opj_input_memory_stream_skip);
    opj_stream_set_user_data(l_stream, &fsrc, NULL);
    opj_stream_set_user_data_length(l_stream, fsrc.dataSize);

    opj_dparameters_t params;
    opj_set_default_decoder_parameters(&params);
    if (!opj_setup_decoder(l_codec, &params)) {
        fprintf(stderr, "ERROR -> opj_decompress: failed to setup the decoder\n");
        opj_stream_destroy(l_stream);
        opj_destroy_codec(l_codec);
        return NULL;
    }

    opj_image_t *image;
    if (!opj_read_header(l_stream, l_codec, &image)) {
        fprintf(stderr, "ERROR -> j2k_to_image: failed to read the header\n");
        opj_stream_destroy(l_stream);
        opj_destroy_codec(l_codec);
        opj_image_destroy(image);
        return NULL;
    }
    if (!opj_decode(l_codec, l_stream, image)) {
        fprintf(stderr, "ERROR -> opj_decompress: failed to decode image!\n");
        opj_destroy_codec(l_codec);
        opj_stream_destroy(l_stream);
        opj_image_destroy(image);
        return NULL;
    }

    void *vOutBits = NULL;
    OPJ_INT32 * ptr;
    int adjustR, bitsAllocated;
    adjustR = 0;
    bitsAllocated = image->comps[0].prec;
    if (bitsAllocated == 8) {
        BYTE v;
        BYTE *newdata = new BYTE[image->x1*image->y1*image->numcomps];
        for (unsigned int c=0; c<image->numcomps; c++) {
            ptr = (OPJ_INT32 *)image->comps[c].data;
            for (unsigned int i = 0; i< image->x1*image->y1; i++){
                v = (BYTE) (*ptr + adjustR);
                ptr++;
                newdata[i*image->numcomps + c] = v;
            }
        }
        vOutBits = (void *) newdata;
    } else { // 12 or 16 bits
        OPJ_UINT16 v;
        OPJ_UINT16 *newdata = new OPJ_UINT16[image->x1*image->y1*image->numcomps];
        for (unsigned int c=0; c<image->numcomps; c++) {
            ptr = (OPJ_INT32 *)image->comps[c].data;
            for (unsigned int i = 0; i< image->x1*image->y1; i++){
                v = (OPJ_UINT16) (*ptr + adjustR);
                ptr++;
                newdata[i*image->numcomps + c] = v;
            }
        }
        vOutBits = (void *) newdata;
    }

    *lenOut = image->x1*image->y1 * (bitsAllocated==8?1:2) * image->numcomps;

    opj_image_destroy(image);
    opj_stream_destroy(l_stream);
    opj_destroy_codec(l_codec);

    return (BYTE *) vOutBits;
}

unsigned char * J2K_Encode(unsigned char * dataSrc, int bitsAllocated, int width, int height, int samplesPerPixel, int planarConfiguration, size_t * lenOut)
{
    opj_cparameters_t params;
    opj_set_default_encoder_parameters(&params);
    params.cod_format = OPJ_CODEC_J2K;
    params.irreversible = 0;
    params.tcp_numlayers = 1;
    params.cp_disto_alloc = 1;

    opj_image_cmptparm_t *cmptparm = new opj_image_cmptparm_t[samplesPerPixel];

    for (int i=0; i<samplesPerPixel; i++)
    {
        cmptparm[i].prec = bitsAllocated;
        cmptparm[i].bpp = bitsAllocated;
        cmptparm[i].sgnd = 0;
        cmptparm[i].dx = 1;
        cmptparm[i].dy = 1;
        cmptparm[i].x0 = 0;
        cmptparm[i].y0 = 0;
        cmptparm[i].w = width;
        cmptparm[i].h = height;
    }

    opj_image_t *jp2_image = opj_image_create(samplesPerPixel, cmptparm, (samplesPerPixel == 1) ? OPJ_CLRSPC_GRAY : OPJ_CLRSPC_SRGB);

    jp2_image->x0 = 0;
    jp2_image->y0 = 0;
    jp2_image->x1 = width;
    jp2_image->y1 = height;

    //dump the source image into the components
    if (samplesPerPixel == 1)
    {
        if (bitsAllocated == 8)
        {
            for (int p=0; p<width*height; p++)
                jp2_image->comps[0].data[p] = dataSrc[p];
        }
        else
        {
            //2-byte
            for (int p=0; p<width*height; p++)
                jp2_image->comps[0].data[p] = ((OPJ_UINT16 *)dataSrc)[p];
        }
    }
    else
    {
        //always 8-bit, to my understanding

        if (planarConfiguration == 0)	//0 = R1G1B1 R2G2B2, 1 = R1R2R3..G1G2G3...B1B2B3...
        {
            for (int i=0; i<samplesPerPixel; i++)
                for (int p=0; p<width*height; p++)
                    jp2_image->comps[i].data[p] = dataSrc[p*samplesPerPixel + i];
        }
        else
        {
            for (int i=0; i<samplesPerPixel; i++)
                for (int p=0; p<width*height; p++)
                    jp2_image->comps[i].data[p] = dataSrc[i*width*height + p];
        }
    }

    opj_codec_t * codec = opj_create_compress(OPJ_CODEC_J2K);
    opj_setup_encoder(codec, &params, jp2_image);


    opj_output_memory_stream streamInfo;
    streamInfo.dataSize = 0;

    opj_stream_t * stream = opj_stream_default_create(0);
    if (!stream){
        return NULL;
    }
    opj_stream_set_write_function(stream, opj_output_memory_stream_write);
    opj_stream_set_seek_function(stream, opj_output_memory_stream_seek);
    opj_stream_set_skip_function(stream, opj_output_memory_stream_skip);
    opj_stream_set_user_data(stream, &streamInfo, NULL);

    opj_start_compress(codec, jp2_image, stream);
    opj_encode(codec, stream);
    opj_end_compress(codec, stream);

    delete []cmptparm;

    *lenOut = streamInfo.dataSize;
    BYTE * vOutBits = (BYTE *) malloc(streamInfo.dataSize), *fillPtr = vOutBits;
    for (std::vector<opj_output_memory_stream_chunk>::iterator i = streamInfo.chunks.begin(); i != streamInfo.chunks.end(); i++)
    {
        memcpy(fillPtr, i->data, i->dataLen);
        delete i->data;
        fillPtr += i->dataLen;
    }

    opj_stream_destroy(stream);
    opj_destroy_codec(codec);
    opj_image_destroy(jp2_image);

    return (BYTE *) vOutBits;
}

// https://groups.google.com/forum/#!topic/openjpeg/8cebr0u7JgY
OPJ_SIZE_T opj_input_memory_stream_read(void * p_buffer, OPJ_SIZE_T p_nb_bytes, void * p_user_data)
{
    opj_input_memory_stream* l_stream = (opj_input_memory_stream*)p_user_data;
    OPJ_SIZE_T l_nb_bytes_read = p_nb_bytes;

    if (l_stream->offset >= l_stream->dataSize) {
        return (OPJ_SIZE_T)-1;
    }
    if (p_nb_bytes > (l_stream->dataSize - l_stream->offset)) {
        l_nb_bytes_read = l_stream->dataSize - l_stream->offset;
    }
    memcpy(p_buffer, &(l_stream->pData[l_stream->offset]), l_nb_bytes_read);
    l_stream->offset += l_nb_bytes_read;
    return l_nb_bytes_read;
}

OPJ_OFF_T opj_input_memory_stream_skip(OPJ_OFF_T p_nb_bytes, void * p_user_data)
{
    opj_input_memory_stream* l_stream = (opj_input_memory_stream*)p_user_data;

    if (p_nb_bytes < 0) {
        return -1;
    }

    l_stream->offset += (OPJ_SIZE_T)p_nb_bytes;

    return p_nb_bytes;
}

OPJ_BOOL opj_input_memory_stream_seek(OPJ_OFF_T p_nb_bytes, void * p_user_data)
{
    opj_input_memory_stream* l_stream = (opj_input_memory_stream*)p_user_data;

    if (p_nb_bytes < 0) {
        return OPJ_FALSE;
    }

    l_stream->offset = (OPJ_SIZE_T)p_nb_bytes;

    return OPJ_TRUE;
}


// Writer adapted from above and made it up

OPJ_BOOL opj_output_memory_stream_seek(OPJ_OFF_T p_nb_bytes, void * p_user_data)
{
    opj_output_memory_stream* stream = (opj_output_memory_stream*)p_user_data;

    return OPJ_TRUE;
}

OPJ_OFF_T opj_output_memory_stream_skip(OPJ_OFF_T p_nb_bytes, void * p_user_data)
{
    opj_output_memory_stream* stream = (opj_output_memory_stream*)p_user_data;

    return p_nb_bytes;
}

OPJ_SIZE_T opj_output_memory_stream_write(void * p_buffer, OPJ_SIZE_T p_nb_bytes, void * p_user_data)
{
    opj_output_memory_stream* stream = (opj_output_memory_stream*)p_user_data;

    opj_output_memory_stream_chunk newChunk;
    newChunk.dataLen = p_nb_bytes;
    newChunk.data = new BYTE[p_nb_bytes];
    memcpy(newChunk.data, p_buffer, p_nb_bytes);

    stream->dataSize += p_nb_bytes;

    stream->chunks.push_back(newChunk);

    return p_nb_bytes;
}
