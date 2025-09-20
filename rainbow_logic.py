#!/usr/bin/env python3
import math

freq = 0.1
spread = 3.0
seed = 1

red_s = 0
blue_s = 4 * (math.pi / 3)
green_s = 2 * (math.pi / 3)

max_lines = 100
max_line_len = 100
line_num = 0

offset = 0
for i in range(max_lines):
    for line_pos in range(max_line_len):
        print(
            f"i: {i}\n\
    r: {math.sin(freq*(line_pos+offset/spread) + red_s) * 127 + 127}\n\
    g: {math.sin(freq*(offset/spread) + green_s) * 127 + 127}\n\
    b: {math.sin(freq*(offset/spread) + blue_s) * 127 + 127}\n\n"
        )
