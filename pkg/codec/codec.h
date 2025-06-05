// Copyright (c) 2023 Segmed Inc.
#include "openjpeg.h"
#include <vector>

typedef unsigned char BYTE;
typedef struct
{
    OPJ_UINT8* pData;
    OPJ_SIZE_T dataSize;
    OPJ_SIZE_T offset;
} opj_input_memory_stream;

OPJ_BOOL opj_input_memory_stream_seek(OPJ_OFF_T p_nb_bytes, void * p_user_data);
OPJ_OFF_T opj_input_memory_stream_skip(OPJ_OFF_T p_nb_bytes, void * p_user_data);
OPJ_SIZE_T opj_input_memory_stream_read(void * p_buffer, OPJ_SIZE_T p_nb_bytes, void * p_user_data);

typedef struct
{
    OPJ_UINT8* data;
    OPJ_SIZE_T dataLen;
} opj_output_memory_stream_chunk;

typedef struct
{
    OPJ_SIZE_T dataSize;
    std::vector<opj_output_memory_stream_chunk> chunks;
} opj_output_memory_stream;

OPJ_BOOL opj_output_memory_stream_seek(OPJ_OFF_T p_nb_bytes, void * p_user_data);
OPJ_OFF_T opj_output_memory_stream_skip(OPJ_OFF_T p_nb_bytes, void * p_user_data);
OPJ_SIZE_T opj_output_memory_stream_write(void * p_buffer, OPJ_SIZE_T p_nb_bytes, void * p_user_data);

