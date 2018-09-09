PyObject *
_Optimize(PyObject *PYBINDGEN_UNUSED(dummy), PyObject *args, PyObject *kwargs, PyObject **return_exception)
{
    PyObject *py_retval;
    struct PiumaResult retval;
    char *path;
    int width = 0;
    int height = 0;
    const char *keywords[] = {"path", "width", "height", NULL};

    if (!PyArg_ParseTupleAndKeywords(args, kwargs, (char *)"s|ii", (char **)keywords, &path, &width, &height))
    {
        return NULL;
    }

    retval = OptimizeWrapper(path, width, height);
    if (retval.message != NULL)
    {
        PyErr_SetString(PyPiumaError_Type, retval.message);
        return NULL;
    }
    py_retval = Py_BuildValue((char *)"s", retval.path);
    return py_retval;
}
