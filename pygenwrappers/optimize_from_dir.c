PyObject *
_OptimizeFromDir(PyObject * PYBINDGEN_UNUSED(dummy), PyObject *args, PyObject *kwargs, PyObject **return_exception)
{
    PyObject *py_retval;
    struct PiumaResult retval;
    char *path;
    int width;
    int height;

    if (!PyArg_ParseTuple(args, (char *) "sii", &path, &width, &height)) {
        return NULL;
    }
    
    retval = OptimizeFromDirWrapper(path, width, height);
    if (retval.message != NULL) {
        PyErr_SetString(PyPiumaError_Type, retval.message);
        return NULL;
    }
    Py_RETURN_NONE;
}
