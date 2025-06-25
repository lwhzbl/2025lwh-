function compose(functions) {
    if (functions.length === 0) {
        return function(x) {
            return x;
        };
    }

    return functions.reduceRight((prevFn, nextFn) => {
        return function(x) {
            return nextFn(prevFn(x));
        };
    });
}