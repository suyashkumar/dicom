#!/usr/bin/env python3.6
import json
import logging
import urllib.request
from typing import IO, NamedTuple, List

logging.basicConfig(level=logging.DEBUG)

INNOLITICS_VERSION_HASH = "8670abdd9ad16c61af5146ef857899699bdd9c5f" # rev2024b (please, update this comment when changing hash value)

INNOLITICS_CREDITS = f'''// This file's contents are derived from the innolitics json representation of the dicom standard. 
// The innolitics source is licensed as follows:
// https://github.com/innolitics/dicom-standard/blob/{INNOLITICS_VERSION_HASH}/LICENSE.txt
// 
// Copyright (c) 2017 Innolitics, LLC.
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is furnished to do
// so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.'''

Tag = NamedTuple('Tag', [
    ('group', int),
    ('elem', int),
    ('vr', List[str]),
    ('name', str),
    ('vm', str),
    ('keyword', str),
    ('retired', bool)])

def read_tags_from_json(json_string: str) -> List[Tag]:
    attrs = json.loads(json_string)
    allowed_vrs_separator = " or "
    return [
        Tag(
            # The id field should always follow format "ggggeeee", so this should be safe.
            group=int(resolvable_tag_id[:4], 16),
            elem=int(resolvable_tag_id[4:], 16),
            # To understand this ternary expression, see: https://dicom.nema.org/medical/dicom/2024a/output/html/part06.html#note_6_2.
            vr=(e["valueRepresentation"] if resolvable_tag_id[:4].lower() != 'fffe' else "NA").split(allowed_vrs_separator),
            name=e["name"],
            vm=e["valueMultiplicity"],
            keyword=e["keyword"],
            retired=e["retired"] == "Y")
        for e in attrs
        if (resolvable_tag_id := e["id"].replace("x", "0")) and len(e["keyword"]) > 0
    ]

def read_tags_from_innolitics(version_hash: str) -> List[Tag]:
    response = urllib.request.urlopen(f"https://raw.githubusercontent.com/innolitics/dicom-standard/{version_hash}/standard/attributes.json")
    json_string = response.read().decode("utf-8")
    return read_tags_from_json(json_string)

def read_tags_from_file(filename: str) -> List[Tag]:
    with open(filename, "r") as response:
        json_string = response.read()
        return read_tags_from_json(json_string)


def tagDictEntry(t: Tag) -> str:
    wrap_in_quotes = lambda s: f'\"{s}\"'
    start_indent = '	'
    return f'{start_indent}{t.keyword}: Info{{{t.keyword}, []string{{{", ".join(map(wrap_in_quotes, t.vr))}}}, "{t.name}", "{t.keyword}", "{t.vm}", {str(t.retired).lower()}}},'

def generate(out: IO[str]):
    tags = read_tags_from_file("command_group_tags.json") + read_tags_from_innolitics(INNOLITICS_VERSION_HASH)
    new_line = chr(10)
    print(f'''// AUTO-GENERATED from generate_tag_definitions.py. DO NOT EDIT.
{INNOLITICS_CREDITS}
package tag

{new_line.join(f"var {t.keyword} = Tag{{0x{t.group:04x}, 0x{t.elem:04x}}}" for t in tags)}

var tagDict = map[Tag]Info{{
{new_line.join(tagDictEntry(t) for t in tags)}
}}''', file=out)

def main():
    with open("tag_definitions.go", "w") as out:
        generate(out)

main()