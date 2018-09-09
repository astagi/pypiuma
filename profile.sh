#!/bin/bash

rm *.profile
rm *.grind

python -m cProfile -o piuma.profile piuma.py
pyprof2calltree -i piuma.profile -o piuma.grind

echo "Profile file 'piuma.grind' created!"


python -m cProfile -o purepiuma.profile purepiuma.py
pyprof2calltree -i purepiuma.profile -o purepiuma.grind

echo "Profile file 'purepiuma.grind' created!"

rm *.profile
