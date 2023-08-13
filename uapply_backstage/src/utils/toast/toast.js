/*
* Used to call the toast out
* @name : name of the toast you want to show
*/
export function showToast(name){
    const currentToast = name;
    console.log(currentToast);
}

// show or hide the toast
export function isToast(isShow){
    return !isShow
}
