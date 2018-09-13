#!/bin/bash

rm *.profile
rm piuma.grind

python -m cProfile -o piuma.profile piuma.py
pyprof2calltree -i piuma.profile -o piuma.grind

echo "Profile file 'piuma.grind' created!"

rm *.profile
