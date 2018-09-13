#!/bin/bash

rm *.profile
rm prof.grind

python -m cProfile -o prof.profile profileall.py
pyprof2calltree -i prof.profile -o profile.grind

echo "Profile file 'profile.grind' created!"

rm *.profile
