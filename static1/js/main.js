function istel(str) {
    var reg = /^1[3-9]\d{9}$/;
    return reg.test(str);
}

function thareIs(str) {
    if (str == "" || str == null) {
        return false
    } else {
        return true;
    }
}