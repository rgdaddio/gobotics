// Selector functions take a slice of the application state and return some data based on that. 
// They never introduce any changes to the application state.

function assetListLength( assetList ) {
    return assetList.length;
}

export default {
    assetListLength
};