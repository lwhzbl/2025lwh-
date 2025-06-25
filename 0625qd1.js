function isInstance(value, classOrSuperclass) {
    if (value === undefined || value === null || classOrSuperclass === undefined || classOrSuperclass === null) {
        return false;
    }

    if (typeof classOrSuperclass !== 'function') {
        return false;
    }

    try {
        return value instanceof classOrSuperclass;
    } catch (e) {
        return false;
    }
}    