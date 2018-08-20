import sys

from pybindgen import Module


mod = Module('pypiuma')
mod.add_include('"piuma.h"')
with open('./pygenwrappers/optimize.c') as optimize_file:
    mod.add_custom_function_wrapper('optimize',
                                        '_Optimize',
                                        optimize_file.read())
mod.add_exception('PiumaError')
mod.generate(sys.stdout)
