#!/usr/bin/env python3.6
import requests
import logging
from typing import IO, NamedTuple, List

logging.basicConfig(level=logging.DEBUG)

INNOLITICS_VERSION_HASH = "8670abdd9ad16c61af5146ef857899699bdd9c5f" # rev2024b (please, update this comment when changing hash value)

Tag = NamedTuple('Tag', [
    ('group', int),
    ('elem', int),
    ('vr', str),
    ('name', str),
    ('vm', str),
    ('keyword', str),
    ('retired', bool)])

def read_tags_from_innolitics(version_hash: str) -> List[Tag]:
    response = requests.get(f"https://raw.githubusercontent.com/innolitics/dicom-standard/{version_hash}/standard/attributes.json")
    response.raise_for_status()
    return [
        Tag(
            # The id field should always follow format "ggggeeee", so this should be safe.
            group=int(resolvable_tag_id[:4], 16),
            elem=int(resolvable_tag_id[4:], 16),
            # To understand this ternary expression, see: https://dicom.nema.org/medical/dicom/current/output/html/part06.html#note_6_2.
            vr=e["valueRepresentation"] if resolvable_tag_id[:4].lower() != 'fffe' else "NA",
            name=e["name"],
            vm=e["valueMultiplicity"],
            keyword=e["keyword"],
            retired=e["retired"] == "Y")
        for e in response.json()
        if (resolvable_tag_id := e["id"].replace("x", "0")) and len(e["keyword"]) > 0
    ]

def generate(out: IO[str]):
    tags = read_tags_from_innolitics(INNOLITICS_VERSION_HASH)

    print("// AUTO-GENERATED from generate_tag_definitions.py. DO NOT EDIT.", file=out)
    print("package tag", file=out)
    print(file=out)
    for t in tags:
        print(f"var {t.keyword} = Tag{{0x{t.group:04x}, 0x{t.elem:04x}}}", file=out)

    print(file=out)
    print("var tagDict map[Tag]Info", file=out)
    print("", file=out)
    print("func init() {", file=out)
    print("	maybeInitTagDict()", file=out)
    print("}", file=out)
    print("func maybeInitTagDict() {", file=out)
    print("	if len(tagDict) > 0 {", file=out)
    print("		return", file=out)
    print("	}", file=out)
    print("	tagDict = make(map[Tag]Info)", file=out)
    for t in tags:
        retired = f"{t.retired}".lower()
        print(f'	tagDict[{t.keyword}] = Info{{{t.keyword}, "{t.vr}", "{t.name}", "{t.keyword}", "{t.vm}", {retired}}}', file=out)
    print("}", file=out)

def main():
    with open("tag_definitions.go", "w") as out:
        generate(out)

main()