import sys

from pybindgen import Module


mod = Module('pypiuma')
mod.add_include('"libpiuma.h"')
with open('./pygenwrappers/optimize.c') as optimize_file:
    mod.add_custom_function_wrapper('optimize',
                                        '_Optimize',
                                        optimize_file.read())
with open('./pygenwrappers/optimize_list.c') as optimize_file:
    mod.add_custom_function_wrapper('optimize_list',
                                        '_OptimizeList',
                                        optimize_file.read())
with open('./pygenwrappers/optimize_from_dir.c') as optimize_file:
    mod.add_custom_function_wrapper('optimize_dir',
                                        '_OptimizeFromDir',
                                        optimize_file.read())
mod.add_exception('PiumaError')
mod.generate(sys.stdout)
