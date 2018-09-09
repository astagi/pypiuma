PyObject *
_OptimizeFromDir(PyObject *PYBINDGEN_UNUSED(dummy), PyObject *args, PyObject *kwargs, PyObject **return_exception)
{
    PyObject *py_retval;
    struct PiumaResult retval;
    char *path;
    int width = 0;
    int height = 0;

    if (!PyArg_ParseTuple(args, (char *)"s|ii", &path, &width, &height))
    {
        return NULL;
    }

    retval = OptimizeFromDirWrapper(path, width, height);
    if (retval.message != NULL)
    {
        PyErr_SetString(PyPiumaError_Type, retval.message);
        return NULL;
    }
    Py_RETURN_NONE;
}
