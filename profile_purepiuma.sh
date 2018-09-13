#!/bin/bash

rm *.profile
rm purepiuma.grind

python -m cProfile -o purepiuma.profile purepiuma.py
pyprof2calltree -i purepiuma.profile -o purepiuma.grind

echo "Profile file 'purepiuma.grind' created!"

rm *.profile
