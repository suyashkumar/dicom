#!/usr/bin/env python3.6
import json
import logging
import urllib.request
from typing import IO, NamedTuple, List

logging.basicConfig(level=logging.DEBUG)

INNOLITICS_VERSION_HASH = "8670abdd9ad16c61af5146ef857899699bdd9c5f" # rev2024b (please, update this comment when changing hash value)

Tag = NamedTuple('Tag', [
    ('group', int),
    ('elem', int),
    ('vr', List[str]),
    ('name', str),
    ('vm', str),
    ('keyword', str),
    ('retired', bool)])

def read_tags_from_innolitics(version_hash: str) -> List[Tag]:
    response = urllib.request.urlopen(f"https://raw.githubusercontent.com/innolitics/dicom-standard/{version_hash}/standard/attributes.json")
    attrs = json.loads(response.read().decode())
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

def generate(out: IO[str]):
    tags = read_tags_from_innolitics(INNOLITICS_VERSION_HASH)
    new_line = chr(10)
    wrap_in_quotes = lambda s: f'\"{s}\"'
    print(f'''// AUTO-GENERATED from generate_tag_definitions.py. DO NOT EDIT.
package tag

{new_line.join(f"var {t.keyword} = Tag{{0x{t.group:04x}, 0x{t.elem:04x}}}" for t in tags)}

var tagDict map[Tag]Info

func init() {{
	maybeInitTagDict()
}}

func maybeInitTagDict() {{
	if len(tagDict) > 0 {{
		return
	}}
	tagDict = make(map[Tag]Info)
{new_line.join(f'	tagDict[{t.keyword}] = Info{{{t.keyword}, []string{{{", ".join(map(wrap_in_quotes, t.vr))}}}, "{t.name}", "{t.keyword}", "{t.vm}", {str(t.retired).lower()}}}' for t in tags)}
}}''', file=out)

def main():
    with open("tag_definitions.go", "w") as out:
        generate(out)

main()