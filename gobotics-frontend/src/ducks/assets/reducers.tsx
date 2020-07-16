import { combineReducers } from "redux";

import actions from "./actions";
import {AddAssetAction, AddAssetsAction} from "./actions"
import Asset from "./types";

interface AssetState {
    assets: Asset[]
}

const initialState : AssetState = {
    assets: []
}

// typescript type for the reducer's action
type AddAssetType = AddAssetAction
type AddAssetsType = AddAssetsAction
type AssetActionType = AddAssetType | AddAssetsType

// Reducers
const addAssetListReducer = (state = initialState, action: AssetActionType) => {
    switch(action.type) {
        case actions.ADD_ASSET: return {
            ...state,
            assets: state.assets.concat(action.payload)
        }
        case actions.ADD_ASSETS: return {
            ...state,
            assets: state.assets.concat(action.payload)
        }
        default: return state
    }
}   

const reducer = combineReducers( {
    assetList: addAssetListReducer,
    // In the future if the module needs more reducers
    //distance: distanceReducer
} );

export default reducer;
