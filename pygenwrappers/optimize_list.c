PyObject *
_OptimizeList(PyObject * PYBINDGEN_UNUSED(dummy), PyObject *args, PyObject *kwargs, PyObject **return_exception)
{
    PyObject *py_retval;
    struct PiumaResult retval;
    PyObject *pathList;
    char **paths;
    int width;
    int height;

    if (!PyArg_ParseTuple(args, (char *) "Oii", &pathList, &width, &height)) {
        return NULL;
    }
    
    int pathsLen = PyObject_Length(pathList);

    paths = (char**) malloc(pathsLen  * sizeof(char*));
    for (int i = 0 ; i < pathsLen ; i++) {
        paths[i] = PyBytes_AS_STRING(PyUnicode_AsEncodedString((PyList_GetItem(pathList, i)), "utf-8", "~E~"));
    }
    retval = OptimizeListWrapper(paths, pathsLen, width, height);
    if (retval.message != NULL) {
        PyErr_SetString(PyPiumaError_Type, retval.message);
        return NULL;
    }
    Py_RETURN_NONE;
}
